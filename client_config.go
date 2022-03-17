package client

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
		baseURL:    "http://34.86.5.170:8086",
	}
}
