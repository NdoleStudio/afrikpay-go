package afrikpay

type canalPlusService struct {
	*paymentService[CanalPlusPaymentRequest, TransactionStatusResponse]
	*transactionStatusService[TransactionStatusResponse]
}

// CanalPlusPaymentRequest is used to create a prepaid payment request for Canal Plus
type CanalPlusPaymentRequest struct {
	ReferenceNumber string `json:"referenceNumber"`
	OptionSlug      string `json:"option_slug"`
	Amount          int    `json:"amount"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	ExternalID      string `json:"externalId"`
	Description     string `json:"description,omitempty"`
}
