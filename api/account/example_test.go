// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package account_test

import (
	"context"
	"fmt"
	"reflect"

	"github.com/mellis/ynab.go"
	"github.com/mellis/ynab.go/api"
)

func ExampleService_GetAccount() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	account, _ := c.Account().GetAccount(context.Background(), "<valid_budget_id>", "<valid_account_id>")
	fmt.Println(reflect.TypeOf(account))

	// Output: *account.Account
}

func ExampleService_GetAccounts() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	f := &api.Filter{LastKnowledgeOfServer: 10}
	snapshot, _ := c.Account().GetAccounts(context.Background(), "<valid_budget_id>", f)
	fmt.Println(reflect.TypeOf(snapshot))

	// Output: *account.SearchResultSnapshot
}
