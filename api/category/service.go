// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package category

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mellis/ynab.go/api"
)

const currentMonthID = "current"

// NewService facilitates the creation of a new category service instance
func NewService(c api.ClientReaderWriter) *Service {
	return &Service{c}
}

// Service wraps YNAB category API endpoints
type Service struct {
	c api.ClientReaderWriter
}

// GetCategories fetches the list of category groups for a budget
// https://api.youneedabudget.com/v1#/Categories/getCategories
func (s *Service) GetCategories(ctx context.Context, budgetID string, f *api.Filter) (*SearchResultSnapshot, error) {
	resModel := struct {
		Data struct {
			CategoryGroups  []*GroupWithCategories `json:"category_groups"`
			ServerKnowledge uint64                 `json:"server_knowledge"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/categories", budgetID)
	if f != nil {
		url = fmt.Sprintf("%s?%s", url, f.ToQuery())
	}
	if err := s.c.Get(ctx, url, &resModel); err != nil {
		return nil, err
	}

	return &SearchResultSnapshot{
		GroupWithCategories: resModel.Data.CategoryGroups,
		ServerKnowledge:     resModel.Data.ServerKnowledge,
	}, nil
}

// GetCategory fetches a specific category from a budget
// https://api.youneedabudget.com/v1#/Categories/getCategoryById
func (s *Service) GetCategory(ctx context.Context, budgetID string, categoryID string) (*Category, error) {
	resModel := struct {
		Data struct {
			Category *Category `json:"category"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/categories/%s", budgetID, categoryID)
	if err := s.c.Get(ctx, url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.Category, nil
}

// GetCategoryForMonth fetches a specific category from a budget month
// https://api.youneedabudget.com/v1#/Categories/getMonthCategoryById
func (s *Service) GetCategoryForMonth(ctx context.Context, budgetID string, categoryID string,
	month api.Date) (*Category, error) {

	return s.getCategoryForMonth(ctx, budgetID, categoryID, api.DateFormat(month))
}

// GetCategoryForCurrentMonth fetches a specific category from the current budget month
// https://api.youneedabudget.com/v1#/Categories/getMonthCategoryById
func (s *Service) GetCategoryForCurrentMonth(ctx context.Context, budgetID string, categoryID string) (*Category, error) {
	return s.getCategoryForMonth(ctx, budgetID, categoryID, currentMonthID)
}

func (s *Service) getCategoryForMonth(ctx context.Context, budgetID string, categoryID, month string) (*Category, error) {
	resModel := struct {
		Data struct {
			Category *Category `json:"category"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/months/%s/categories/%s", budgetID, month, categoryID)
	if err := s.c.Get(ctx, url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.Category, nil
}

// UpdateCategoryForMonth updates a category for a month
// https://api.youneedabudget.com/v1#/Categories/updateMonthCategory
func (s *Service) UpdateCategoryForMonth(ctx context.Context, budgetID string, categoryID string, month api.Date,
	p PayloadMonthCategory) (*Category, error) {

	return s.updateCategoryForMonth(ctx, budgetID, categoryID, api.DateFormat(month), p)
}

// UpdateCategoryForCurrentMonth updates a category for the current month
// https://api.youneedabudget.com/v1#/Categories/updateMonthCategory
func (s *Service) UpdateCategoryForCurrentMonth(ctx context.Context, budgetID string, categoryID string,
	p PayloadMonthCategory) (*Category, error) {

	return s.updateCategoryForMonth(ctx, budgetID, categoryID, currentMonthID, p)
}

func (s *Service) updateCategoryForMonth(ctx context.Context, budgetID string, categoryID string, month string,
	p PayloadMonthCategory) (*Category, error) {

	payload := struct {
		MonthCategory *PayloadMonthCategory `json:"month_category"`
	}{
		&p,
	}

	buf, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}

	resModel := struct {
		Data struct {
			Category *Category `json:"category"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/months/%s/categories/%s", budgetID,
		month, categoryID)

	if err := s.c.Put(ctx, url, &resModel, buf); err != nil {
		return nil, err
	}
	return resModel.Data.Category, nil
}
