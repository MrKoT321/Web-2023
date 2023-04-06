package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type indexPageData struct {
	Title           string
	Subtitle        string
	FeaturedPosts   []featuredPostData
	MostRecentPosts []mostRecentPostData
}

type postPageData struct {
	Title           string
	Subtitle        string
}

type featuredPostData struct {
	Title          string `db:"title"`
	Subtitle       string `db:"subtitle"`
	ImgModifier    string `db:"img_modifier"`
	Author         string `db:"author"`
	AuthorModifier string `db:"author_modifier"`
	PublishDate    string `db:"publish_date"`
}

type mostRecentPostData struct {
	Title          string `db:"title"`
	Subtitle       string `db:"subtitle"`
	ImgModifier    string `db:"img_modifier"`
	Author         string `db:"author"`
	AuthorModifier string `db:"author_modifier"`
	PublishDate    string `db:"publish_date"`
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

func post(w http.ResponseWriter, r *http.Request){
	ts, err := template.ParseFiles("pages/post.html") 
		if err != nil {
			http.Error(w, "Internal Server Error", 500) 
			log.Println(err)
			return
		}

		data := postPageData{
			Title:           "The Road Ahead",
			Subtitle:        "The road ahead might be paved - it might not be.",
		}
		
		err = ts.Execute(w, data)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}
		
		log.Println("Request completed successfully")
}
func featuredPosts(db *sqlx.DB) ([]featuredPostData, error) {
	const query = `
		SELECT
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

	var posts []mostRecentPostData 

	err := db.Select(&posts, query) 
	if err != nil {                           
		return nil, err
	}

	return posts, nil
}
