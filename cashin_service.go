package afrikpay

import (
	"context"
	"encoding/json"
	"net/http"
)

type cashinService[T, V any] struct {
	*service
	slug string
}

// MakeDeposit on wallet
//
// https://developers.afrikpay.com/devportal/apis/cdb09c3d-5128-4fa7-bf09-4d0fd1cf2948/documents/default
func (payment *cashinService[T, V]) MakeDeposit(ctx context.Context, payload *T) (*V, *Response, error) {
	request, err := payment.client.newRequest(ctx, http.MethodPost, "/api/oss/cashin/partner/v1", payload)
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
