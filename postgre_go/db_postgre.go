package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "192.168.1.8"
	port     = 5432
	user     = "postgres"
	password = "gigih"
	dbname   = "myproject"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `select * from mhsw_get($1::int)`
	row, err := db.Query(query, 5)
	if err != nil {
		panic(err)
	}
	var name string
	for row.Next() {
		err = row.Scan(&name)
		println(name)
	}

	// sqlStatement := `
	//  INSERT INTO users (age, email, first_name, last_name)
	//  VALUES ($1, $2, $3, $4)
	//  RETURNING id`
	// sqlStatement := `select first_name, last_name from users where email=$1`
	// id := 2
	// var age int
	// var first_name, last_name string
	// err = db.QueryRow(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun").Scan(&id, &email)
	// row := db.QueryRow(sqlStatement, "jon@calhoun.io")
	// err = row.Scan(&first_name, &last_name)
	//  switch err = row.Scan(&first_name, &last_name); err {
	// case sql.ErrNoRows:
	// 	println("No Row Return")
	// case nil:
	// 	println("This name is: ", first_name, "", last_name)
	// default:
	// 	panic(err)
	// }
	// if err != nil {
	// 	panic(err)
	// }
	// println("data: ", first_name, last_name)

	// rows, err := db.Query("SELECT id, first_name FROM users LIMIT $1", 5)
	// if err != nil {
	// 	// handle this error better than this
	// 	panic(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	var id int
	// 	var firstName string
	// 	err = rows.Scan(&id, &firstName)
	// 	if err != nil {
	// 		// handle this error
	// 		panic(err)
	// 	}
	// 	fmt.Println(id, firstName)
	// }
	// // get any error encountered during iteration
	// err = rows.Err()
	// if err != nil {
	// 	panic(err)
	// }
}
