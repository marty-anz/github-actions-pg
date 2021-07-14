package primes

import "math"

func GetPrimesUpto(number uint64) []uint64 {
	if number < 2 {
		return []uint64{}
	}

	numbers := make([]uint64, number-1)

	for i := uint64(0); i+2 <= number; i++ {
		numbers[i] = i + 2
	}

	for i := uint64(2); i <= uint64(math.Ceil(math.Sqrt(float64(number)))); i++ {
		for j, v := range numbers {
			if v != i && v%i == 0 {
				numbers[j] = 0
			}
		}

		newNumbers := []uint64{}

		for _, v := range numbers {
			if v != 0 {
				newNumbers = append(newNumbers, v)
			}
		}

		numbers = newNumbers
	}

	return numbers
}
