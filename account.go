package client

import "net/http"

// AccountBalanceResponse is the response when querying the account balance
type AccountBalanceResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Result  *AccountBalance `json:"result,omitempty"`
}

// AccountBalance contains details about the account
type AccountBalance struct {
	Name        string `json:"name"`
	MainBalance string `json:"mainbalance"`
	MainDeposit string `json:"maindeposit"`
}

// IsSuccessfull determines if the transaction was successful
func (response AccountBalanceResponse) IsSuccessfull() bool {
	return response.Code == http.StatusOK
}
