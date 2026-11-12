# OpenSkin Go SDK

Typed Go client for the OpenSkin Platform API.

## Install

```sh
go get github.com/wlchrist/openskin-go@v0.2.0
```

## Example

```go
package main

import (
	"context"
	"log"
	"os"

	openskin "github.com/wlchrist/openskin-go"
)

func main() {
	apiKey := os.Getenv("OPENSKIN_API_KEY")
	if apiKey == "" {
		log.Fatal("set OPENSKIN_API_KEY before running")
	}

	client := openskin.NewPlatformClient("")
	ctx := openskin.WithAPIKey(context.Background(), apiKey)

	account, _, err := client.PlatformV2API.GetV2Principal(ctx).Execute()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("tier: %s", account.Tier)
}
```

The client uses `https://api-v2.openskin.dev` by default. Create an API key in
the [OpenSkin dashboard](https://platform.openskin.dev/dashboard).

See the [API reference](https://platform.openskin.dev/docs/platform) for the
available operations. Endpoints that require `Idempotency-Key` or `If-Match`
expose those headers as required builder methods.

Licensed under the MIT License.
