package homework

func average(input [15]float32) (result float32) {
	var res float32 = 0
	for _, v := range input {
		res += v
	}
	res = res / 15
	return res
}

/*package main

import "fmt"

func main() {
	results := [15]float32{1, 2, 3, 4, 5, 6}
	grim := average(results)
	fmt.Println(grim)
	//fmt.Println(len(results))
	//fmt.Println(results)
}

func average(input [15]float32) (result float32) {
	var res float32 = 0
	//var uid float32
	for _, v := range input {
		res += v
	}
	//(len(input))
	res = res / 15
	return res
}*/
