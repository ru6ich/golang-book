package main

import "fmt"

type Movie struct {
	Title  string
	Year   int
	Color  bool
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humprey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"PaulNewman"}},
	}
	fmt.Println(movies)

}
