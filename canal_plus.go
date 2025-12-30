package afrikpay

type canalPlusService struct {
	*paymentService[CanalPlusPaymentRequest, TransactionStatusResponse]
	*optionService[CanalPlusOptionRequest, CanalPlusOptionResponse]
	*transactionStatusService[TransactionStatusResponse]
}

// CanalPlusPaymentRequest is used to create a prepaid payment request for Canal Plus
type CanalPlusPaymentRequest struct {
	ReferenceNumber string `json:"referenceNumber"`
	OptionSlug      string `json:"optionSlug"`
	Amount          int    `json:"amount"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	ExternalID      string `json:"externalId"`
	Description     string `json:"description,omitempty"`
}

// CanalPlusOptionRequest represents the request structure for Canal Plus payment options
type CanalPlusOptionRequest struct {
	ReferenceNumber string `json:"referenceNumber"`
}

// CanalPlusOptionResponse represents the response structure for Canal Plus payment options
type CanalPlusOptionResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  []struct {
		Voucher         any    `json:"voucher"`
		OptionID        int    `json:"optionId"`
		Name            string `json:"name"`
		Slug            string `json:"slug"`
		Amount          int    `json:"amount"`
		ReferenceNumber string `json:"referenceNumber"`
		Date            string `json:"date"`
	} `json:"result"`
}
