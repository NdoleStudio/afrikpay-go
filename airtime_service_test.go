package afrikpay

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/NdoleStudio/afrikpay-go/internal/helpers"
	"github.com/NdoleStudio/afrikpay-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestAirtimeService_TransferRequest(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	apiKey := "api-key"
	agentID := "agent-id"
	agentPlatform := "agent-platform"
	agentPassword := "agent-password"
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, [][]byte{stubs.TransferWithError()}, &requests)
	client := New(
		WithBaseURL(server.URL),
		WithAgentID(agentID),
		WithAPIKey(apiKey),
		WithAgentPassword(agentPassword),
		WithAgentPlatform(agentPlatform),
	)
	params := AirtimeTransferParams{
		Operator:          "mtn",
		PurchaseReference: "test-ref",
		Amount:            "987",
		PhoneNumber:       "00000000",
		Mode:              AirtimeModeAccount,
	}
	expectedRequest := map[string]interface{}{
		"operator":      params.Operator,
		"reference":     params.PhoneNumber,
		"amount":        params.Amount,
		"mode":          params.Mode.String(),
		"purchaseref":   params.PurchaseReference,
		"agentid":       agentID,
		"agentplatform": agentPlatform,
		"agentpwd":      agentPassword,
		"hash":          fmt.Sprintf("%x", md5.Sum([]byte(params.Operator+params.PhoneNumber+params.Amount+apiKey))),
	}

	// Act
	_, _, err := client.Airtime.Transfer(context.Background(), params)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(requests))
	buf, err := ioutil.ReadAll(requests[0].Body)
	assert.NoError(t, err)

	requestBody := map[string]interface{}{}
	err = json.Unmarshal(buf, &requestBody)
	assert.NoError(t, err)

	// Assert
	assert.Equal(t, expectedRequest, requestBody)

	// Teardown
	server.Close()
}

func TestAirtimeService_TransferWithoutError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.Transfer())
	client := New(WithBaseURL(server.URL))
	params := AirtimeTransferParams{}

	// Act
	transaction, response, err := client.Airtime.Transfer(context.Background(), params)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, &AirtimeResponse{
		Code:    200,
		Message: "Success",
		Result: &AirtimeTransaction{
			OperatorID:       "1647539307",
			TransactionID:    "1069",
			Status:           "SUCCESS",
			Date:             "2022-03-17 18:48:26",
			Ticket:           nil,
			ReferenceID:      "18360",
			ProcessingNumber: "aaba045a-d571-41e9-9ea4-54cd78782e03",
		},
	}, transaction)

	assert.True(t, transaction.IsSuccessfull())

	// Teardown
	server.Close()
}

func TestAirtimeService_TransferWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.TransferWithError())
	client := New(WithBaseURL(server.URL))
	params := AirtimeTransferParams{}

	// Act
	transaction, _, err := client.Airtime.Transfer(context.Background(), params)

	// Assert
	assert.NoError(t, err)
	assert.Nil(t, transaction.Result)

	assert.Equal(t, "412: bad password", transaction.Message)
	assert.Equal(t, http.StatusInternalServerError, transaction.Code)
	assert.Nil(t, transaction.Result)
	assert.False(t, transaction.IsSuccessfull())

	// Teardown
	server.Close()
}

func TestAirtimeService_StatusRequest(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	apiKey := "api-key"
	agentID := "agent-id"
	agentPlatform := "agent-platform"
	transactionID := "transaction-id"
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, [][]byte{stubs.TransferWithError()}, &requests)
	client := New(
		WithBaseURL(server.URL),
		WithAgentID(agentID),
		WithAPIKey(apiKey),
		WithAgentPlatform(agentPlatform),
	)
	expectedRequest := map[string]interface{}{
		"processingnumber": transactionID,
		"agentid":          agentID,
		"agentplatform":    agentPlatform,
		"hash":             fmt.Sprintf("%x", md5.Sum([]byte(transactionID+apiKey))),
	}

	// Act
	_, _, err := client.Airtime.Status(context.Background(), transactionID)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(requests))
	buf, err := ioutil.ReadAll(requests[0].Body)
	assert.NoError(t, err)

	requestBody := map[string]interface{}{}
	err = json.Unmarshal(buf, &requestBody)
	assert.NoError(t, err)

	// Assert
	assert.Equal(t, expectedRequest, requestBody)

	// Teardown
	server.Close()
}

func TestAirtimeService_StatusWithoutError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.Transfer())
	client := New(WithBaseURL(server.URL))
	transactionID := "transaction-id"

	// Act
	transaction, response, err := client.Airtime.Status(context.Background(), transactionID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, &AirtimeResponse{
		Code:    200,
		Message: "Success",
		Result: &AirtimeTransaction{
			OperatorID:       "1647539307",
			TransactionID:    "1069",
			Status:           "SUCCESS",
			Date:             "2022-03-17 18:48:26",
			Ticket:           nil,
			ReferenceID:      "18360",
			ProcessingNumber: "aaba045a-d571-41e9-9ea4-54cd78782e03",
		},
	}, transaction)

	assert.True(t, transaction.IsSuccessfull())

	// Teardown
	server.Close()
}

func TestAirtimeService_StatusWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.TransferWithError())
	client := New(WithBaseURL(server.URL))
	transactionID := "transaction-id"

	// Act
	transaction, _, err := client.Airtime.Status(context.Background(), transactionID)

	// Assert
	assert.NoError(t, err)
	assert.Nil(t, transaction.Result)

	assert.Equal(t, "412: bad password", transaction.Message)
	assert.Equal(t, http.StatusInternalServerError, transaction.Code)
	assert.Nil(t, transaction.Result)
	assert.False(t, transaction.IsSuccessfull())

	// Teardown
	server.Close()
}
