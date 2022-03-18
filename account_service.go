package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// accountService is the API client for the `/api/account/` endpoint
type accountService service

// Balance is used for the partner to get the Balance of his account
//
// API Docs: https://developer.afrikpay.com/documentation/account/agent/balance/v2/
func (service *accountService) Balance(ctx context.Context) (*AccountBalanceResponse, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodPost, "/api/account/agent/balance/v2/", map[string]string{
		"agentid":       service.client.agentID,
		"agentplatform": service.client.agentPlatform,
		"hash":          service.client.hash(service.client.agentID, service.client.apiKey),
	})
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	balance := new(AccountBalanceResponse)
	if err = json.Unmarshal(*response.Body, &balance); err != nil {
		return nil, response, err
	}

	return balance, response, nil
}
