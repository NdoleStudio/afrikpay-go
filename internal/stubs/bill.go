package stubs

// BillPay is a dummy json response for the `/api/bill/v2/` endpoint
func BillPay() []byte {
	return []byte(`
		{
		   "code":200,
		   "message":"success",
		   "result":{
			  "operatorid":"xxxx-xxxx-xxxx-xxxx-5286 : 0000000000068 : 8.8 Kwh",
			  "txnid":"5xxxx",
			  "status":"PENDING",
			  "date":"2022-04-19 18:00:06",
			  "referenceid":null,
			  "processingnumber":"aaba045a-d571-41e9-9ea4-54cd78782e03"
		   }
		}
`)
}

// BillPayWithError is a dummy json response for the `/api/airtime/v2/` endpoint with an error
func BillPayWithError() []byte {
	return []byte(`
		{
		   "code":500,
		   "message":"General Failure",
		   "result":null
		}
`)
}
