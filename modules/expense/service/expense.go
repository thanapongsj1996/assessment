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

func (s expenseService) AddExpense(req dto.ExpenseReq) (*dto.ExpenseRes, error) {
	expense := model.Expense{}
	copier.Copy(&expense, &req)

	result, err := s.expenseRepo.SaveExpense(expense)
	if err != nil {
		return nil, err
	}

	response := dto.ExpenseRes{}
	copier.Copy(&response, result)

	return &response, nil
}

func (s expenseService) GetExpenseByID(id int) (*dto.ExpenseRes, error) {
	result, err := s.expenseRepo.GetExpenseByID(id)
	if err != nil {
		return nil, err
	}

	response := dto.ExpenseRes{}
	copier.Copy(&response, result)

	return &response, nil
}

func (s expenseService) UpdateExpense(id int, req dto.ExpenseReq) (*dto.ExpenseRes, error) {
	expense := model.Expense{}
	copier.Copy(&expense, &req)

	result, err := s.expenseRepo.UpdateExpense(id, expense)
	if err != nil {
		return nil, err
	}

	response := dto.ExpenseRes{}
	copier.Copy(&response, result)

	return &response, nil
}

func (s expenseService) GetAllExpenses() (*[]dto.ExpenseRes, error) {
	result, err := s.expenseRepo.GetAllExpenses()
	if err != nil {
		return nil, err
	}

	var response []dto.ExpenseRes
	copier.Copy(&response, result)

	return &response, nil
}
