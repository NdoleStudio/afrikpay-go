package afrikpay

// TransactionStatusRequest is used to request the status of a transaction
type TransactionStatusRequest struct {
	ReferenceNumber string `json:"referenceNumber,omitempty"`
	Amount          int    `json:"amount,omitempty"`
	ExternalID      string `json:"externalId,omitempty"`
	RequestID       string `json:"requestId,omitempty"`
	TransactionID   int    `json:"transactionId,omitempty"`
	FinancialID     string `json:"financialId,omitempty"`
	ProviderID      string `json:"providerId,omitempty"`
}
