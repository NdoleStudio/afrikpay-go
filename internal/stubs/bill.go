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

// BillPayPending is a dummy json response for the `/api/bill/v2/` endpoint
func BillPayPending() []byte {
	return []byte(`
		{
		   "code":200,
		   "message":"success",
		   "result":{
			  "operatorid":null,
			  "txnid":"63854",
			  "status":"PENDING",
			  "date":"2022-06-11 14:37:31",
			  "referenceid":null,
			  "processingnumber":"123456"
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

// BillAmount is a dummy json response for the `/api/bill/getamount/` endpoint
func BillAmount() []byte {
	return []byte(`
		{
		   "code":200,
		   "message":"success",
		   "result": 1100
		}
`)
}
