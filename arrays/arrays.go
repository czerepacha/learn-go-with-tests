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

func SumAllTails(s ...[]int) (tails []int) {
	for _, num := range s {
		if len(num) == 0 {
			tails = append(tails, 0)
		} else {
			tails = append(tails, Sum(num[1:]))
		}
	}
	return
}
