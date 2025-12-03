package afrikpay

type airtimeService struct {
	*paymentService[AirtimePaymentRequest, TransactionStatusResponse]
	*transactionStatusService[TransactionStatusResponse]
}

// AirtimePaymentRequest Airtime purchase request
type AirtimePaymentRequest struct {
	ReferenceNumber string `json:"referenceNumber"`
	Amount          int    `json:"amount"`
	Email           string `json:"email"`
	ExternalID      string `json:"externalId"`
	Description     string `json:"description"`
}
