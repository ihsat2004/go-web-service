package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// db configuration details
const (
	postgres_host     = "dpg-cot6bcsf7o1s73e1na60-a.singapore-postgres.render.com"
	postgres_port     = 5432
	postgres_user     = "postgres_admin"
	postgres_password = "mpVh1OvK6X9Y6RyFrbHZMvKYDBH4mMyK"
	postgres_dbname   = "my_db_htpg"
)

// create pointer variable Db which points to sql driver
var Db *sql.DB

// init() is always called before main()
func init() {
	// creating the connection string
	db_info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", postgres_host, postgres_port, postgres_user, postgres_password, postgres_dbname)

	var err error
	// open connection to database
	Db, err = sql.Open("postgres", db_info)

	if err != nil {
		panic(err)
	} else {
		log.Println("Database successfully configured")
	}
}
