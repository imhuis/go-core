package main

import (
	"fmt"
	json2 "go-core/complex-type/json-type"
	"log"
	"os"
)

var movies = []json2.Movie{
	{
		Title:  "Casablanca",
		Year:   1942,
		Color:  true,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
	},
	{
		Title:  "Cool Hand Luke",
		Year:   1967,
		Actors: []string{"Paul Newman"},
	},
}

func main() {
	//data1, err := json-type.Marshal(movies)
	//if err != nil {
	//	log.Fatalf("Could not marshal: %s", err)
	//}
	//
	//data2, err := json-type.MarshalIndent(movies, "", "    ")
	//fmt.Printf("%s\n", data1)
	//fmt.Printf("%s\n", data2)

	// repo:golang/go is:open json decoder
	SearchIssues()
}

func SearchIssues() {
	result, err := json2.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
