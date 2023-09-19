package main

import "gopkg.in/mgo.v2/bson"

// Movie represents a movie entity
type Movie struct {
    ID       bson.ObjectId `bson:"_id"`
    Name     string        `bson:"name"`
    Year     int           `bson:"year"`
    Director string        `bson:"director"`
    Actors   []string      `bson:"actors"`
}

// Implement the fetchMoviesFromDB function to fetch movies from MongoDB

// Alternatively, you can read movies data from a JSON file:
//
func fetchMoviesFromDB() []Movie {
    data, err := ioutil.ReadFile("db/movies.json")
    if err != nil {
        // Handle error
    }

    var movies []Movie
    if err := json.Unmarshal(data, &movies); err != nil {
        // Handle error
    }

    return movies
}
