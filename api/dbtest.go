package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func MyDBHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := sql.Open("pgx", "host=ec2-34-242-89-204.eu-west-1.compute.amazonaws.com port=5432 dbname=dc6sc248lptkch user=wweeunqhtbukan password=63eb2d71cecda037bc505a347fa215681f21692d2ac9d03f6dbb2194093af577")
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

	query := `insert into users (first_name, last_name, created_at, updated_at) values ($1, $2, $3, $4)`
	_, err = conn.Exec(query, "Andi", "Mehrfach", time.Now(), time.Now())
	if err != nil {
		log.Fatal(err)
	}

	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}
}

func getAllRows(conn *sql.DB) error {
	// rows, err := conn.Query("select id, first_name, last_name from users")
	q := `
		select u.id, u.first_name, u.last_name, e.email from users u
		inner join emails e on (e.user_id = u.id)
	`

	rows, err := conn.Query(q)
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	var firstName, lastName, email string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName, &email)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("Record is", id, firstName, lastName, email)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}

	return nil
}