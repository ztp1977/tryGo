package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/k0kubun/pp"
)

var schema = `
CREATE TABLE person (
    first_name text,
    last_name text,
    email text
);

CREATE TABLE place (
    country text,
    city text NULL,
    telcode int
);`

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}

func main() {
	// this Pings the database trying to connect, panics on error
	// use sqlx.Open() for sql.Open() semantics
	//db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
	db, err := sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}

	pp.Println(db)

	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv
	//db.MustExec(schema)

	//tx := db.MustBegin()
	//tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	//tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
	//tx.MustExec("INSERT INTO place (country, city, telcode) VALUES ($1, $2, $3)", "United States", "New York", "1")
	//tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Hong Kong", "852")
	//tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Singapore", "65")
	//// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	//tx.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &Person{"Jane", "Citizen", "jane.citzen@example.com"})
	//tx.Commit()

	// Query the database, storing results in a []Person (wrapped in []interface{})
	people := []Person{}
	db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")
	pp.Println(people)
	jason, john := people[0], people[1]

	fmt.Printf("%#v\n%#v", jason, john)
	// Person{FirstName:"Jason", LastName:"Moiron", Email:"jmoiron@jmoiron.net"}
	// Person{FirstName:"John", LastName:"Doe", Email:"johndoeDNE@gmail.net"}

	// You can also get a single result, a la QueryRow
	jason = Person{}
	err = db.Get(&jason, "SELECT * FROM person WHERE first_name=$1", "Jason")
	fmt.Printf("%#v\n", jason)
	// Person{FirstName:"Jason", LastName:"Moiron", Email:"jmoiron@jmoiron.net"}

	// if you have null fields and use SELECT *, you must use sql.Null* in your struct
	places := []Place{}
	err = db.Select(&places, "SELECT * FROM place ORDER BY telcode ASC")
	if err != nil {
		fmt.Println(err)
		return
	}
	usa, singsing, honkers := places[0], places[1], places[2]

	fmt.Printf("%#v\n%#v\n%#v\n", usa, singsing, honkers)
	// Place{Country:"United States", City:sql.NullString{String:"New York", Valid:true}, TelCode:1}
	// Place{Country:"Singapore", City:sql.NullString{String:"", Valid:false}, TelCode:65}
	// Place{Country:"Hong Kong", City:sql.NullString{String:"", Valid:false}, TelCode:852}

	// Loop through rows using only one struct
	place := Place{}
	rows, err := db.Queryx("SELECT * FROM place")
	for rows.Next() {
		err := rows.StructScan(&place)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", place)
	}
	// Place{Country:"United States", City:sql.NullString{String:"New York", Valid:true}, TelCode:1}
	// Place{Country:"Hong Kong", City:sql.NullString{String:"", Valid:false}, TelCode:852}
	// Place{Country:"Singapore", City:sql.NullString{String:"", Valid:false}, TelCode:65}

	// Named queries, using `:name` as the bindvar.  Automatic bindvar support
	// which takes into account the dbtype based on the driverName on sqlx.Open/Connect
	_, err = db.NamedExec(`INSERT INTO person (first_name,last_name,email) VALUES (:first,:last,:email)`,
		map[string]interface{}{
			"first": "Bin",
			"last":  "Smuth",
			"email": "bensmith@allblacks.nz",
		})

	// Selects Mr. Smith from the database
	rows, err = db.NamedQuery(`SELECT * FROM person WHERE first_name=:fn`, map[string]interface{}{"fn": "Bin"})

	// Named queries can also use structs.  Their bind names follow the same rules
	// as the name -> db mapping, so struct fields are lowercased and the `db` tag
	// is taken into consideration.
	rows, err = db.NamedQuery(`SELECT * FROM person WHERE first_name=:first_name`, jason)
}
