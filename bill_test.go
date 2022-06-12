package afrikpay

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/NdoleStudio/afrikpay-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestBillPayResponse_IsPending(t *testing.T) {
	t.Run("pending transaction is true", func(t *testing.T) {
		// Arrange
		response := new(BillPayResponse)
		err := json.Unmarshal(stubs.BillPayPending(), response)
		assert.Nil(t, err)

		// Act
		isPending := response.IsPending()

		// Assert
		assert.True(t, isPending)
	})

	t.Run("failed transaction is false", func(t *testing.T) {
		// Arrange
		response := new(BillPayResponse)
		err := json.Unmarshal(stubs.BillPayWithError(), response)
		assert.Nil(t, err)

		// Act
		isPending := response.IsPending()

		// Assert
		assert.False(t, isPending)
	})

	t.Run("successfull transaction is false", func(t *testing.T) {
		// Arrange
		response := new(BillPayResponse)
		err := json.Unmarshal(stubs.BillPay(), response)
		assert.Nil(t, err)

		// Act
		isPending := response.IsPending()

		// Assert
		assert.False(t, isPending)
	})
}

func TestBillPayResponse_IsSuccessful(t *testing.T) {
	t.Run("pending transaction is false", func(t *testing.T) {
		// Arrange
		response := new(BillPayResponse)
		err := json.Unmarshal(stubs.BillPayPending(), response)
		assert.Nil(t, err)

		// Act
		isSuccessful := response.IsSuccessful()

		// Assert
		assert.False(t, isSuccessful)
	})

	t.Run("failed transaction is false", func(t *testing.T) {
		// Arrange
		response := new(BillPayResponse)
		err := json.Unmarshal(stubs.BillPayWithError(), response)
		assert.Nil(t, err)

		// Act
		isSuccessful := response.IsSuccessful()

		// Assert
		assert.False(t, isSuccessful)
	})

	t.Run("successfull transaction is true", func(t *testing.T) {
		// Arrange
		response := new(BillPayResponse)
		err := json.Unmarshal(stubs.BillPay(), response)
		assert.Nil(t, err)

		// Act
		isSuccessful := response.IsSuccessful()

		// Assert
		assert.True(t, isSuccessful)
	})
}

func TestBillTransaction_GetDate(t *testing.T) {
	t.Run("pending transaction returns correct date", func(t *testing.T) {
		// Arrange
		response := new(BillPayResponse)
		err := json.Unmarshal(stubs.BillPayPending(), response)
		expected := time.Date(2022, 6, 11, 13, 37, 31, 0, time.UTC)
		assert.Nil(t, err)

		// Act
		date, err := response.Result.GetDate()

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, expected.Unix(), date.Unix())
	})

	t.Run("failed transaction returns error", func(t *testing.T) {
		// Arrange
		response := new(BillPayResponse)
		err := json.Unmarshal(stubs.BillPayWithError(), response)
		assert.Nil(t, err)

		// Act
		_, err = response.Result.GetDate()

		// Assert
		assert.NotNil(t, err)
	})

	t.Run("successfull transaction returns correct date", func(t *testing.T) {
		// Arrange
		response := new(BillPayResponse)
		err := json.Unmarshal(stubs.BillPay(), response)
		expected := time.Date(2022, 4, 19, 17, 0, 6, 0, time.UTC)
		assert.Nil(t, err)

		// Act
		date, err := response.Result.GetDate()

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, expected.Unix(), date.Unix())
	})
}
