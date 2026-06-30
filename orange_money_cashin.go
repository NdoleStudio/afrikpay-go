package afrikpay

type orangeMoneyCashinService struct {
	*cashinService[OrangeMoneyCashinPaymentRequest, TransactionStatusResponse]
	*transactionStatusService[TransactionStatusResponse]
}

// OrangeMoneyCashinPaymentRequest Cashin via Orange Money
type OrangeMoneyCashinPaymentRequest struct {
	ReferenceNumber string `json:"referenceNumber"`
	Amount          int    `json:"amount"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	ExternalID      string `json:"externalId"`
	Description     string `json:"description"`
}
