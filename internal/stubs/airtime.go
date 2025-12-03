package stubs

// AirtimePaymentResponse returns a stubbed response for Airtime payment
func AirtimePaymentResponse() []byte {
	return []byte(`
{
    "code": 200,
    "message": "success",
    "result": {
        "acceptNotificationUrl": null,
        "acceptUrl": null,
        "accountName": "uname",
        "accountNumber": "4831069263065402",
        "adminValidaterUsername": null,
        "amount": 1000,
        "bankAccountId": null,
        "callbackUrl": "https://mp9097f0b0590e076c79.free.beeceptor.com/callback",
        "cancelNotificationUrl": null,
        "cancelUrl": null,
        "code": null,
        "commissionId": "000000033765",
        "data": null,
        "date": "2025-12-03 13:19:44",
        "declineNotificationUrl": null,
        "declineUrl": null,
        "description": "Orange Airtime Purchase",
        "email": "mail@domain.com",
        "errorCode": null,
        "errorMessage": null,
        "errorType": null,
        "externalId": "224169cd-caa6-46d3-8262-eb95adb6b1d9",
        "financialCommission": 50,
        "financialFees": 0,
        "financialId": "000000033764",
        "ipAddress": "169.254.169.126",
        "isWalletPayment": false,
        "noFees": null,
        "notificationUrl": null,
        "optionSlug": null,
        "paymentLink": null,
        "paymentServiceFeature": null,
        "paymentWallet": null,
        "phone": "237689120974",
        "providerFees": 0,
        "providerId": "R910452.6263.2358",
        "reason": null,
        "reference": null,
        "referenceNumber": "659683157",
        "requestId": "1850493358573419",
        "requestStatus": "SUCCESS",
        "rollbackId": null,
        "service": "orange-airtime-service-feature",
        "serviceName": null,
        "signature": null,
        "status": "SUCCESS",
        "terminalId": 1831267794013435,
        "terminalName": "uname (api)",
        "terminalUserAgent": "Go-http-client/1.1",
        "transactionId": 1850493354204578,
        "type": "payment",
        "username": "uname",
        "voucher": {
            "value": "659683157, 1000"
        }
    }
}
`)
}
