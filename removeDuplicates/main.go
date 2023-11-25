package main

import "fmt"

type Developer struct {
	Name string
	Age  int
}

func FilterUnique(developers []Developer) []string {
	// TODO Implement

	output := []string{}
	mp := make(map[string]int)
	for _, k := range developers {
		_, exists := mp[k.Name]
		if !exists {
			mp[k.Name] = 1
			output = append(output, k.Name)
		}
	}
	return output
}

func main() {
	fmt.Println("Filter Unique Challenge")
	developers := []Developer{
		Developer{Name: "Elliot"},
		Developer{Name: "Alan"},
		Developer{Name: "Jennifer"},
		Developer{Name: "Graham"},
		Developer{Name: "Paul"},
		Developer{Name: "Alan"},
	}
	FilterUnique(developers)
}
