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

func TestAccountService_BalanceRequest(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	apiKey := "api-key"
	agentID := "agent-id"
	agentPlatform := "agent-platform"
	agentPassword := "agent-password"
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, [][]byte{stubs.AccountBalanceWithError()}, &requests)
	client := New(
		WithBaseURL(server.URL),
		WithAgentID(agentID),
		WithAPIKey(apiKey),
		WithAgentPassword(agentPassword),
		WithAgentPlatform(agentPlatform),
	)

	expectedRequest := map[string]interface{}{
		"agentid":       agentID,
		"agentplatform": agentPlatform,
		"hash":          fmt.Sprintf("%x", md5.Sum([]byte(agentID+apiKey))),
	}

	// Act
	_, _, err := client.Account.Balance(context.Background())
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

func TestAccountService_BalanceWithoutError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.AccountBalance())
	client := New(WithBaseURL(server.URL))

	// Act
	transaction, response, err := client.Account.Balance(context.Background())

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, &AccountBalanceResponse{
		Code:    200,
		Message: "Success",
		Result: &AccountBalance{
			Name:        "COMPANY_NAME",
			MainBalance: "100",
			MainDeposit: "200",
		},
	}, transaction)

	assert.True(t, transaction.IsSuccessfull())

	// Teardown
	server.Close()
}

func TestAccountService_BalanceRequestWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.TransferWithError())
	client := New(WithBaseURL(server.URL))

	// Act
	transaction, _, err := client.Account.Balance(context.Background())

	// Assert
	assert.NoError(t, err)
	assert.Nil(t, transaction.Result)

	assert.Equal(t, "412: bad password", transaction.Message)
	assert.Equal(t, http.StatusInternalServerError, transaction.Code)
	assert.Nil(t, transaction.Result)

	// Teardown
	server.Close()
}
