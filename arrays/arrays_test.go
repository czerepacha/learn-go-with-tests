package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	checkSum := func(t testing.TB, numbers []int, got, want int) {
		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		checkSum(t, numbers, got, want)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		checkSum(t, numbers, got, want)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
	want := []int{6, 15}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	}

	t.Run("make tails sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
		want := []int{5, 14}
		checkSums(t, got, want)
	})

	t.Run("safely add empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}
		checkSums(t, got, want)
	})

}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 2, 3})
	}
}
