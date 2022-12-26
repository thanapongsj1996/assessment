package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/thanapongsj1996/assessment/modules/expense/service"
	"net/http"
)

type ExpenseHandler struct {
	expenseService service.ExpenseService
}

func NewExpenseHandler(service service.ExpenseService) ExpenseHandler {
	return ExpenseHandler{expenseService: service}
}

func (h *ExpenseHandler) Test(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"TEST": "OK",
	})
}
