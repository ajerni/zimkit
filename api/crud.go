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

func MyCrudHandler(w http.ResponseWriter, r *http.Request) {
	// use like https://zimkit.vercel.app/api/crud?op=c,r,u,d&id=xyz&first=xyz&last=xyz

	operation := r.URL.Query().Get("op")
	id_req := r.URL.Query().Get("id")
	first := r.URL.Query().Get("first")
	last := r.URL.Query().Get("last")

	conn, err := sql.Open("pgx", "host=ec2-34-242-89-204.eu-west-1.compute.amazonaws.com port=5432 dbname="+os.Getenv("DB_NAME")+" user=wweeunqhtbukan password="+os.Getenv("DB_PW"))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect: %v\n", err))
	}
	defer conn.Close()

	log.Println("Connected to database!")

	err = conn.Ping()
	if err != nil {
		log.Fatal("Cannot ping database!")
	}

	log.Println("Pinged database!")

	// switch throug CRUD operations
	switch operation {
	case "c":
		// Create (insert)
		query := `insert into users (first_name, last_name, created_at, updated_at) values ($1, $2, $3, $4)`
		_, err = conn.Exec(query, first, last, time.Now(), time.Now())
		if err != nil {
			log.Fatal(err)
		}
	case "r":
		// Read one (select)
		query := `select id, first_name, last_name from users where id = $1`

		var firstName, lastName string
		var id int

		row := conn.QueryRow(query, id_req)
		err = row.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, fmt.Sprintf("QueryRow returns %d %s %s", id, firstName, lastName))
	case "u":
		// Update
		query := `
		UPDATE users
		SET first_name = $1, last_name = $2
		WHERE id = $3;
		`
		_, err = conn.Exec(query, first, last, id_req)
		if err != nil {
			log.Fatal(err)
		}
	case "d":
		// Delete
		query := `delete from users where id = $1`
		_, err = conn.Exec(query, id_req)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Default()
	}

	err = getAllRows(conn, w)
	if err != nil {
		log.Fatal(err)
	}
}

func getAllRows(conn *sql.DB, w http.ResponseWriter) error {
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

		// TODO: retun html
		fmt.Fprintf(w, fmt.Sprintf("Record is %d %s %s\n", id, firstName, lastName))
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}

	return nil
}
