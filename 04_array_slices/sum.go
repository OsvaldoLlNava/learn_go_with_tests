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

func SumAllTails(slices ...[]int) []int {

	var sumas []int

	for _, v := range slices {
		if len(v) > 1 {
			sumas = append(sumas, Sum(v[1:]))
		} else {
			sumas = append(sumas, 0)
		}

	}

	return sumas
}
