package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/thanapongsj1996/assessment/modules/expense/dto"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockExpenseService struct{}

func (s mockExpenseService) AddExpense(req dto.AddExpenseReq) (*dto.AddExpenseRes, error) {
	return &dto.AddExpenseRes{
		ID:     1,
		Title:  "strawberry smoothie",
		Amount: 98,
		Note:   "note",
		Tags:   []string{"food", "beverage"},
	}, nil
}

func TestAddExpenseSuccess(t *testing.T) {
	addExpenseRequestJsonBody := `{"title":"strawberry smoothie","amount":98,"note":"note","tags":["food","beverage"]}`
	expectResponseBody := `{"id":1,"title":"strawberry smoothie","amount":98,"note":"note","tags":["food","beverage"]}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(addExpenseRequestJsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockService := mockExpenseService{}
	h := NewExpenseHandler(mockService)

	// Assertions
	if assert.NoError(t, h.AddExpense(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expectResponseBody, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestAddExpenseBadRequest(t *testing.T) {
	addExpenseRequestJsonBody := `{"title":"strawberry smoothie","amount":98,"note":"qqq","tags":"food"}`
	expectResponseBody := "\"Can not bind data\""

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(addExpenseRequestJsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockService := mockExpenseService{}
	h := NewExpenseHandler(mockService)

	// Assertions
	if assert.NoError(t, h.AddExpense(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, expectResponseBody, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}
