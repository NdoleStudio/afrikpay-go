package afrikpay

// PointerToString converts a string to *string
func PointerToString(input *string) string {
	if input == nil {
		return ""
	}
	return *input
}

// StringToPointer converts a string to *string
func StringToPointer(input string) *string {
	return &input
}
