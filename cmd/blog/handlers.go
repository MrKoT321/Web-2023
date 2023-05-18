package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
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
	Title    string `db:"title"`
	Subtitle string `db:"subtitle"`
	ImgURL   string `db:"post_img"`
	Content  string `db:"content"`
}

type adminPageData struct {
	Title string
}

type loginPageData struct {
	Title string
}

type createPostRequest struct {
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	PostIMG     string `json:"postIMG"`
	PostName    string `json:"postIMGName"`
	Author      string `json:"authorName"`
	AuthorIMG   string `json:"authorIMG"`
	AuthorName  string `json:"authorIMGName"`
	PreviewIMG  string `json:"previewIMG"`
	PreviewName string `json:"previewIMGName"`
	PublishDate string `json:"publishDate"`
	Content     string `json:"content"`
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
	ImgURL         string `db:"preview_img"`
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
			http.Error(w, "Invalid post id", 404)
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

func admin(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ts, err := template.ParseFiles("pages/admin.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}
		data := adminPageData{
			Title: "Admin",
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

func login(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ts, err := template.ParseFiles("pages/login.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		data := loginPageData{
			Title: "Login",
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
		post.PostURL = "/post/" + post.PostID
	}

	return posts, nil
}

func mostRecentPosts(db *sqlx.DB) ([]*mostRecentPostData, error) {
	const query = `
		SELECT
		 	post_id,
			title,
			subtitle,
			preview_img,
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
		post.PostURL = "/post/" + post.PostID
	}

	return posts, nil
}

func postByID(db *sqlx.DB, postID int) (postData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			post_img,
			content
		FROM
			post
		WHERE
			post_id = ?
	`

	var post postData

	err := db.Get(&post, query, postID)
	if err != nil {
		return postData{}, err
	}

	return post, nil
}

func createPost(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		var req createPostRequest

		err = json.Unmarshal(reqData, &req)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		err = savePost(db, req)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		log.Println("Save post request completed successfully")
	}
}

func savePost(db *sqlx.DB, req createPostRequest) error {
	
	authorImgName, err := makeImg(req.AuthorName, req.AuthorIMG)
	if err != nil {
		log.Println(err)
		return err
	}
	postImgName, err := makeImg(req.PostName, req.PostIMG)
	if err != nil {
		log.Println(err)
		return err
	}
	previewImgName, err := makeImg(req.PreviewName, req.PreviewIMG)
	if err != nil {
		log.Println(err)
		return err
	}

	const query = `
       INSERT INTO post
       (
           title,
           subtitle,
		   preview_img,
		   post_img,
		   author,
		   author_modifier,
		   publish_date,
		   content
       )
       VALUES
       (
           ?,
           ?,
		   ?,
		   ?,
		   ?,
		   ?,
		   ?,
		   ?
       )
	`

	_, err = db.Exec(query, req.Title, req.Subtitle, previewImgName, postImgName, req.Author, authorImgName, req.PublishDate, req.Content)

	return err

}

func makeImg(string imgName, string imgContent) (string, err) {
	decodedAuthorIMG, err := base64.StdEncoding.DecodeString(imgContent)
	if err != nil {
		log.Println(err)
		return "", err
	}

	fileAuthorIMG, err := os.Create("static/img/" + imgName)
	if err != nil {
		log.Println(err)
		return "", err
	}

	_, err = fileAuthorIMG.Write(decodedAuthorIMG)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return "static/img/" + imgName, err
}