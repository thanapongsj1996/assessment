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

func (s mockExpenseService) AddExpense(req dto.ExpenseReq) (*dto.ExpenseRes, error) {
	return &dto.ExpenseRes{
		ID:     1,
		Title:  "strawberry smoothie",
		Amount: 98,
		Note:   "note",
		Tags:   []string{"food", "beverage"},
	}, nil
}

func (s mockExpenseService) GetExpenseByID(id int) (*dto.ExpenseRes, error) {
	return &dto.ExpenseRes{
		ID:     1,
		Title:  "strawberry smoothie",
		Amount: 98,
		Note:   "note",
		Tags:   []string{"food", "beverage"},
	}, nil
}

func (s mockExpenseService) UpdateExpense(id int, req dto.ExpenseReq) (*dto.ExpenseRes, error) {
	return &dto.ExpenseRes{
		ID:     1,
		Title:  "strawberry smoothie",
		Amount: 98,
		Note:   "updated note",
		Tags:   []string{"food", "beverage"},
	}, nil
}

func (s mockExpenseService) GetAllExpenses() (*[]dto.ExpenseRes, error) {
	return &[]dto.ExpenseRes{
		{
			ID:     1,
			Title:  "strawberry smoothie",
			Amount: 98,
			Note:   "note",
			Tags:   []string{"food", "beverage"},
		},
		{
			ID:     2,
			Title:  "apple smoothie",
			Amount: 89,
			Note:   "no discount",
			Tags:   []string{"beverage"},
		},
	}, nil
}

func TestHandlerAddExpenseSuccess(t *testing.T) {
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

func TestHandlerAddExpenseBadRequest(t *testing.T) {
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

func TestHandlerGetExpenseSuccess(t *testing.T) {
	expectResponseBody := `{"id":1,"title":"strawberry smoothie","amount":98,"note":"note","tags":["food","beverage"]}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/expenses", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	mockService := mockExpenseService{}
	h := NewExpenseHandler(mockService)

	// Assertions
	if assert.NoError(t, h.GetExpenseByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectResponseBody, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestHandlerUpdateExpenseSuccess(t *testing.T) {
	updateExpenseRequestJsonBody := `{"title":"strawberry smoothie","amount":98,"note":"updated note","tags":["food","beverage"]}`
	expectResponseBody := `{"id":1,"title":"strawberry smoothie","amount":98,"note":"updated note","tags":["food","beverage"]}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "/expenses", strings.NewReader(updateExpenseRequestJsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	mockService := mockExpenseService{}
	h := NewExpenseHandler(mockService)

	// Assertions
	if assert.NoError(t, h.UpdateExpense(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectResponseBody, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestHandlerGetAllExpensesSuccess(t *testing.T) {
	expectResponseBody := `[{"id":1,"title":"strawberry smoothie","amount":98,"note":"note","tags":["food","beverage"]},{"id":2,"title":"apple smoothie","amount":89,"note":"no discount","tags":["beverage"]}]`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/expenses", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockService := mockExpenseService{}
	h := NewExpenseHandler(mockService)

	// Assertions
	if assert.NoError(t, h.GetAllExpenses(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectResponseBody, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}
