# OpenSkin Go SDK

Typed Go client for the OpenSkin Platform API.

## Install

```sh
go get github.com/wlchrist/openskin-go@v0.1.0
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
	client := openskin.NewPlatformClient("")
	ctx := openskin.WithAPIKey(context.Background(), os.Getenv("OPENSKIN_API_KEY"))

	health, _, err := client.PlatformV2API.GetV2Health(ctx).Execute()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("status: %s", health.Status)
}
```

The client uses `https://api-v2.openskin.dev` by default. Create an API key in
the [OpenSkin dashboard](https://platform.openskin.dev/dashboard).

API operations use typed request builders. Required headers such as
`Idempotency-Key` and `If-Match` are required builder methods.

Licensed under the MIT License.
