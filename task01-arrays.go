package homework

func average(input [15]float32) (result float32) {
	for _, v := range input {
		result += v
	}
	result = result / 6
	return result
}

/*package main

import "fmt"

func main() {
	results := []float64{1, 2, 3, 4, 5, 6}
	grim := score(results)
	fmt.Println(grim)
}

func score(g []float64) float64 {
	var res float64 = 0
	for _, v := range g {
		res += v
	}
	res = res / 6
	return res
}*/
