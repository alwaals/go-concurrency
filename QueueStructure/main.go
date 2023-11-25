package main

import (
	"errors"
	"fmt"
)

type Flight struct {
	Origin      string
	Destination string
	Price       int
}
type Queue struct {
	FlightList []Flight
}

func (s *Queue) Push(s1 Flight) {
	s.FlightList = append(s.FlightList, s1)
}
func (s *Queue) Pop() (Flight, error) {
	if s.IsEmpty() {
		return Flight{}, errors.New("Flight list is empty, cannot perform pop operation")
	} else {
		firstEle := s.FlightList[0]
		s.FlightList = s.FlightList[1:]
		return firstEle, nil
	}
}
func (s *Queue) Peek() (Flight, error) {
	if s.IsEmpty() {
		return Flight{}, errors.New("Flight list is empty, cannot perform Peek operation")
	} else {
		return s.FlightList[0], nil
	}
}
func (s *Queue) IsEmpty() bool {
	return len(s.FlightList) == 0
}
func main() {
	flights:=[]Flight{
		Flight{Origin: "India",Destination: "Texas",Price: 13000},
		Flight{Origin: "Texas",Destination: "Dallas",Price: 3000},
		Flight{Origin: "Dallas",Destination: "Lewisville",Price: 1000},
	}
	q:=Queue{FlightList: flights}
	fl,_:=q.Peek()
	fmt.Println(fl)
	q.Push(Flight{Origin: "Lewisville",Destination: "Highland",Price: 900})
	fmt.Println(q.FlightList)
	q.Pop()
	fmt.Println("After Pop:",q.FlightList)
}
