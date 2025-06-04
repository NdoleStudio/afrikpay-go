package afrikpay

// WebhookStatus is the status of a webhook response
type WebhookStatus string

const (
	// WebhookStatusRejected is used tell Afrikpay to cancel the transaction
	WebhookStatusRejected WebhookStatus = "REJECTED"
	// WebhookStatusAccepted is used tell Afrikpay to proceed with the transaction
	WebhookStatusAccepted WebhookStatus = "ACCEPTED"
)

// WebhookStatusResponse is the response from Afrikpay when checking the status of a transaction
type WebhookStatusResponse struct {
	Code    int                         `json:"code"`
	Message string                      `json:"message"`
	Result  WebhookStatusResponseResult `json:"result"`
}

// WebhookStatusResponseResult contains the details of the transaction status response
type WebhookStatusResponseResult struct {
	Status          WebhookStatus `json:"status"`
	Amount          int           `json:"amount"`
	ReferenceNumber *string       `json:"referenceNumber"`
	ErrorMessage    *string       `json:"errorMessage"`
	ErrorCode       *int          `json:"errorCode"`
	Description     *string       `json:"description"`
}

// WebhookStatusRequest is used to check the status of a transaction by afrikpay
type WebhookStatusRequest struct {
	ExternalID string `json:"externalId"`
}
