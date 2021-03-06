# afrikpay-go

[![Build](https://github.com/NdoleStudio/afrikpay-go/actions/workflows/main.yml/badge.svg)](https://github.com/NdoleStudio/afrikpay-go/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/NdoleStudio/afrikpay-go/branch/main/graph/badge.svg)](https://codecov.io/gh/NdoleStudio/afrikpay-go)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/NdoleStudio/afrikpay-go/badges/quality-score.png?b=main)](https://scrutinizer-ci.com/g/NdoleStudio/afrikpay-go/?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/NdoleStudio/afrikpay-go)](https://goreportcard.com/report/github.com/NdoleStudio/afrikpay-go)
[![GitHub contributors](https://img.shields.io/github/contributors/NdoleStudio/afrikpay-go)](https://github.com/NdoleStudio/afrikpay-go/graphs/contributors)
[![GitHub license](https://img.shields.io/github/license/NdoleStudio/afrikpay-go?color=brightgreen)](https://github.com/NdoleStudio/afrikpay-go/blob/master/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/NdoleStudio/afrikpay-go)](https://pkg.go.dev/github.com/NdoleStudio/afrikpay-go)


This package provides a Go client for the AfrikPay HTTP API https://developer.afrikpay.com

## Installation

`afrikpay-go` is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/NdoleStudio/afrikpay-go
```

Alternatively the same can be achieved if you use `import` in a package:

```go
import "github.com/NdoleStudio/afrikpay-go"
```


## Implemented

- [Airtime](#airtime)
  - `POST /api/airtime/v2/`: Transfer airtime
  - `POST /api/airtime/status/v2/`: Airtime status
- [Account](#account)
  - `POST /api/account/agent/balance/v2/`: Account Balance
- [Bill](#bill)
  - `POST /api/bill/v2/`: Pay bills or subscriptions
  - `POST /api/bill/status/v2/`: Bill transaction status
  - `POST /api/bill/getamount/`: Get bill amount

## Usage

### Initializing the Client

An instance of the client can be created using `New()`.

```go
package main

import (
	"github.com/NdoleStudio/afrikpay-go"
)

func main()  {
  client := afrikpay.New(
    afrikpay.WithAgentID(""/* agent id */),
    afrikpay.WithAPIKey(""/* api key */),
    afrikpay.WithAgentPassword(""/* agent password */),
    afrikpay.WithAgentPlatform(""/* agent platform */),
  )
}
```

### Error handling

All API calls return an `error` as the last return object. All successful calls will return a `nil` error.

```go
status, response, err := afrikpay.Airtime.Transfer(context.Background())
if err != nil {
    //handle error
}
```

### Airtime

#### `POST /api/airtime/v2/`: Transfer airtime

Transfer is intended for communication / Internet credit transfer operations to telephone numbers.

```go
transaction, _, err := afrikpay.Airtime.Transfer(context.Background(), &AirtimeTransferParams{
    Operator:          "mtn",
    PurchaseReference: "test-ref",
    Amount:            "987",
    PhoneNumber:       "00000000",
    Mode:              "cash",
})

if err != nil {
    log.Fatal(err)
}

log.Println(status.Code) // 200
```

#### `POST /api/airtime/status/v2/`: Airtime Status

The Airtime Status API is intended for getting the status of a airtime transaction

```go
transaction, _, err := afrikpay.Airtime.Status(context.Background(), ""/* Transaction ID */)

if err != nil {
    log.Fatal(err)
}

log.Println(status.Code) // 200
```


### Account

#### `POST /api/account/agent/balance/v2/`: Account Balance

The Balance API is used for the partner to get the Balance of their account

```go
balance, _, err := afrikpay.Account.Balance(context.Background())
if err != nil {
    log.Fatal(err)
}

log.Println(status.Code) // 200
```

### Bill

#### `POST /api/bill/v2/`: Pay bills or subscriptions

The Bill API is intended for bill payment operations.

```go
transaction, _, err := client.Bill.Pay(context.Background(), BillPayParams{
    Biller:           BillerEneoPostpay,
    BillID:           billerID,
    Mode:             ModeCash
})
if err != nil {
    log.Fatal(err)
}

log.Println(transaction.Code) // 200
```

#### `POST /api/bill/status/v2/`: Bill transaction status

The Bill Status API is intended for getting the status of a bill transaction

```go
transaction, _ , err := client.Bill.Status(context.Background(), "transaction-id")
if err != nil {
    log.Fatal(err)
}

log.Println(transaction.Code) //200
```

#### `POST /api/bill/getamount/`: Get bill amount

The Bill Amount API is used to get the amount of a specific bill

```go
amount, _ , err := client.Bill.Amount(context.Background(), afrikpay.BillerEneoPostpay, "bill-number")
if err != nil {
    log.Fatal(err)
}

log.Println(transaction.Code) //200
```


## Testing

You can run the unit tests for this client from the root directory using the command below:

```bash
go test -v
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
