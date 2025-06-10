package afrikpay

import "net/http"

type clientConfig struct {
	httpClient     *http.Client
	baseURL        string
	apiKey         string
	walletUsername string
	walletPassword string
	walletPin      string
}

func defaultClientConfig() *clientConfig {
	return &clientConfig{
		httpClient: http.DefaultClient,
		baseURL:    "https://api.developers.afrikpay.com",
	}
}
