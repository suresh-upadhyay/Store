package cmd

import (
	"database/sql"
	"fmt"
	"github.com/dunzoit/Store/config"
	_ "github.com/lib/pq" // here
	"log"
)

func mustPrepareDB() (*sql.DB, error) {
	var connectionString  = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.DBUSERNAME, config.DBPASSWORD, config.DBNAME)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
