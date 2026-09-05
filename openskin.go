package openskin

import "context"

const (
	ProductionBaseURL = "https://api-v2.openskin.dev"
	V2TestBaseURL     = "https://api-v2-v2-test.up.railway.app"
)

// NewPlatformClient creates a client for the selected OpenSkin API host.
// Pass an empty baseURL to use the production API.
func NewPlatformClient(baseURL string) *APIClient {
	configuration := NewConfiguration()
	if baseURL == "" {
		baseURL = ProductionBaseURL
	}
	configuration.Servers = ServerConfigurations{{URL: baseURL}}
	return NewAPIClient(configuration)
}

// WithAPIKey attaches a Platform API key to requests made with ctx.
func WithAPIKey(ctx context.Context, apiKey string) context.Context {
	return context.WithValue(ctx, ContextAPIKeys, map[string]APIKey{
		"PlatformAPIKey": {Key: apiKey},
	})
}
