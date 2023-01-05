package service

import (
	"github.com/jinzhu/copier"
	"github.com/thanapongsj1996/assessment/modules/expense/dto"
	"github.com/thanapongsj1996/assessment/modules/expense/model"
	"github.com/thanapongsj1996/assessment/modules/expense/repository"
)

type expenseService struct {
	expenseRepo repository.ExpenseRepository
}

func NewExpenseService(repo repository.ExpenseRepository) ExpenseService {
	return expenseService{expenseRepo: repo}
}

func (s expenseService) AddExpense(req dto.AddExpenseReq) (*dto.AddExpenseRes, error) {
	expense := model.Expense{}
	copier.Copy(&expense, &req)

	result, err := s.expenseRepo.SaveExpense(expense)
	if err != nil {
		return nil, err
	}

	response := dto.AddExpenseRes{}
	copier.Copy(&response, result)

	return &response, nil
}
