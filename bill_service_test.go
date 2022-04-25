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

func TestBillService_Pay_Request(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	apiKey := "api-key"
	agentID := "agent-id"
	agentPlatform := "agent-platform"
	agentPassword := "agent-password"
	processingNumber := "processing-number"
	billerID := "999999999"
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, [][]byte{stubs.BillPay()}, &requests)
	client := New(
		WithBaseURL(server.URL),
		WithAgentID(agentID),
		WithAPIKey(apiKey),
		WithAgentPassword(agentPassword),
		WithAgentPlatform(agentPlatform),
	)
	params := BillPayParams{
		Biller:           BillerEneoPostpay,
		BillID:           billerID,
		Mode:             ModeCash,
		ProcessingNumber: StringToPointer(processingNumber),
	}
	expectedRequest := map[string]interface{}{
		"biller":           params.Biller.string(),
		"billid":           params.BillID,
		"mode":             params.Mode.String(),
		"processingnumber": PointerToString(params.ProcessingNumber),
		"agentid":          agentID,
		"agentplatform":    agentPlatform,
		"agentpwd":         agentPassword,
		"hash":             fmt.Sprintf("%x", md5.Sum([]byte(params.Biller.string()+params.BillID+apiKey))),
	}

	// Act
	_, _, err := client.Bill.Pay(context.Background(), params)
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

func TestBillService_Pay(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.BillPay())
	client := New(WithBaseURL(server.URL))
	params := BillPayParams{}

	// Act
	transaction, response, err := client.Bill.Pay(context.Background(), params)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, &BillPayResponse{
		Code:    200,
		Message: "success",
		Result: &BillTransaction{
			OperatorID:       StringToPointer("xxxx-xxxx-xxxx-xxxx-5286 : 0000000000068 : 8.8 Kwh"),
			TransactionID:    "5xxxx",
			Status:           "PENDING",
			Date:             "2022-04-19 18:00:06",
			ReferenceID:      nil,
			ProcessingNumber: "aaba045a-d571-41e9-9ea4-54cd78782e03",
		},
	}, transaction)

	assert.True(t, transaction.IsSuccessful())

	// Teardown
	server.Close()
}

func TestBillService_Pay_WithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.BillPayWithError())
	client := New(WithBaseURL(server.URL))
	params := AirtimeTransferParams{}

	// Act
	transaction, _, err := client.Airtime.Transfer(context.Background(), params)

	// Assert
	assert.NoError(t, err)
	assert.Nil(t, transaction.Result)

	assert.Equal(t, "General Failure", transaction.Message)
	assert.Equal(t, http.StatusInternalServerError, transaction.Code)
	assert.Nil(t, transaction.Result)
	assert.False(t, transaction.IsSuccessful())

	// Teardown
	server.Close()
}

func TestBillService_Status_Request(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	apiKey := "api-key"
	agentID := "agent-id"
	agentPlatform := "agent-platform"
	transactionID := "aaba045a-d571-41e9-9ea4-54cd78782e03"
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, [][]byte{stubs.BillPay()}, &requests)
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
	_, _, err := client.Bill.Status(context.Background(), transactionID)
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

func TestBillService_Status(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.BillPay())
	client := New(WithBaseURL(server.URL))
	transactionID := "aaba045a-d571-41e9-9ea4-54cd78782e03"

	// Act
	transaction, response, err := client.Bill.Status(context.Background(), transactionID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, &BillPayResponse{
		Code:    200,
		Message: "success",
		Result: &BillTransaction{
			OperatorID:       StringToPointer("xxxx-xxxx-xxxx-xxxx-5286 : 0000000000068 : 8.8 Kwh"),
			TransactionID:    "5xxxx",
			Status:           "PENDING",
			Date:             "2022-04-19 18:00:06",
			ReferenceID:      nil,
			ProcessingNumber: "aaba045a-d571-41e9-9ea4-54cd78782e03",
		},
	}, transaction)

	assert.True(t, transaction.IsSuccessful())

	// Teardown
	server.Close()
}

func TestBillService_Status_WithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.BillPayWithError())
	client := New(WithBaseURL(server.URL))
	transactionID := "transaction-id"

	// Act
	transaction, _, err := client.Airtime.Status(context.Background(), transactionID)

	// Assert
	assert.NoError(t, err)
	assert.Nil(t, transaction.Result)

	assert.Equal(t, "General Failure", transaction.Message)
	assert.Equal(t, http.StatusInternalServerError, transaction.Code)
	assert.Nil(t, transaction.Result)
	assert.False(t, transaction.IsSuccessful())

	// Teardown
	server.Close()
}

func TestBillService_Amount_Request(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	apiKey := "api-key"
	agentID := "agent-id"
	agentPlatform := "agent-platform"
	billID := "618442737"
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, [][]byte{stubs.BillAmount()}, &requests)
	client := New(
		WithBaseURL(server.URL),
		WithAgentID(agentID),
		WithAPIKey(apiKey),
		WithAgentPlatform(agentPlatform),
	)

	expectedRequest := map[string]interface{}{
		"biller":        BillerEneoPostpay.string(),
		"billid":        billID,
		"agentid":       agentID,
		"agentplatform": agentPlatform,
		"hash":          fmt.Sprintf("%x", md5.Sum([]byte(BillerEneoPostpay.string()+billID+apiKey))),
	}

	// Act
	_, _, err := client.Bill.Amount(context.Background(), BillerEneoPostpay, billID)
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

func TestBillService_Amount(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.BillAmount())
	client := New(WithBaseURL(server.URL))
	billID := "618442737"
	amount := 1100

	// Act
	transaction, response, err := client.Bill.Amount(context.Background(), BillerEneoPostpay, billID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, &BillAmountResponse{
		Code:    200,
		Message: "success",
		Result:  &amount,
	}, transaction)

	assert.True(t, transaction.IsSuccessful())

	// Teardown
	server.Close()
}

func TestBillService_Amount_WithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.BillPayWithError())
	client := New(WithBaseURL(server.URL))
	billID := "618442737"

	// Act
	transaction, _, err := client.Bill.Amount(context.Background(), BillerEneoPostpay, billID)

	// Assert
	assert.NoError(t, err)
	assert.Nil(t, transaction.Result)

	assert.Equal(t, "General Failure", transaction.Message)
	assert.Equal(t, http.StatusInternalServerError, transaction.Code)
	assert.Nil(t, transaction.Result)
	assert.False(t, transaction.IsSuccessful())

	// Teardown
	server.Close()
}
