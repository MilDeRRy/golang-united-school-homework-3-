package homework

func reverse(input [4]int64) (result []int64) {
	var resul []int64
	for i := len(input) - 1; i >= 0; i-- {
		t := (input[i])
		if i >= 0 {
			resul = append(resul, t)
		}
	}
	return resul
}

/*package main

import "fmt"

func main() {
	input := []int64{1, 2, 5, 15, 16, 17}
	grim := reverse(input)
	fmt.Println(grim)
}

func reverse(input []int64) (result []int64) {
	var resul []int64
	for i := len(input) - 1; i >= 0; i-- {
		t := (input[i])
		if i >= 0 {
			resul = append(resul, t)
		}
	}
	return resul
}*/
