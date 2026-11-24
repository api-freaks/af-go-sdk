# Apifreaks Go SDK

[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-Built%20with%20Fern-brightgreen)](https://buildwithfern.com?utm_source=github&utm_medium=github&utm_campaign=readme&utm_source=Apifreaks%2FGo)
[![Go Reference](https://pkg.go.dev/badge/github.com/api-freaks/af-go-sdk/v2.svg)](https://pkg.go.dev/github.com/api-freaks/af-go-sdk/v2)

The Apifreaks Go library provides convenient access to the Apifreaks APIs from Go.

## Table of Contents

* [Installation](#installation)
* [Reference](#reference)
* [Usage](#usage)
* [Environments](#environments)
* [Errors](#errors)
* [Request Types](#request-types)
* [Raw Responses](#raw-responses)
* [Advanced](#advanced)

  * [Retries](#retries)
  * [Timeouts](#timeouts)
  * [Additional Headers](#additional-headers)
  * [Additional Query String Parameters](#additional-query-string-parameters)
* [Contributing](#contributing)

## Installation

Install the SDK with:

```sh
go get github.com/api-freaks/af-go-sdk/v2
```

Then import the SDK in your Go code:

```go
import (
    sdk "github.com/api-freaks/af-go-sdk/v2"
    "github.com/api-freaks/af-go-sdk/v2/client"
)
```

## Reference

A full reference for this library is available [here](./reference.md).

> **Note:** `GetCreditsUsage` (`get_credits_usage`) requires an organization admin API key. Other keys get a `403` (`*sdk.ForbiddenError`).

## Usage

Instantiate and use the client with the following:

```go
package main

import (
    "context"
    "fmt"
    "log"

    sdk "github.com/api-freaks/af-go-sdk/v2"
    "github.com/api-freaks/af-go-sdk/v2/client"
)

func main() {
    c := client.NewClient()

    response, err := c.GeolocationLookup(
        context.Background(),
        &sdk.GeolocationLookupRequest{
            APIKey: "your_api_key",
            IP:     sdk.String("8.8.8.8"),
        },
    )
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%+v\n", response)
}
```

### Domain WHOIS example

```go
package main

import (
    "context"
    "fmt"
    "log"

    sdk "github.com/api-freaks/af-go-sdk/v2"
    "github.com/api-freaks/af-go-sdk/v2/client"
)

func main() {
    c := client.NewClient()

    response, err := c.DomainWhoisLookup(
        context.Background(),
        &sdk.DomainWhoisLookupRequest{
            APIKey:     "your_api_key",
            DomainName: "example.com",
        },
    )
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%+v\n", response)
}
```

### Bulk geolocation example

```go
package main

import (
    "context"
    "fmt"
    "log"

    sdk "github.com/api-freaks/af-go-sdk/v2"
    "github.com/api-freaks/af-go-sdk/v2/client"
)

func main() {
    c := client.NewClient()

    response, err := c.BulkGeolocationLookup(
        context.Background(),
        &sdk.BulkGeolocationLookupRequest{
            APIKey: "your_api_key",
            Ips: []string{
                "8.8.8.8",
                "1.1.1.1",
            },
        },
    )
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%+v\n", response)
}
```

## Environments

This SDK allows you to configure different environments for API requests.

```go
package main

import (
    sdk "github.com/api-freaks/af-go-sdk/v2"
    "github.com/api-freaks/af-go-sdk/v2/client"
    "github.com/api-freaks/af-go-sdk/v2/option"
)

func main() {
    c := client.NewClient(
        option.WithBaseURL(sdk.Environments.Default),
    )

    _ = c
}
```

## Errors

When the API returns a non-success status code, the SDK returns an error. You can inspect specific Apifreaks errors or the generic API error type.

```go
package main

import (
    "context"
    "errors"
    "fmt"
    "log"

    sdk "github.com/api-freaks/af-go-sdk/v2"
    "github.com/api-freaks/af-go-sdk/v2/client"
    "github.com/api-freaks/af-go-sdk/v2/core"
)

func main() {
    c := client.NewClient()

    _, err := c.GeolocationLookup(
        context.Background(),
        &sdk.GeolocationLookupRequest{
            APIKey: "your_api_key",
            IP:     sdk.String("8.8.8.8"),
        },
    )
    if err == nil {
        return
    }

    var badRequest *sdk.BadRequestError
    if errors.As(err, &badRequest) {
        // Body is the parsed error payload (timestamp, status, error, message, path).
        if body := badRequest.GetBody(); body != nil && body.Message != nil {
            fmt.Printf("Bad request: %s\n", *body.Message)
        }
        return
    }

    var apiErr *core.APIError
    if errors.As(err, &apiErr) {
        fmt.Printf("API error %d: %s\n", apiErr.StatusCode, apiErr.Error())
        return
    }

    log.Fatal(err)
}
```

If an error response body isn't a JSON object of that shape (for example `format=xml`), the SDK returns the plain `*core.APIError` instead of the typed error; its `Error()` still contains the raw body.

If a successful (2xx) response body can't be decoded into the expected response type, the SDK returns a `*core.APIError` that carries the status code and headers and wraps a `*core.ResponseDecodeError`. That error holds the raw response `Body` and the underlying decoding error (for example `*json.UnmarshalTypeError`), and `errors.As` can reach each of them.

## Request Types

The SDK exports all request types as Go structs from the root package. Import the root package and pass request structs into the client methods.

```go
import sdk "github.com/api-freaks/af-go-sdk/v2"

request := &sdk.GeolocationLookupRequest{
    APIKey: "your_api_key",
    IP:     sdk.String("8.8.8.8"),
}
```

Pointer helper functions are also exported from the root package, for example `sdk.String`, `sdk.Int`, `sdk.Float64`, `sdk.Bool`, `sdk.Time`, `sdk.MustParseDate`, and `sdk.MustParseDateTime`.

## Raw Responses

Use `WithRawResponse` when you need the response body together with the HTTP status code and headers.

```go
raw, err := c.WithRawResponse.GeolocationLookup(
    context.Background(),
    &sdk.GeolocationLookupRequest{
        APIKey: "your_api_key",
        IP:     sdk.String("8.8.8.8"),
    },
)
if err != nil {
    log.Fatal(err)
}

fmt.Println(raw.StatusCode)
fmt.Println(raw.Header)
fmt.Printf("%+v\n", raw.Body)
```

## Advanced

### Retries

The SDK is instrumented with automatic retries. Use `option.WithMaxAttempts` to configure retry attempts, or `option.WithoutRetries` to disable retries for a client or individual request.

```go
c := client.NewClient(
    option.WithMaxAttempts(3),
)
```

For a single request:

```go
response, err := c.GeolocationLookup(
    context.Background(),
    &sdk.GeolocationLookupRequest{
        APIKey: "your_api_key",
        IP:     sdk.String("8.8.8.8"),
    },
    option.WithMaxAttempts(3),
)
```

Disable retries:

```go
response, err := c.GeolocationLookup(
    context.Background(),
    &sdk.GeolocationLookupRequest{
        APIKey: "your_api_key",
        IP:     sdk.String("8.8.8.8"),
    },
    option.WithoutRetries(),
)
```

### Timeouts

Use a custom `http.Client` to configure request timeouts.

```go
package main

import (
    "net/http"
    "time"

    "github.com/api-freaks/af-go-sdk/v2/client"
    "github.com/api-freaks/af-go-sdk/v2/option"
)

func main() {
    httpClient := &http.Client{
        Timeout: 30 * time.Second,
    }

    c := client.NewClient(
        option.WithHTTPClient(httpClient),
    )

    _ = c
}
```

### Additional Headers

You can add custom headers to requests using `option.WithHTTPHeader`.

```go
headers := http.Header{}
headers.Set("X-Custom-Header", "custom-value")

response, err := c.GeolocationLookup(
    context.Background(),
    &sdk.GeolocationLookupRequest{
        APIKey: "your_api_key",
        IP:     sdk.String("8.8.8.8"),
    },
    option.WithHTTPHeader(headers),
)
```

### Additional Query String Parameters

You can add custom query parameters to requests using `option.WithQueryParameters`.

```go
query := url.Values{}
query.Set("fields", "location,currency")

response, err := c.GeolocationLookup(
    context.Background(),
    &sdk.GeolocationLookupRequest{
        APIKey: "your_api_key",
        IP:     sdk.String("8.8.8.8"),
    },
    option.WithQueryParameters(query),
)
```

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically.
Additions made directly to this library would have to be moved over to our generation code,
otherwise they would be overwritten upon the next generated release. Feel free to open a PR as
a proof of concept, but know that we will not be able to merge it as-is. We suggest opening
an issue first to discuss with us!

On the other hand, contributions to the README are always very welcome!
