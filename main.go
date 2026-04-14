package main

import (
	"database/sql"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/merkurtran/goblog/app/http/middlewares"
	"github.com/merkurtran/goblog/bootstrap"
	"github.com/merkurtran/goblog/pkg/database"
	"github.com/merkurtran/goblog/pkg/logger"
)

var router *mux.Router
var db *sql.DB

type ArticlesFormData struct {
	Title, Body string
	URL         *url.URL
	Errors      map[string]string
}

type Article struct {
	Title, Body string
	ID          int64
}

func RouteName2URL(routename string, pairs ...string) string {
	url, err := router.Get(routename).URL(pairs...)
	if err != nil {
		logger.LogError(err)
		return ""
	}
	return url.String()
}

func saveArticleToDB(title string, body string) (int64, error) {
	var (
		id   int64
		err  error
		rs   sql.Result
		stmt *sql.Stmt
	)
	stmt, err = db.Prepare("INSERT INTO articles (title, body) VALUE(?,?)")
	if err != nil {
		return 0, nil
	}

	defer stmt.Close()

	rs, err = stmt.Exec(title, body)
	if err != nil {
		return 0, nil
	}
	if id, err = rs.LastInsertId(); id > 0 {
		return id, nil
	}
	return 0, err
}

func getArticleByID(id string) (Article, error) {
	article := Article{}
	query := "SELECT * FROM articles WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Body)
	return article, err
}

func (a Article) Delete() (rowsAffected int64, err error) {
	rs, err := db.Exec("DELETE FROM articles WHERE id = " + strconv.FormatInt(a.ID, 10))
	if err != nil {
		return 0, nil
	}
	if n, _ := rs.RowsAffected(); n > 0 {
		return n, nil
	}
	return 0, nil
}

func main() {
	database.Initialize()
	db = database.DB

	bootstrap.SetupDB()
	router = bootstrap.SetRoute()

	err := http.ListenAndServe(":3002", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
