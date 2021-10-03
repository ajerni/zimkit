package handler

import (
	"database/sql"
    "fmt"
	"net/http"
	"log"
    "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func MyDbHandler(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "andiroot:LetMeIn1976@tcp(xnmxviali589.db.solnet.ch)/andigodb")

    // if there is an error opening the connection, handle it
    if err != nil {
        log.Print(err.Error())
    }
    defer db.Close()

    // Execute the query
    results, err := db.Query("SELECT * FROM gotabel WHERE id = 2")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

	fmt.Fprintf(w, fmt.Sprintf("Result: %s", results))

	defer db.Close()
   
// // Capture connection properties.
// cfg := mysql.Config{
// 	// User:   os.Getenv("DBUSER"),
// 	// Passwd: os.Getenv("DBPASS"),
// 	User: "andiroot",
// 	Passwd: "LetMeIn1976",
// 	Net:    "tcp",
// 	Addr:   "xnmxviali589.db.solnet.ch",
// 	DBName: "andigodb",
// }
// // Get a database handle.
// var err error
// db, err = sql.Open("mysql", cfg.FormatDSN())
// if err != nil {
// 	log.Fatal(err)
// }

// pingErr := db.Ping()
// if pingErr != nil {
// 	log.Fatal(pingErr)
// }
// fmt.Println("Connected!")

// row := db.QueryRow("SELECT * FROM gotable WHERE id = 2")
// fmt.Fprintf(w, fmt.Sprintf("Result %s", row))

}