package iteration

const repeatCount = 5

func Repeat(s string, n int) (r string) {
	for i := 0; i < n; i++ {
		r += s
	}
	return
}

