package main

import (
	"fmt"
	"sort"
)

// Flight - a struct that
// contains information about flights
type Flight struct {
	Origin      string
	Destination string
	Price       int
}
type FlightSlice []Flight

func (f FlightSlice) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f FlightSlice) Less(i, j int) bool { return f[i].Price < f[j].Price }
func (f FlightSlice) Len() int           { return len(f) }

// SortByPrice sorts flights from highest to lowest
func SortByPrice(flights FlightSlice) []Flight {
	// implement
	//newFlights := flights
	sort.Sort(sort.Reverse(flights))
	return flights
}

func printFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Println(fmt.Sprintf("Origin: %s, Destination: %s, Price: %d", flight.Origin, flight.Destination, flight.Price))
	}
}

func main() {
	// an empty slice of flights
	flights := []Flight{Flight{Origin: "StationA", Destination: "StationB", Price: 2000},
		Flight{Origin: "StationC", Destination: "StationD", Price: 5000},
		Flight{Origin: "StationE", Destination: "Texas", Price: 1000}}

	sortedList := SortByPrice(flights)
	printFlights(sortedList)
}
