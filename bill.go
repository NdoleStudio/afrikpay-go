package afrikpay

import (
	"errors"
	"net/http"
	"time"
)

// Biller is the type of bill we want to pay
type Biller string

const (
	// BillerEneoPostpay is used for postpaid bills of ENEO
	BillerEneoPostpay = Biller("eneopostpay")
	// BillerEneoPrepay is used for postpaid bills of ENEO
	BillerEneoPrepay = Biller("eneoprepay")
	// BillerCamwater is used for postpaid bills of Cameroon Water Corporation
	BillerCamwater = Biller("camwater")
	// BillerCanal is used for canal+ subscriptions
	BillerCanal = Biller("canal")
	// BillerUDS is used for bills of the University of Dschang
	BillerUDS = Biller("uds")
)

// string converts the Biller to a string
func (biller Biller) string() string {
	return string(biller)
}

// BillSMS determines if we want to receive a notification via SMS
type BillSMS string

const (
	// BillSMSEnabled means you will receive a notification via SMS
	BillSMSEnabled = BillSMS("yes")
	// BillSMSDisabled means you will not receive a notification via SMS
	BillSMSDisabled = BillSMS("no")
)

// string converts the BillSMS to a string
func (sms *BillSMS) string() string {
	if sms == nil {
		return string(BillSMSDisabled)
	}
	return string(*sms)
}

// Pointer converts BillSMS to *BillSMS
func (sms BillSMS) Pointer() *BillSMS {
	return &sms
}

// BillPayParams parameters for bill payment request
type BillPayParams struct {
	Biller           Biller
	BillID           string
	Mode             Mode
	Amount           *string
	Provider         *string
	Account          *string
	Mobile           *string
	Code             *string
	SMS              *BillSMS
	ProcessingNumber *string
}

// BillPayResponse is returned from bill pay/status api requests
type BillPayResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Result  *BillTransaction `json:"result,omitempty"`
}

// BillTransaction is the details for a bill payment transaction
type BillTransaction struct {
	OperatorID       *string     `json:"operatorid"`
	TransactionID    string      `json:"txnid"`
	Status           string      `json:"status"`
	Date             string      `json:"date"`
	ReferenceID      interface{} `json:"referenceid"`
	ProcessingNumber string      `json:"processingnumber"`
}

// GetDate returns the date as time.Time
func (transaction *BillTransaction) GetDate() (time.Time, error) {
	if transaction == nil {
		return time.Time{}, errors.New("the transaction is nil")
	}

	loc, err := time.LoadLocation("Africa/Douala")
	if err != nil {
		return time.Time{}, err
	}

	return time.ParseInLocation("2006-01-02 15:04:05", transaction.Date, loc)
}

// IsPending determines if the transaction is pending
func (response BillPayResponse) IsPending() bool {
	return response.Code == http.StatusOK && response.Result != nil && response.Result.OperatorID == nil && response.Result.Status == "PENDING"
}

// IsSuccessful determines if the transaction was successful
func (response BillPayResponse) IsSuccessful() bool {
	return response.Code == http.StatusOK && response.Result != nil && response.Result.OperatorID != nil
}

// BillAmountResponse is returned from bill amount api requests
type BillAmountResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  *int   `json:"result,omitempty"`
}

// IsSuccessful determines if the transaction was successful
func (response BillAmountResponse) IsSuccessful() bool {
	return response.Code == http.StatusOK && response.Result != nil && *response.Result > 1000
}
