package afrikpay

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/afrikpay-go/internal/helpers"
	"github.com/NdoleStudio/afrikpay-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestOptionService_GetOptions(t *testing.T) {
	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusOK, stubs.CanalPlusOptionResponse())
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	// Act
	options, response, err := client.CanalPlus.GetOptions(context.Background(), &CanalPlusOptionRequest{
		ReferenceNumber: "11111111111111",
	})

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, 6, len(options.Result))
}
