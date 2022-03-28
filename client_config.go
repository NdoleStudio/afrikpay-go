package afrikpay

import "net/http"

type clientConfig struct {
	httpClient    *http.Client
	baseURL       string
	apiKey        string
	agentID       string
	agentPlatform string
	agentPassword string
}

func defaultClientConfig() *clientConfig {
	return &clientConfig{
		httpClient: http.DefaultClient,
		baseURL:    "https://api.afrikpay.com",
	}
}
