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

func (r expenseRepository) GetExpenseByID(id int) (*model.Expense, error) {
	expense := model.Expense{}
	tx := r.db.Table(expenseTable).Where("id = ?", id).First(&expense)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &expense, nil
}

func (r expenseRepository) UpdateExpense(id int, expense model.Expense) (*model.Expense, error) {
	_, err := r.GetExpenseByID(id)
	if err != nil {
		return nil, err
	}

	expense.ID = id
	tx := r.db.Table(expenseTable).Save(&expense)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &expense, nil
}

func (r expenseRepository) GetAllExpenses() (*[]model.Expense, error) {
	var expense []model.Expense
	tx := r.db.Table(expenseTable).Find(&expense)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &expense, nil
}
