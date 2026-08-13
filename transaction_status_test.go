package afrikpay

import (
	"encoding/json"
	"testing"

	"github.com/NdoleStudio/afrikpay-go/internal/stubs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactionStatusResponse_DecodesReferenceToken(t *testing.T) {
	response := new(TransactionStatusResponse)

	err := json.Unmarshal(stubs.ENEOPrepaidResponse(), response)

	require.NoError(t, err)
	require.NotNil(t, response.Result)
	assert.Equal(t, "8844-9878-2415-5207-7373", response.Result.Reference.Token)
}

func TestTransactionStatusResponse_IsNotPerformed_WithCode51314(t *testing.T) {
	// Arrange
	response := &TransactionStatusResponse{
		Code:    51314,
		Message: "System response mismatch",
	}

	// Act & Assert
	assert.True(t, response.IsNotPerformed())
}

func TestTransactionStatusResponse_IsNotPerformed_WithCode40618(t *testing.T) {
	// Arrange
	response := &TransactionStatusResponse{
		Code:    40618,
		Message: "Transaction not found",
	}

	// Act & Assert
	assert.True(t, response.IsNotPerformed())
}

func TestTransactionStatusResponse_IsNotPerformed_WithMessageSystemResponseMismatch(t *testing.T) {
	// Arrange
	response := &TransactionStatusResponse{
		Code:    0,
		Message: "System response mismatch",
	}

	// Act & Assert
	assert.True(t, response.IsNotPerformed())
}

func TestTransactionStatusResponse_IsNotPerformed_WithMessageTransactionNotFound(t *testing.T) {
	// Arrange
	response := &TransactionStatusResponse{
		Code:    0,
		Message: "Transaction not found",
	}

	// Act & Assert
	assert.True(t, response.IsNotPerformed())
}

func TestTransactionStatusResponse_IsNotPerformed_ReturnsFalseForSuccessfulResponse(t *testing.T) {
	// Arrange
	response := &TransactionStatusResponse{
		Code:    200,
		Message: "success",
	}

	// Act & Assert
	assert.False(t, response.IsNotPerformed())
}
