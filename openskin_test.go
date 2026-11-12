package openskin

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
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

func TestProxyDiagnosticsRedactCredentialsAndWireJSONRetainsThem(t *testing.T) {
	username := "proxy-user-sensitive"
	password := "proxy-password-sensitive"
	token := "partner-token-sensitive"
	proxy := V2InventoryRequestProxy{
		Scheme: "https", Host: "proxy-host-sensitive", Port: 8443,
		Username: &username, Password: &password,
	}
	request := V2SteamPartnerInventoryRequest{
		PartnerSteamId64: "76561198000000000", TradeOfferToken: &token, Proxy: &proxy,
	}

	for name, value := range map[string]interface{}{
		"proxy":   &proxy,
		"request": &request,
	} {
		for _, format := range []string{"%v", "%+v", "%#v"} {
			diagnostic := fmt.Sprintf(format, value)
			for _, secret := range []string{username, password, token, "proxy-host-sensitive", "8443", "https"} {
				if strings.Contains(diagnostic, secret) {
					t.Errorf("%s with %s exposed %q: %s", name, format, secret, diagnostic)
				}
			}
		}
	}

	for name, value := range map[string]interface{}{"proxy": &proxy, "request": &request} {
		wire, err := json.Marshal(value)
		if err != nil {
			t.Fatalf("marshal %s: %v", name, err)
		}
		expectedValues := []string{username, password, "proxy-host-sensitive", "8443", "https"}
		if name == "request" {
			expectedValues = append(expectedValues, token)
		}
		for _, expected := range expectedValues {
			if !strings.Contains(string(wire), expected) {
				t.Errorf("%s JSON omitted %q: %s", name, expected, wire)
			}
		}
	}
}
