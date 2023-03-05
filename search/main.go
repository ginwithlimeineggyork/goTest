package main

import (
	"net/http"
	"os"
	"net/url"
	"log"
	"math"
	"time"
	"strconv"
	"gobl/news"
	"bytes"
	"html/template"
	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))
func (s *Search) IsLastPage() bool {
	return s.NextPage >= s.TotalPages
}

func (s *Search) CurrentPage() int {
	if s.NextPage == 1 {
		return s.NextPage
	}

	return s.NextPage - 1
}

func (s *Search) PreviousPage() int {
	return s.CurrentPage() - 1
}
var newsapi *news.Client

func indexHandler(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}

type Search struct {
	Query      string
	NextPage   int
	TotalPages int
	Results    *news.Results
}

func searchHandler(newsapi *news.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u,err := url.Parse(r.URL.String())
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
		params:=u.Query()
		searchQuery := params.Get("q")
		page := params.Get("page")
		if page == "" {
			page = "1"
		}
		results, err := newsapi.FetchEverything(searchQuery, page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		nextPage, err := strconv.Atoi(page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		search := &Search{
			Query:      searchQuery,
			NextPage:   nextPage,
			TotalPages: int(math.Ceil(float64(results.TotalResults) / float64(newsapi.PageSize))),
			Results:    results,
		}
		if ok := !search.IsLastPage(); ok {
			search.NextPage++
		}
		buf := &bytes.Buffer{}
		err = tpl.Execute(buf, search)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		buf.WriteTo(w)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == "" {
		log.Fatal("Env: apiKey must be set")
	}
	myClient := &http.Client{Timeout:10*time.Second}
	newsapi:=news.NewClient(myClient,apiKey,20)
	fs := http.FileServer(http.Dir("static"))
	mux := http.NewServeMux()
	mux.Handle("/static/",http.StripPrefix("/static/",fs))
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/search",searchHandler(newsapi))
	http.ListenAndServe(":"+port,mux)
}
