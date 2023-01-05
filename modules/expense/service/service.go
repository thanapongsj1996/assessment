package service

import "github.com/thanapongsj1996/assessment/modules/expense/dto"

type ExpenseService interface {
	AddExpense(req dto.AddExpenseReq) (*dto.AddExpenseRes, error)
}
