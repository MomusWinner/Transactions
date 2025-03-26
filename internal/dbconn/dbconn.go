package dbconn

import (
	"Transactions/database"
	"Transactions/internal/config"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DBQueries *database.Queries
var DB *sql.DB

func Init(conf *config.Config) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.PostgresHost,
		conf.PostgresPort,
		conf.PostgresUser,
		conf.PostgresPassword,
		conf.PostgresDatabase,
	)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("Couldn't connect to database. %v", err)
		os.Exit(0)
	}

	DBQueries = database.New(DB)
}
