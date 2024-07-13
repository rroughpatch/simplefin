package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Index renders the index page
func Index(c echo.Context) error {
	log.Println("Index handler called") // Add logging
	data := map[string]interface{}{
		"title": "SimpleFin",
	}
	return c.Render(http.StatusOK, "index.html", data)
}

// Budget renders the budget page
func Budget(c echo.Context) error {
	log.Println("Budget handler called") // Add logging
	data := map[string]interface{}{
		"title": "Budget Tracking",
	}
	return c.Render(http.StatusOK, "budget.html", data)
}
