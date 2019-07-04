package request

import (
	"database/sql"
	"fmt"
	"log"
)

func dbConnect() *sql.DB {
	resp := Option().Database
	db, err := sql.Open("mysql", resp.User+":"+resp.Passwd+"@tcp("+resp.Server+")/"+resp.DB)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection Established")
	}
	return db
}

func GetUser(id int) {
	var email string
	db := dbConnect()
	rows, err := db.Query("select request_id, email from gdpr_requests where request_id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	for rows.Next() {
		err := rows.Scan(&id, &email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, email)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
