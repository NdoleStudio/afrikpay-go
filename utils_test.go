package afrikpay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointerToString(t *testing.T) {
	// Arrange
	input := "test-input"

	// Act
	inputString := PointerToString(&input)

	// Assert
	assert.Equal(t, input, inputString)
}

func TestStringToPointer(t *testing.T) {
	// Arrange
	input := "test-input"

	// Act
	inputPtr := StringToPointer(input)

	// Assert
	assert.Equal(t, &input, inputPtr)
}
