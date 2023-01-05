package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/thanapongsj1996/assessment/modules/expense/dto"
	"github.com/thanapongsj1996/assessment/modules/expense/model"
	"testing"
)

type mockExpenseRepository struct{}

func (r mockExpenseRepository) SaveExpense(expense model.Expense) (*model.Expense, error) {
	return &model.Expense{
		ID:     1,
		Title:  "strawberry smoothie",
		Amount: 98,
		Note:   "note",
		Tags:   []string{"food", "beverage"},
	}, nil
}

func TestServiceAddExpenseSuccess(t *testing.T) {
	mockRepo := mockExpenseRepository{}
	service := NewExpenseService(mockRepo)

	addExpenseReq := dto.AddExpenseReq{
		Title:  "strawberry smoothie",
		Amount: 98,
		Note:   "note",
		Tags:   []string{"food", "beverage"},
	}
	want := &dto.AddExpenseRes{
		ID:     1,
		Title:  "strawberry smoothie",
		Amount: 98,
		Note:   "note",
		Tags:   []string{"food", "beverage"},
	}

	got, _ := service.AddExpense(addExpenseReq)
	assert.Equal(t, want, got)
}
