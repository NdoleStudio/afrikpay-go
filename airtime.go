package afrikpay

import "net/http"

// AirtimeTransferParams parameters for airtime transfer request
type AirtimeTransferParams struct {
	Operator          string
	PhoneNumber       string
	PurchaseReference string
	Amount            string
	Mode              AirtimeMode
}

// AirtimeMode is used to determine the payment mode
type AirtimeMode string

const (
	// AirtimeModeCash payment from agent partner account
	AirtimeModeCash = AirtimeMode("cash")
	// AirtimeModeMoney send payment request
	AirtimeModeMoney = AirtimeMode("money")
	// AirtimeModeAccount make payment operation from afrikpay client
	AirtimeModeAccount = AirtimeMode("account")
)

// String converts the AirtimeMode to a string
func (mode AirtimeMode) String() string {
	return string(mode)
}

// AirtimeResponse is returned from airtime api requests
type AirtimeResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Result  *AirtimeTransaction `json:"result,omitempty"`
}

// AirtimeTransaction is the details for an aitime transaction
type AirtimeTransaction struct {
	OperatorID       string      `json:"operatorid"`
	TransactionID    string      `json:"txnid"`
	Status           string      `json:"status"`
	Date             string      `json:"date"`
	Ticket           interface{} `json:"ticket,omitempty"`
	ReferenceID      string      `json:"referenceid"`
	ProcessingNumber string      `json:"processingnumber"`
}

// IsSuccessfull determines if the transaction was successful
func (response AirtimeResponse) IsSuccessfull() bool {
	return response.Code == http.StatusOK
}
