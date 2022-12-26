package repository

import "gorm.io/gorm"

type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return expenseRepository{db: db}
}
