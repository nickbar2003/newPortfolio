package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/go-sql-driver/mysql"
)

type HelloWorld struct {
	Message string `json:"message"`
}

type Figure struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
	Deathdate string `json:"deathdate"`
}

var quotesDB *sql.DB

func main() {
	connStr := "root:mauFJcuf5dhRMQrjj@/quotes"

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	quotesDB = db

	quotesDB.SetConnMaxLifetime(time.Minute * 3)
	quotesDB.SetMaxOpenConns(10)
	quotesDB.SetMaxIdleConns(10)

	// rows, err := db.Query("SELECT * FROM figure")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/fetchFigures", fetchFigures)
	e.GET("/hello", Greetings)
	e.Logger.Fatal(e.Start(":3000"))

}

func Greetings(c echo.Context) error {
	return c.JSON(http.StatusOK, HelloWorld{
		Message: "Hello World",
	})
}

// This function adapts code for querying from the GO website
// See Here: https://go.dev/doc/database/querying
func fetchFigures(c echo.Context) error {
	log.Println("made it to function")
	name := c.QueryParam("name")

	// query :=
	var searchFigure Figure
	// var figure sql.row
	err := quotesDB.QueryRow("SELECT * FROM figure WHERE name = ?", name).Scan(&searchFigure.Id, &searchFigure.Name, &searchFigure.Birthdate, &searchFigure.Deathdate)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("no results found", err)
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Figure not found"})
		}
		log.Println(err)
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Encounterd error in query"})
	}

	birth := searchFigure.Birthdate
	// death := searchFigure.Deathdate

	// look it up again you can definitly do what you want with this query, use multiple paramters and conidtions
	rows, err := quotesDB.Query("SELECT * FROM figure WHERE birthdate <= ?", birth)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("no results found", err)
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Figure not found"})
		}
		log.Println(err)
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Encounterd error in query"})
	}
	defer rows.Close()

	var figures []Figure

	for rows.Next() {
		var figure Figure
		if err := rows.Scan(&figure.Id, &figure.Name, &figure.Birthdate, &figure.Deathdate); err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Encounterd error in query"})
		}
		figures = append(figures, figure)
	}
	if err = rows.Err(); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Encounterd error in query"})
	}

	log.Println(figures)
	return c.JSON(http.StatusOK, figures)
}
