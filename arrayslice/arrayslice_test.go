package arrayslice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("Expected '%d' but got '%d'", want, got)
		}
	})

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7}

		got := Sum(numbers)
		want := 28

		if got != want {
			t.Errorf("Expected '%d' but got '%d'", want, got)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{1, 2, 3, 4})
	want := []int{6, 10}

	checkSums(t, got, want)
}

func TestSumAllTails(t *testing.T) {

	t.Run("make the sum of tail elements", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{1, 2, 3, 4})
		want := []int{5, 9}

		checkSums(t, got, want)
	})

	t.Run("call Sum All Tails with empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}

		checkSums(t, got, want)
	})
}

func checkSums(t *testing.T, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
