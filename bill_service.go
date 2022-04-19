package afrikpay

import (
	"context"
	"encoding/json"
	"net/http"
)

// billService is the API client for the `/api/bill/` endpoint
type billService service

// Pay Bills or Subscriptions
//
// API Docs: https://developer.afrikpay.com/documentation/bill/v2/
func (service *billService) Pay(ctx context.Context, params BillPayParams) (*BillResponse, *Response, error) {
	request, err := service.client.newRequest(
		ctx,
		http.MethodPost,
		"/api/bill/v2/",
		service.billPayParamsToPayload(params),
	)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	status := new(BillResponse)
	if err = json.Unmarshal(*response.Body, status); err != nil {
		return nil, response, err
	}

	return status, response, nil
}

// Status is intended for getting the status of an airtime transaction
//
// API Docs: https://developer.afrikpay.com/documentation/airtime/status/v2/
func (service *billService) Status(ctx context.Context, transactionID string) (*BillResponse, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodPost, "/api/bill/status/v2/", map[string]string{
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

	status := new(BillResponse)
	if err = json.Unmarshal(*response.Body, status); err != nil {
		return nil, response, err
	}

	return status, response, nil
}

func (service *billService) billPayParamsToPayload(params BillPayParams) map[string]string {
	payload := map[string]string{
		"biller":        params.Biller.string(),
		"billid":        params.BillID,
		"mode":          params.Mode.String(),
		"agentid":       service.client.agentID,
		"agentplatform": service.client.agentPlatform,
		"agentpwd":      service.client.agentPassword,
		"hash":          service.client.hash(params.Biller.string(), params.BillID, PointerToString(params.Amount), service.client.apiKey),
	}

	if params.Amount != nil {
		payload["amount"] = PointerToString(params.Amount)
	}
	if params.Provider != nil {
		payload["provider"] = PointerToString(params.Provider)
	}
	if params.Account != nil {
		payload["account"] = PointerToString(params.Account)
	}
	if params.Mobile != nil {
		payload["mobile"] = PointerToString(params.Mobile)
	}
	if params.Code != nil {
		payload["code"] = PointerToString(params.Code)
	}
	if params.SMS != nil {
		payload["sms"] = params.SMS.string()
	}
	if params.ProcessingNumber != nil {
		payload["processingnumber"] = PointerToString(params.ProcessingNumber)
	}

	return payload
}
