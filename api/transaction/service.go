// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package transaction

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mellis/ynab.go/api"
)

// NewService facilitates the creation of a new transaction service instance
func NewService(c api.ClientReaderWriter) *Service {
	return &Service{c}
}

// Service wraps YNAB transaction API endpoints
type Service struct {
	c api.ClientReaderWriter
}

// GetTransactions fetches the list of transactions from
// a budget with filtering capabilities
// https://api.youneedabudget.com/v1#/Transactions/getTransactions
func (s *Service) GetTransactions(ctx context.Context, budgetID string, f *Filter) ([]*Transaction, uint64, error) {
	resModel := struct {
		Data struct {
			Transactions    []*Transaction `json:"transactions"`
			ServerKnowledge uint64         `json:"server_knowledge"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/transactions", budgetID)
	if f != nil {
		url = fmt.Sprintf("%s?%s", url, f.ToQuery())
	}

	if err := s.c.Get(ctx, url, &resModel); err != nil {
		return nil, 0, err
	}

	return resModel.Data.Transactions, resModel.Data.ServerKnowledge, nil
}

// GetTransaction fetches a specific transaction from a budget
// https://api.youneedabudget.com/v1#/Transactions/getTransactionsById
func (s *Service) GetTransaction(ctx context.Context, budgetID string, transactionID string) (*Transaction, error) {
	resModel := struct {
		Data struct {
			Transaction *Transaction `json:"transaction"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/transactions/%s", budgetID, transactionID)
	if err := s.c.Get(ctx, url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.Transaction, nil
}

// CreateTransaction creates a new transaction for a budget
// https://api.youneedabudget.com/v1#/Transactions/createTransaction
func (s *Service) CreateTransaction(ctx context.Context, budgetID string,
	p PayloadTransaction) (*OperationSummary, error) {

	return s.CreateTransactions(ctx, budgetID, []PayloadTransaction{p})
}

// CreateTransactions creates one or more new transactions for a budget
// https://api.youneedabudget.com/v1#/Transactions/createTransaction
func (s *Service) CreateTransactions(ctx context.Context, budgetID string,
	p []PayloadTransaction) (*OperationSummary, error) {

	payload := struct {
		Transactions []PayloadTransaction `json:"transactions"`
	}{
		p,
	}

	buf, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}

	resModel := struct {
		Data *OperationSummary `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/transactions", budgetID)
	err = s.c.Post(ctx, url, &resModel, buf)
	if err != nil {
		return nil, err
	}
	return resModel.Data, nil
}

// BulkCreateTransactions creates multiple transactions for a budget
// https://api.youneedabudget.com/v1#/Transactions/bulkCreateTransactions
// Deprecated: Use transaction.CreateTransactions instead.
func (s *Service) BulkCreateTransactions(ctx context.Context, budgetID string,
	ps []PayloadTransaction) (*Bulk, error) {

	payload := struct {
		Transactions []PayloadTransaction `json:"transactions"`
	}{
		ps,
	}

	buf, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}

	resModel := struct {
		Data struct {
			Bulk *Bulk `json:"bulk"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/transactions/bulk", budgetID)
	if err := s.c.Post(ctx, url, &resModel, buf); err != nil {
		return nil, err
	}
	return resModel.Data.Bulk, nil
}

// UpdateTransaction updates a whole transaction for a replacement
// https://api.youneedabudget.com/v1#/Transactions/updateTransaction
func (s *Service) UpdateTransaction(ctx context.Context, budgetID string, transactionID string,
	p PayloadTransaction) (*Transaction, error) {

	payload := struct {
		Transaction *PayloadTransaction `json:"transaction"`
	}{
		&p,
	}

	buf, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}

	resModel := struct {
		Data struct {
			Transaction *Transaction `json:"transaction"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/transactions/%s", budgetID, transactionID)
	if err := s.c.Put(ctx, url, &resModel, buf); err != nil {
		return nil, err
	}
	return resModel.Data.Transaction, nil
}

// UpdateTransactions creates one or more new transactions for a budget
// https://api.youneedabudget.com/v1#/Transactions/updateTransactions
func (s *Service) UpdateTransactions(ctx context.Context, budgetID string,
	p []PayloadTransaction) (*OperationSummary, error) {

	payload := struct {
		Transactions []PayloadTransaction `json:"transactions"`
	}{
		p,
	}

	buf, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}

	resModel := struct {
		Data *OperationSummary `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/transactions", budgetID)
	err = s.c.Patch(ctx, url, &resModel, buf)
	if err != nil {
		return nil, err
	}
	return resModel.Data, nil
}

// GetTransactionsByAccount fetches the list of transactions of a specific account
// from a budget with filtering capabilities
// https://api.youneedabudget.com/v1#/Transactions/getTransactionsByAccount
func (s *Service) GetTransactionsByAccount(ctx context.Context, budgetID string, accountID string,
	f *Filter) ([]*Transaction, error) {

	resModel := struct {
		Data struct {
			Transactions []*Transaction `json:"transactions"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/accounts/%s/transactions", budgetID, accountID)
	if f != nil {
		url = fmt.Sprintf("%s?%s", url, f.ToQuery())
	}

	if err := s.c.Get(ctx, url, &resModel); err != nil {
		return nil, err
	}

	return resModel.Data.Transactions, nil
}

// GetTransactionsByCategory fetches the list of transactions of a specific category
// from a budget with filtering capabilities
// https://api.youneedabudget.com/v1#/Transactions/getTransactionsByCategory
func (s *Service) GetTransactionsByCategory(ctx context.Context, budgetID string, categoryID string,
	f *Filter) ([]*Hybrid, error) {

	resModel := struct {
		Data struct {
			Transactions []*Hybrid `json:"transactions"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/categories/%s/transactions", budgetID, categoryID)
	if f != nil {
		url = fmt.Sprintf("%s?%s", url, f.ToQuery())
	}

	if err := s.c.Get(ctx, url, &resModel); err != nil {
		return nil, err
	}

	return resModel.Data.Transactions, nil
}

// GetTransactionsByPayee fetches the list of transactions of a specific payee
// from a budget with filtering capabilities
// https://api.youneedabudget.com/v1#/Transactions/getTransactionsByPayee
func (s *Service) GetTransactionsByPayee(ctx context.Context, budgetID string, payeeID string,
	f *Filter) ([]*Hybrid, error) {

	resModel := struct {
		Data struct {
			Transactions []*Hybrid `json:"transactions"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/payees/%s/transactions", budgetID, payeeID)
	if f != nil {
		url = fmt.Sprintf("%s?%s", url, f.ToQuery())
	}

	if err := s.c.Get(ctx, url, &resModel); err != nil {
		return nil, err
	}

	return resModel.Data.Transactions, nil
}

// GetScheduledTransactions fetches the list of scheduled transactions from
// a budget
// https://api.youneedabudget.com/v1#/Scheduled_Transactions/getScheduledTransactions
func (s *Service) GetScheduledTransactions(ctx context.Context, budgetID string) ([]*Scheduled, error) {
	resModel := struct {
		Data struct {
			ScheduledTransactions []*Scheduled `json:"scheduled_transactions"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/scheduled_transactions", budgetID)
	if err := s.c.Get(ctx, url, &resModel); err != nil {
		return nil, err
	}

	return resModel.Data.ScheduledTransactions, nil
}

// GetScheduledTransaction fetches a specific scheduled transaction from a budget
// https://api.youneedabudget.com/v1#/Scheduled_Transactions/getScheduledTransactionById
func (s *Service) GetScheduledTransaction(ctx context.Context, budgetID, scheduledTransactionID string) (*Scheduled, error) {
	resModel := struct {
		Data struct {
			ScheduledTransactions *Scheduled `json:"scheduled_transaction"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/scheduled_transactions/%s", budgetID, scheduledTransactionID)
	if err := s.c.Get(ctx, url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.ScheduledTransactions, nil
}

// Filter represents the optional filter while fetching transactions
type Filter struct {
	Since *api.Date
	Type  *Status
	// LastKnowledgeOfServer The starting server knowledge. If provided,
	// only entities that have changed since last_knowledge_of_server
	// will be included
	LastKnowledgeOfServer *uint64
}

// ToQuery returns the filters as a HTTP query string
func (f *Filter) ToQuery() string {
	pairs := make([]string, 0, 2)
	if f.Since != nil && !f.Since.IsZero() {
		pairs = append(pairs, fmt.Sprintf("since_date=%s",
			api.DateFormat(*f.Since)))
	}
	if f.Type != nil {
		pairs = append(pairs, fmt.Sprintf("type=%s", string(*f.Type)))
	}
	if f.LastKnowledgeOfServer != nil {
		pairs = append(pairs, fmt.Sprintf("last_knowledge_of_server=%d", *f.LastKnowledgeOfServer))
	}
	return strings.Join(pairs, "&")
}
