package main

import (
	"errors"
	"fmt"
	"log"
)

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func GetMinMax(flights []Flight) (int, int, error) {
	// Implement me :)
	min, max := 0, 0
	if len(flights) > 0 {
		min, max = flights[0].Price, flights[0].Price
		fmt.Println(min,max)
		for _, v := range flights {
			if v.Price < min {
				min = v.Price
			} else if v.Price > max {
				max = v.Price
			}
		}
	} else {
		return min, max, errors.New("Flight is of Empty size")
	}
	return min, max, nil
}

func main() {
	fmt.Println("Getting the Minimum and Maximum Flight Prices")
	fl := []Flight{
		Flight{Origin: "India", Destination: "Texas", Price: 1000},
		Flight{Origin: "Texas", Destination: "Lewisville", Price: 5000},
		Flight{Origin: "Lewisville", Destination: "Highland", Price: 3000},
	}
	min, max, err := GetMinMax(fl)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("min,max:", min, max)
}
