# YNAB API Go Library

[![Go Report Card](https://goreportcard.com/badge/github.com/mellis/ynab.go)](https://goreportcard.com/report/github.com/mellis/ynab.go) [![GoDoc Reference](https://godoc.org/github.com/mellis/ynab.go?status.svg)](https://godoc.org/github.com/mellis/ynab.go)

This is an UNOFFICIAL Go client for the YNAB API. It covers 100% of the resources made available by the [YNAB API](https://api.youneedabudget.com).

## Installation

```
go get github.com/mellis/ynab.go
```

## Usage

To use this client you must [obtain an access token](https://api.youneedabudget.com/#authentication-overview) from your [My Account](https://app.youneedabudget.com/settings) page of the YNAB web app.

```go
package main

import (
	"fmt"

	"github.com/mellis/ynab.go"
)

const accessToken = "bf0cbb14b4330-not-real-3de12e66a389eaafe2"

func main() {
	c := ynab.NewClient(accessToken)
	budgets, err := c.Budget().GetBudgets()
	if err != nil {
		panic(err)
	}

	for _, budget := range budgets {
		fmt.Println(budget.Name)
		// ...
	}
}
```

See the [godoc](https://godoc.org/github.com/mellis/ynab.go) to see all the available methods with example usage.

## Development

- Make sure you have Go 1.19 or later installed
- Run tests with `go test -race ./...`

## License

BSD-2-Clause
