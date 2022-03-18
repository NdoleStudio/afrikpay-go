package stubs

// AccountBalanceWithError is a dummy json response for the `/api/account/agent/balance/v2/` endpoint with an error
func AccountBalanceWithError() []byte {
	return []byte(`
		{
			"code": 500,
			"message": "412: bad password",
			"result": null
		}
`)
}
