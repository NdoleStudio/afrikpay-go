package afrikpay

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/NdoleStudio/afrikpay-go/internal/stubs"
	"github.com/davecgh/go-spew/spew"
)

func TestPaymentService_Payment_EneoPrepaid(t *testing.T) {
	result := new(ENEOPrepaidPaymentResponse)
	if err := json.Unmarshal(stubs.ENEOPrepaidResponse(), result); err != nil {
		log.Fatal(err)
	}
	spew.Dump(result.Result)
}
