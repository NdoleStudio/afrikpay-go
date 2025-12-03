package stubs

// OrangeMoneyCashinResponse returns a stubbed response for Orange Money Cashin payment
func OrangeMoneyCashinResponse() []byte {
	return []byte(`
{
    "code": 200,
    "message": "success",
    "result": {
        "acceptUrl": null,
        "accountName": "uname",
        "accountNumber": "4831069263065402",
        "amount": 1000,
        "bankAccountId": null,
        "callbackUrl": "https://mp9097f0b0590e076c79.free.beeceptor.com/callback",
        "cancelUrl": null,
        "code": null,
        "commissionId": null,
        "data": null,
        "date": "2025-12-03 12:19:12",
        "declineUrl": null,
        "description": "Orange Money Cashin Pay",
        "email": "support@emai.com",
        "errorCode": null,
        "errorMessage": null,
        "errorType": null,
        "externalId": "69267a601a761d64cd6f6f2f",
        "financialCommission": 1,
        "financialFees": 0,
        "financialId": "000000033763",
        "ipAddress": "169.254.169.126",
        "isWalletPayment": false,
        "noFees": false,
        "optionSlug": null,
        "paymentLink": null,
        "paymentServiceFeature": null,
        "paymentWallet": null,
        "phone": "673978334",
        "providerFees": 0,
        "providerId": null,
        "reference": null,
        "referenceNumber": "699999999",
        "requestId": "1850489561706069",
        "requestStatus": "PAYED",
        "rollbackId": null,
        "service": "orange-money-cashin-service-feature",
        "serviceName": null,
        "signature": null,
        "status": "PROGRESS",
        "terminalId": 1831267794013435,
        "terminalName": "uname (api)",
        "terminalUserAgent": "Go-http-client/1.1",
        "transactionId": 1850489545302716,
        "type": "cashin",
        "username": "uname",
        "voucher": {
            "payToken": "CI3579286593S51037",
            "providerTime": null,
            "value": "ProviderId, "
        }
    }
}`)
}

// OrangeMoneyCashinErrorResponse returns a stubbed error response for Orange Money Cashin payment
func OrangeMoneyCashinErrorResponse() []byte {
	return []byte(`
{
    "code": 5001,
    "message": "Transaction 69267a601a761d64cd6f6f2f has been rejected because of the difference of amounts: 500, 1000",
    "result": null
}`)
}
