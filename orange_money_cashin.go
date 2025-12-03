package afrikpay

type orangeMoneyCashinService struct {
	*cashinService[OrangeMoneyCashinPaymentRequest, OrangeMoneyCashinPaymentResponse]
	*transactionStatusService[OrangeMoneyCashinPaymentResponse]
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

// OrangeMoneyCashinPaymentResponse is the response from Afrikpay when creating a cashin payment request via Orange Money
type OrangeMoneyCashinPaymentResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  *struct {
		ErrorCode    any    `json:"errorCode"`
		ErrorMessage any    `json:"errorMessage"`
		ErrorType    any    `json:"errorType"`
		Status       string `json:"status"`
		CallbackURL  string `json:"callbackUrl"`
		Voucher      struct {
			PayToken     string `json:"payToken"`
			ProviderTime string `json:"providerTime"`
			Value        string `json:"value"`
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
	} `json:"result"`
}

// IsFailed checks if the Orange Money Cashin payment response indicates a failure
func (response *OrangeMoneyCashinPaymentResponse) IsFailed() bool {
	return response.Code != 200 || (response.Result != nil && response.Result.Status == "FAILED")
}

// IsInProgress checks if the Orange Money Cashin payment is still in progress
func (response *OrangeMoneyCashinPaymentResponse) IsInProgress() bool {
	return response.Code == 200 && response.Result != nil && (response.Result.Status == "PROGRESS" || response.Result.Status == "PENDING" || response.Result.Status == "ACCEPTED" || response.Result.Status == "PAYED")
}
