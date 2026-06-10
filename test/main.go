package main

import (
	"context"
	"fmt"

	sdk "github.com/api-freaks/sdk"
	"github.com/api-freaks/sdk/client"
)

func main() {
	c := client.NewClient()

	req := &sdk.GeolocationLookupRequest{
		APIKey: "YOUR_API_KEY",
		IP:     sdk.String("8.8.8.8"),
	}

	resp, err := c.GeolocationLookup(
		context.Background(),
		req,
	)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
