package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thanapongsj1996/assessment/modules/expense/handler"
)

func NewExpenseRoute(e *echo.Echo, handler handler.ExpenseHandler) {
	routes := e.Group("expenses")

	routes.POST("", handler.AddExpense)
	routes.GET("/:id", handler.GetExpenseByID)
}
