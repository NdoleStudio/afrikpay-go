package afrikpay

import (
	"encoding/json"
	"testing"
)

func TestPaymentService_Payment_EneoPrepaid(t *testing.T) {
	result := new(ENEOPrepaidPaymentResponse)
	if err = json.Unmarshal(stub, request); err != nil {
		return nil, response, err
	}
}
