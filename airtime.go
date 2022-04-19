package afrikpay

import "net/http"

// AirtimeTransferParams parameters for airtime transfer request
type AirtimeTransferParams struct {
	Operator          string
	PhoneNumber       string
	PurchaseReference string
	Amount            string
	Mode              Mode
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

// IsSuccessful determines if the transaction was successful
func (response AirtimeResponse) IsSuccessful() bool {
	return response.Code == http.StatusOK
}
