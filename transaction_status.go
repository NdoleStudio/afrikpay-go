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

// TransactionStatusResponse is the response from Afrikpay for a transaction status
type TransactionStatusResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  *struct {
		ErrorCode    *int   `json:"errorCode"`
		ErrorMessage any    `json:"errorMessage"`
		ErrorType    any    `json:"errorType"`
		Status       string `json:"status"`
		CallbackURL  string `json:"callbackUrl"`
		Voucher      struct {
			Value string `json:"value"`
		} `json:"voucher"`
		TransactionID         int      `json:"transactionId"`
		AccountName           string   `json:"accountName"`
		AccountNumber         string   `json:"accountNumber"`
		Username              string   `json:"username"`
		ReferenceNumber       string   `json:"referenceNumber"`
		Amount                int      `json:"amount"`
		Type                  string   `json:"type"`
		Service               string   `json:"service"`
		ServiceName           string   `json:"serviceName"`
		FinancialFees         int      `json:"financialFees"`
		FinancialCommission   int      `json:"financialCommission"`
		ProviderFees          int      `json:"providerFees"`
		Phone                 string   `json:"phone"`
		Email                 string   `json:"email"`
		Code                  string   `json:"code"`
		OptionSlug            string   `json:"optionSlug"`
		Description           string   `json:"description"`
		ExternalID            string   `json:"externalId"`
		FinancialID           string   `json:"financialId"`
		ProviderID            string   `json:"providerId"`
		RequestID             string   `json:"requestId"`
		Data                  struct{} `json:"data"`
		RequestStatus         string   `json:"requestStatus"`
		CommissionID          string   `json:"commissionId"`
		RollbackID            string   `json:"rollbackId"`
		TerminalID            int      `json:"terminalId"`
		TerminalName          string   `json:"terminalName"`
		TerminalUserAgent     string   `json:"terminalUserAgent"`
		Reference             struct{} `json:"reference"`
		IPAddress             string   `json:"ipAddress"`
		Date                  string   `json:"date"`
		Signature             string   `json:"signature"`
		PaymentServiceFeature string   `json:"paymentServiceFeature"`
		PaymentWallet         string   `json:"paymentWallet"`
		NoFees                bool     `json:"noFees"`
		PaymentLink           string   `json:"paymentLink"`
		AcceptURL             string   `json:"acceptUrl"`
		DeclineURL            string   `json:"declineUrl"`
		CancelURL             string   `json:"cancelUrl"`
	} `json:"result"`
}

// IsFailed checks if the CANAL+ payment has failed
func (response *TransactionStatusResponse) IsFailed() bool {
	return response.Code != 200 ||
		(response.Result != nil &&
			(response.Result.Status == "FAILED" || response.Result.ErrorType == "TransactionExternalIdNotFoundException") ||
			(response.Result.ErrorCode != nil && (*response.Result.ErrorCode == 40633 || *response.Result.ErrorCode == 40614)))
}

// IsInProgress checks if the CANAL+ payment is still in progress
func (response *TransactionStatusResponse) IsInProgress() bool {
	return response.Code == 200 && response.Result != nil && (response.Result.Status == "PROGRESS" || response.Result.Status == "PENDING" || response.Result.Status == "ACCEPTED" || response.Result.Status == "PAYED")
}
