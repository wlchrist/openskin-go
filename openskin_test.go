package openskin

import (
	"context"
	"testing"
)

func TestNewPlatformClientSelectsHost(t *testing.T) {
	client := NewPlatformClient(V2TestBaseURL)
	server, err := client.cfg.Servers.URL(0, nil)
	if err != nil {
		t.Fatal(err)
	}
	if server != V2TestBaseURL {
		t.Fatalf("server = %q, want %q", server, V2TestBaseURL)
	}
}

func TestWithAPIKey(t *testing.T) {
	ctx := WithAPIKey(context.Background(), "test-key")
	keys, ok := ctx.Value(ContextAPIKeys).(map[string]APIKey)
	if !ok || keys["PlatformAPIKey"].Key != "test-key" {
		t.Fatal("PlatformAPIKey was not attached to the request context")
	}
}
