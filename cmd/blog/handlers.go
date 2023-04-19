package main

import (
<<<<<<< HEAD
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
=======
	"html/template"
	"log"
	"net/http"

>>>>>>> 2933629f94622529b12de65018a845ea35fee356
	"github.com/jmoiron/sqlx"
)

type indexPageData struct {
	Title           string
	Subtitle        string
<<<<<<< HEAD
	FeaturedPosts   []*featuredPostData
	MostRecentPosts []*mostRecentPostData
}

type postData struct {
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	ImgModifier string `db:"img_modifier"`
	Content     string `db:"content"`
}

type featuredPostData struct {
	PostID         string `db:"post_id"`
=======
	FeaturedPosts   []featuredPostData
	MostRecentPosts []mostRecentPostData
}

type postPageData struct {
	Title           string
	Subtitle        string
}

type featuredPostData struct {
>>>>>>> 2933629f94622529b12de65018a845ea35fee356
	Title          string `db:"title"`
	Subtitle       string `db:"subtitle"`
	ImgModifier    string `db:"img_modifier"`
	Author         string `db:"author"`
	AuthorModifier string `db:"author_modifier"`
	PublishDate    string `db:"publish_date"`
<<<<<<< HEAD
	PostURL        string
}

type mostRecentPostData struct {
	PostID         string `db:"post_id"`
=======
}

type mostRecentPostData struct {
>>>>>>> 2933629f94622529b12de65018a845ea35fee356
	Title          string `db:"title"`
	Subtitle       string `db:"subtitle"`
	ImgModifier    string `db:"img_modifier"`
	Author         string `db:"author"`
	AuthorModifier string `db:"author_modifier"`
	PublishDate    string `db:"publish_date"`
<<<<<<< HEAD
	PostURL        string
=======
>>>>>>> 2933629f94622529b12de65018a845ea35fee356
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		featuredPostsData, err := featuredPosts(db)
		if err != nil {
<<<<<<< HEAD
			http.Error(w, "Internal Server Error", 500)
=======
			http.Error(w, "Internal Server Error", 500) 
>>>>>>> 2933629f94622529b12de65018a845ea35fee356
			log.Println(err)
			return
		}
		mostRecentPostsData, err := mostRecentPosts(db)
		if err != nil {
<<<<<<< HEAD
			http.Error(w, "Internal Server Error", 500)
=======
			http.Error(w, "Internal Server Error", 500) 
>>>>>>> 2933629f94622529b12de65018a845ea35fee356
			log.Println(err)
			return
		}

<<<<<<< HEAD
		ts, err := template.ParseFiles("pages/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
=======
		ts, err := template.ParseFiles("pages/index.html") 
		if err != nil {
			http.Error(w, "Internal Server Error", 500) 
>>>>>>> 2933629f94622529b12de65018a845ea35fee356
			log.Println(err)
			return
		}

		data := indexPageData{
			Title:           "Let's do it together.",
			Subtitle:        "We travel the world in search of stories. Come along for the ride. ",
			FeaturedPosts:   featuredPostsData,
			MostRecentPosts: mostRecentPostsData,
		}

		err = ts.Execute(w, data)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		log.Println("Request completed successfully")
	}
}

<<<<<<< HEAD
func post(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		postIDStr := mux.Vars(r)["postID"]
		postID, err := strconv.Atoi(postIDStr)
		if err != nil {
			http.Error(w, "Invalid post id", 403)
=======
func post(w http.ResponseWriter, r *http.Request){
	ts, err := template.ParseFiles("pages/post.html") 
		if err != nil {
			http.Error(w, "Internal Server Error", 500) 
>>>>>>> 2933629f94622529b12de65018a845ea35fee356
			log.Println(err)
			return
		}

<<<<<<< HEAD
		post, err := postByID(db, postID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Post not found", 404)
				log.Println(err)
				return
			}

			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/post.html")
=======
		data := postPageData{
			Title:           "The Road Ahead",
			Subtitle:        "The road ahead might be paved - it might not be.",
		}
		
		err = ts.Execute(w, data)
>>>>>>> 2933629f94622529b12de65018a845ea35fee356
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}
<<<<<<< HEAD

		err = ts.Execute(w, post)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		log.Println("Request completed successfully")
	}
}

func featuredPosts(db *sqlx.DB) ([]*featuredPostData, error) {
	const query = `
		SELECT
			post_id,
=======
		
		log.Println("Request completed successfully")
}
func featuredPosts(db *sqlx.DB) ([]featuredPostData, error) {
	const query = `
		SELECT
>>>>>>> 2933629f94622529b12de65018a845ea35fee356
			title,
			subtitle,
			img_modifier,
			author,
			author_modifier,
			publish_date
		FROM
			post
		WHERE featured = 1
	`

<<<<<<< HEAD
	var posts []*featuredPostData

	err := db.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		post.PostURL = "/post/" + post.PostID // Формируем исходя из ID ордера в базе
	}

	fmt.Println(posts)

	return posts, nil
}

func mostRecentPosts(db *sqlx.DB) ([]*mostRecentPostData, error) {
	const query = `
		SELECT
		 	post_id,
=======
	var posts []featuredPostData 

	err := db.Select(&posts, query) 
	if err != nil {                        
		return nil, err
	}

	return posts, nil
}

func mostRecentPosts(db *sqlx.DB) ([]mostRecentPostData, error) {
	const query = `
		SELECT
>>>>>>> 2933629f94622529b12de65018a845ea35fee356
			title,
			subtitle,
			img_modifier,
			author,
			author_modifier,
			publish_date
		FROM
			post
		WHERE featured = 0
	`

<<<<<<< HEAD
	var posts []*mostRecentPostData

	err := db.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		post.PostURL = "/post/" + post.PostID // Формируем исходя из ID ордера в базе
	}

	fmt.Println(posts)

	return posts, nil
}

// Получает информацию о конкретном ордере из базы данных
func postByID(db *sqlx.DB, postID int) (postData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			img_modifier,
			content
		FROM
			post
		WHERE
			post_id = ?
	`
	// В SQL-запросе добавились параметры, как в шаблоне. ? означает параметр, который мы передаем в запрос ниже

	var post postData

	// Обязательно нужно передать в параметрах orderID
	err := db.Get(&post, query, postID)
	if err != nil {
		return postData{}, err
	}

	return post, nil
}
=======
	var posts []mostRecentPostData 

	err := db.Select(&posts, query) 
	if err != nil {                           
		return nil, err
	}

	return posts, nil
}
>>>>>>> 2933629f94622529b12de65018a845ea35fee356
