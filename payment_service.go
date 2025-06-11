package afrikpay

import (
	"context"
	"encoding/json"
	"net/http"
)

type paymentService[T, V any] struct {
	*service
	slug string
}

// Payment carries out a transaction (airtime, bill, taxes, school, ENEO)
//
// https://developers.afrikpay.com/devportal/apis/5ef8e2f6-0765-43fa-b6d2-ddd34de7ef1a/documents/default
func (payment *paymentService[T, V]) Payment(ctx context.Context, payload *T) (*V, *Response, error) {
	request, err := payment.client.newRequest(ctx, http.MethodPost, "/api/oss/payment/partner/v1", payload)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Add("Service", payment.slug)
	request.Header.Add("Pin", payment.client.walletPin)

	response, err := payment.client.do(request)
	if err != nil {
		return nil, response, err
	}

	result := new(V)
	if err = json.Unmarshal(*response.Body, result); err != nil {
		return nil, response, err
	}

	return result, response, nil
}
