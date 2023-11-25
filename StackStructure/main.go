package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	Flights []Flight
}

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

// func (s *Stack) Pop() Flight {
// 	// TODO Implement
// 	if !s.IsEmpty() {
// 		//fmt.Println("Before:", s.Flights)
// 		poppedItem:=s.Flights[0]
// 		s.Flights = s.Flights[1:len(s.Flights)]
// 		return poppedItem
// 	}
// 	return Flight{}
// }
func (s *Stack) Pop() (Flight, error) {
	if s.IsEmpty() {
		return Flight{}, errors.New("Stack is Empty")
	} else {
		lastElemIndex := len(s.Flights) - 1
		flight := s.Flights[lastElemIndex]
		s.Flights = s.Flights[:lastElemIndex]
		return flight, nil
	}
}

func (s *Stack) Push(f Flight) {
	// TODO Implement
	if !s.IsEmpty() {
		s.Flights = append(s.Flights, f)
	}
}

func (s *Stack) Peek() (Flight,error) {
	// TODO Implement
	if s.IsEmpty() {
		return Flight{}, errors.New("Stack is Empty")
	} else {
		lastElemIndex := len(s.Flights) - 1
		flight := s.Flights[lastElemIndex]
		return flight, nil
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.Flights) == 0
}

func main() {
	fmt.Println("Go Stack Implementation")
	flights := []Flight{
		Flight{Origin: "India", Destination: "Texas", Price: 20000},
		Flight{Origin: "Texas", Destination: "Lewisville", Price: 50000},
		Flight{Origin: "Lewisville", Destination: "Highland", Price: 10000},
	}
	s := Stack{Flights: flights}
	f,_:=s.Peek()
	fmt.Println("Peek:", f)
	s.Push(Flight{Origin: "Highland", Destination: "1224", Price: 1224})
	fmt.Println("After Push:", s.Flights)
	f,_=s.Pop()
	fmt.Println("After Pop:", f)

}
