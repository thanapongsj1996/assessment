//go:build unit

package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thanapongsj1996/assessment/modules/expense/dto"
	"github.com/thanapongsj1996/assessment/modules/expense/model"
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
func (r mockExpenseRepository) GetExpenseByID(id int) (*model.Expense, error) {
	return &model.Expense{
		ID:     id,
		Title:  "strawberry smoothie",
		Amount: 98,
		Note:   "note",
		Tags:   []string{"food", "beverage"},
	}, nil
}
func (r mockExpenseRepository) UpdateExpense(id int, expense model.Expense) (*model.Expense, error) {
	return &model.Expense{
		ID:     1,
		Title:  "strawberry smoothie",
		Amount: 98,
		Note:   "updated note",
		Tags:   []string{"food", "beverage"},
	}, nil
}
func (r mockExpenseRepository) GetAllExpenses() (*[]model.Expense, error) {
	return &[]model.Expense{
		{
			ID:     1,
			Title:  "strawberry smoothie",
			Amount: 98,
			Note:   "note",
			Tags:   []string{"food", "beverage"},
		},
		{
			ID:     2,
			Title:  "apple smoothie",
			Amount: 89,
			Note:   "no discount",
			Tags:   []string{"beverage"},
		},
	}, nil
}

func TestServiceAddExpenseSuccess(t *testing.T) {
	mockRepo := mockExpenseRepository{}
	service := NewExpenseService(mockRepo)

	addExpenseReq := dto.ExpenseReq{
		Title:  "strawberry smoothie",
		Amount: 98,
		Note:   "note",
		Tags:   []string{"food", "beverage"},
	}
	want := &dto.ExpenseRes{
		ID:     1,
		Title:  "strawberry smoothie",
		Amount: 98,
		Note:   "note",
		Tags:   []string{"food", "beverage"},
	}

	got, _ := service.AddExpense(addExpenseReq)
	assert.Equal(t, want, got)
}

func TestServiceGetExpenseSuccess(t *testing.T) {
	mockRepo := mockExpenseRepository{}
	service := NewExpenseService(mockRepo)

	want := &dto.ExpenseRes{
		ID:     1,
		Title:  "strawberry smoothie",
		Amount: 98,
		Note:   "note",
		Tags:   []string{"food", "beverage"},
	}

	got, err := service.GetExpenseByID(1)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expect %v got %v", want, got)
	}
	if !errors.Is(err, nil) {
		t.Errorf("expect %v got %v", nil, err)
	}
}

func TestServiceUpdateExpenseSuccess(t *testing.T) {
	mockRepo := mockExpenseRepository{}
	service := NewExpenseService(mockRepo)

	updateExpenseReq := dto.ExpenseReq{
		Title:  "strawberry smoothie",
		Amount: 98,
		Note:   "updated note",
		Tags:   []string{"food", "beverage"},
	}
	want := &dto.ExpenseRes{
		ID:     1,
		Title:  "strawberry smoothie",
		Amount: 98,
		Note:   "updated note",
		Tags:   []string{"food", "beverage"},
	}

	got, err := service.UpdateExpense(1, updateExpenseReq)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expect %v got %v", want, got)
	}
	if !errors.Is(err, nil) {
		t.Errorf("expect %v got %v", nil, err)
	}
}

func TestServiceGetAllExpensesSuccess(t *testing.T) {
	mockRepo := mockExpenseRepository{}
	service := NewExpenseService(mockRepo)

	want := &[]dto.ExpenseRes{
		{
			ID:     1,
			Title:  "strawberry smoothie",
			Amount: 98,
			Note:   "note",
			Tags:   []string{"food", "beverage"},
		},
		{
			ID:     2,
			Title:  "apple smoothie",
			Amount: 89,
			Note:   "no discount",
			Tags:   []string{"beverage"},
		},
	}

	got, err := service.GetAllExpenses()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expect %v got %v", want, got)
	}
	if !errors.Is(err, nil) {
		t.Errorf("expect %v got %v", nil, err)
	}
}
