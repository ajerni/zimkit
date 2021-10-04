package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func MyDBHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := sql.Open("pgx", "host=ec2-34-242-89-204.eu-west-1.compute.amazonaws.com port=5432 dbname="+os.Getenv("DB_NAME")+" user=wweeunqhtbukan password="+os.Getenv("DB_PW"))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect: %v\n", err))
	}
	defer conn.Close()

	log.Println("Connected to database!")

	// test my connection
	err = conn.Ping()
	if err != nil {
		log.Fatal("Cannot ping database!")
	}

	log.Println("Pinged database!")

	query := `delete from users where first_name = 'Andi'`
	_, err = conn.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	err = getAllRows(conn, w)
	if err != nil {
		log.Fatal(err)
	}
}

func getAllRows(conn *sql.DB, w http.ResponseWriter) error {
	// rows, err := conn.Query("select id, first_name, last_name from users")
	// q := `
	// 	select u.id, u.first_name, u.last_name, e.email from users u
	// 	inner join emails e on (e.user_id = u.id)
	// `
	q := `select id, first_name, last_name from users`

	rows, err := conn.Query(q)
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	var firstName, lastName string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}
		// fmt.Println("Record is", id, firstName, lastName, email)
		fmt.Fprintf(w, fmt.Sprintf("Record is %d %s %s\n", id, firstName, lastName))
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}

	return nil
}
