# Reference
<details><summary><code>client.GeolocationLookup() -> *sdk.GeolocationLookupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get detailed geolocation data for an IP address including country, city, timezone, currency, and optional security and user-agent information
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GeolocationLookupRequest{
        APIKey: "apiKey",
    }
client.GeolocationLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.GeolocationLookupRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**ip:** `*string` — IPv4, IPv6, or hostname for geolocation lookup
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*sdk.GeolocationLookupRequestLang` — Response language for location fields
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — Comma separated list of fields to include in response
    
</dd>
</dl>

<dl>
<dd>

**excludes:** `*string` — Comma separated list of fields to exclude from response
    
</dd>
</dl>

<dl>
<dd>

**include:** `*string` — Additional data to include (location, network, security, currency, time_zone, user_agent, country_metadata , hostname, liveHostname, hostnameFallbackLivet)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkGeolocationLookup(request) -> []*sdk.BulkGeolocationLookupResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve detailed geolocation data for multiple IP addresses in a single request.
Supports up to `50,000` IP-addresses/host-names per request.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.BulkGeolocationLookupRequest{
        APIKey: "apiKey",
        Ips: []string{
            "ips",
        },
    }
client.BulkGeolocationLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.BulkGeolocationLookupRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*string` — Language of the response.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include in the response. Can include "geo".
    
</dd>
</dl>

<dl>
<dd>

**excludes:** `*string` — Comma-separated list of fields to exclude from the response (except "ip").
    
</dd>
</dl>

<dl>
<dd>

**include:** `*string` — Comma-separated list of additional information to include in the response.
    
</dd>
</dl>

<dl>
<dd>

**ips:** `[]string` — List of IP addresses or hostnames to lookup
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.IPSecurityLookup() -> *sdk.IPSecurityLookupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get comprehensive security information for a given IP address. Detects VPNs, proxies, Tor nodes, and other security threats.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.IPSecurityLookupRequest{
        APIKey: "apiKey",
    }
client.IPSecurityLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.IPSecurityLookupRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**ip:** `*string` — A valid IPv4 or IPv6 address to look up. If omitted, the API uses the public IP of the requesting client.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to return. Supports dot notation (e.g. security.threat_score).
    
</dd>
</dl>

<dl>
<dd>

**excludes:** `*string` — Comma-separated list of fields to remove from the response. Supports dot notation (e.g. security.is_tor).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkIPSecurityLookup(request) -> []*sdk.BulkIPSecurityLookupResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The Bulk IP Security Lookup API allows you to retrieve security details for up to `50,000` IP-addresses in a single request.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.BulkIPSecurityLookupRequest{
        APIKey: "apiKey",
        Ips: []string{
            "ips",
        },
    }
client.BulkIPSecurityLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.BulkIPSecurityLookupRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to return. Supports dot notation (e.g. security.threat_score).
    
</dd>
</dl>

<dl>
<dd>

**excludes:** `*string` — Comma-separated list of fields to remove from the response. Supports dot notation (e.g. security.is_tor).
    
</dd>
</dl>

<dl>
<dd>

**ips:** `[]string` — List of IP addresses to lookup
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GeocoderSearch() -> []*sdk.GeocoderSearchResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Convert a given address or place name into geographic coordinates (latitude and longitude).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GeocoderSearchRequest{
        APIKey: "apiKey",
        Query: "query",
    }
client.GeocoderSearch(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.GeocoderSearchRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**query:** `string` — Free-form search query, e.g. Wembley Stadium, London
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Max number of results to return (1–40). May return fewer if matches are weak.
    
</dd>
</dl>

<dl>
<dd>

**minLat:** `*float64` — Minimum latitude for the viewbox. Must be ≤ max_lat and between -90 and 90.
    
</dd>
</dl>

<dl>
<dd>

**maxLat:** `*float64` — Maximum latitude for the viewbox. Must be ≥ min_lat and between -90 and 90.
    
</dd>
</dl>

<dl>
<dd>

**minLon:** `*float64` — Minimum longitude for the viewbox. Must be ≤ max_lon and between -180 and 180.
    
</dd>
</dl>

<dl>
<dd>

**maxLon:** `*float64` — Maximum longitude for the viewbox. Must be ≥ min_lon and between -180 and 180.
    
</dd>
</dl>

<dl>
<dd>

**acceptLanguage:** `*string` — Preferred language order for showing search results. This may either be a simple comma-separated list of language codes or a single entry. The results will be in the 1st language which is matched from the header. As a fallback if the results are not supported in the given language, 'en' will be used.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GeocoderReverse() -> *sdk.GeocoderReverseResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Convert geographic coordinates (latitude and longitude) into a human-readable address or place name.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GeocoderReverseRequest{
        APIKey: "apiKey",
        Lat: 1.1,
        Lon: 1.1,
    }
client.GeocoderReverse(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.GeocoderReverseRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**lat:** `float64` — WGS84 latitude value ranging from -90 to 90.
    
</dd>
</dl>

<dl>
<dd>

**lon:** `float64` — WGS84 longitude value ranging from -180 to 180.
    
</dd>
</dl>

<dl>
<dd>

**acceptLanguage:** `*string` — Preferred language order for showing search results. This may either be a simple comma-separated list of language codes or a single entry. The results will be in the 1st language which is matched from the header. As a fallback if the results are not supported in the given language, en will be used.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DomainWhoisLookup() -> *sdk.DomainWhoisLookupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve current WHOIS information for a domain name.
This endpoint provides detailed registration information including registrar details,
dates, nameservers, and registrant information.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.DomainWhoisLookupRequest{
        APIKey: "apiKey",
        DomainName: "domainName",
    }
client.DomainWhoisLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.DomainWhoisLookupRequestFormat` — Response format (defaults to json)
    
</dd>
</dl>

<dl>
<dd>

**domainName:** `string` — Domain name for WHOIS lookup
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkDomainWhoisLookup(request) -> *sdk.BulkDomainWhoisLookupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve WHOIS information for `100 Domains per Request`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.BulkDomainWhoisLookupRequest{
        APIKey: "apiKey",
        DomainNames: []string{
            "domainNames",
        },
    }
client.BulkDomainWhoisLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.BulkDomainWhoisLookupRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**domainNames:** `[]string` — A list of domain names for which WHOIS data is requested.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.IPWhoisLookup() -> *sdk.IPWhoisLookupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns WHOIS registration details for a specified IP address (IPv4 or IPv6).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.IPWhoisLookupRequest{
        APIKey: "apiKey",
        IP: "ip",
    }
client.IPWhoisLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.IPWhoisLookupRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**ip:** `string` — The IP address (IPv4 or IPv6) for which WHOIS data is requested.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AsnWhoisLookup() -> *sdk.AsnWhoisLookupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns WHOIS registration details for a specified ASN, with or without the 'as' prefix.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.AsnWhoisLookupRequest{
        APIKey: "apiKey",
        Asn: "asn",
    }
client.AsnWhoisLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.AsnWhoisLookupRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**asn:** `string` — The Autonomous System Number (ASN) to retrieve WHOIS data for. Can be prefixed with 'as' or not.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DomainWhoisHistory() -> *sdk.DomainWhoisHistoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve historical WHOIS records for a domain name.
This endpoint provides a timeline of all recorded changes in domain registration information.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.DomainWhoisHistoryRequest{
        APIKey: "apiKey",
        DomainName: "domainName",
    }
client.DomainWhoisHistory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.DomainWhoisHistoryRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**domainName:** `string` — Domain name for historical WHOIS lookup
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DomainWhoisReverse() -> *sdk.DomainWhoisReverseResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Performs a reverse WHOIS search using one or more search parameters like keyword, email, owner, or company.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.DomainWhoisReverseRequest{
        APIKey: "apiKey",
    }
client.DomainWhoisReverse(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.DomainWhoisReverseRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**keyword:** `*string` — Keyword search term for reverse WHOIS by keyword (case-insensitive pattern matching).
    
</dd>
</dl>

<dl>
<dd>

**email:** `*string` — Email search term for reverse WHOIS by email address (case-insensitive exact or regex match; * wildcard supported).
    
</dd>
</dl>

<dl>
<dd>

**owner:** `*string` — Registrant or owner name for reverse WHOIS (a full-text search phrase matching technique to retrieve results).
    
</dd>
</dl>

<dl>
<dd>

**company:** `*string` — Organization or company name for reverse WHOIS (full-text search phrase matching technique to retrieve results).
    
</dd>
</dl>

<dl>
<dd>

**exact:** `*bool` — Accepts 'true' or 'false'. "true" returns only records that exactly match the input (keyword, owner/registrant, or company). "false" returns all matches and is the default when omitted.
    
</dd>
</dl>

<dl>
<dd>

**mode:** `*sdk.DomainWhoisReverseRequestMode` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page number for paginated results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DomainDNSLookup() -> *sdk.DomainDNSLookupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve real-time DNS records for any hostname. Supports multiple record types including A, AAAA, MX, NS, SOA, SPF, TXT, and CNAME records.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.DomainDNSLookupRequest{
        APIKey: "apiKey",
        Type: []*string{
            sdk.String(
                "type",
            ),
        },
    }
client.DomainDNSLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.DomainDNSLookupRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**hostName:** `*string` — Hostname or URL whose DNS records are required.
    
</dd>
</dl>

<dl>
<dd>

**ipAddress:** `*string` — The IP address for requested DNS's PTR record. 'type' parameter must be set to 'all'.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` — A comma-separated list of DNS record types for lookup. Possible values: A, AAAA, MX, NS, SOA, SPF, TXT, CNAME, or all. When ipAddress is provided, type must be "all".
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkDomainDNSLookup(request) -> *sdk.BulkDomainDNSLookupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Perform DNS lookups for multiple hostnames in a single request. Supports up to `100 host-names per request`
and returns DNS records including A, AAAA, MX, NS, SOA, SPF, TXT, and CNAME records.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.BulkDomainDNSLookupRequest{
        APIKey: "apiKey",
        Type: []*string{
            sdk.String(
                "type",
            ),
        },
        DomainNames: []string{
            "domainNames",
        },
    }
client.BulkDomainDNSLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.BulkDomainDNSLookupRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` 

A comma-separated list of DNS record types for lookup.
Possible values: A, AAAA, MX, NS, SOA, SPF, TXT, CNAME, or all
    
</dd>
</dl>

<dl>
<dd>

**domainNames:** `[]string` — List of hostnames to lookup DNS records for
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DomainDNSHistory() -> *sdk.DomainDNSHistoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve historical DNS records for any hostname. Access unique historical data for A, AAAA, MX, NS, SOA, SPF, TXT, and CNAME records,
including subdomains. Results are paginated with up to 100 unique records per page.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.DomainDNSHistoryRequest{
        APIKey: "apiKey",
        HostName: "host-name",
        Type: []*string{
            sdk.String(
                "type",
            ),
        },
    }
client.DomainDNSHistory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.DomainDNSHistoryRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**hostName:** `string` — Hostname or URL whose historical DNS records are required
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` 

A comma-separated list of DNS record types for lookup.
Possible values: A, AAAA, MX, NS, SOA, SPF, TXT, CNAME, or all
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page number for paginated results
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DomainDNSReverse() -> *sdk.DomainDNSReverseResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all the hostnames associated with any particular A, AAAA, MX, NS, SOA, SPF, TXT, and CNAME DNS records. For instance, you can access all the hostnames hosted on any IP/CIDR notation, all the domain names using Cloudflare name servers, and all the domain names using Google Mailbox
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.DomainDNSReverseRequest{
        APIKey: "apiKey",
        Type: sdk.DomainDNSReverseRequestTypeA,
        Value: "value",
    }
client.DomainDNSReverse(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.DomainDNSReverseRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*sdk.DomainDNSReverseRequestType` 

The type of reverse DNS lookup to perform. Determines how the value parameter is interpreted:
- A: IPv4 CIDR block
- AAAA: IPv6 CIDR block
- MX: Mail provider domain
- NS: Name server provider hostname
- SOA: SOA record admin domain
- SPF/TXT: Target verification strings
- CNAME: Target hostname
    
</dd>
</dl>

<dl>
<dd>

**value:** `string` — Provide an IP or CIDR for A/AAAA lookups, or a hostname/selector for MX, NS, SOA, SPF, TXT, and CNAME queries. Wildcard regex patterns are also supported (e.g., mail.google.com, m*.google.com, _spf.g*.com, s*.g*.com).
    
</dd>
</dl>

<dl>
<dd>

**exact:** `*bool` — Accepts 'true' or 'false'. "true" returns only records that exactly match the input (NS, MX, CNAME, SOA, SPF, TXT). "false" returns all matches (default when omitted).
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page number to paginate through results (defaults to 1).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.WebScrape(request) -> *sdk.WebScrapeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Execute a series of web scraping instructions on a target URL. 
Supports various operations like form filling, clicking, data extraction, and CAPTCHA solving.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.WebScrapeRequest{
        APIKey: "apiKey",
        URL: "https://example.com",
        Body: &sdk.WebScrapeRequestBody{
            WebScrapeRequestBodyBlockURL: &sdk.WebScrapeRequestBodyBlockURL{
                BlockURL: []string{
                    "https://example.com/ads.js",
                    "https://tracker.example.com/*",
                },
                Cookies: []*sdk.WebScrapeRequestBodyBlockURLCookiesItem{
                    &sdk.WebScrapeRequestBodyBlockURLCookiesItem{
                        Name: "sessionid",
                        Value: "abc123",
                    },
                    &sdk.WebScrapeRequestBodyBlockURLCookiesItem{
                        Name: "user_pref",
                        Value: "darkmode",
                    },
                },
                Instructions: []*sdk.WebScrapeRequestBodyBlockURLInstructionsItem{
                    &sdk.WebScrapeRequestBodyBlockURLInstructionsItem{
                        WebScrapeRequestBodyBlockURLInstructionsItemFill: &sdk.WebScrapeRequestBodyBlockURLInstructionsItemFill{
                            Fill: &sdk.WebScrapeRequestBodyBlockURLInstructionsItemFillFill{
                                Place: "#username",
                                Value: "myuser",
                            },
                        },
                    },
                    &sdk.WebScrapeRequestBodyBlockURLInstructionsItem{
                        WebScrapeRequestBodyBlockURLInstructionsItemFill: &sdk.WebScrapeRequestBodyBlockURLInstructionsItemFill{
                            Fill: &sdk.WebScrapeRequestBodyBlockURLInstructionsItemFillFill{
                                Place: "#password",
                                Value: "mypassword",
                            },
                        },
                    },
                    &sdk.WebScrapeRequestBodyBlockURLInstructionsItem{
                        WebScrapeRequestBodyBlockURLInstructionsItemClick: &sdk.WebScrapeRequestBodyBlockURLInstructionsItemClick{
                            Click: sdk.String(
                                "#loginButton",
                            ),
                        },
                    },
                    &sdk.WebScrapeRequestBodyBlockURLInstructionsItem{
                        WebScrapeRequestBodyBlockURLInstructionsItemWait: &sdk.WebScrapeRequestBodyBlockURLInstructionsItemWait{
                            Wait: sdk.Int(
                                2000,
                            ),
                        },
                    },
                    &sdk.WebScrapeRequestBodyBlockURLInstructionsItem{
                        WebScrapeRequestBodyBlockURLInstructionsItemExtract: &sdk.WebScrapeRequestBodyBlockURLInstructionsItemExtract{
                            Extract: &sdk.WebScrapeRequestBodyBlockURLInstructionsItemExtractExtract{
                                HTML: sdk.String(
                                    "#profile",
                                ),
                                Text: sdk.String(
                                    "#welcome-message",
                                ),
                                UserData: sdk.String(
                                    "#user-info",
                                ),
                            },
                        },
                    },
                    &sdk.WebScrapeRequestBodyBlockURLInstructionsItem{
                        WebScrapeRequestBodyBlockURLInstructionsItemBlockElement: &sdk.WebScrapeRequestBodyBlockURLInstructionsItemBlockElement{
                            BlockElement: []string{
                                ".ad-banner",
                                "//div[@class='popup']",
                            },
                        },
                    },
                    &sdk.WebScrapeRequestBodyBlockURLInstructionsItem{
                        WebScrapeRequestBodyBlockURLInstructionsItemGeneralImageCaptcha: &sdk.WebScrapeRequestBodyBlockURLInstructionsItemGeneralImageCaptcha{
                            GeneralImageCaptcha: []*sdk.WebScrapeRequestBodyBlockURLInstructionsItemGeneralImageCaptchaGeneralImageCaptchaItem{
                                &sdk.WebScrapeRequestBodyBlockURLInstructionsItemGeneralImageCaptchaGeneralImageCaptchaItem{
                                    ImagePath: sdk.String(
                                        "#captcha-img",
                                    ),
                                    TextField: sdk.String(
                                        "#captcha-input",
                                    ),
                                    ImageUpdatePath: sdk.String(
                                        "#refresh-captcha",
                                    ),
                                    CaptchaFailedPath: sdk.String(
                                        "#captcha-error",
                                    ),
                                    Model: sdk.WebScrapeRequestBodyBlockURLInstructionsItemGeneralImageCaptchaGeneralImageCaptchaItemModelMiniOcrV1.Ptr(),
                                },
                            },
                        },
                    },
                },
            },
        },
    }
client.WebScrape(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.WebScrapeRequestFormat` — Response format returned by the API.
    
</dd>
</dl>

<dl>
<dd>

**url:** `string` — Target URL to scrape
    
</dd>
</dl>

<dl>
<dd>

**text:** `*bool` — Set to `true` to return the data in text format else `false` for data in html format with tags.
    
</dd>
</dl>

<dl>
<dd>

**jsEnabled:** `*bool` 

Set  `true` to handle websites with JavaScript. Set `false` to handle static html websites.


 Default value is `true`.
    
</dd>
</dl>

<dl>
<dd>

**proxy:** `*sdk.WebScrapeRequestProxy` — Use proxy for requests
    
</dd>
</dl>

<dl>
<dd>

**sslIgnore:** `*bool` 

Ignore SSL certificate errors.


 Only works if **jsEnabled** is **true**.
    
</dd>
</dl>

<dl>
<dd>

**windowSize:** `*string` 

Specify the browser window size in the format 'width,height' (e.g., "1920w,1080h"). Default value is the default resolutions provided by web/browser.


 Only works if **jsEnabled** is **true**.
    
</dd>
</dl>

<dl>
<dd>

**adBlock:** `*bool` 

Set to `true` to apply ad-blocker to the specified URL else false or ignore to not apply.


 Only works if **jsEnabled** is **true**.
    
</dd>
</dl>

<dl>
<dd>

**captcha:** `*bool` 

if true user can provide captcha instructions in the instructions to solve image captchas.


  Only works if **jsEnabled** is **true**.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*sdk.WebScrapeRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.EmailValidate(request) -> *sdk.EmailValidateResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Validates a single email address and returns result.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.EmailValidateRequest{
        APIKey: "apiKey",
        Email: "email",
    }
client.EmailValidate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.EmailValidateRequestFormat` — Format of the response
    
</dd>
</dl>

<dl>
<dd>

**email:** `string` — Email address to validate
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Name of the email address
    
</dd>
</dl>

<dl>
<dd>

**ip:** `*string` — IP address of the email address
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkEmailValidate(request) -> *sdk.BulkEmailValidateResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Validates a bulk of email addresses and returns result for each. Maximum `10` email addresses per request.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.BulkEmailValidateRequest{
        APIKey: "apiKey",
        EmailData: []*sdk.BulkEmailValidateRequestEmailDataItem{
            &sdk.BulkEmailValidateRequestEmailDataItem{
                Email: "email",
            },
        },
    }
client.BulkEmailValidate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.BulkEmailValidateRequestFormat` — Format of the response
    
</dd>
</dl>

<dl>
<dd>

**emailData:** `[]*sdk.BulkEmailValidateRequestEmailDataItem` — Array of email objects for bulk validation
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PhoneValidate(request) -> *sdk.PhoneValidateResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Validates a single phone number and returns detailed metadata including carrier, line type, geolocation, time zones, and standardized formats.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PhoneValidateRequest{
        APIKey: "apiKey",
        Number: "+14155552671",
    }
client.PhoneValidate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PhoneValidateRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object. If not provided, the API defaults to JSON format.
    
</dd>
</dl>

<dl>
<dd>

**number:** `string` — Phone number to validate. Accepts international format (+14155552671), local format (4155552671) with region, or IDD format (0014155552671) with dialer_region.
    
</dd>
</dl>

<dl>
<dd>

**region:** `*string` — Two-letter ISO country code (e.g., US, GB). Required when number is in local format without + prefix. Cannot be used together with dialer_region.
    
</dd>
</dl>

<dl>
<dd>

**dialerRegion:** `*string` — Two-letter ISO country code indicating the country the number is being dialed from. Required when number uses IDD exit code. Cannot be used together with region.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkPhoneValidate(request) -> []*sdk.BulkPhoneValidateResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Validates up to 100 phone numbers in a single request. Each number is processed independently — invalid entries return per-number errors without affecting the rest of the batch.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.BulkPhoneValidateRequest{
        APIKey: "apiKey",
        Numbers: []*sdk.BulkPhoneValidateRequestNumbersItem{
            &sdk.BulkPhoneValidateRequestNumbersItem{
                Number: "+14155552671",
            },
            &sdk.BulkPhoneValidateRequestNumbersItem{
                Number: "+447911123456",
            },
            &sdk.BulkPhoneValidateRequestNumbersItem{
                Number: "+919876543210",
            },
        },
    }
client.BulkPhoneValidate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.BulkPhoneValidateRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object. If not provided, the API defaults to JSON format.
    
</dd>
</dl>

<dl>
<dd>

**numbers:** `[]*sdk.BulkPhoneValidateRequestNumbersItem` — Array of phone number objects. Maximum 100 per request.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DomainSslLookup() -> *sdk.DomainSslLookupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve comprehensive SSL certificate information without the certificate chain.
This endpoint provides detailed information about the SSL certificate including expiry dates, issuer details, and encryption methods.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.DomainSslLookupRequest{
        APIKey: "apiKey",
        DomainName: "domainName",
    }
client.DomainSslLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.DomainSslLookupRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**domainName:** `string` — Domain name or URL whose SSL certificate lookup is required
    
</dd>
</dl>

<dl>
<dd>

**sslRaw:** `*bool` — Set to true to get the raw openSSL response of the domain
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DomainSslChainLookup() -> *sdk.DomainSslChainLookupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the complete SSL certificate chain from root Certificate Authority (CA) to end-user certificate.
This endpoint provides comprehensive information about each certificate in the chain.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.DomainSslChainLookupRequest{
        APIKey: "apiKey",
        DomainName: "domainName",
    }
client.DomainSslChainLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.DomainSslChainLookupRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**domainName:** `string` — Domain name or URL whose SSL certificate chain lookup is required
    
</dd>
</dl>

<dl>
<dd>

**sslRaw:** `*bool` — Set to true to get the raw openSSL response for each certificate in the chain
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DomainAvailabilityCheck() -> *sdk.DomainAvailabilityCheckResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The Domain Search API is designed to simplify the process of finding available domain names across all top-level domains (TLDs) and second-level domains (SLDs).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.DomainAvailabilityCheckRequest{
        APIKey: "apiKey",
        Domain: "domain",
    }
client.DomainAvailabilityCheck(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.DomainAvailabilityCheckRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**domain:** `string` — Domain name whose availability is to be checked.
    
</dd>
</dl>

<dl>
<dd>

**source:** `*sdk.DomainAvailabilityCheckRequestSource` — Specify the data source for domain availability checks. Use "dns" for DNS-based lookups or "whois" for WHOIS-based lookups. By default, "dns" is used.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkDomainAvailabilityCheck(request) -> *sdk.BulkDomainAvailabilityCheckResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Perform Bulk Domain Availability checks using a list of domains. Supports upto `100 Domains Per Request`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.BulkDomainAvailabilityCheckRequest{
        APIKey: "apiKey",
        DomainNames: []string{
            "domainNames",
        },
    }
client.BulkDomainAvailabilityCheck(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.BulkDomainAvailabilityCheckRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**source:** `*sdk.BulkDomainAvailabilityCheckRequestSource` — Specify the data source for domain availability checks. Use "dns" for DNS-based lookups or "whois" for WHOIS-based lookups. By default, "dns" is used.
    
</dd>
</dl>

<dl>
<dd>

**domainNames:** `[]string` — List of domain names to check.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DomainAvailabilitySuggestions() -> *sdk.DomainAvailabilitySuggestionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The Domain Search API is designed to simplify the process of finding available domain names across all top-level domains (TLDs) and second-level domains (SLDs).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.DomainAvailabilitySuggestionsRequest{
        APIKey: "apiKey",
        Domain: "domain",
    }
client.DomainAvailabilitySuggestions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.DomainAvailabilitySuggestionsRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**domain:** `string` — Domain name for availability and suggestions.
    
</dd>
</dl>

<dl>
<dd>

**source:** `*sdk.DomainAvailabilitySuggestionsRequestSource` — Specify the data source for domain availability checks. Use "dns" for DNS-based lookups or "whois" for WHOIS-based lookups. By default, "dns" is used.
    
</dd>
</dl>

<dl>
<dd>

**count:** `*int` — Number of suggestions to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SubdomainsLookup() -> *sdk.SubdomainsLookupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The Subdomain Lookup API is designed to retrieve subdomains related to the given domain name. It helps you explore subdomains that are available for registration or usage.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.SubdomainsLookupRequest{
        APIKey: "apiKey",
        Domain: "domain",
    }
client.SubdomainsLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.SubdomainsLookupRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**domain:** `string` — Domain name for availability and suggestions.
    
</dd>
</dl>

<dl>
<dd>

**after:** `*time.Time` — Filter subdomains seen after this date (format YYYY-MM-DD).
    
</dd>
</dl>

<dl>
<dd>

**before:** `*time.Time` — Filter subdomains seen before this date( format YYYY-MM-DD).
    
</dd>
</dl>

<dl>
<dd>

**status:** `*sdk.SubdomainsLookupRequestStatus` — Filter subdomains by status (active or inactive).
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Page number for paginated results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfMerge(request) -> *sdk.PdfMergeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API merges multiple PDF files into a single PDF, in the order they are provided
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfMergeRequest{
        APIKey: "apiKey",
    }
client.PdfMerge(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfMergeRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `*string` — An array of unique file IDs referencing PDF files previously uploaded to the API Freaks server. Use this parameter to merge existing files without re-uploading them. Provide multiple IDs to merge files in the specified order.
    
</dd>
</dl>

<dl>
<dd>

**destroy:** `*bool` — If set to `true`, the input file(s) will be permanently deleted from the server immediately after the output PDF is generated.
    
</dd>
</dl>

<dl>
<dd>

**output:** `*string` — Specifies the desired name for the resulting merged PDF file. If not provided, a default name will be assigned.
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` — The URL to which the webhook notification will be sent after the task is completed.
    
</dd>
</dl>

<dl>
<dd>

**webhookFailureNotification:** `*bool` — If true, a notification will also be sent by email in case the webhook request fails all the retries.  The email notification will be sent to the requesting user or their organization’s admin if part of one.
    
</dd>
</dl>

<dl>
<dd>

**webhookAuthorization:** `*string` — Optional custom header for webhook requests. Format: `Key:Value` (e.g., `Authorization:Bearer token123`). This will be sent as an HTTP header in the webhook call.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfRemovePages(request) -> *sdk.PdfRemovePagesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API removes a selection or range of pages from a PDF file.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfRemovePagesRequest{
        APIKey: "apiKey",
        Pages: "pages",
        File: strings.NewReader(
            "",
        ),
    }
client.PdfRemovePages(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfRemovePagesRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `*string` — The unique identifier of a PDF file already uploaded to the API Freaks server. Use this as an alternative to uploading a new file directly.
    
</dd>
</dl>

<dl>
<dd>

**destroy:** `*bool` — If set to `true`, the input file(s) will be permanently deleted from the server immediately after the output PDF is generated.
    
</dd>
</dl>

<dl>
<dd>

**output:** `*string` — The desired name for the output PDF file after pages have been removed. If not provided, a default name will be assigned.
    
</dd>
</dl>

<dl>
<dd>

**pages:** `string` — Specifies which pages to remove from the PDF. Accepts individual page numbers (e.g., '1,7') and/or ascending page ranges (e.g., '3-5'). Use commas to separate entries and hyphens for ranges. Reverse ranges (e.g., '5-3') are not allowed. Alternatively, you may provide only one of the following keywords: 'even' (removes all even-numbered pages), 'odd' (removes all odd-numbered pages), or 'last' (removes only the last page). The keyword 'all' is not supported for this operation. Examples: '1,3-5', 'even'. Mixing special keywords with specific pages/ranges is not allowed.
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` — The URL to which the webhook notification will be sent after the task is completed.
    
</dd>
</dl>

<dl>
<dd>

**webhookFailureNotification:** `*bool` — If true, a notification will also be sent by email in case the webhook request fails all the retries.  The email notification will be sent to the requesting user or their organization’s admin if part of one.
    
</dd>
</dl>

<dl>
<dd>

**webhookAuthorization:** `*string` — Optional custom header for webhook requests. Format: `Key:Value` (e.g., `Authorization:Bearer token123`). This will be sent as an HTTP header in the webhook call.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfSplit(request) -> *sdk.PdfSplitResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API splits a PDF into multiple parts based on specified page numbers or ranges.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfSplitRequest{
        APIKey: "apiKey",
        File: strings.NewReader(
            "",
        ),
    }
client.PdfSplit(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfSplitRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `*string` — The unique ID of a PDF file already uploaded to the API Freaks server. Use this as an alternative to uploading a new file directly.
    
</dd>
</dl>

<dl>
<dd>

**destroy:** `*bool` — If set to `true`, the input file(s) will be permanently deleted from the server immediately after the output PDF is generated.
    
</dd>
</dl>

<dl>
<dd>

**output:** `*string` — The desired base name for the output PDF files after splitting. If not provided, a default naming convention will be used.
    
</dd>
</dl>

<dl>
<dd>

**pages:** `*string` 

Defines the page numbers or ranges where the PDF should be split. Provide individual pages and/or ranges in any order (for example: "1-4,9-5,16-last"). Separate entries with commas and use hyphens for ranges.

Special keywords (use alone):

• `even` — split at every even-numbered page

• `odd` — split at every odd-numbered page

• `all` — split the PDF into single-page files

The keyword `last` can be used anywhere in the string, in combination with page numbers or ranges (for example: "5-last", "last-2", "1,last,9").

Examples:
- "1,4-2,last"
- "odd"
- "all"
- "last,2-5"

Invalid example: "1,odd" (mixing a keyword other than "last" with specific pages/ranges is not allowed). You can pass multiple pages entries to produce multiple output files.
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` — The URL to which the webhook notification will be sent after the task is completed.
    
</dd>
</dl>

<dl>
<dd>

**webhookFailureNotification:** `*bool` — If true, a notification will also be sent by email in case the webhook request fails all the retries.  The email notification will be sent to the requesting user or their organization’s admin if part of one.
    
</dd>
</dl>

<dl>
<dd>

**webhookAuthorization:** `*string` — Optional custom header for webhook requests. Format: `Key:Value` (e.g., `Authorization:Bearer token123`). This will be sent as an HTTP header in the webhook call.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfRotate(request) -> *sdk.PdfRotateResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API rotates pages of a PDF by a specified angle (in multiples of 90 degrees).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfRotateRequest{
        APIKey: "apiKey",
        Rotate: 1,
        File: strings.NewReader(
            "",
        ),
    }
client.PdfRotate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfRotateRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `*string` — The unique ID of a PDF file already uploaded to the API Freaks server. Use this as an alternative to uploading a new file directly.
    
</dd>
</dl>

<dl>
<dd>

**destroy:** `*bool` — If set to `true`, the input file(s) will be permanently deleted from the server immediately after the output PDF is generated.
    
</dd>
</dl>

<dl>
<dd>

**output:** `*string` — The desired name for the output PDF file after rotation. If not provided, a default name will be assigned.
    
</dd>
</dl>

<dl>
<dd>

**pages:** `*string` — Specifies which pages to rotate. Accepts individual page numbers (e.g., '1,7') and/or ascending page ranges (e.g., '3-5'). Use commas to separate entries and hyphens for ranges. Reverse ranges (e.g., '5-3') are not allowed. Alternatively, provide only one of the following keywords: 'even' (rotate all even-numbered pages), 'odd' (rotate all odd-numbered pages), 'last' (rotate only the last page), or 'all' (rotate all pages). Examples: '1,3-5', 'odd', 'all'. Mixing special keywords with specific pages/ranges is not allowed.
    
</dd>
</dl>

<dl>
<dd>

**rotate:** `int` — The angle in degrees to rotate the selected pages. Must be one of the following values: 0, 90, 180, 270, -90, -180, or -270. All rotations are applied clockwise.
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` — The URL to which the webhook notification will be sent after the task is completed.
    
</dd>
</dl>

<dl>
<dd>

**webhookFailureNotification:** `*bool` — If true, a notification will also be sent by email in case the webhook request fails all the retries.  The email notification will be sent to the requesting user or their organization’s admin if part of one.
    
</dd>
</dl>

<dl>
<dd>

**webhookAuthorization:** `*string` — Optional custom header for webhook requests. Format: `Key:Value` (e.g., `Authorization:Bearer token123`). This will be sent as an HTTP header in the webhook call.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfCompress(request) -> *sdk.PdfCompressResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API compresses a given PDF file to reduce its file size.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfCompressRequest{
        APIKey: "apiKey",
        CompressionLevel: sdk.PdfCompressRequestCompressionLevelLow,
        File: strings.NewReader(
            "",
        ),
    }
client.PdfCompress(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfCompressRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `*string` — The unique ID of a PDF file already uploaded to the API Freaks server. Use this as an alternative to uploading a new file.
    
</dd>
</dl>

<dl>
<dd>

**output:** `*string` — Name of the output PDF.
    
</dd>
</dl>

<dl>
<dd>

**compressionLevel:** `*sdk.PdfCompressRequestCompressionLevel` — Controls how aggressively the PDF is compressed. Lower levels preserve more quality, while higher levels reduce file size more.
    
</dd>
</dl>

<dl>
<dd>

**destroy:** `*bool` — If set to true, the input file(s) will be deleted from the server immediately after the output is generated.
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` — The URL to which the webhook notification will be sent after the task is completed.
    
</dd>
</dl>

<dl>
<dd>

**webhookFailureNotification:** `*bool` — If true, a notification will also be sent by email in case the webhook request fails all the retries.  The email notification will be sent to the requesting user or their organization’s admin if part of one.
    
</dd>
</dl>

<dl>
<dd>

**webhookAuthorization:** `*string` — Optional custom header for webhook requests. Format: `Key:Value` (e.g., `Authorization:Bearer token123`). This will be sent as an HTTP header in the webhook call.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfExtractPages(request) -> *sdk.PdfExtractPagesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API extracts specific pages or page ranges from a PDF file and returns them as a new PDF.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfExtractPagesRequest{
        APIKey: "apiKey",
        Pages: "pages",
        File: strings.NewReader(
            "",
        ),
    }
client.PdfExtractPages(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfExtractPagesRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `*string` — The unique ID of a PDF file already uploaded to the API Freaks server. Use this as an alternative to uploading a new file directly.
    
</dd>
</dl>

<dl>
<dd>

**destroy:** `*bool` — If set to `true`, the input file(s) will be permanently deleted from the server immediately after the output PDF is generated.
    
</dd>
</dl>

<dl>
<dd>

**output:** `*string` — The desired name for the output PDF file after pages have been extracted. If not provided, a default name will be assigned.
    
</dd>
</dl>

<dl>
<dd>

**pages:** `string` — Specifies which pages to extract from the PDF. You can provide individual page numbers (e.g., '2') and/or page ranges in any order, including descending (e.g., '9-5', '16-last'). Use commas to separate entries and hyphens for ranges. You may alternatively pass only one of the special keywords: 'even' (extracts all even-numbered pages), 'odd' (extracts all odd-numbered pages), 'last' (extracts only the last page), or 'all' (extracts all pages into individual files). Examples: '2,6-3', 'even', 'all'. Mixing special keywords with specific pages/ranges is not allowed.
    
</dd>
</dl>

<dl>
<dd>

**separated:** `*bool` — If set to `true`, each of the specified pages will be extracted and returned as a separate PDF file.
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` — The URL to which the webhook notification will be sent after the task is completed.
    
</dd>
</dl>

<dl>
<dd>

**webhookFailureNotification:** `*bool` — If true, a notification will also be sent by email in case the webhook request fails all the retries.  The email notification will be sent to the requesting user or their organization’s admin if part of one.
    
</dd>
</dl>

<dl>
<dd>

**webhookAuthorization:** `*string` — Optional custom header for webhook requests. Format: `Key:Value` (e.g., `Authorization:Bearer token123`). This will be sent as an HTTP header in the webhook call.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfLinearize(request) -> *sdk.PdfLinearizeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

API endpoint that linearizes any given PDF, restructuring it for faster loading and page-by-page viewing in web browsers.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfLinearizeRequest{
        APIKey: "apiKey",
        File: strings.NewReader(
            "",
        ),
    }
client.PdfLinearize(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfLinearizeRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `*string` — The unique ID of a PDF file already uploaded to the API Freaks server. Use this as an alternative to uploading a new file directly.
    
</dd>
</dl>

<dl>
<dd>

**destroy:** `*bool` — If set to `true`, the input file(s) will be permanently deleted from the server immediately after the output PDF is generated.
    
</dd>
</dl>

<dl>
<dd>

**output:** `*string` — The desired name for the output PDF file after pages have been extracted. If not provided, a default name will be assigned.
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` — The URL to which the webhook notification will be sent after the task is completed.
    
</dd>
</dl>

<dl>
<dd>

**webhookFailureNotification:** `*bool` — If true, a notification will also be sent by email in case the webhook request fails all the retries.  The email notification will be sent to the requesting user or their organization’s admin if part of one.
    
</dd>
</dl>

<dl>
<dd>

**webhookAuthorization:** `*string` — Optional custom header for webhook requests. Format: `Key:Value` (e.g., `Authorization:Bearer token123`). This will be sent as an HTTP header in the webhook call.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfEncrypt(request) -> *sdk.PdfEncryptResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API encrypts a PDF file by setting a password required to open it.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfEncryptRequest{
        APIKey: "apiKey",
        UserPassword: "user_password",
        File: strings.NewReader(
            "",
        ),
    }
client.PdfEncrypt(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfEncryptRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `*string` — The unique ID of a PDF file already uploaded to the API Freaks server. Use this as an alternative to uploading a new file directly.
    
</dd>
</dl>

<dl>
<dd>

**destroy:** `*bool` — If set to `true`, the input file(s) will be permanently deleted from the server immediately after the output PDF is generated.
    
</dd>
</dl>

<dl>
<dd>

**output:** `*string` — The desired name for the output encrypted PDF file. If not provided, a default name will be assigned.
    
</dd>
</dl>

<dl>
<dd>

**filePassword:** `*string` — The password to unlock the input file if it is already protected. Either the owner password or user password can be provided. The owner password takes precedence. Password Length should be between 6 and 128 characters.
    
</dd>
</dl>

<dl>
<dd>

**userPassword:** `string` — Sets the user password required to open and view the encrypted PDF file. Password Length should be between 6 and 128 characters.
    
</dd>
</dl>

<dl>
<dd>

**ownerPassword:** `*string` — Sets the owner password for the PDF file. This password provides full access, including the ability to remove restrictions. If not provided, the `user_password` will also be used as the owner password. Password Length should be between 6 and 128 characters.
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` — The URL to which the webhook notification will be sent after the task is completed.
    
</dd>
</dl>

<dl>
<dd>

**webhookFailureNotification:** `*bool` — If true, a notification will also be sent by email in case the webhook request fails all the retries.  The email notification will be sent to the requesting user or their organization’s admin if part of one.
    
</dd>
</dl>

<dl>
<dd>

**webhookAuthorization:** `*string` — Optional custom header for webhook requests. Format: `Key:Value` (e.g., `Authorization:Bearer token123`). This will be sent as an HTTP header in the webhook call.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfDecrypt(request) -> *sdk.PdfDecryptResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API decrypts PDF files, removing all encryption, including open passwords and permission restrictions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfDecryptRequest{
        APIKey: "apiKey",
        FilePassword: "file_password",
        File: strings.NewReader(
            "",
        ),
    }
client.PdfDecrypt(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfDecryptRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `*string` — The unique ID of a PDF file already uploaded to the API Freaks server. Use this as an alternative to uploading a new file directly.
    
</dd>
</dl>

<dl>
<dd>

**destroy:** `*bool` — If set to `true`, the input file(s) will be permanently deleted from the server immediately after the output PDF is generated.
    
</dd>
</dl>

<dl>
<dd>

**output:** `*string` — The desired name for the output decrypted PDF file. If not provided, a default name will be assigned.
    
</dd>
</dl>

<dl>
<dd>

**filePassword:** `string` — The password to unlock the input file if it is protected. Either the owner password or user password can be provided. The owner password takes precedence. Password Length should be between 6 and 128 characters.
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` — The URL to which the webhook notification will be sent after the task is completed.
    
</dd>
</dl>

<dl>
<dd>

**webhookFailureNotification:** `*bool` — If true, a notification will also be sent by email in case the webhook request fails all the retries.  The email notification will be sent to the requesting user or their organization’s admin if part of one.
    
</dd>
</dl>

<dl>
<dd>

**webhookAuthorization:** `*string` — Optional custom header for webhook requests. Format: `Key:Value` (e.g., `Authorization:Bearer token123`). This will be sent as an HTTP header in the webhook call.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfRestrict(request) -> *sdk.PdfRestrictResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API applies permission restrictions on a PDF file, such as disabling printing, copying, or editing. This can include password protection to enforce restrictions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfRestrictRequest{
        APIKey: "apiKey",
        UserPassword: "user_password",
        Restrictions: []*sdk.PdfRestrictRequestRestrictionsItem{
            sdk.PdfRestrictRequestRestrictionsItemPrintHigh.Ptr(),
        },
        File: strings.NewReader(
            "",
        ),
    }
client.PdfRestrict(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfRestrictRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `*string` — The unique ID of a PDF file already uploaded to the API Freaks server. Use this as an alternative to uploading a new file directly.
    
</dd>
</dl>

<dl>
<dd>

**destroy:** `*bool` — If set to `true`, the input file(s) will be permanently deleted from the server immediately after the output PDF is generated.
    
</dd>
</dl>

<dl>
<dd>

**output:** `*string` — The desired name for the output restricted PDF file. If not provided, a default name will be assigned.
    
</dd>
</dl>

<dl>
<dd>

**filePassword:** `*string` — The password to unlock the input file if it is already secured. Provide the owner password if available; otherwise, the user password. The owner password takes precedence. Password Length should be between 6 and 128 characters.
    
</dd>
</dl>

<dl>
<dd>

**userPassword:** `string` — Sets the password users will use to open the PDF. If this is not set, only the owner password will be configured, and anyone can open the PDF file with the provided restrictions enabled. Password Length should be between 6 and 128 characters.
    
</dd>
</dl>

<dl>
<dd>

**ownerPassword:** `*string` — Sets the password that allows full access to the PDF (e.g., removing restrictions). If not provided, the `user_password` (if set) will also be used as the owner password. Password Length should be between 6 and 128 characters.
    
</dd>
</dl>

<dl>
<dd>

**restrictions:** `*sdk.PdfRestrictRequestRestrictionsItem` 

A comma-separated list of restrictions to apply to the PDF. These define what the end-user is *not* allowed to do with the PDF. Available options are:


* **print_high** – Disables high-quality printing.
* **print_low** – Disables low-resolution printing.
* **edit_document_assembly** – Prevents reordering or inserting pages.
* **fill_form_fields** – Disallows filling in PDF form fields.
* **edit_annotations** – Disables adding or modifying annotations or comments.
* **modify_content** – Prevents modifying existing content in the PDF.
* **copy_and_extract_content** – Disables copying text or images from the PDF.
* **use_accessibility** – Prevents screen readers or accessibility tools from accessing content.
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` — The URL to which the webhook notification will be sent after the task is completed.
    
</dd>
</dl>

<dl>
<dd>

**webhookFailureNotification:** `*bool` — If true, a notification will also be sent by email in case the webhook request fails all the retries.  The email notification will be sent to the requesting user or their organization’s admin if part of one.
    
</dd>
</dl>

<dl>
<dd>

**webhookAuthorization:** `*string` — Optional custom header for webhook requests. Format: `Key:Value` (e.g., `Authorization:Bearer token123`). This will be sent as an HTTP header in the webhook call.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfUnrestrict(request) -> *sdk.PdfUnrestrictResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API removes permission restrictions from a PDF while keeping it encrypted. If you want to remove all security (including encryption), use the `/pdf/decrypt` endpoint instead.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfUnrestrictRequest{
        APIKey: "apiKey",
        FilePassword: "file_password",
        File: strings.NewReader(
            "",
        ),
    }
client.PdfUnrestrict(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfUnrestrictRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `*string` — The unique ID of a PDF file already uploaded to the API Freaks server. Use this as an alternative to uploading a new file directly.
    
</dd>
</dl>

<dl>
<dd>

**destroy:** `*bool` — If set to `true`, the input file(s) will be permanently deleted from the server immediately after the output PDF is generated.
    
</dd>
</dl>

<dl>
<dd>

**output:** `*string` — The desired name for the output unrestricted PDF file. If not provided, a default name will be assigned.
    
</dd>
</dl>

<dl>
<dd>

**filePassword:** `string` — The password to unlock the input file. Either the owner password or user password can be provided. The owner password takes precedence. Password Length should be between 6 and 128 characters.
    
</dd>
</dl>

<dl>
<dd>

**userPassword:** `*string` — Sets the user password for the PDF file. Password Length should be between 6 and 128 characters.
    
</dd>
</dl>

<dl>
<dd>

**ownerPassword:** `*string` — Sets the owner password for the PDF file. If the owner password is not provided, the `user_password` will also be used as the owner password. Password Length should be between 6 and 128 characters.
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` — The URL to which the webhook notification will be sent after the task is completed.
    
</dd>
</dl>

<dl>
<dd>

**webhookFailureNotification:** `*bool` — If true, a notification will also be sent by email in case the webhook request fails all the retries.  The email notification will be sent to the requesting user or their organization’s admin if part of one.
    
</dd>
</dl>

<dl>
<dd>

**webhookAuthorization:** `*string` — Optional custom header for webhook requests. Format: `Key:Value` (e.g., `Authorization:Bearer token123`). This will be sent as an HTTP header in the webhook call.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfConvertToPng(request) -> *sdk.PdfConvertToPngResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API converts a given PDF file into a sequence of PNG images.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfConvertToPngRequest{
        APIKey: "apiKey",
        File: strings.NewReader(
            "",
        ),
    }
client.PdfConvertToPng(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfConvertToPngRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `*string` — The unique ID of a PDF file already uploaded to the API Freaks server. Use this as an alternative to uploading a new file directly.
    
</dd>
</dl>

<dl>
<dd>

**destroy:** `*bool` — If set to `true`, the input file(s) will be permanently deleted from the server immediately after the output PDF is generated.
    
</dd>
</dl>

<dl>
<dd>

**output:** `*string` — The desired name for the output unrestricted PDF file. If not provided, a default name will be assigned.
    
</dd>
</dl>

<dl>
<dd>

**pages:** `*string` — Specifies the pages or ranges at which to split the PDF. Accepts individual page numbers (e.g., '1') and/or page ranges (e.g., '4-2', 'last'). Ranges can be ascending or descending. Use commas to separate entries and hyphens for ranges. Alternatively, provide only one of the following keywords: 'even' (split at every even-numbered page), 'odd' (split at every odd-numbered page), 'last' (split at the last page only), or 'all' (split into single pages). Examples: '1,4-2,last', 'odd', 'all'. Mixing special keywords with specific pages/ranges is not allowed.
    
</dd>
</dl>

<dl>
<dd>

**resolution:** `*int` — Specifies the resolution (in DPI) for the output images. Acceptable Range is from 20 to 1200.
    
</dd>
</dl>

<dl>
<dd>

**imageSmoothing:** `*string` — Determines the smoothing options to apply during image conversion. Valid values are 'none', 'all' or a combination of 'text', 'line', and 'image' (comma-separated).If not provided, no smoothing will be applied.
    
</dd>
</dl>

<dl>
<dd>

**profile:** `*sdk.PdfConvertToPngRequestProfile` — Specifies the color profile for the output PNG images. Acceptable values: bw (1-bit black & white, smallest size, no grayscale or color), gray (8-bit grayscale), rgb (24-bit RGB color, default), rgba (32-bit RGB color with alpha channel for transparency), 4-bit (4-bit indexed color, up to 16 colors, smaller size), or 8-bit (8-bit indexed color, up to 256 colors).
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` — The URL to which the webhook notification will be sent after the task is completed.
    
</dd>
</dl>

<dl>
<dd>

**webhookFailureNotification:** `*bool` — If true, a notification will also be sent by email in case the webhook request fails all the retries.  The email notification will be sent to the requesting user or their organization’s admin if part of one.
    
</dd>
</dl>

<dl>
<dd>

**webhookAuthorization:** `*string` — Optional custom header for webhook requests. Format: `Key:Value` (e.g., `Authorization:Bearer token123`). This will be sent as an HTTP header in the webhook call.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfConvertToJpg(request) -> *sdk.PdfConvertToJpgResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API converts a given PDF file into a sequence of JPG images.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfConvertToJpgRequest{
        APIKey: "apiKey",
        File: strings.NewReader(
            "",
        ),
    }
client.PdfConvertToJpg(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfConvertToJpgRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `*string` — The unique ID of a PDF file already uploaded to the API Freaks server. Use this as an alternative to uploading a new file directly.
    
</dd>
</dl>

<dl>
<dd>

**destroy:** `*bool` — If set to `true`, the input file(s) will be permanently deleted from the server immediately after the output PDF is generated.
    
</dd>
</dl>

<dl>
<dd>

**output:** `*string` — The desired name for the output unrestricted PDF file. If not provided, a default name will be assigned.
    
</dd>
</dl>

<dl>
<dd>

**quality:** `*int` — Controls JPG compression quality. Higher values yield sharper images with larger file sizes.
    
</dd>
</dl>

<dl>
<dd>

**pages:** `*string` — Specifies the pages or ranges at which to split the PDF. Accepts individual page numbers (e.g., '1') and/or page ranges (e.g., '4-2', 'last'). Ranges can be ascending or descending. Use commas to separate entries and hyphens for ranges. Alternatively, provide only one of the following keywords: 'even' (split at every even-numbered page), 'odd' (split at every odd-numbered page), 'last' (split at the last page only), or 'all' (split into single pages). Examples: '1,4-2,last', 'odd', 'all'. Mixing special keywords with specific pages/ranges is not allowed.
    
</dd>
</dl>

<dl>
<dd>

**resolution:** `*int` — Specifies the resolution (in DPI) for the output images. Acceptable Range is from 20 to 1200.
    
</dd>
</dl>

<dl>
<dd>

**imageSmoothing:** `*string` — Determines the smoothing options to apply during image conversion. Valid values are 'none', 'all' or a combination of 'text', 'line', and 'image' (comma-separated).If not provided, no smoothing will be applied.
    
</dd>
</dl>

<dl>
<dd>

**profile:** `*sdk.PdfConvertToJpgRequestProfile` — Specifies the color profile for the output PNG images. Acceptable values: bw (1-bit black & white, smallest size, no grayscale or color), gray (8-bit grayscale), rgb (24-bit RGB color, default), rgba (32-bit RGB color with alpha channel for transparency), 4-bit (4-bit indexed color, up to 16 colors, smaller size), or 8-bit (8-bit indexed color, up to 256 colors).
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` — The URL to which the webhook notification will be sent after the task is completed.
    
</dd>
</dl>

<dl>
<dd>

**webhookFailureNotification:** `*bool` — If true, a notification will also be sent by email in case the webhook request fails all the retries.  The email notification will be sent to the requesting user or their organization’s admin if part of one.
    
</dd>
</dl>

<dl>
<dd>

**webhookAuthorization:** `*string` — Optional custom header for webhook requests. Format: `Key:Value` (e.g., `Authorization:Bearer token123`). This will be sent as an HTTP header in the webhook call.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfConvertToTiff(request) -> *sdk.PdfConvertToTiffResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API converts a given PDF file into a sequence of TIFF images. The output images can be saved as a single TIFF file, or as a sequence of TIFF files.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfConvertToTiffRequest{
        APIKey: "apiKey",
        File: strings.NewReader(
            "",
        ),
    }
client.PdfConvertToTiff(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfConvertToTiffRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `*string` — The unique ID of a PDF file already uploaded to the API Freaks server. Use this as an alternative to uploading a new file directly.
    
</dd>
</dl>

<dl>
<dd>

**destroy:** `*bool` — If set to `true`, the input file(s) will be permanently deleted from the server immediately after the output PDF is generated.
    
</dd>
</dl>

<dl>
<dd>

**output:** `*string` — The desired name for the output unrestricted PDF file. If not provided, a default name will be assigned.
    
</dd>
</dl>

<dl>
<dd>

**pages:** `*string` — Specifies the pages or ranges at which to split the PDF. Accepts individual page numbers (e.g., '1') and/or page ranges (e.g., '4-2', 'last'). Ranges can be ascending or descending. Use commas to separate entries and hyphens for ranges. Alternatively, provide only one of the following keywords: 'even' (split at every even-numbered page), 'odd' (split at every odd-numbered page), 'last' (split at the last page only), or 'all' (split into single pages). Examples: '1,4-2,last', 'odd', 'all'. Mixing special keywords with specific pages/ranges is not allowed.
    
</dd>
</dl>

<dl>
<dd>

**resolution:** `*int` — Specifies the resolution (in DPI) for the output images. Acceptable Range is from 20 to 1200.
    
</dd>
</dl>

<dl>
<dd>

**imageSmoothing:** `*string` — Determines the smoothing options to apply during image conversion. Valid values are 'none', 'all' or a combination of 'text', 'line', and 'image' (comma-separated).If not provided, no smoothing will be applied.
    
</dd>
</dl>

<dl>
<dd>

**profile:** `*sdk.PdfConvertToTiffRequestProfile` — Specifies the color profile for the output PNG images. Acceptable values: bw (1-bit black & white, smallest size, no grayscale or color), gray (8-bit grayscale), rgb (24-bit RGB color, default), rgba (32-bit RGB color with alpha channel for transparency), 4-bit (4-bit indexed color, up to 16 colors, smaller size), or 8-bit (8-bit indexed color, up to 256 colors).
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` — The URL to which the webhook notification will be sent after the task is completed.
    
</dd>
</dl>

<dl>
<dd>

**webhookFailureNotification:** `*bool` — If true, a notification will also be sent by email in case the webhook request fails all the retries.  The email notification will be sent to the requesting user or their organization’s admin if part of one.
    
</dd>
</dl>

<dl>
<dd>

**webhookAuthorization:** `*string` — Optional custom header for webhook requests. Format: `Key:Value` (e.g., `Authorization:Bearer token123`). This will be sent as an HTTP header in the webhook call.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfConvertToBmp(request) -> *sdk.PdfConvertToBmpResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Converts a PDF file to a BMP image.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfConvertToBmpRequest{
        APIKey: "apiKey",
        File: strings.NewReader(
            "",
        ),
    }
client.PdfConvertToBmp(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfConvertToBmpRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `*string` — The unique ID of a PDF file already uploaded to the API Freaks server. Use this as an alternative to uploading a new file directly.
    
</dd>
</dl>

<dl>
<dd>

**destroy:** `*bool` — If set to `true`, the input file(s) will be permanently deleted from the server immediately after the output PDF is generated.
    
</dd>
</dl>

<dl>
<dd>

**output:** `*string` — The desired name for the output unrestricted PDF file. If not provided, a default name will be assigned.
    
</dd>
</dl>

<dl>
<dd>

**pages:** `*string` — Specifies the pages or ranges at which to split the PDF. Accepts individual page numbers (e.g., '1') and/or page ranges (e.g., '4-2', 'last'). Ranges can be ascending or descending. Use commas to separate entries and hyphens for ranges. Alternatively, provide only one of the following keywords: 'even' (split at every even-numbered page), 'odd' (split at every odd-numbered page), 'last' (split at the last page only), or 'all' (split into single pages). Examples: '1,4-2,last', 'odd', 'all'. Mixing special keywords with specific pages/ranges is not allowed.
    
</dd>
</dl>

<dl>
<dd>

**resolution:** `*int` — Specifies the resolution (in DPI) for the output images. Acceptable Range is from 20 to 1200.
    
</dd>
</dl>

<dl>
<dd>

**imageSmoothing:** `*string` — Determines the smoothing options to apply during image conversion. Valid values are 'none', 'all' or a combination of 'text', 'line', and 'image' (comma-separated).If not provided, no smoothing will be applied.
    
</dd>
</dl>

<dl>
<dd>

**profile:** `*sdk.PdfConvertToBmpRequestProfile` — Specifies the color profile for the output PNG images. Acceptable values: bw (1-bit black & white, smallest size, no grayscale or color), gray (8-bit grayscale), rgb (24-bit RGB color, default), rgba (32-bit RGB color with alpha channel for transparency), 4-bit (4-bit indexed color, up to 16 colors, smaller size), or 8-bit (8-bit indexed color, up to 256 colors).
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` — The URL to which the webhook notification will be sent after the task is completed.
    
</dd>
</dl>

<dl>
<dd>

**webhookFailureNotification:** `*bool` — If true, a notification will also be sent by email in case the webhook request fails all the retries.  The email notification will be sent to the requesting user or their organization’s admin if part of one.
    
</dd>
</dl>

<dl>
<dd>

**webhookAuthorization:** `*string` — Optional custom header for webhook requests. Format: `Key:Value` (e.g., `Authorization:Bearer token123`). This will be sent as an HTTP header in the webhook call.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfConvertToGif(request) -> *sdk.PdfConvertToGifResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API converts a given PDF file into a sequence of GIF images.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfConvertToGifRequest{
        APIKey: "apiKey",
        File: strings.NewReader(
            "",
        ),
    }
client.PdfConvertToGif(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfConvertToGifRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `*string` — The unique ID of a PDF file already uploaded to the API Freaks server. Use this as an alternative to uploading a new file directly.
    
</dd>
</dl>

<dl>
<dd>

**destroy:** `*bool` — If set to `true`, the input file(s) will be permanently deleted from the server immediately after the output PDF is generated.
    
</dd>
</dl>

<dl>
<dd>

**output:** `*string` — The desired name for the output unrestricted PDF file. If not provided, a default name will be assigned.
    
</dd>
</dl>

<dl>
<dd>

**pages:** `*string` — Specifies the pages or ranges at which to split the PDF. Accepts individual page numbers (e.g., '1') and/or page ranges (e.g., '4-2', 'last'). Ranges can be ascending or descending. Use commas to separate entries and hyphens for ranges. Alternatively, provide only one of the following keywords: 'even' (split at every even-numbered page), 'odd' (split at every odd-numbered page), 'last' (split at the last page only), or 'all' (split into single pages). Examples: '1,4-2,last', 'odd', 'all'. Mixing special keywords with specific pages/ranges is not allowed.
    
</dd>
</dl>

<dl>
<dd>

**resolution:** `*int` — Specifies the resolution (in DPI) for the output images. Acceptable Range is from 20 to 1200.
    
</dd>
</dl>

<dl>
<dd>

**imageSmoothing:** `*string` — Determines the smoothing options to apply during image conversion. Valid values are 'none', 'all' or a combination of 'text', 'line', and 'image' (comma-separated).If not provided, no smoothing will be applied.
    
</dd>
</dl>

<dl>
<dd>

**profile:** `*sdk.PdfConvertToGifRequestProfile` — Specifies the color profile for the output PNG images. Acceptable values: bw (1-bit black & white, smallest size, no grayscale or color), gray (8-bit grayscale), rgb (24-bit RGB color, default), rgba (32-bit RGB color with alpha channel for transparency), 4-bit (4-bit indexed color, up to 16 colors, smaller size), or 8-bit (8-bit indexed color, up to 256 colors).
    
</dd>
</dl>

<dl>
<dd>

**webhookURL:** `*string` — The URL to which the webhook notification will be sent after the task is completed.
    
</dd>
</dl>

<dl>
<dd>

**webhookFailureNotification:** `*bool` — If true, a notification will also be sent by email in case the webhook request fails all the retries.  The email notification will be sent to the requesting user or their organization’s admin if part of one.
    
</dd>
</dl>

<dl>
<dd>

**webhookAuthorization:** `*string` — Optional custom header for webhook requests. Format: `Key:Value` (e.g., `Authorization:Bearer token123`). This will be sent as an HTTP header in the webhook call.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfUploadResources(request) -> *sdk.PdfUploadResourcesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API uploads multiple PDF files to the API Freaks server and generates their unique file IDs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfUploadResourcesRequest{
        APIKey: "apiKey",
    }
client.PdfUploadResources(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfUploadResourcesRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfDownloadResource() -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API downloads PDF files or ZIP archives from the server using their unique resource ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfDownloadResourceRequest{
        APIKey: "apiKey",
        ResourceID: "resource_id",
    }
client.PdfDownloadResource(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfDownloadResourceRequestFormat` 
    
</dd>
</dl>

<dl>
<dd>

**resourceID:** `string` — The unique identifier of the file or ZIP archive to download.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfGetTaskStatus() -> *sdk.PdfGetTaskStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API checks the status of a previously initiated PDF processing task using its unique task ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfGetTaskStatusRequest{
        APIKey: "apiKey",
        TaskID: "task_id",
    }
client.PdfGetTaskStatus(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfGetTaskStatusRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>

<dl>
<dd>

**taskID:** `string` — The unique ID of the PDF processing task for which the status is requested.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfGetFileStatus() -> *sdk.PdfGetFileStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API checks the status of a PDF file using its unique file ID, providing information about its creation and potential deletion time.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfGetFileStatusRequest{
        APIKey: "apiKey",
        FileID: "file_id",
    }
client.PdfGetFileStatus(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfGetFileStatusRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `string` — The unique ID of the file whose status is requested.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfListFiles() -> *sdk.PdfListFilesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API retrieves a list of all PDF files uploaded and generated by a specific user. Please note that if the user is part of an organization, only the Organization Administrator can access this endpoint. Organization Members cannot access this endpoint.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfListFilesRequest{
        APIKey: "apiKey",
    }
client.PdfListFiles(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfListFilesRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PdfDeleteFile() -> *sdk.PdfDeleteFileResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This API deletes a PDF file using its unique file ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.PdfDeleteFileRequest{
        APIKey: "apiKey",
        FileID: "file_id",
    }
client.PdfDeleteFile(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.PdfDeleteFileRequestFormat` — Specifies the desired format for the API response. Choose 'json' for a JSON object or 'xml' for an XML structure.
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `string` — The unique ID of the file to be deleted.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ScreenshotCapture() -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Capture full-page screenshots and videos of websites with advanced options like device simulation, custom code injection, cookie banner blocking, and scrollable content recording.
Supports multiple output formats including JSON, image, GIF, MP4, and WebM.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.ScreenshotCaptureRequest{
        APIKey: "apiKey",
        URL: "url",
    }
client.ScreenshotCapture(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**output:** `*sdk.ScreenshotCaptureRequestOutput` — Output format for screenshot results
    
</dd>
</dl>

<dl>
<dd>

**fileType:** `*sdk.ScreenshotCaptureRequestFileType` — File type for screenshot output
    
</dd>
</dl>

<dl>
<dd>

**url:** `string` — URLs to capture screenshots of
    
</dd>
</dl>

<dl>
<dd>

**width:** `*int` — Browser viewport width in pixels
    
</dd>
</dl>

<dl>
<dd>

**height:** `*int` — Browser viewport height in pixels
    
</dd>
</dl>

<dl>
<dd>

**fullPage:** `*bool` — Capture a full-page screenshot
    
</dd>
</dl>

<dl>
<dd>

**fresh:** `*bool` — Bypass cache and take a fresh screenshot
    
</dd>
</dl>

<dl>
<dd>

**noCookieBanners:** `*bool` — Remove cookie banners from the screenshot
    
</dd>
</dl>

<dl>
<dd>

**enableCaching:** `*bool` — Enable caching for repeated requests
    
</dd>
</dl>

<dl>
<dd>

**blockAds:** `*bool` — Block advertisements on the page
    
</dd>
</dl>

<dl>
<dd>

**blockChatWidgets:** `*bool` — Block chat widget scripts from loading
    
</dd>
</dl>

<dl>
<dd>

**extractText:** `*bool` — Extract visible text from the page
    
</dd>
</dl>

<dl>
<dd>

**extractHTML:** `*bool` — Extract HTML content of the page
    
</dd>
</dl>

<dl>
<dd>

**destroyScreenshot:** `*bool` — Auto-destroy screenshot after fetch
    
</dd>
</dl>

<dl>
<dd>

**lazyLoad:** `*bool` — Enable lazy-loading content before screenshot
    
</dd>
</dl>

<dl>
<dd>

**retina:** `*bool` — Capture screenshot in high-DPI (Retina) mode
    
</dd>
</dl>

<dl>
<dd>

**darkMode:** `*bool` — Render page in dark mode
    
</dd>
</dl>

<dl>
<dd>

**blockTracking:** `*bool` — Block common user-tracking scripts
    
</dd>
</dl>

<dl>
<dd>

**enableIncognito:** `*bool` — Enable private/incognito mode for browser session
    
</dd>
</dl>

<dl>
<dd>

**omitBackground:** `*bool` — Omit background color (transparent background)
    
</dd>
</dl>

<dl>
<dd>

**thumbnailWidth:** `*int` — Thumbnail width in pixels
    
</dd>
</dl>

<dl>
<dd>

**adjustTop:** `*int` — Adjust top in pixels
    
</dd>
</dl>

<dl>
<dd>

**waitForEvent:** `*sdk.ScreenshotCaptureRequestWaitForEvent` — Wait for a specific load event before capturing the screenshot.
    
</dd>
</dl>

<dl>
<dd>

**grayscale:** `*int` — Range:0 to 100 for grayscale filter
    
</dd>
</dl>

<dl>
<dd>

**delay:** `*int` — How many milliseconds to wait before taking the screenshot
    
</dd>
</dl>

<dl>
<dd>

**timeout:** `*int` — Maximum timeout in milliseconds. Defalut is `10,000`
    
</dd>
</dl>

<dl>
<dd>

**ttl:** `*int` — Number of seconds the screenshot should be cached
    
</dd>
</dl>

<dl>
<dd>

**clipX:** `*int` — X position of the clipping rectangle in pixels
    
</dd>
</dl>

<dl>
<dd>

**clipY:** `*int` — Y position of the clipping rectangle in pixels
    
</dd>
</dl>

<dl>
<dd>

**clipWidth:** `*int` — Width of the clipping rectangle in pixels
    
</dd>
</dl>

<dl>
<dd>

**clipHeight:** `*int` — Height of the clipping rectangle in pixels
    
</dd>
</dl>

<dl>
<dd>

**cssUrl:** `*string` — URL to CSS file
    
</dd>
</dl>

<dl>
<dd>

**css:** `*string` — Your custom CSS code
    
</dd>
</dl>

<dl>
<dd>

**jsURL:** `*string` — URL to JS file
    
</dd>
</dl>

<dl>
<dd>

**js:** `*string` — Your JS code
    
</dd>
</dl>

<dl>
<dd>

**blockJs:** `*bool` — Block Scripts
    
</dd>
</dl>

<dl>
<dd>

**blockStylesheets:** `*bool` — Block Stylesheets
    
</dd>
</dl>

<dl>
<dd>

**blockImages:** `*bool` — Block Images
    
</dd>
</dl>

<dl>
<dd>

**blockMedia:** `*bool` — Block Media
    
</dd>
</dl>

<dl>
<dd>

**blockFont:** `*bool` — Block Fonts
    
</dd>
</dl>

<dl>
<dd>

**blockTextTrack:** `*bool` — Block Text Tracks
    
</dd>
</dl>

<dl>
<dd>

**blockXhr:** `*bool` — Block XHR Requests
    
</dd>
</dl>

<dl>
<dd>

**blockFetch:** `*bool` — Block Fetch Requests
    
</dd>
</dl>

<dl>
<dd>

**blockEventSource:** `*bool` — Block Event Source
    
</dd>
</dl>

<dl>
<dd>

**blockWebSocket:** `*bool` — Block Web Sockets
    
</dd>
</dl>

<dl>
<dd>

**blockManifest:** `*bool` — Block Manifest
    
</dd>
</dl>

<dl>
<dd>

**blockSpecificRequests:** `*string` — Comma- or newline-separated list of specific requests to block. Each line and comma are treated as separate requests for processing. Example: https://example.com, https://example.js
    
</dd>
</dl>

<dl>
<dd>

**blurSelector:** `*string` 

Comma-separated list of indexed CSS selectors to blur.
Format: `index:<selector>`, e.g., `0:.banner,1:#ads`.
    
</dd>
</dl>

<dl>
<dd>

**removeSelector:** `*string` 

Comma-separated list of indexed CSS selectors to blur.
Format: `index:<selector>`, e.g., `0:.banner,1:#ads`.
    
</dd>
</dl>

<dl>
<dd>

**resultFileName:** `*string` 

Specify a meaningful & unique file name to easily identify the screenshot result.
Avoid using spaces or special characters; use hyphens or underscores to separate words.
    
</dd>
</dl>

<dl>
<dd>

**scrollingScreenshot:** `*bool` — **`Scrolling Screenshot`**: Capture a long scrolling screenshot. When true, disable `fullPage` and `freshScreenshot`.
    
</dd>
</dl>

<dl>
<dd>

**scrollSpeed:** `*sdk.ScreenshotCaptureRequestScrollSpeed` — Speed of scrolling during the screenshot.
    
</dd>
</dl>

<dl>
<dd>

**scrollBack:** `*bool` — If true, the scroll will reverse back to the top after reaching the bottom.
    
</dd>
</dl>

<dl>
<dd>

**startImmediately:** `*bool` — If true, the scrolling capture will start immediately upon page load.
    
</dd>
</dl>

<dl>
<dd>

**multipleScrolling:** `*bool` — If true, multiple scrolling screenshots will be taken at different viewport sizes.
    
</dd>
</dl>

<dl>
<dd>

**sizes:** `*string` — Comma-separated list of viewport sizes in the format index:XXw:YYh. Example: sizes=0:120w:300h,1:240w:500h
    
</dd>
</dl>

<dl>
<dd>

**duration:** `*float64` — Duration in seconds for the scrolling capture. Acceptable range: 0 to 100 seconds.
    
</dd>
</dl>

<dl>
<dd>

**failOnError:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**longitude:** `*float64` 
    
</dd>
</dl>

<dl>
<dd>

**latitude:** `*float64` 
    
</dd>
</dl>

<dl>
<dd>

**proxy:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**headers:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**cookies:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**scrollToElement:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**selector:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**acceptLanguages:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**customHTML:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**imageQuality:** `*float64` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkScreenshotCapture(request) -> *sdk.BulkScreenshotCaptureResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Our Bulk Screenshot API allows you to capture screenshots of multiple webpages simultaneously, saving you time and effort. Instead of manually capturing each page one by one, you can batch process URLs and receive high-quality screenshots in the format you choose.
 Maximum `50 URLs` per request.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.BulkScreenshotCaptureRequest{
        APIKey: "apiKey",
        URLs: []*sdk.BulkScreenshotCaptureRequestURLsItem{
            &sdk.BulkScreenshotCaptureRequestURLsItem{
                URL: "url",
            },
        },
    }
client.BulkScreenshotCapture(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.BulkScreenshotCaptureRequestFormat` 
    
</dd>
</dl>

<dl>
<dd>

**urls:** `[]*sdk.BulkScreenshotCaptureRequestURLsItem` — List of website URLs to capture screenshots of
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CurrencyLatestRates() -> *sdk.CurrencyLatestRatesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get live forex rates for all world currencies with customizable update frequency
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CurrencyLatestRatesRequest{
        APIKey: "apiKey",
    }
client.CurrencyLatestRates(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.CurrencyLatestRatesRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**base:** `*string` — Base currency for rate calculations
    
</dd>
</dl>

<dl>
<dd>

**symbols:** `*string` — Comma separated list of desired currency codes
    
</dd>
</dl>

<dl>
<dd>

**updates:** `*sdk.CurrencyLatestRatesRequestUpdates` — Exchange rates update period (1d=daily, 1h=hourly, 10m=10 minutes, 1m=1 minute)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CurrencyHistoricalRates() -> *sdk.CurrencyHistoricalRatesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get historical exchange rates for any specific date
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CurrencyHistoricalRatesRequest{
        APIKey: "apiKey",
        Date: sdk.MustParseDate(
            "2023-01-15",
        ),
    }
client.CurrencyHistoricalRates(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.CurrencyHistoricalRatesRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**base:** `*string` — Base currency for rate calculations
    
</dd>
</dl>

<dl>
<dd>

**symbols:** `*string` — Comma separated list of desired currency codes
    
</dd>
</dl>

<dl>
<dd>

**date:** `time.Time` — Specific date in YYYY-MM-DD format
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CurrencyConvertLatest() -> *sdk.CurrencyConvertLatestResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Convert amount between currencies using the latest exchange rates
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CurrencyConvertLatestRequest{
        APIKey: "apiKey",
        From: "from",
        To: "to",
    }
client.CurrencyConvertLatest(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.CurrencyConvertLatestRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**from:** `string` — Source currency code
    
</dd>
</dl>

<dl>
<dd>

**to:** `string` — Target currency code
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*float64` — Amount to convert
    
</dd>
</dl>

<dl>
<dd>

**updates:** `*sdk.CurrencyConvertLatestRequestUpdates` — Exchange rates update period (1d=daily, 1h=hourly, 10m=10 minutes, 1m=1 minute)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CurrencyConvertHistorical() -> *sdk.CurrencyConvertHistoricalResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Convert amount between currencies using historical rates
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CurrencyConvertHistoricalRequest{
        APIKey: "apiKey",
        From: "from",
        To: "to",
        Date: sdk.MustParseDate(
            "2023-01-15",
        ),
    }
client.CurrencyConvertHistorical(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.CurrencyConvertHistoricalRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**from:** `string` — From currency symbol
    
</dd>
</dl>

<dl>
<dd>

**to:** `string` — To currency symbol
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*float64` — The Amount to be converted
    
</dd>
</dl>

<dl>
<dd>

**date:** `time.Time` — specific date (format YYYY-MM-DD) of which exchange rates is used.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CurrencyTimeSeries() -> *sdk.CurrencyTimeSeriesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get exchange rates for a time range
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CurrencyTimeSeriesRequest{
        APIKey: "apiKey",
        StartDate: sdk.MustParseDate(
            "2023-01-15",
        ),
    }
client.CurrencyTimeSeries(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.CurrencyTimeSeriesRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `time.Time` — Start date (format YYYY-MM-DD) of the preferred time frame
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*time.Time` — End date (format YYYY-MM-DD) of the preferred time frame
    
</dd>
</dl>

<dl>
<dd>

**base:** `*string` — Base currency
    
</dd>
</dl>

<dl>
<dd>

**symbols:** `*string` — comma separated list of desired currencies/ commodities symbols
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CurrencyFluctuation() -> *sdk.CurrencyFluctuationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get currency fluctuation data for a time period
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CurrencyFluctuationRequest{
        APIKey: "apiKey",
        StartDate: sdk.MustParseDate(
            "2023-01-15",
        ),
        Base: sdk.String(
            "USD",
        ),
    }
client.CurrencyFluctuation(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.CurrencyFluctuationRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `time.Time` — Start date (format YYYY-MM-DD) of the preferred time frame
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*time.Time` — End date (format YYYY-MM-DD) of the preferred time frame
    
</dd>
</dl>

<dl>
<dd>

**base:** `*string` — Base currency
    
</dd>
</dl>

<dl>
<dd>

**symbols:** `*string` — comma separated list of desired currencies/ commodities symbols
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CurrencyConvertByIP() -> *sdk.CurrencyConvertByIPResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Convert amount using user's location
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CurrencyConvertByIPRequest{
        APIKey: "apiKey",
        From: "from",
    }
client.CurrencyConvertByIP(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.CurrencyConvertByIPRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**updates:** `*sdk.CurrencyConvertByIPRequestUpdates` — Exchange rates update period (1d=daily, 1h=hourly, 10m=10 minutes, 1m=1 minute)
    
</dd>
</dl>

<dl>
<dd>

**from:** `string` — From currency symbol
    
</dd>
</dl>

<dl>
<dd>

**ip:** `*string` — IPv4 or IPv6 geolocated currency
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*float64` — Amount to convert
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CurrencySupported() -> *sdk.CurrencySupportedResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get list of all supported currencies with their metadata
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CurrencySupportedRequest{
        APIKey: "apiKey",
    }
client.CurrencySupported(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.CurrencySupportedRequestFormat` — Format of the response.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CurrencySymbols() -> *sdk.CurrencySymbolsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get currency symbols and codes
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CurrencySymbolsRequest{
        APIKey: "apiKey",
    }
client.CurrencySymbols(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.CurrencySymbolsRequestFormat` — Format of the response.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CurrencyHistoricalLimits() -> *sdk.CurrencyHistoricalLimitsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about historical data availability and limits
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CurrencyHistoricalLimitsRequest{
        APIKey: "apiKey",
    }
client.CurrencyHistoricalLimits(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.CurrencyHistoricalLimitsRequestFormat` — Format of the response.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CommodityLatestRates() -> *sdk.CommodityLatestRatesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get live commodity rates with customizable update frequency
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CommodityLatestRatesRequest{
        APIKey: "apiKey",
        Symbols: []*string{
            sdk.String(
                "symbols",
            ),
        },
        Updates: sdk.CommodityLatestRatesRequestUpdatesTenM,
    }
client.CommodityLatestRates(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.CommodityLatestRatesRequestFormat` — Format of the Response
    
</dd>
</dl>

<dl>
<dd>

**symbols:** `*string` — Comma separated list of desired commodities symbols *(e.g. XAU,XAG,WTI,BRENT)* **Required**
    
</dd>
</dl>

<dl>
<dd>

**updates:** `*sdk.CommodityLatestRatesRequestUpdates` — Exchange rates update period. Possible values are: (1) `10m` - 10 minute update (2) `1m` - 1 minute update **Required**
    
</dd>
</dl>

<dl>
<dd>

**quote:** `*string` — Specifies the target currency for the exchange rate; default quote currency is the market currency of commodity *(e.g. USD, EUR, INR)*
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CommodityHistoricalRates() -> *sdk.CommodityHistoricalRatesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get historical commodity rates for a specific date
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CommodityHistoricalRatesRequest{
        APIKey: "apiKey",
        Date: sdk.MustParseDate(
            "2023-01-15",
        ),
        Symbols: []*string{
            sdk.String(
                "symbols",
            ),
        },
    }
client.CommodityHistoricalRates(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.CommodityHistoricalRatesRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**date:** `time.Time` — Historical date (YYYY-MM-DD)
    
</dd>
</dl>

<dl>
<dd>

**symbols:** `*string` — Comma-separated list of commodity symbols
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CommodityFluctuation() -> *sdk.CommodityFluctuationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get commodity price fluctuation data for a time period
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CommodityFluctuationRequest{
        APIKey: "apiKey",
        Symbols: []*string{
            sdk.String(
                "symbols",
            ),
        },
        StartDate: sdk.MustParseDate(
            "2023-01-15",
        ),
        EndDate: sdk.MustParseDate(
            "2023-01-15",
        ),
    }
client.CommodityFluctuation(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.CommodityFluctuationRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**symbols:** `*string` — Comma-separated list of commodity symbols
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `time.Time` — Start date (YYYY-MM-DD)
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `time.Time` — End date (YYYY-MM-DD)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CommodityTimeSeries() -> *sdk.CommodityTimeSeriesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get commodity rates for a time range
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CommodityTimeSeriesRequest{
        APIKey: "apiKey",
        Symbols: []*string{
            sdk.String(
                "symbols",
            ),
        },
        StartDate: sdk.MustParseDate(
            "2023-01-15",
        ),
        EndDate: sdk.MustParseDate(
            "2023-01-15",
        ),
    }
client.CommodityTimeSeries(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.CommodityTimeSeriesRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**symbols:** `*string` — Comma-separated list of commodity symbols
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `time.Time` — Start date (YYYY-MM-DD)
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `time.Time` — End date (YYYY-MM-DD)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CommoditySymbols() -> *sdk.CommoditySymbolsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get list of supported commodities
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CommoditySymbolsRequest{
        APIKey: "apiKey",
    }
client.CommoditySymbols(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.CommoditySymbolsRequestFormat` — Format of the response.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.VatSupportedCountries() -> *sdk.VatSupportedCountriesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a list of supported countries.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.VatSupportedCountriesRequest{
        APIKey: "apiKey",
    }
client.VatSupportedCountries(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.VatSupportedCountriesRequestFormat` — Format of the response. Default is JSON.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*sdk.VatSupportedCountriesRequestType` — Type of supported country. Supported values: IBAN, SWIFT, VAT. By default, it returns all supported countries for all types.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.VatRateByIP() -> []*sdk.VatRateByIPResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Fetches VAT rate based on the specified or originating IP address.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.VatRateByIPRequest{
        APIKey: "apiKey",
    }
client.VatRateByIP(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.VatRateByIPRequestFormat` — Specify the desired response format. Options: 'json' (default) or 'xml'.
    
</dd>
</dl>

<dl>
<dd>

**ipAddress:** `*string` — IPv4 or IPv6 address to look up VAT rate for. If omitted, the originating IP address will be used.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.VatRateByCountry() -> []*sdk.VatRateByCountryResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Fetches VAT rates for a single country or state provided via query parameters.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.VatRateByCountryRequest{
        APIKey: "apiKey",
        Country: "country",
    }
client.VatRateByCountry(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.VatRateByCountryRequestFormat` — Specify the desired response format. Options: 'json' (default) or 'xml'.
    
</dd>
</dl>

<dl>
<dd>

**country:** `string` — Country identifier in Alpha-2 (PK), Alpha-3 (PAK), or full name (Pakistan). Combine with the optional "state" query for sub-national VAT; values are case-insensitive and may use underscores instead of spaces.
    
</dd>
</dl>

<dl>
<dd>

**state:** `*string` — Optional state or region in Alpha-2 (NY) or full name (New_York). Use with "country" for state-level VAT; values are case-insensitive and may use underscores.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkVatRateByCountry(request) -> *sdk.BulkVatRateByCountryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves VAT details for multiple countries or country-state combinations in a single request. Maximum of `100` entries per request are allowed.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.BulkVatRateByCountryRequest{
        APIKey: "apiKey",
        Countries: []*sdk.BulkVatRateByCountryRequestCountriesItem{
            &sdk.BulkVatRateByCountryRequestCountriesItem{
                Country: "PAK",
            },
            &sdk.BulkVatRateByCountryRequestCountriesItem{
                Country: "United_States",
                State: sdk.String(
                    "New_York",
                ),
            },
        },
    }
client.BulkVatRateByCountry(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.BulkVatRateByCountryRequestFormat` — Specify the desired response format. Options: 'json' (default) or 'xml'.
    
</dd>
</dl>

<dl>
<dd>

**countries:** `[]*sdk.BulkVatRateByCountryRequestCountriesItem` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.VatValidate() -> *sdk.VatValidateResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Validates an EU or UK VAT number and returns registration status details.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.VatValidateRequest{
        APIKey: "apiKey",
        VatNumber: "vatNumber",
    }
client.VatValidate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.VatValidateRequestFormat` — Specify the desired response format. Options: 'json' (default) or 'xml'.
    
</dd>
</dl>

<dl>
<dd>

**vatNumber:** `string` — EU or UK VAT number to validate.
    
</dd>
</dl>

<dl>
<dd>

**requesterVatNumber:** `*string` — Requester EU or UK VAT number.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.IbanValidate() -> *sdk.IbanValidateResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Checks an IBAN for structural validity, checksum accuracy, and bank metadata.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.IbanValidateRequest{
        APIKey: "apiKey",
        Iban: "iban",
    }
client.IbanValidate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.IbanValidateRequestFormat` — Specify the desired response format. Options: 'json' (default) or 'xml'.
    
</dd>
</dl>

<dl>
<dd>

**iban:** `string` — IBAN to validate.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SwiftCodeFind() -> []string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Fetches SWIFT codes for a given country, bank, and city.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.SwiftCodeFindRequest{
        APIKey: "apiKey",
    }
client.SwiftCodeFind(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.SwiftCodeFindRequestFormat` — Specify the desired response format. Options: 'json' (default) or 'xml'.
    
</dd>
</dl>

<dl>
<dd>

**country:** `*string` — Country name (accepts full name, e.g., Pakistan, United States). If only the country parameter is supplied, lists all banks in the country.
    
</dd>
</dl>

<dl>
<dd>

**bank:** `*string` — Bank name (upper case) used to filter SWIFT codes. Should be used together with the country parameter. If only country and bank are provided (without city), returns the list of cities for that bank.
    
</dd>
</dl>

<dl>
<dd>

**city:** `*string` — Gives SWIFT codes for a bank. Optionally specify the city (upper case) to narrow results to a specific city for that bank.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SwiftCodeLookup() -> *sdk.SwiftCodeLookupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Fetches detailed information about a SWIFT code.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.SwiftCodeLookupRequest{
        APIKey: "apiKey",
        SwiftCode: "swiftCode",
    }
client.SwiftCodeLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.SwiftCodeLookupRequestFormat` — Specify the desired response format. Options: 'json' (default) or 'xml'.
    
</dd>
</dl>

<dl>
<dd>

**swiftCode:** `string` — SWIFT/BIC code to lookup (must be 8 or 11 characters).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ZipcodeLookup() -> *sdk.ZipcodeLookupResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.ZipcodeLookupRequest{
        APIKey: "apiKey",
        Code: "code",
    }
client.ZipcodeLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.ZipcodeLookupRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**code:** `string` — Comma separated list of postal / zip codes. Max. 100 values.
    
</dd>
</dl>

<dl>
<dd>

**country:** `*string` — Country code in ISO 3166-1 alpha-2 format. If not provided, search results will be returned from all countries.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkZipcodeLookup(request) -> *sdk.BulkZipcodeLookupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Validates a bulk of ZIP/postal codes and returns result for each. Maximum `100` ZIP/postal codes per request.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.BulkZipcodeLookupRequest{
        APIKey: "apiKey",
        Codes: []string{
            "codes",
        },
    }
client.BulkZipcodeLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.BulkZipcodeLookupRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**codes:** `[]string` — Comma separated list of postal / zip codes. Max. 100 values.
    
</dd>
</dl>

<dl>
<dd>

**country:** `*string` — Country code in ISO 3166-1 alpha-2 format. If not provided, search results will be returned from all countries.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ZipcodeSearchByCity() -> *sdk.ZipcodeSearchByCityResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.ZipcodeSearchByCityRequest{
        APIKey: "apiKey",
        City: "city",
        Country: "country",
    }
client.ZipcodeSearchByCity(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.ZipcodeSearchByCityRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**city:** `string` — Name of the city in which we want to find zipcodes in.
    
</dd>
</dl>

<dl>
<dd>

**country:** `string` — Country code in ISO 3166-1 alpha-2 format.
    
</dd>
</dl>

<dl>
<dd>

**stateName:** `*string` — Name of the state or province associated with the country.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page number to retrieve paginated results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ZipcodeSearchByRegion() -> *sdk.ZipcodeSearchByRegionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.ZipcodeSearchByRegionRequest{
        APIKey: "apiKey",
        Country: "country",
        Region: "region",
    }
client.ZipcodeSearchByRegion(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.ZipcodeSearchByRegionRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**country:** `string` — Country code in ISO 3166-1 alpha-2 format.
    
</dd>
</dl>

<dl>
<dd>

**region:** `string` — Name of the region, state or province associated with the country.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page no. to retrieve paginated results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ZipcodeSearchByRadius() -> *sdk.ZipcodeSearchByRadiusResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.ZipcodeSearchByRadiusRequest{
        APIKey: "apiKey",
        Radius: 1.1,
    }
client.ZipcodeSearchByRadius(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.ZipcodeSearchByRadiusRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**code:** `*string` — Postal/Zip code to be used as the center point for the search.
    
</dd>
</dl>

<dl>
<dd>

**lat:** `*float64` — Latitude coordinate for the base location.
    
</dd>
</dl>

<dl>
<dd>

**long:** `*float64` — Longitude coordinate for the base location.
    
</dd>
</dl>

<dl>
<dd>

**country:** `*string` — Country code in ISO 3166-1 alpha-2 format. Required only when using the code parameter.
    
</dd>
</dl>

<dl>
<dd>

**radius:** `float64` — Search radius for the query. The maximum allowed values are: - 100 km - 100 mi - 109361 yd - 100000 m - 328084 ft - 3937007.75 in
    
</dd>
</dl>

<dl>
<dd>

**unit:** `*sdk.ZipcodeSearchByRadiusRequestUnit` — Supported distance units are m, km, mi, ft, yd, in.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page no. to retrieve paginated results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ZipcodeDistance(request) -> *sdk.ZipcodeDistanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get distance between postal codes. Maximum `100` postal codes per request.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.ZipcodeDistanceRequest{
        APIKey: "apiKey",
        Compare: []string{
            "compare",
        },
        Country: "country",
    }
client.ZipcodeDistance(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.ZipcodeDistanceRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**compare:** `[]string` — Comma separated list of postal / zip codes with which base point is compared w.r.t. Max 100 zip codes can be provided.
    
</dd>
</dl>

<dl>
<dd>

**code:** `*string` — Postal/Zip code to be used as the base point.
    
</dd>
</dl>

<dl>
<dd>

**lat:** `*float64` — Latitude coordinate for the base location.
    
</dd>
</dl>

<dl>
<dd>

**long:** `*float64` — Longitude coordinate for the base location.
    
</dd>
</dl>

<dl>
<dd>

**country:** `string` — Country code in ISO 3166-1 alpha-2 format.
    
</dd>
</dl>

<dl>
<dd>

**unit:** `*sdk.ZipcodeDistanceRequestUnit` — Supported distance units are m, km, mi, ft, yd, in.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ZipcodeDistanceMatch(request) -> *sdk.ZipcodeDistanceMatchResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get matching ZIP/postal code pairs within a specified distance. Maximum `100` postal codes per request.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.ZipcodeDistanceMatchRequest{
        APIKey: "apiKey",
        Codes: []string{
            "codes",
        },
        Country: "country",
    }
client.ZipcodeDistanceMatch(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.ZipcodeDistanceMatchRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**codes:** `[]string` — Comma-separated list of postal/zip codes. Maximum 100 values allowed.
    
</dd>
</dl>

<dl>
<dd>

**country:** `string` — Country code in ISO 3166-1 alpha-2 format.
    
</dd>
</dl>

<dl>
<dd>

**distance:** `*float64` — Maximum allowed distance between postal code pairs.
    
</dd>
</dl>

<dl>
<dd>

**unit:** `*sdk.ZipcodeDistanceMatchRequestUnit` — Supported distance units are m, km, mi, ft, yd, in.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CurrentWeather() -> *sdk.CurrentWeatherResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get current weather data including temperature, humidity, precipitation, wind conditions, atmospheric pressure, and air quality for any location. Accepts city names, coordinates, or IP addresses. Also includes astronomy data and timezone-aware timestamps.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.CurrentWeatherRequest{
        APIKey: "apiKey",
    }
client.CurrentWeather(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.CurrentWeatherRequestFormat` — Response format returned by the API.
    
</dd>
</dl>

<dl>
<dd>

**location:** `*string` — City name, place name, or full address.
    
</dd>
</dl>

<dl>
<dd>

**lat:** `*float64` — Latitude of the location.
    
</dd>
</dl>

<dl>
<dd>

**long:** `*float64` — Longitude of the location.
    
</dd>
</dl>

<dl>
<dd>

**ip:** `*string` — IP(v4 or v6) address for location inference.
    
</dd>
</dl>

<dl>
<dd>

**timezone:** `*string` — Timezone for the results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkCurrentWeather(request) -> *sdk.BulkCurrentWeatherResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve current weather conditions for up to `50 locations` in a single request. A maximum of 50 locations (city names, IP addresses, or geographic coordinates) can be included in the request body.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.BulkCurrentWeatherRequest{
        APIKey: "apiKey",
        Locations: []*sdk.BulkCurrentWeatherRequestLocationsItem{
            &sdk.BulkCurrentWeatherRequestLocationsItem{
                Location: sdk.String(
                    "lahore",
                ),
            },
            &sdk.BulkCurrentWeatherRequestLocationsItem{
                Lat: sdk.Float64(
                    32.5,
                ),
                Long: sdk.Float64(
                    74.5,
                ),
            },
            &sdk.BulkCurrentWeatherRequestLocationsItem{
                IP: sdk.String(
                    "8.8.8.8",
                ),
            },
            &sdk.BulkCurrentWeatherRequestLocationsItem{
                Location: sdk.String(
                    "seoul",
                ),
            },
        },
    }
client.BulkCurrentWeather(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.BulkCurrentWeatherRequestFormat` — Response format returned by the API.
    
</dd>
</dl>

<dl>
<dd>

**timezone:** `*string` — Timezone for the results.
    
</dd>
</dl>

<dl>
<dd>

**locations:** `[]*sdk.BulkCurrentWeatherRequestLocationsItem` — Array of locations to fetch weather data for
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.WeatherForecast() -> *sdk.WeatherForecastResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Access comprehensive weather forecasts with customizable precision - choose from daily overviews, hourly breakdowns, or even minute-by-minute data. Configure your date ranges or use the default 7-day forecast for standard weather planning.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.WeatherForecastRequest{
        APIKey: "apiKey",
    }
client.WeatherForecast(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.WeatherForecastRequestFormat` — Response format returned by the API.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*time.Time` — Start date for the forecast in YYYY-MM-DD format. Forecast dates must be current or future dates only. Past dates are not allowed for forecast data. The difference between startDate and endDate must not exceed 16 days.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*time.Time` — End date for the forecast in YYYY-MM-DD format. Forecast dates must be current or future dates only. Past dates are not allowed for forecast data. The difference between startDate and endDate must not exceed 16 days.
    
</dd>
</dl>

<dl>
<dd>

**forecastDays:** `*int` — Number of days for the forecast, from 1 to 16. Default is 7. Maximum value is 16.
    
</dd>
</dl>

<dl>
<dd>

**location:** `*string` — City name, place name, or full address.
    
</dd>
</dl>

<dl>
<dd>

**lat:** `*float64` — Latitude of the location.
    
</dd>
</dl>

<dl>
<dd>

**long:** `*float64` — Longitude of the location.
    
</dd>
</dl>

<dl>
<dd>

**ip:** `*string` — IP(v4 or v6) address for location inference.
    
</dd>
</dl>

<dl>
<dd>

**precision:** `*sdk.WeatherForecastRequestPrecision` — Precision of the forecast data.
    
</dd>
</dl>

<dl>
<dd>

**timezone:** `*string` — Timezone for the results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.HistoricalWeather() -> *sdk.HistoricalWeatherResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Access past weather conditions for specific dates with records going back to 1940. Retrieve comprehensive historical data with both daily and hourly precision options.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.HistoricalWeatherRequest{
        APIKey: "apiKey",
        Date: sdk.MustParseDate(
            "2023-01-15",
        ),
    }
client.HistoricalWeather(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.HistoricalWeatherRequestFormat` — Response format returned by the API.
    
</dd>
</dl>

<dl>
<dd>

**date:** `time.Time` — Specific date for which to fetch weather data in YYYY-MM-DD format. Historical dates must be past dates only. Current or future dates are not allowed for historical data. Data available from 1940 onwards.
    
</dd>
</dl>

<dl>
<dd>

**location:** `*string` — City name, place name, or full address.
    
</dd>
</dl>

<dl>
<dd>

**lat:** `*float64` — Latitude of the location.
    
</dd>
</dl>

<dl>
<dd>

**long:** `*float64` — Longitude of the location.
    
</dd>
</dl>

<dl>
<dd>

**ip:** `*string` — IP(v4 or v6) address for location inference.
    
</dd>
</dl>

<dl>
<dd>

**precision:** `*sdk.HistoricalWeatherRequestPrecision` — Precision of the historical data. **Note:** 'daily' returns daily aggregates, 'hourly' returns hourly data.
    
</dd>
</dl>

<dl>
<dd>

**timezone:** `*string` — Timezone for the results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.WeatherTimeSeries() -> *sdk.WeatherTimeSeriesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pull historical weather information for date ranges up to 90 days (daily data) or 7 days (hourly data). Get consistent formatting across your specified date range with reliable historical weather patterns.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.WeatherTimeSeriesRequest{
        APIKey: "apiKey",
        StartDate: sdk.MustParseDate(
            "2023-01-15",
        ),
        EndDate: sdk.MustParseDate(
            "2023-01-15",
        ),
    }
client.WeatherTimeSeries(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.WeatherTimeSeriesRequestFormat` — Response format returned by the API.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `time.Time` — Starting date for the data in YYYY-MM-DD format. Historical dates must be past dates only. Current or future dates are not allowed for historical data. Data available from 1940 onwards. For precision=daily, the difference between endDate and startDate must not exceed 90 days. For precision=hourly, the difference must not exceed 7 days.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `time.Time` — End date for the data in YYYY-MM-DD format. Historical dates must be past dates only. Current or future dates are not allowed for historical data. Data available from 1940 onwards. For precision=daily, the difference between endDate and startDate must not exceed 90 days. For precision=hourly, the difference must not exceed 7 days.
    
</dd>
</dl>

<dl>
<dd>

**location:** `*string` — City name, place name, or full address.
    
</dd>
</dl>

<dl>
<dd>

**lat:** `*float64` — Latitude of the location.
    
</dd>
</dl>

<dl>
<dd>

**long:** `*float64` — Longitude of the location.
    
</dd>
</dl>

<dl>
<dd>

**ip:** `*string` — IP(v4 or v6) address for location inference.
    
</dd>
</dl>

<dl>
<dd>

**precision:** `*sdk.WeatherTimeSeriesRequestPrecision` — Precision of the data.
    
</dd>
</dl>

<dl>
<dd>

**timezone:** `*string` — Timezone for the results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarineWeather() -> *sdk.MarineWeatherResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Provides hourly forecasts of marine conditions including wave heights, wave directions, wave periods, swell info, sea surface temperatures, and ocean currents. Supports multiple geographical points and returns daily max wave statistics for up to 7 days. Ideal for maritime planning, navigation, and coastal activities.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.MarineWeatherRequest{
        APIKey: "apiKey",
    }
client.MarineWeather(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.MarineWeatherRequestFormat` — Response format returned by the API.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*time.Time` — Starting date for marine forecast data in YYYY-MM-DD format. Forecast dates must be current or future dates only. Past dates are not allowed for forecast data. The difference between endDate and startDate must not exceed 16 days.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*time.Time` — End date for marine forecast data in YYYY-MM-DD format. Forecast dates must be current or future dates only. Past dates are not allowed for forecast data. The difference between endDate and startDate must not exceed 16 days.
    
</dd>
</dl>

<dl>
<dd>

**location:** `*string` — City name, place name, or full address.
    
</dd>
</dl>

<dl>
<dd>

**lat:** `*float64` — Latitude of the location.
    
</dd>
</dl>

<dl>
<dd>

**long:** `*float64` — Longitude of the location.
    
</dd>
</dl>

<dl>
<dd>

**ip:** `*string` — IP(v4 or v6) address for location inference.
    
</dd>
</dl>

<dl>
<dd>

**precision:** `*sdk.MarineWeatherRequestPrecision` — Precision of the marine data.
    
</dd>
</dl>

<dl>
<dd>

**timezone:** `*string` — Timezone for the results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AirQuality() -> *sdk.AirQualityResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Monitor and predict air quality conditions using European and US AQI standards. Track pollutant concentrations including PM10, PM2.5, carbon monoxide, nitrogen dioxide, sulfur dioxide, ozone, and dust particles. Get current readings plus hourly forecasts up to 5 days ahead, complete with UV index and aerosol measurements for comprehensive air quality assessment.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.AirQualityRequest{
        APIKey: "apiKey",
    }
client.AirQuality(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.AirQualityRequestFormat` — Response format returned by the API.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*time.Time` — Starting date for AQI forecast data in YYYY-MM-DD format. Forecast dates must be current or future dates only. Past dates are not allowed for forecast data. The difference between endDate and startDate must not exceed 5 days.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*time.Time` — End date for AQI forecast data in YYYY-MM-DD format. Forecast dates must be current or future dates only. Past dates are not allowed for forecast data. The difference between endDate and startDate must not exceed 5 days.
    
</dd>
</dl>

<dl>
<dd>

**location:** `*string` — City name, place name, or full address.
    
</dd>
</dl>

<dl>
<dd>

**lat:** `*float64` — Latitude of the location.
    
</dd>
</dl>

<dl>
<dd>

**long:** `*float64` — Longitude of the location.
    
</dd>
</dl>

<dl>
<dd>

**ip:** `*string` — IP(v4 or v6) address for location inference.
    
</dd>
</dl>

<dl>
<dd>

**precision:** `*sdk.AirQualityRequestPrecision` — Only hourly precision is supported; returns hourly AQI data for the selected date range.
    
</dd>
</dl>

<dl>
<dd>

**timezone:** `*string` — Timezone for the results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FloodForecast() -> *sdk.FloodForecastResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Provides flood forecast data for a given location, including river discharge metrics such as mean, median, maximum, minimum, and percentile values (p25, p75). Requires a startDate and endDate, with the date range limited to 16 days. Location can be specified using city name, latitude/longitude, or IP address.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.FloodForecastRequest{
        APIKey: "apiKey",
        StartDate: sdk.MustParseDate(
            "2023-01-15",
        ),
        EndDate: sdk.MustParseDate(
            "2023-01-15",
        ),
    }
client.FloodForecast(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.FloodForecastRequestFormat` — Response format returned by the API.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `time.Time` — Starting date for flood forecast data in YYYY-MM-DD format. Forecast dates must be current or future dates only. Past dates are not allowed for forecast data. The difference between endDate and startDate must not exceed 16 days.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `time.Time` — End date for flood forecast data in YYYY-MM-DD format. Forecast dates must be current or future dates only. Past dates are not allowed for forecast data. The difference between endDate and startDate must not exceed 16 days.
    
</dd>
</dl>

<dl>
<dd>

**location:** `*string` — City name, place name, or full address.
    
</dd>
</dl>

<dl>
<dd>

**lat:** `*float64` — Latitude of the location.
    
</dd>
</dl>

<dl>
<dd>

**long:** `*float64` — Longitude of the location.
    
</dd>
</dl>

<dl>
<dd>

**ip:** `*string` — IP(v4 or v6) address for location inference.
    
</dd>
</dl>

<dl>
<dd>

**precision:** `*sdk.FloodForecastRequestPrecision` — Only daily precision is supported; returns flood forecast data for the selected date range.
    
</dd>
</dl>

<dl>
<dd>

**timezone:** `*string` — Timezone for the results.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GetCountries() -> *sdk.GetCountriesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve countries, optionally filtered by region or subregion.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetCountriesRequest{
        APIKey: "apiKey",
    }
client.GetCountries(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.GetCountriesRequestFormat` — Format of the response
    
</dd>
</dl>

<dl>
<dd>

**region:** `*string` — Optional filter to return countries within a specific region from the region endpoint.
    
</dd>
</dl>

<dl>
<dd>

**subregion:** `*string` — Optional filter to return countries within a specific subregion from the subregion endpoint.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GetCountryDetails() -> *sdk.GetCountryDetailsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetCountryDetailsRequest{
        APIKey: "apiKey",
        Country: "country",
    }
client.GetCountryDetails(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.GetCountryDetailsRequestFormat` — Format of the response
    
</dd>
</dl>

<dl>
<dd>

**country:** `string` — Country code in ISO 3166-1 alpha-2 format.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GetRegions() -> *sdk.GetRegionsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetRegionsRequest{
        APIKey: "apiKey",
    }
client.GetRegions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.GetRegionsRequestFormat` — Format of the response
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GetSubregions() -> *sdk.GetSubregionsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetSubregionsRequest{
        APIKey: "apiKey",
        Region: "region",
    }
client.GetSubregions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.GetSubregionsRequestFormat` — Format of the response
    
</dd>
</dl>

<dl>
<dd>

**region:** `string` — Name of the region.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GetAdminLevels() -> *sdk.GetAdminLevelsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve administrative units based on ISO 3166-1 alpha-2 country code.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetAdminLevelsRequest{
        APIKey: "apiKey",
        Country: "country",
    }
client.GetAdminLevels(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.GetAdminLevelsRequestFormat` — Format of the response
    
</dd>
</dl>

<dl>
<dd>

**country:** `string` — Country code in ISO 3166-1 alpha-2 format
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GetAdminUnits() -> *sdk.GetAdminUnitsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve administrative divisions for a given country using ISO 3166-1 alpha-2 country codes. You can optionally filter by administrative levels.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetAdminUnitsRequest{
        APIKey: "apiKey",
        Country: "country",
    }
client.GetAdminUnits(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.GetAdminUnitsRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**country:** `string` — Country code in ISO 3166-1 alpha-2 format.
    
</dd>
</dl>

<dl>
<dd>

**adminLevels:** `*string` — Comma-separated list to filter results by one or more administrative levels.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GetAdminUnitDetails() -> *sdk.GetAdminUnitDetailsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve detailed administrative unit information by country and optionally filtered by admin code.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetAdminUnitDetailsRequest{
        APIKey: "apiKey",
        Country: "country",
        AdminUnit: "admin_unit",
    }
client.GetAdminUnitDetails(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.GetAdminUnitDetailsRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**country:** `string` — Country code in ISO 3166-1 alpha-2 format.
    
</dd>
</dl>

<dl>
<dd>

**adminUnit:** `string` — Optional admin code to fetch details for a specific administrative unit.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GetCities() -> *sdk.GetCitiesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of cities within a country, optionally filtered by an administrative unit code.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetCitiesRequest{
        APIKey: "apiKey",
        Country: "country",
    }
client.GetCities(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.GetCitiesRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**country:** `string` — Country code in ISO 3166-1 alpha-2 format.
    
</dd>
</dl>

<dl>
<dd>

**adminUnit:** `*string` — Administrative unit code used to filter cities within a specific region.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GetSupportedFlags() -> []*sdk.GetSupportedFlagsResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get list of all supported flags with their metadata
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetSupportedFlagsRequest{
        APIKey: "apiKey",
    }
client.GetSupportedFlags(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GetFlags() -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the flag for a specific country
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GetFlagsRequest{
        APIKey: "apiKey",
        Name: "name",
        Shape: sdk.GetFlagsRequestShapeFlat,
        Type: sdk.GetFlagsRequestTypeCountry,
    }
client.GetFlags(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — Country code in ISO 3166-1 alpha-2 format.
    
</dd>
</dl>

<dl>
<dd>

**shape:** `*sdk.GetFlagsRequestShape` — Flag shape. One of: `'flat'` or `'round'`.
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.GetFlagsRequestFormat` — Flag format. Applicable only for PNG or WEBP formats. Default is png.
    
</dd>
</dl>

<dl>
<dd>

**size:** `*sdk.GetFlagsRequestSize` — Flag size in pixels. Valid options: `16px`, `24px`, `32px`, `48px`, `64px`. Applicable only for PNG or WEBP formats.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*sdk.GetFlagsRequestType` — Type of flag. One of: `country` or `organization`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TimezoneLookup() -> *sdk.TimezoneLookupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve current time, date, and timezone-related information by specifying a timezone name, location address, location coordinates, IP address, or use the client IP address if no parameter is passed.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.TimezoneLookupRequest{
        APIKey: "apiKey",
    }
client.TimezoneLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.TimezoneLookupRequestFormat` — Format of the response
    
</dd>
</dl>

<dl>
<dd>

**ip:** `*string` — IPv4 or IPv6 address to extract timezone information.
    
</dd>
</dl>

<dl>
<dd>

**tz:** `*string` — Timezone name (e.g., "Asia/Kolkata") to retrieve information directly.
    
</dd>
</dl>

<dl>
<dd>

**location:** `*string` — Location string (preferably city and country) to extract timezone.
    
</dd>
</dl>

<dl>
<dd>

**lat:** `*float64` — Latitude for geolocation lookup.
    
</dd>
</dl>

<dl>
<dd>

**long:** `*float64` — Longitude for geolocation lookup.
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*sdk.TimezoneLookupRequestLang` — Language code for response localization (default is "en").
    
</dd>
</dl>

<dl>
<dd>

**iataCode:** `*string` — 3-letter IATA airport code (e.g., JFK).
    
</dd>
</dl>

<dl>
<dd>

**icaoCode:** `*string` — 4-letter ICAO airport code (e.g., KJFK).
    
</dd>
</dl>

<dl>
<dd>

**loCode:** `*string` — 5-letter UN/LO city code.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TimezoneConvert() -> *sdk.TimezoneConvertResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Converts a given time from one timezone to another using various input types like timezone name, coordinates, location, or codes.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.TimezoneConvertRequest{
        APIKey: "apiKey",
    }
client.TimezoneConvert(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.TimezoneConvertRequestFormat` — Format of the response .
    
</dd>
</dl>

<dl>
<dd>

**time:** `*string` — Time to convert in `yyyy-MM-dd HH:mm` or `yyyy-MM-dd HH:mm:ss` format.
    
</dd>
</dl>

<dl>
<dd>

**tzFrom:** `*string` — Source timezone name (e.g., `Asia/Kolkata`).
    
</dd>
</dl>

<dl>
<dd>

**tzTo:** `*string` — Target timezone name (e.g., `America/New_York`).
    
</dd>
</dl>

<dl>
<dd>

**latFrom:** `*float64` — Latitude of source location.
    
</dd>
</dl>

<dl>
<dd>

**longFrom:** `*float64` — Longitude of source location.
    
</dd>
</dl>

<dl>
<dd>

**latTo:** `*float64` — Latitude of target location.
    
</dd>
</dl>

<dl>
<dd>

**longTo:** `*float64` — Longitude of target location.
    
</dd>
</dl>

<dl>
<dd>

**locationFrom:** `*string` — From location (city/country).
    
</dd>
</dl>

<dl>
<dd>

**locationTo:** `*string` — To location (city/country).
    
</dd>
</dl>

<dl>
<dd>

**iataFrom:** `*string` — From IATA airport code (e.g., JFK).
    
</dd>
</dl>

<dl>
<dd>

**iataTo:** `*string` — To IATA airport code.
    
</dd>
</dl>

<dl>
<dd>

**icaoFrom:** `*string` — From ICAO airport code (e.g., KJFK).
    
</dd>
</dl>

<dl>
<dd>

**icaoTo:** `*string` — To ICAO airport code.
    
</dd>
</dl>

<dl>
<dd>

**locodeFrom:** `*string` — From UN/LO CODE.
    
</dd>
</dl>

<dl>
<dd>

**locodeTo:** `*string` — To UN/LO CODE.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserAgentLookup() -> *sdk.UserAgentLookupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Parse User Agent string to get detailed browser, device, and operating system information
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.UserAgentLookupRequest{
        APIKey: "apiKey",
    }
client.UserAgentLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.UserAgentLookupRequestFormat` — Format of the response
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BulkUserAgentLookup(request) -> []*sdk.BulkUserAgentLookupResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Parse up to `50,000 User-Agent strings` at once in a single request.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.BulkUserAgentLookupRequest{
        APIKey: "apiKey",
        UaStrings: []string{
            "uaStrings",
        },
    }
client.BulkUserAgentLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.BulkUserAgentLookupRequestFormat` — Format of the response
    
</dd>
</dl>

<dl>
<dd>

**uaStrings:** `[]string` — List of user agent strings to parse
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.OcrPredict(request) -> *sdk.OcrPredictResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Perform Optical Character Recognition (OCR) on images, PDFs, or ZIP archives. Supports two models: `mini-ocr-v1` for CAPTCHA-optimized OCR and `ocr-v1` for general-purpose document text extraction. Supports zonal OCR to extract text from specific regions of an image.

**Notes:**
- The `zone` query parameter cannot be given with .pdf and .zip types as it can only be applied to single image query.
- The `page_range` query parameter cannot be given in any other type except .pdf types.
- PDFs containing images in them are allowed only for processing.
- The `mini-ocr-v1` model doesn’t support the following query parameters:
    - `page_range` (.pdf types)
    - `zone`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.OcrPredictRequest{
        APIKey: "apiKey",
        Model: sdk.OcrPredictRequestModelMiniOcrV1,
        OcrPredictRequestModel: sdk.OcrPredictRequestModelMiniOcrV1,
    }
client.OcrPredict(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**url:** `*string` — URL of the image or PDF (required if `file` not provided)
    
</dd>
</dl>

<dl>
<dd>

**model:** `*sdk.OcrPredictRequestModel` — OCR model to use.
    
</dd>
</dl>

<dl>
<dd>

**pageRange:** `*string` — Specify page range for multi-page PDFs (e.g., '1,3,5-10' or 'allpages'). **Note:** This parameter can only be used with .pdf file types.
    
</dd>
</dl>

<dl>
<dd>

**zone:** `*string` — Define OCR zones using coordinates (top:left:height:width). Multiple zones can be defined using commas. Only available for model 'ocr-v1'. **Note:** This parameter cannot be used with .pdf and .zip file types as it can only be applied to single image queries.
    
</dd>
</dl>

<dl>
<dd>

**newLine:** `*int` — Set to 1 to split output text into individual lines (default: 0)
    
</dd>
</dl>

<dl>
<dd>

**ocrPredictRequestURL:** `*string` — URL of the image or PDF (required if `file` not provided)
    
</dd>
</dl>

<dl>
<dd>

**ocrPredictRequestModel:** `*sdk.OcrPredictRequestModel` — OCR model to use. `mini-ocr-v1` for CAPTCHA OCR, `ocr-v1` for general OCR
    
</dd>
</dl>

<dl>
<dd>

**ocrPredictRequestPageRange:** `*string` — Specify page range for multi-page PDFs (e.g., '1,3,5-10' or 'allpages'). **Note:** This parameter can only be used with .pdf file types.
    
</dd>
</dl>

<dl>
<dd>

**ocrPredictRequestZone:** `*string` — Define OCR zones using coordinates (top:left:height:width). Multiple zones can be defined using commas. Only available for model 'ocr-v1'. **Note:** This parameter cannot be used with .pdf and .zip file types as it can only be applied to single image queries.
    
</dd>
</dl>

<dl>
<dd>

**ocrPredictRequestNewLine:** `*int` — Set to 1 to split output text into individual lines (default: 0)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GrammarDetect(request) -> *sdk.GrammarDetectResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Analyze text for grammar errors and return the exact words flagged as grammatically incorrect with zero-based word positions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GrammarDetectRequest{
        APIKey: "apiKey",
        Text: "The global mental is health crisis is now a serious and compelex problem. It need quick and ongoing action from policymakers, healthcare workers, and the whole society.",
    }
client.GrammarDetect(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**text:** `string` — Text to analyze for grammar errors
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GrammarCorrect(request) -> *sdk.GrammarCorrectResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Submit text with grammatical issues and receive a clean grammar-corrected result for proofreading and content workflows.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.GrammarCorrectRequest{
        APIKey: "apiKey",
        Text: "The global mental is health crisis is now a serious and compelex problem. It need quick and ongoing action from policymakers, healthcare workers, and the whole society.",
    }
client.GrammarCorrect(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**text:** `string` — Text to correct
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.WeakWordsDetect(request) -> *sdk.WeakWordsDetectResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Analyze text and return weak, vague, or filler words with zero-based word positions to help writers produce clearer and more concise content.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.WeakWordsDetectRequest{
        APIKey: "apiKey",
        Text: "Many people cannot get the support they need to handle their conditions well.",
    }
client.WeakWordsDetect(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**text:** `string` — Text to analyze for weak words
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ReadabilityScore(request) -> *sdk.ReadabilityScoreResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Analyze text readability using industry-standard formulas including Flesch Reading Ease, Flesch-Kincaid Grade Level, Gunning Fog Index, SMOG Index, Coleman-Liau Index, and Automated Readability Index.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.ReadabilityScoreRequest{
        APIKey: "apiKey",
        Text: "The global mental is health crisis is now a serious and compelex problem. It needs quick and ongoing action from policymakers, healthcare workers, and the whole society.",
    }
client.ReadabilityScore(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**target:** `*sdk.ReadabilityScoreRequestTarget` — Target audience used to tune sentence difficulty levels
    
</dd>
</dl>

<dl>
<dd>

**exclude:** `*string` — Comma-separated response sections to omit. Possible values are readability_scores, sentence_readability, readability_grade
    
</dd>
</dl>

<dl>
<dd>

**text:** `string` — Text to analyze for readability
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AstronomyLookup() -> *sdk.AstronomyLookupResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve sunrise and sunset times, current position of the moon, and other related information by specifying a location address, location coordinates, IP address, or using the client IP address if no parameter is passed.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &sdk.AstronomyLookupRequest{
        APIKey: "apiKey",
    }
client.AstronomyLookup(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — Your API key
    
</dd>
</dl>

<dl>
<dd>

**format:** `*sdk.AstronomyLookupRequestFormat` — Format of the response.
    
</dd>
</dl>

<dl>
<dd>

**location:** `*string` — Location name or address
    
</dd>
</dl>

<dl>
<dd>

**lat:** `*float64` — Latitude for location coordinates
    
</dd>
</dl>

<dl>
<dd>

**long:** `*float64` — Longitude for location coordinates
    
</dd>
</dl>

<dl>
<dd>

**ip:** `*string` — IP address for location detection
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**date:** `*time.Time` — Date for astronomy data (YYYY-MM-DD)
    
</dd>
</dl>

<dl>
<dd>

**elevation:** `*float64` — Timezone of the location for which astronomy data is required
    
</dd>
</dl>

<dl>
<dd>

**timeZone:** `*string` — 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

