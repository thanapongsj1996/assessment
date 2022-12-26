package service

import "github.com/thanapongsj1996/assessment/modules/expense/repository"

type expenseService struct {
	expenseRepo repository.ExpenseRepository
}

func NewExpenseService(repo repository.ExpenseRepository) ExpenseService {
	return expenseService{expenseRepo: repo}
}
