package arrays

func Sum(n []int) (sum int) {
	for _, n := range n {
		sum += n
	}
	return
}

func SumAll(s ...[]int) (sums []int) {
	for _, num := range s {
		sums = append(sums, Sum(num))
	}
	return
}

