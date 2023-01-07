package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/thanapongsj1996/assessment/modules/expense/dto"
	"github.com/thanapongsj1996/assessment/modules/expense/service"
	"net/http"
	"strconv"
)

type ExpenseHandler struct {
	expenseService service.ExpenseService
}

func NewExpenseHandler(service service.ExpenseService) ExpenseHandler {
	return ExpenseHandler{expenseService: service}
}

func (h *ExpenseHandler) AddExpense(c echo.Context) error {
	addExpenseReq := dto.ExpenseReq{}
	if err := c.Bind(&addExpenseReq); err != nil {
		return c.JSON(http.StatusBadRequest, "Can not bind data")
	}

	result, err := h.expenseService.AddExpense(addExpenseReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *ExpenseHandler) GetExpenseByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Can not bind data")
	}

	result, err := h.expenseService.GetExpenseByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
