// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package payee_test

import (
	"context"
	"fmt"
	"reflect"

	"github.com/mellis/ynab.go/api"

	"github.com/mellis/ynab.go"
)

func ExampleService_GetPayee() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	p, _ := c.Payee().GetPayee(context.Background(), "<valid_budget_id>", "<valid_payee_id>")
	fmt.Println(reflect.TypeOf(p))

	// Output: *payee.Payee
}

func ExampleService_GetPayees() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	f := &api.Filter{LastKnowledgeOfServer: 10}
	payees, _ := c.Payee().GetPayees(context.Background(), "<valid_budget_id>", f)
	fmt.Println(reflect.TypeOf(payees))

	// Output: *payee.SearchResultSnapshot
}

func ExampleService_GetPayeeLocation() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	l, _ := c.Payee().GetPayeeLocation(context.Background(), "<valid_budget_id>", "<valid_payee_location_id>")
	fmt.Println(reflect.TypeOf(l))

	// Output: *payee.Location
}

func ExampleService_GetPayeeLocations() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	locations, _ := c.Payee().GetPayeeLocations(context.Background(), "<valid_budget_id>")
	fmt.Println(reflect.TypeOf(locations))

	// Output: []*payee.Location
}

func ExampleService_GetPayeeLocationsByPayee() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	locations, _ := c.Payee().GetPayeeLocationsByPayee(context.Background(), "<valid_budget_id>", "<valid_payee_id>")
	fmt.Println(reflect.TypeOf(locations))

	// Output: []*payee.Location
}
