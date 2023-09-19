package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")
	http.Handle("/", r)

	// Serve static files from the "static" directory
	staticDir := "./static/"
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Load movies from MongoDB or a data source
	movies := fetchMoviesFromDB() // Implement this function

	// Render the HTML template
	tmpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Movies []Movie
	}{
		Movies: movies,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type Movie struct {
	ID       bson.ObjectId `bson:"_id"`
	Name     string        `bson:"name"`
	Year     int           `bson:"year"`
	Director string        `bson:"director"`
	Actors   []string      `bson:"actors"`
}

func fetchMoviesFromDB() []Movie {
	data, err := os.ReadFile("db/movies.json")
	if err != nil {
		// Handle error
	}

	var movies []Movie
	if err := json.Unmarshal(data, &movies); err != nil {
		// Handle error
	}

	return movies
}
