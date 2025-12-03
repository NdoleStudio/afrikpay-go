package afrikpay

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/afrikpay-go/internal/helpers"
	"github.com/NdoleStudio/afrikpay-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestAirtime_Pay(t *testing.T) {
	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusOK, stubs.AirtimePaymentResponse())
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))
	input := &AirtimePaymentRequest{
		ReferenceNumber: "659683157",
		Amount:          1000,
		Email:           "mail@domain.com",
		ExternalID:      "224169cd-caa6-46d3-8262-eb95adb6b1d9",
		Description:     "Orange Airtime Purchase",
	}

	// Act
	txn, response, err := client.OrangeAirtime.Pay(context.Background(), input)

	// Assert
	assert.Nil(t, err)
	assert.False(t, txn.IsFailed())
	assert.Equal(t, 1000, txn.Result.Amount)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
}

func TestAirtime_TransactionStatus(t *testing.T) {
	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusOK, stubs.AirtimePaymentResponse())
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))
	input := &TransactionStatusRequest{
		ReferenceNumber: "659683157",
		Amount:          1000,
		ExternalID:      "224169cd-caa6-46d3-8262-eb95adb6b1d9",
	}

	// Act
	txn, response, err := client.OrangeAirtime.TransactionStatus(context.Background(), input)

	// Assert
	assert.Nil(t, err)
	assert.False(t, txn.IsFailed())
	assert.Equal(t, 1000, txn.Result.Amount)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
}
