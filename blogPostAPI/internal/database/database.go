package database // import "github.com/mauritt/blog/blogPostApi/internal/database"

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// BlogPost containts the data for blog posts
type BlogPost struct {
	ID       int    `json:"id"`
	Headline string `json:"headline"`
	Content  string `json:"content"`
}

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

// GetBlogPosts gets all posts
func GetBlogPosts(db *sql.DB) ([]BlogPost, error) {

	var post BlogPost
	var posts []BlogPost

	query := "SELECT * FROM post"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("Couldn't prepare statement: %v\v%v", query, err.Error())
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Printf("Query failed\n%v", err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&post.ID, &post.Headline, &post.Content)
		if err != nil {
			log.Printf(err.Error())
			return nil, err
		}
		posts = append(posts, post)

	}
	err = rows.Err()
	if err = rows.Err(); err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	return posts, nil

}
