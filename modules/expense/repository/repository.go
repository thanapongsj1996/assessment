package repository

import "github.com/thanapongsj1996/assessment/modules/expense/model"

type ExpenseRepository interface {
	SaveExpense(expense model.Expense) (*model.Expense, error)
}
