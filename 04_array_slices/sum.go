package main

func Sum(numeros []int) int {
	var sum int
	for _, v := range numeros {
		sum += v
	}
	return sum
}
