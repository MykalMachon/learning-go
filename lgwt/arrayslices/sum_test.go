package arrayslices

import (
	"reflect"
	"testing"
	// can use this to compare slices with slices.Equal
)

// run coverage with go test -cover

func TestSum(t *testing.T) {
	// could also do the following
	// numbers := [...]int{1,2,3,4,5}
	t.Run("Collection of 5 numbers with int type", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// go doesn't allow you to use == != to compare slices
	// you have to use reflect.DeepEqual()
	// can use slices.Equal()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
