package afrikpay

import (
	"context"
	"encoding/json"
	"net/http"
)

// BalanceResponse is the response when querying the account balance
type BalanceResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Result  *BalanceResponseResult `json:"result"`
}

// BalanceResponseResult contains the result of the balance response
type BalanceResponseResult struct {
	AccountDepositNumber     string  `json:"accountDepositNumber"`
	AccountCommissionNumber  string  `json:"accountCommissionNumber"`
	AccountDepositBalance    float64 `json:"accountDepositBalance"`
	AccountCommissionBalance float64 `json:"accountCommissionBalance"`
}

// Balance lists all visible accounts of the given user, and for each account show the current deposit balance, and current commission.
//
// API Docs: https://developers.afrikpay.com/devportal/apis/cc20a007-0c4f-4de8-ba1a-17ab6dec5e63/documents/default
func (client *Client) Balance(ctx context.Context) (*BalanceResponse, *Response, error) {
	request, err := client.newRequest(ctx, http.MethodPost, "/api/oss/balance/partner/v1", nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := client.do(request)
	if err != nil {
		return nil, response, err
	}

	balance := new(BalanceResponse)
	if err = json.Unmarshal(*response.Body, &balance); err != nil {
		return nil, response, err
	}

	return balance, response, nil
}
