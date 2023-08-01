package main

import "fmt"

func location(city string) (region string, continent string) {
	//var region string
	//var continent string
	switch city {
	case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
		region, continent = "India", "Asia"
	case "Lafayette", "Louisville", "Boulder":
		region, continent = "Colorado", "USA"
	case "Irvine", "Los Angeles", "San Diego":
		region, continent = "California", "USA"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return
}

func main1() {
	region, continent := location("Irvine")
	fmt.Printf("John works in %s, %s\n", region, continent)
	a, b := location("china")
	fmt.Printf("Alex live in %s, %s\n", a, b)
}
