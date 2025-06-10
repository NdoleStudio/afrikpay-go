package afrikpay

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

// WithAPIKey sets the AfrikPay API Key
func WithAPIKey(apiKey string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.apiKey = apiKey
	})
}

// WithWalletUsername sets the AfrikPay wallet username
func WithWalletUsername(walletUsername string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.walletUsername = walletUsername
	})
}

// WithWalletPassword sets the AfrikPay wallet password
func WithWalletPassword(walletPassword string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.walletPassword = walletPassword
	})
}

// WithWalletPin sets the AfrikPay wallet pin
func WithWalletPin(walletPin string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.walletPin = walletPin
	})
}
