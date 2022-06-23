package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var ref []int
	for k := range input {
		ref = append(ref, k)
	}
	sort.Ints(ref)

	var eys []string
	for _, k := range ref {
		eys = append(eys, input[k])
	}
	return eys
}

/*package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[int]string{
		0:   "aa",
		8:   "bb",
		500: "cc"}
	grim := sortMapValues(m)
	fmt.Println(grim)
}

func sortMapValues(input map[int]string) (result []string) {
	var ref []int
	for k := range input {
		ref = append(ref, k)
	}
	sort.Ints(ref)

	var eys []string
	for _, k := range ref {
	//	fmt.Println(k, input[k])
		eys = append(eys, input[k])
	}
	return eys
}*/
