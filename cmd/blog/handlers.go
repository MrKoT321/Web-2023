package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type indexPageData struct {
	Title           string
	Subtitle        string
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
	Title          string `db:"title"`
	Subtitle       string `db:"subtitle"`
	ImgModifier    string `db:"img_modifier"`
	Author         string `db:"author"`
	AuthorModifier string `db:"author_modifier"`
	PublishDate    string `db:"publish_date"`
	PostURL        string
}

type mostRecentPostData struct {
	PostID         string `db:"post_id"`
	Title          string `db:"title"`
	Subtitle       string `db:"subtitle"`
	ImgModifier    string `db:"img_modifier"`
	Author         string `db:"author"`
	AuthorModifier string `db:"author_modifier"`
	PublishDate    string `db:"publish_date"`
	PostURL        string
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		featuredPostsData, err := featuredPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}
		mostRecentPostsData, err := mostRecentPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
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

func post(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		postIDStr := mux.Vars(r)["postID"]
		postID, err := strconv.Atoi(postIDStr)
		if err != nil {
			http.Error(w, "Invalid post id", 403)
			log.Println(err)
			return
		}

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
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

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
