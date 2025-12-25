package afrikpay

import (
	"context"
	"encoding/json"
	"net/http"
)

type optionService[T, V any] struct {
	*service
	slug string
}

// GetOptions returns the payment options for a service e.g. CANAL+
//
// https://developers.afrikpay.com/devportal/apis/929f94b6-58df-49d3-969f-dbdb842c064c/documents/default
func (service *optionService[T, V]) GetOptions(ctx context.Context, payload *T) (*V, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodPost, "/api/oss/option/partner/v1", payload)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Add("Service", service.slug)

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	result := new(V)
	if err = json.Unmarshal(*response.Body, result); err != nil {
		return nil, response, err
	}

	return result, response, nil
}
