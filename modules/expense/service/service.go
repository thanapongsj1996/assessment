package service

import "github.com/thanapongsj1996/assessment/modules/expense/dto"

type ExpenseService interface {
	AddExpense(req dto.ExpenseReq) (*dto.ExpenseRes, error)
	GetExpenseByID(id int) (*dto.ExpenseRes, error)
	UpdateExpense(id int, req dto.ExpenseReq) (*dto.ExpenseRes, error)
}
