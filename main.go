package main

import (
	"database/sql"
	"getCovid/cases"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
)

var DatabaseConnection *sql.DB

func main() {

	operation, err := sql.Open("postgres", "postgres://postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("Error Opening Database: %q", err)
	}

	DatabaseConnection = operation

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/patient", func(c *gin.Context) {
		location := c.Query("statecode")
		gender := c.Query("gender") // shortcut for c.Request.URL.Query().Get("lastname")
		age := c.Query("agebracket")
		status := c.Query("currentstatus")

		c.String(http.StatusOK, "Hello %s %s %s %s", location, gender, age, status)
	})

	router.GET("/cases", cases.Getcovidcases)

	router.Run()

}
