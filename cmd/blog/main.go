package main

import (
	"database/sql"
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> 2933629f94622529b12de65018a845ea35fee356
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
<<<<<<< HEAD
	// "github.com/gorilla/mux"
=======
>>>>>>> 2933629f94622529b12de65018a845ea35fee356
	"github.com/jmoiron/sqlx"
)

const (
	port         = ":8000"
	dbDriverName = "mysql"
)

func main() {
	db, err := openDB()
	if err != nil {
		log.Fatal(err)
	}

	dbx := sqlx.NewDb(db, dbDriverName)

<<<<<<< HEAD
	// mux := mux.NewRouter()
	mux := http.NewServeMux()
	mux.HandleFunc("/home", index(dbx))

	mux.HandleFunc("/post/{postID}", post(dbx))

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	fmt.Println("Start server" + port)
=======
	mux := http.NewServeMux()
	mux.HandleFunc("/home", index(dbx))
	mux.HandleFunc("/post", post)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Start server " + port)
>>>>>>> 2933629f94622529b12de65018a845ea35fee356
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func openDB() (*sql.DB, error) {
<<<<<<< HEAD
	return sql.Open(dbDriverName, "root:1234@tcp(localhost:3306)/blog?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true")
}
=======
	return sql.Open(dbDriverName, "root:qwerty15122004@tcp(localhost:3306)/blog?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true")
}
>>>>>>> 2933629f94622529b12de65018a845ea35fee356
