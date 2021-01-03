package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"rest-api.jishnu.net/models"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var err error

	username := os.Getenv("USERNAME")
	password := os.Getenv("MYSQL_DB_PASSWORD")
	host := os.Getenv("MYSQL_DB_HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	dbname := os.Getenv("DBNAME")

	// Initalize the sql.DB connection and assign the same to the global variable
	//Taking the values from the configmap and secret implemented via kubernetes objects
	models.DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbname))
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/hello", helloFromDB)
	http.ListenAndServe(":8080", nil)
}

// helloFromDB will receive the http request and send back the message string as response
func helloFromDB(w http.ResponseWriter, r *http.Request) {
	msgs, err := models.Hello()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, msg := range msgs {
		fmt.Fprintf(w, "%s, %d \n", msg.Greeting, msg.Year)
	}
}
