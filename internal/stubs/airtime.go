package stubs

// Transfer is a dummy json response for the `/api/airtime/v2/` endpoint
func Transfer() []byte {
	return []byte(`
		{
			"code": 200,
			"message": "Success",
			"result": {
				"operatorid": "1647539307",
				"txnid": "1069",
				"status": "SUCCESS",
				"date": "2022-03-17 18:48:26",
				"ticket": null,
				"referenceid": "18360",
				"processingnumber": "aaba045a-d571-41e9-9ea4-54cd78782e03"
			}
		}
`)
}

// TransferWithError is a dummy json response for the `/api/airtime/v2/` endpoint with an error
func TransferWithError() []byte {
	return []byte(`
		{
			"code": 500,
			"message": "412: bad password",
			"result": null
		}
`)
}
