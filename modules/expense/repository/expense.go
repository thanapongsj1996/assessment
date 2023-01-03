package repository

import (
	"github.com/thanapongsj1996/assessment/modules/expense/model"
	"gorm.io/gorm"
)

const expenseTable = "expenses"

type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return expenseRepository{db: db}
}

func (r expenseRepository) SaveExpense(expense model.Expense) (*model.Expense, error) {
	tx := r.db.Table(expenseTable).Create(&expense)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &expense, nil
}
