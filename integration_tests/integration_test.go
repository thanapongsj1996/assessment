//go:build integration

package integration_tests

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/thanapongsj1996/assessment/database"
	"github.com/thanapongsj1996/assessment/modules/expense/handler"
	"github.com/thanapongsj1996/assessment/modules/expense/model"
	"github.com/thanapongsj1996/assessment/modules/expense/repository"
	"github.com/thanapongsj1996/assessment/modules/expense/service"
	"github.com/thanapongsj1996/assessment/routes"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

const serverPort = 80

func setupServer(eh *echo.Echo) {
	go func(e *echo.Echo) {

		database.InitDB("postgresql://root:root@db/assessment-integration?sslmode=disable")

		db := database.GetDB()
		defer db.Table("expenses").Delete(&model.Expense{})
		defer database.CloseDB()

		expenseRepo := repository.NewExpenseRepository(db)
		expenseService := service.NewExpenseService(expenseRepo)
		expenseHandler := handler.NewExpenseHandler(expenseService)
		routes.NewExpenseRoute(e, expenseHandler)

		e.Start(fmt.Sprintf(":%d", serverPort))
	}(eh)
	for {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%d", serverPort), 30*time.Second)
		if err != nil {
			log.Println(err)
		}
		if conn != nil {
			conn.Close()
			break
		}
	}
}

func TestAddExpense(t *testing.T) {
	// Setup server
	eh := echo.New()
	setupServer(eh)

	// Arrange
	reqBody := `{"title":"banana smoothie","amount":79,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}`
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%d/expenses", serverPort), strings.NewReader(reqBody))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	client := http.Client{}

	// Act
	resp, err := client.Do(req)
	assert.NoError(t, err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	resp.Body.Close()

	// Assertions
	expected := `{"id":2,"title":"banana smoothie","amount":79,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}`

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		assert.Equal(t, expected, strings.TrimSpace(string(byteBody)))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)
}

func TestGetExpenseById(t *testing.T) {
	// Setup server
	eh := echo.New()
	setupServer(eh)

	// Arrange
	reqBody := ``
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/expenses/1", serverPort), strings.NewReader(reqBody))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	client := http.Client{}

	// Act
	resp, err := client.Do(req)
	assert.NoError(t, err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	resp.Body.Close()

	// Assertions
	expected := `{"id":1,"title":"strawberry smoothie","amount":79,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}`

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, expected, strings.TrimSpace(string(byteBody)))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)
}

func TestUpdateExpense(t *testing.T) {
	// Setup server
	eh := echo.New()
	setupServer(eh)

	// Arrange
	reqBody := `{"title":"apple smoothie","amount":79,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}`
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:%d/expenses/1", serverPort), strings.NewReader(reqBody))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	client := http.Client{}

	// Act
	resp, err := client.Do(req)
	assert.NoError(t, err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	resp.Body.Close()

	// Assertions
	expected := `{"id":1,"title":"apple smoothie","amount":79,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}`

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, expected, strings.TrimSpace(string(byteBody)))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)
}

func TestGetAllExpenses(t *testing.T) {
	// Setup server
	eh := echo.New()
	setupServer(eh)

	// Arrange
	reqBody := ``
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/expenses", serverPort), strings.NewReader(reqBody))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	client := http.Client{}

	// Act
	resp, err := client.Do(req)
	assert.NoError(t, err)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)
}
