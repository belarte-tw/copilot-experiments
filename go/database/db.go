package database

// package database contains functions to access the database
// and retrieve information about actors and movies from db.json

// embed db.json into a byte slice
// use go embed

import (
	_ "embed"
	"encoding/json"
	"errors"
)

//go:embed db.json
var dbjson []byte

type DB struct {
	Actors      []Actor       `json:"actors"`
	Movies      []Movie       `json:"movies"`
	Filmography []Filmography `json:"filmography"`
}

type Actor struct {
	ID   string `json:"imdb_id"`
	Name string `json:"name"`
}

type Movie struct {
	ID    string `json:"imdb_id"`
	Title string `json:"title"`
}

type Filmography struct {
	Actor string `json:"actor"`
	Movie string `json:"movie"`
}

// implement the New function
// it should return a DB struct with the data from db.json
// use the dbjson byte slice
func New() DB {
	//unmarshal dbjson into db
	var db DB
	json.Unmarshal(dbjson, &db)
	return db
}

// implement the ActorFromId function
// it should return an Actor struct with the id and name of the actor
// and an error if the actor is not found
// use the db to retrieve the data
func (db *DB) ActorFromId(id string) (Actor, error) {
	for _, a := range db.Actors {
		if a.ID == id {
			return a, nil
		}
	}

	return Actor{}, errors.New("actor not found:" + id)
}

// implement the MovieFromId function
// it should return a Movie struct with the id and title of the movie
// and an error if the movie is not found
// use the db to retrieve the data
func (db *DB) MovieFromId(id string) (Movie, error) {
	for _, m := range db.Movies {
		if m.ID == id {
			return m, nil
		}
	}

	return Movie{}, errors.New("movie not found:" + id)
}

// implement the ActorsIdsFromMovieId function
// it should return a slice of strings with the ids of the actors
// and an error if the movie is not found
// use the db to retrieve the data
func (db *DB) ActorsIdsFromMovieId(id string) ([]string, error) {
	if _, err := db.MovieFromId(id); err != nil {
		return nil, err
	}

	var actors []string
	for _, f := range db.Filmography {
		if f.Movie == id {
			actors = append(actors, f.Actor)
		}
	}
	return actors, nil
}

func (db *DB) MoviesIdsFromActorId(id string) ([]string, error) {
	if _, err := db.ActorFromId(id); err != nil {
		return nil, err
	}

	movies := []string{}
	for _, f := range db.Filmography {
		if f.Actor == id {
			movies = append(movies, f.Movie)
		}
	}
	return movies, nil
}
