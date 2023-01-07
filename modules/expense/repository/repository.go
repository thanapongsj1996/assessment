package repository

import "github.com/thanapongsj1996/assessment/modules/expense/model"

type ExpenseRepository interface {
	SaveExpense(expense model.Expense) (*model.Expense, error)
	GetExpenseByID(id int) (*model.Expense, error)
	UpdateExpense(id int, expense model.Expense) (*model.Expense, error)
	GetAllExpenses() (*[]model.Expense, error)
}
