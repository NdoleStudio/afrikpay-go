package afrikpay

import (
	"context"
	"encoding/json"
	"net/http"
)

// airtimeService is the API client for the `/` endpoint
type airtimeService service

// Transfer is intended for communication / Internet credit transfer operations to telephone numbers.
//
// API Docs: https://developer.afrikpay.com/documentation/airtime/v2/
func (service *airtimeService) Transfer(ctx context.Context, params AirtimeTransferParams) (*AirtimeResponse, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodPost, "/api/airtime/v2/", map[string]string{
		"operator":      params.Operator,
		"reference":     params.PhoneNumber,
		"amount":        params.Amount,
		"mode":          params.Mode.String(),
		"purchaseref":   params.PurchaseReference,
		"agentid":       service.client.agentID,
		"agentplatform": service.client.agentPlatform,
		"agentpwd":      service.client.agentPassword,
		"hash":          service.client.hash(params.Operator, params.PhoneNumber, params.Amount, service.client.apiKey),
	})
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	status := new(AirtimeResponse)
	if err = json.Unmarshal(*response.Body, status); err != nil {
		return nil, response, err
	}

	return status, response, nil
}

// Status is intended for getting the status of an airtime transaction
//
// API Docs: https://developer.afrikpay.com/documentation/airtime/status/v2/
func (service *airtimeService) Status(ctx context.Context, transactionID string) (*AirtimeResponse, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodPost, "/api/airtime/status/v2/", map[string]string{
		"processingnumber": transactionID,
		"agentid":          service.client.agentID,
		"agentplatform":    service.client.agentPlatform,
		"hash":             service.client.hash(transactionID, service.client.apiKey),
	})
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	status := new(AirtimeResponse)
	if err = json.Unmarshal(*response.Body, status); err != nil {
		return nil, response, err
	}

	return status, response, nil
}
