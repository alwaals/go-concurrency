package main

import (
	"fmt"
)

// Flight struct which contains
// the origin, destination and price of a flight
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

// IsSubset checks to see if the first set of
// flights is a subset of the second set of flights.
func IsSubset(first, second []Flight) bool {
	// implement
	//res:=reflect.DeepEqual(first,second)

	if len(second) != len(first) {
		return false
	}
	mp := make(map[Flight]int)
	for _, v := range second {
		if _, exists := mp[v]; exists {
			mp[v] += 1
		} else {
			mp[v] = 1
		}
	}
	fmt.Println("Before:", mp)
	for _, p := range first {
		value, exists := mp[p]
		if exists {
			if value == 1 {
				delete(mp, p)
			} else {
				mp[p]--
			}
		} else {
			return false
		}
	}
	fmt.Println("After:", mp)
	return len(mp) == 0
}

func main() {
	fmt.Println("Sets and Subsets Challenge")
	firstFlights := []Flight{
		Flight{Origin: "GLA", Destination: "CDG", Price: 1000},
		Flight{Origin: "GLA", Destination: "JFK", Price: 5000},
		Flight{Origin: "GLA", Destination: "SNG", Price: 3000},
	}

	secondFlights := []Flight{
		Flight{Origin: "GLA", Destination: "CDG", Price: 1000},
		Flight{Origin: "GLA", Destination: "JFK", Price: 5000},
		Flight{Origin: "GLA", Destination: "SNG", Price: 3000},
		Flight{Origin: "GLA", Destination: "AMS", Price: 500},
	}

	subset := IsSubset(firstFlights, secondFlights)
	fmt.Println(subset)
}
