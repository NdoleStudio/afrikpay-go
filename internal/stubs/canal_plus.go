package stubs

// CanalPlusOptionResponse returns a stubbed response for Canal Plus options
func CanalPlusOptionResponse() []byte {
	return []byte(`
{
  "code": 200,
  "message": "success",
  "result": [
    {
      "voucher": null,
      "optionId": 1803057796021245,
      "name": "Access + 1 mois 15000",
      "slug": "access-plus",
      "amount": 15000,
      "referenceNumber": null,
      "date": "2024-06-27 23:11:51"
    },
    {
      "voucher": null,
      "optionId": 1803057796025301,
      "name": "Tous Canal 1 mois 45000",
      "slug": "tout-canal",
      "amount": 45000,
      "referenceNumber": null,
      "date": "2024-06-27 23:11:51"
    },
    {
      "voucher": null,
      "optionId": 1803057796017468,
      "name": "Charme 7000",
      "slug": "charme",
      "amount": 7000,
      "referenceNumber": null,
      "date": "2024-06-27 23:11:51"
    },
    {
      "voucher": null,
      "optionId": 1803057796013696,
      "name": "Evasion + 1 mois 22500",
      "slug": "evasion-plus",
      "amount": 22500,
      "referenceNumber": null,
      "date": "2024-06-27 23:11:51"
    },
    {
      "voucher": null,
      "optionId": 1803057796009599,
      "name": "Evasion 1 mois 10000",
      "slug": "evasion",
      "amount": 10000,
      "referenceNumber": null,
      "date": "2024-06-27 23:11:51"
    },
    {
      "voucher": null,
      "optionId": 1803057796003231,
      "name": "Access 1 mois 5000",
      "slug": "access",
      "amount": 5000,
      "referenceNumber": null,
      "date": "2024-06-27 23:11:51"
    }
  ]
}
`)
}
