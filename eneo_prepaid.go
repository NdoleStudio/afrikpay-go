package afrikpay

type eneoPrepaidService struct {
	*paymentService[ENEOPrepaidPaymentRequest, TransactionStatusResponse]
	*transactionStatusService[TransactionStatusResponse]
}

// ENEOPrepaidPaymentRequest is used to create a prepaid payment request for ENEO
type ENEOPrepaidPaymentRequest struct {
	ReferenceNumber string `json:"referenceNumber"`
	Amount          int    `json:"amount"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	ExternalID      string `json:"externalId"`
	Description     string `json:"description,omitempty"`
}
