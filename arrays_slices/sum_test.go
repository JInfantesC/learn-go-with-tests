package arrays_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	esCorrecto := func(t *testing.T, got, expected int, given []int) {
		t.Helper()
		if got != expected {
			t.Errorf("expected %d got %d, given %v", expected, got, given)
		}
	}
	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		expected := 6
		esCorrecto(t, got, expected, numbers)
	})

}
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 6})
	expected := []int{6, 10}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v got %v", expected, got)
	}

}
func TestSumAllTails(t *testing.T) {
    esCorrecto := func(t *testing.T, got, expected []int ) {
		t.Helper()
        if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v got %v", expected, got)
		}
	}

	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 6}, []int{2}, []int{2, 2})
		expected := []int{5, 6, 0, 2}
		esCorrecto(t,got,expected)
	})
	t.Run("Sum of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4, 6})
		expected := []int{0,6}
		esCorrecto(t,got,expected)
	})
}
func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}
func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2, 3}, []int{4, 6})
	}
}
