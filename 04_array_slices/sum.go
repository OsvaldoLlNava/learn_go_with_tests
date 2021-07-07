package main

func Sum(numeros []int) int {
	var sum int
	for _, v := range numeros {
		sum += v
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	var resultados []int

	for _, v := range slices {

		resultados = append(resultados, Sum(v))
	}

	return resultados
}
