package main

import "fmt"

func main() {

	superHero := map[string]map[string] string {

		"superMan" : map[string] string {
		"realName" : "Clark Kent",
		"city" : "Metrpolis",

		},

		"batMan" : map[string] string {
		"realName" : "Bruce Wayne",
		"city" : "Gotham City",
	
		},
	}

	if temp, hero := superHero["batMan"]; hero {
		fmt.Println(temp["realName"], temp["city"])
	}

}