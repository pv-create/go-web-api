package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct {
	id             int     `json:"id"`
	FirstName      string  `json:"FirstName"`
	SecondName     string  `json:"SecondName"`
	balance        float64 `json:"balance"`
	TelephonNumber string  `json:"TelephonNumber"`
}

func getAllUsers(c *gin.Context) {
	connStr := "user=pavelvilkov dbname=bank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from Users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []User{}

	for rows.Next() {
		p := User{}
		err := rows.Scan(&p.id, &p.FirstName, &p.SecondName, &p.balance, &p.TelephonNumber)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	e, err := json.Marshal(products)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(e))
	c.IndentedJSON(http.StatusOK, string(e))
}

func main() {
	router := gin.Default()
	router.GET("/Users", getAllUsers)
	router.Run("localhost:8080")
}
