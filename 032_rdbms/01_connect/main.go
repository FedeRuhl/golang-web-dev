package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB
var err error

func main() {
	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err = sql.Open("mysql", os.Getenv("CONNECTION"))
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	// check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Successfully completed.")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
