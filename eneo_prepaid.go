package afrikpay

// ENEOPrepaidPaymentRequest is used to create a prepaid payment request for ENEO
type ENEOPrepaidPaymentRequest struct {
	ReferenceNumber string `json:"referenceNumber"`
	Amount          int    `json:"amount"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	ExternalID      string `json:"externalId"`
	Description     string `json:"description"`
}

// ENEOPrepaidPaymentResponse is the response from Afrikpay when creating a prepaid payment request for ENEO
type ENEOPrepaidPaymentResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  *struct {
		ErrorCode    any    `json:"errorCode"`
		ErrorMessage any    `json:"errorMessage"`
		ErrorType    any    `json:"errorType"`
		Status       string `json:"status"`
		CallbackURL  string `json:"callbackUrl"`
		Voucher      struct {
			ContactNno         string `json:"contactNno"`
			DaysLastPurchase   string `json:"daysLastPurchase"`
			ValeurKwh          string `json:"valeurKwh"`
			AccNumber          string `json:"accNumber"`
			Token              string `json:"token"`
			DateTime           string `json:"dateTime"`
			AmountDebt         string `json:"amountDebt"`
			BalanceDebt        string `json:"balanceDebt"`
			ReceiptNo          int    `json:"receiptNo"`
			LessRound          string `json:"lessRound"`
			AmountTender       string `json:"amountTender"`
			StatusDateTime     any    `json:"statusDateTime"`
			StatusUniqueNumber any    `json:"statusUniqueNumber"`
			Value              string `json:"value"`
		} `json:"voucher"`
		TransactionID       int      `json:"transactionId"`
		AccountName         string   `json:"accountName"`
		AccountNumber       string   `json:"accountNumber"`
		Username            string   `json:"username"`
		ReferenceNumber     string   `json:"referenceNumber"`
		Amount              int      `json:"amount"`
		Type                string   `json:"type"`
		Service             string   `json:"service"`
		ServiceName         string   `json:"serviceName"`
		FinancialFees       int      `json:"financialFees"`
		FinancialCommission int      `json:"financialCommission"`
		ProviderFees        int      `json:"providerFees"`
		Phone               string   `json:"phone"`
		Email               string   `json:"email"`
		Code                string   `json:"code"`
		OptionSlug          string   `json:"optionSlug"`
		Description         string   `json:"description"`
		ExternalID          string   `json:"externalId"`
		FinancialID         string   `json:"financialId"`
		ProviderID          string   `json:"providerId"`
		RequestID           string   `json:"requestId"`
		Data                struct{} `json:"data"`
		RequestStatus       string   `json:"requestStatus"`
		CommissionID        string   `json:"commissionId"`
		RollbackID          string   `json:"rollbackId"`
		TerminalID          int      `json:"terminalId"`
		TerminalName        string   `json:"terminalName"`
		TerminalUserAgent   string   `json:"terminalUserAgent"`
		Reference           struct {
			ErrorCode       any      `json:"errorCode"`
			ErrorMessage    any      `json:"errorMessage"`
			ErrorType       any      `json:"errorType"`
			Status          string   `json:"status"`
			CallbackURL     any      `json:"callbackUrl"`
			Voucher         struct{} `json:"voucher"`
			ReferenceID     string   `json:"referenceId"`
			ReferenceNumber string   `json:"referenceNumber"`
			Amount          int      `json:"amount"`
			Name            string   `json:"name"`
			GenerationDate  any      `json:"generationDate"`
			ExpirationDate  any      `json:"expirationDate"`
			Date            string   `json:"date"`
			ProviderCode    string   `json:"providerCode"`
			ProviderMessage any      `json:"providerMessage"`
			Options         any      `json:"options"`
			ProviderStatus  any      `json:"providerStatus"`
			Adress          string   `json:"adress"`
			Contact         string   `json:"contact"`
			Reflocal        string   `json:"reflocal"`
			AccNo           string   `json:"accNo"`
			Ti              string   `json:"ti"`
			Krn             string   `json:"krn"`
			Sgc             string   `json:"sgc"`
			At              int      `json:"at"`
			Tt              int      `json:"tt"`
			DateTime        string   `json:"dateTime"`
			Token           string   `json:"token"`
		} `json:"reference"`
		IPAddress             string `json:"ipAddress"`
		Date                  string `json:"date"`
		Signature             string `json:"signature"`
		PaymentServiceFeature string `json:"paymentServiceFeature"`
		PaymentWallet         string `json:"paymentWallet"`
		NoFees                bool   `json:"noFees"`
		PaymentLink           string `json:"paymentLink"`
		AcceptURL             string `json:"acceptUrl"`
		DeclineURL            string `json:"declineUrl"`
		CancelURL             string `json:"cancelUrl"`
	} `json:"result"`
}

type eneoPrepaidService struct {
	*paymentService[ENEOPrepaidPaymentRequest, ENEOPrepaidPaymentResponse]
	*transactionStatusService[ENEOPrepaidPaymentRequest, ENEOPrepaidPaymentResponse]
}
