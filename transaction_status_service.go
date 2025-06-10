package afrikpay

import (
	"context"
	"encoding/json"
	"net/http"
)

type transactionStatusService[T, V any] struct {
	*service
	slug string
}

// TransactionStatus Get status of specific transaction
//
// https://developers.afrikpay.com/devportal/apis/3f939c15-8d37-4497-ad04-fa1ae2481aae/documents/default
func (service *transactionStatusService[T, V]) TransactionStatus(ctx context.Context, payload *T) (*V, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodPost, "/api/oss/transaction/status/partner/v1", payload)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Add("Service", service.slug)

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	result := new(V)
	if err = json.Unmarshal(*response.Body, request); err != nil {
		return nil, response, err
	}

	return result, response, nil
}
