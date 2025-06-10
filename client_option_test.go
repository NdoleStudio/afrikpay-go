package afrikpay

import (
	"net/http"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestWithHTTPClient(t *testing.T) {
	t.Run("httpClient is not set when the httpClient is nil", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()

		// Act
		WithHTTPClient(nil).apply(config)

		// Assert
		assert.NotNil(t, config.httpClient)
	})

	t.Run("httpClient is set when the httpClient is not nil", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()
		newClient := &http.Client{Timeout: 300}

		// Act
		WithHTTPClient(newClient).apply(config)

		// Assert
		assert.NotNil(t, config.httpClient)
		assert.Equal(t, newClient.Timeout, config.httpClient.Timeout)
	})
}

func TestWithBaseURL(t *testing.T) {
	t.Run("baseURL is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		baseURL := "https://example.com"
		config := defaultClientConfig()

		// Act
		WithBaseURL(baseURL).apply(config)

		// Assert
		assert.Equal(t, config.baseURL, config.baseURL)
	})

	t.Run("tailing / is trimmed from baseURL", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		baseURL := "https://example.com/"
		config := defaultClientConfig()

		// Act
		WithBaseURL(baseURL).apply(config)

		// Assert
		assert.Equal(t, "https://example.com", config.baseURL)
	})
}

func TestWithAPIKey(t *testing.T) {
	t.Run("apiKey is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()
		apiKey := "test-key"

		// Act
		WithAPIKey(apiKey).apply(config)

		// Assert
		assert.Equal(t, apiKey, config.apiKey)
	})
}

func TestWithWalletUsername(t *testing.T) {
	t.Run("wallet username is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()
		walletUsername := faker.Username()

		// Act
		WithWalletUsername(walletUsername).apply(config)

		// Assert
		assert.Equal(t, walletUsername, config.walletUsername)
	})
}

func TestWithWalletPassword(t *testing.T) {
	t.Run("wallet password is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()
		walletPassword := faker.Password()

		// Act
		WithWalletPassword(walletPassword).apply(config)

		// Assert
		assert.Equal(t, walletPassword, config.walletPassword)
	})
}

func TestWithWalletPin(t *testing.T) {
	t.Run("wallet pin is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()
		walletPin := faker.YearString()

		// Act
		WithWalletPin(walletPin).apply(config)

		// Assert
		assert.Equal(t, walletPin, config.walletPin)
	})
}
