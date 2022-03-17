package client

import (
	"net/http"
	"strings"
)

// Option is options for constructing a client
type Option interface {
	apply(config *clientConfig)
}

type clientOptionFunc func(config *clientConfig)

func (fn clientOptionFunc) apply(config *clientConfig) {
	fn(config)
}

// WithHTTPClient sets the underlying HTTP client used for API requests.
// By default, http.DefaultClient is used.
func WithHTTPClient(httpClient *http.Client) Option {
	return clientOptionFunc(func(config *clientConfig) {
		if httpClient != nil {
			config.httpClient = httpClient
		}
	})
}

// WithBaseURL set's the base url for the afrikpay API
func WithBaseURL(baseURL string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		if baseURL != "" {
			config.baseURL = strings.TrimRight(baseURL, "/")
		}
	})
}

// WithAPIKey sets the API Key
func WithAPIKey(apiKey string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.apiKey = apiKey
	})
}

// WithAgentID sets the Agent ID for api calls
func WithAgentID(agentID string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.agentID = agentID
	})
}

// WithAgentPlatform sets the agent platform
func WithAgentPlatform(agentPlatform string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.agentPlatform = agentPlatform
	})
}

// WithAgentPassword sets the agent password
func WithAgentPassword(agentPassword string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.agentPassword = agentPassword
	})
}
