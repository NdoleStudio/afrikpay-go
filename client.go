package afrikpay

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type service struct {
	client *Client
}

// Client is the afrikpay API client.
// Exec not instantiate this client with Client{}. Use the New method instead.
type Client struct {
	httpClient          *http.Client
	common              service
	baseURL             string
	apiKey              string
	authorizationHeader string
	walletPin           string

	ENEOPrepaid       *eneoPrepaidService
	CanalPlus         *canalPlusService
	OrangeMoneyCashin *orangeMoneyCashinService
	OrangeAirtime     *airtimeService
	CamtelAirtime     *airtimeService
	MTNAirtime        *airtimeService
}

// New creates and returns a new *Client from a slice of Option.
func New(options ...Option) *Client {
	config := defaultClientConfig()

	for _, option := range options {
		option.apply(config)
	}

	client := &Client{
		httpClient:          config.httpClient,
		baseURL:             config.baseURL,
		apiKey:              config.apiKey,
		authorizationHeader: base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.walletUsername, config.walletPassword))),
		walletPin:           config.walletPin,
	}

	client.common.client = client
	client.ENEOPrepaid = &eneoPrepaidService{
		paymentService: &paymentService[ENEOPrepaidPaymentRequest, ENEOPrepaidPaymentResponse]{
			service: &client.common,
			slug:    "eneo-prepaid-bill-service-feature",
		},
		transactionStatusService: &transactionStatusService[ENEOPrepaidPaymentResponse]{
			service: &client.common,
			slug:    "eneo-prepaid-bill-service-feature",
		},
	}

	client.CanalPlus = &canalPlusService{
		paymentService: &paymentService[CanalPlusPaymentRequest, TransactionStatusResponse]{
			service: &client.common,
			slug:    "canal-auto-bill-service-feature",
		},
		transactionStatusService: &transactionStatusService[TransactionStatusResponse]{
			service: &client.common,
			slug:    "canal-auto-bill-service-feature",
		},
	}

	client.OrangeMoneyCashin = &orangeMoneyCashinService{
		cashinService: &cashinService[OrangeMoneyCashinPaymentRequest, OrangeMoneyCashinPaymentResponse]{
			service: &client.common,
			slug:    "orange-money-cashin-service-feature",
		},
		transactionStatusService: &transactionStatusService[OrangeMoneyCashinPaymentResponse]{
			service: &client.common,
			slug:    "orange-money-cashin-service-feature",
		},
	}

	client.OrangeAirtime = &airtimeService{
		paymentService: &paymentService[AirtimePaymentRequest, TransactionStatusResponse]{
			service: &client.common,
			slug:    "orange-airtime-service-feature",
		},
		transactionStatusService: &transactionStatusService[TransactionStatusResponse]{
			service: &client.common,
			slug:    "orange-airtime-service-feature",
		},
	}

	client.MTNAirtime = &airtimeService{
		paymentService: &paymentService[AirtimePaymentRequest, TransactionStatusResponse]{
			service: &client.common,
			slug:    "mtn-airtime-service-feature",
		},
		transactionStatusService: &transactionStatusService[TransactionStatusResponse]{
			service: &client.common,
			slug:    "mtn-airtime-service-feature",
		},
	}

	client.CamtelAirtime = &airtimeService{
		paymentService: &paymentService[AirtimePaymentRequest, TransactionStatusResponse]{
			service: &client.common,
			slug:    "camtel-manual-airtime-service-feature",
		},
		transactionStatusService: &transactionStatusService[TransactionStatusResponse]{
			service: &client.common,
			slug:    "camtel-manual-airtime-service-feature",
		},
	}

	return client
}

// newRequest creates an API request.
// A relative URL must be provided in uri in which case it is resolved relative to the BaseURL of the Client.
// URI's should always be specified with a preceding slash.
func (client *Client) newRequest(ctx context.Context, method, uri string, body any) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, client.baseURL+uri, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("ApiKey", client.apiKey)
	req.Header.Set("X-Authorization", client.authorizationHeader)

	return req, nil
}

// do carries out an HTTP request and returns a Response
func (client *Client) do(request *http.Request) (*Response, error) {
	httpResponse, err := client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer func() { _ = httpResponse.Body.Close() }()

	resp, err := client.newResponse(httpResponse)
	if err != nil {
		return resp, err
	}

	_, err = io.Copy(io.Discard, httpResponse.Body)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// newResponse converts an *http.Response to *Response
func (client *Client) newResponse(httpResponse *http.Response) (*Response, error) {
	if httpResponse == nil {
		return nil, fmt.Errorf("%T cannot be nil", httpResponse)
	}

	resp := new(Response)
	resp.HTTPResponse = httpResponse

	buf, err := io.ReadAll(resp.HTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	resp.Body = &buf

	return resp, resp.Error()
}
