package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	gowiki "github.com/trietmn/go-wiki"

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

// type Info struct {
// 	Born string `json:"born"`
// 	Died string `json:"died"`
// }

// type SearchFigure struct {
// 	Name  string `json:"name"`
// 	Title string `json:"title"`
// 	Info  Info   `json:"info"`
// }

var figureDB *sql.DB

func main() {
	// connStr := "root:mauFJcuf5dhRMQrjj@/quotes"
	// connStr := "sql3680854:ugjnPiwSB7@tcp(sql3.freesqldatabase.com:3306)/sql3680854"
	connStr := "freedb_sql3680854:5tSKgASdNR2&5eV@tcp(sql.freedb.tech:3306)/freedb_sql3680854"

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	figureDB = db

	figureDB.SetConnMaxLifetime(time.Minute * 3)
	figureDB.SetMaxOpenConns(10)
	figureDB.SetMaxIdleConns(10)

	// rows, err := db.Query("SELECT * FROM figure")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/fetchFigures", fetchFigures)
	e.GET("/hello", Greetings)
	e.GET("/fetchBio", fetchBio)
	e.GET("/fetchImage", fetchImage)
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
	err := figureDB.QueryRow("SELECT * FROM figure WHERE name = ?", name).Scan(&searchFigure.Id, &searchFigure.Name, &searchFigure.Birthdate, &searchFigure.Deathdate)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("no results found", err)
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Figure not found"})
		}
		log.Println(err)
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Encounterd error in query"})
	}

	birth := searchFigure.Birthdate
	death := searchFigure.Deathdate

	// look it up again you can definitly do what you want with this query, use multiple paramters and conidtions
	rows, err := figureDB.Query("SELECT * FROM figure WHERE birthdate <= ? AND deathdate >= ? OR birthdate >= ? AND birthdate <= ? OR birthdate <= ? AND deathdate >= ? OR birthdate >= ? AND deathdate <= ?", birth, birth, birth, death, birth, death, birth, death)
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

	return c.JSON(http.StatusOK, figures)
}

func fetchBio(c echo.Context) error {
	name := c.QueryParam("name")
	log.Println(name)
	res, err := gowiki.Summary(name, 5, -1, false, true)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Summary: %v\n", res)
	return c.JSON(http.StatusOK, res)

}

func fetchImage(c echo.Context) error {
	name := c.QueryParam("name")
	log.Println(name)
	page, err := gowiki.GetPage(name, -1, false, true)
	if err != nil {
		log.Println(err)
	}

	html, err := page.GetHTML()
	if err != nil {
		log.Println(err)
	}

	// image, err := page.GetImagesURL()
	// if err != nil {
	// 	log.Println(err)
	// }

	// return c.JSON(http.StatusOK, image[0])
	return c.JSON(http.StatusOK, html)

}
