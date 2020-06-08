package database // import "github.com/mauritt/blog/blogPostApi/internal/database"

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// ConnectDB connects to the database
func ConnectDB() (*sql.DB, error) {

	dbUser, userExists := os.LookupEnv("MYSQLL_USER")
	if userExists != true {
		log.Printf("Couldn't get database username")
		return nil, fmt.Errorf("Couldn't get database username")
	}

	dbPassword, passwordExists := os.LookupEnv("MYSQL_PASSWORD")
	if passwordExists != true {
		log.Printf("Couldn't get database password")
		return nil, fmt.Errorf("Couldn't get database password")
	}

	connectStatement := fmt.Sprintf("%v:%v@tcp(db:3306)/blog", dbUser, dbPassword)

	db, err := sql.Open("mysql", connectStatement)
	if err != nil {
		log.Printf("Couldn't connect to db:\n%v", err.Error())
		return nil, err
	}
	return db, nil
}
