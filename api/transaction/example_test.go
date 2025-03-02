// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package transaction_test

import (
	"context"
	"fmt"
	"reflect"

	"github.com/mellis/ynab.go"
	"github.com/mellis/ynab.go/api"
	"github.com/mellis/ynab.go/api/transaction"
)

func ExampleService_CreateTransaction() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	p := transaction.PayloadTransaction{
		AccountID: "<valid_account_id>",
		// ...
	}
	tx, _ := c.Transaction().CreateTransaction(context.Background(), "<valid_budget_id>", p)
	fmt.Println(reflect.TypeOf(tx))

	// Output: *transaction.OperationSummary
}

func ExampleService_CreateTransactions() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	p := []transaction.PayloadTransaction{
		{
			AccountID: "<valid_account_id>",
			// ...
		},
	}
	tx, _ := c.Transaction().CreateTransactions(context.Background(), "<valid_budget_id>", p)
	fmt.Println(reflect.TypeOf(tx))

	// Output: *transaction.OperationSummary
}

func ExampleService_BulkCreateTransactions() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	p := []transaction.PayloadTransaction{
		{
			AccountID: "<valid_account_id>",
			// ...
		},
		{
			AccountID: "<another_valid_account_id>",
			// ...
		},
	}
	bulk, _ := c.Transaction().BulkCreateTransactions(context.Background(), "<valid_budget_id>", p)
	fmt.Println(reflect.TypeOf(bulk))

	// Output: *transaction.Bulk
}

func ExampleService_UpdateTransaction() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	p := transaction.PayloadTransaction{
		AccountID: "<valid_account_id>",
		// ...
	}
	tx, _ := c.Transaction().UpdateTransaction(context.Background(), "<valid_budget_id>",
		"<valid_transaction_id>", p)
	fmt.Println(reflect.TypeOf(tx))

	// Output: *transaction.Transaction
}

func ExampleService_GetTransaction() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	tx, _ := c.Transaction().GetTransaction(context.Background(), "<valid_budget_id>",
		"<valid_transaction_id>")
	fmt.Println(reflect.TypeOf(tx))

	// Output: *transaction.Transaction
}

func ExampleService_GetTransactions() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	transactions, _ := c.Transaction().GetTransactions(context.Background(), "<valid_budget_id>", nil)
	fmt.Println(reflect.TypeOf(transactions))

	// Output: []*transaction.Transaction
}

func ExampleService_GetTransactions_filtered() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	date, _ := api.DateFromString("2010-09-09")
	f := &transaction.Filter{
		Since: &date,
		Type:  transaction.StatusUnapproved.Pointer(),
	}
	transactions, _ := c.Transaction().GetTransactions(context.Background(), "<valid_budget_id>", f)
	fmt.Println(reflect.TypeOf(transactions))

	// Output: []*transaction.Transaction
}

func ExampleService_GetTransactionsByAccount() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	transactions, _ := c.Transaction().GetTransactionsByAccount(
		context.Background(), "<valid_budget_id>", "<valid_account_id>", nil)
	fmt.Println(reflect.TypeOf(transactions))

	// Output: []*transaction.Transaction
}

func ExampleService_GetTransactionsByAccount_filtered() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	date, _ := api.DateFromString("2010-09-09")
	f := &transaction.Filter{
		Since: &date,
		Type:  transaction.StatusUnapproved.Pointer(),
	}
	transactions, _ := c.Transaction().GetTransactionsByAccount(
		context.Background(), "<valid_budget_id>", "<valid_account_id>", f)
	fmt.Println(reflect.TypeOf(transactions))

	// Output: []*transaction.Transaction
}

func ExampleService_GetTransactionsByCategory() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	transactions, _ := c.Transaction().GetTransactionsByCategory(
		context.Background(), "<valid_budget_id>", "<valid_category_id>", nil)
	fmt.Println(reflect.TypeOf(transactions))

	// Output: []*transaction.Hybrid
}

func ExampleService_GetTransactionsByCategory_filtered() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	date, _ := api.DateFromString("2010-09-09")
	f := &transaction.Filter{
		Since: &date,
		Type:  transaction.StatusUnapproved.Pointer(),
	}
	transactions, _ := c.Transaction().GetTransactionsByCategory(
		context.Background(), "<valid_budget_id>", "<valid_category_id>", f)
	fmt.Println(reflect.TypeOf(transactions))

	// Output: []*transaction.Hybrid
}

func ExampleService_GetTransactionsByPayee() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	transactions, _ := c.Transaction().GetTransactionsByPayee(
		context.Background(), "<valid_budget_id>", "<valid_payee_id>", nil)
	fmt.Println(reflect.TypeOf(transactions))

	// Output: []*transaction.Hybrid
}

func ExampleService_GetTransactionsByPayee_filtered() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	date, _ := api.DateFromString("2010-09-09")
	f := &transaction.Filter{
		Since: &date,
		Type:  transaction.StatusUnapproved.Pointer(),
	}
	transactions, _ := c.Transaction().GetTransactionsByPayee(
		context.Background(), "<valid_budget_id>", "<valid_payee_id>", f)
	fmt.Println(reflect.TypeOf(transactions))

	// Output: []*transaction.Hybrid
}

func ExampleService_GetScheduledTransaction() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	tx, _ := c.Transaction().GetScheduledTransaction(
		context.Background(),
		"<valid_budget_id>",
		"<valid_scheduled_transaction_id>",
	)
	fmt.Println(reflect.TypeOf(tx))

	// Output: *transaction.Scheduled
}

func ExampleService_GetScheduledTransactions() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	transactions, _ := c.Transaction().GetScheduledTransactions(context.Background(), "<valid_budget_id>")
	fmt.Println(reflect.TypeOf(transactions))

	// Output: []*transaction.Scheduled
}
