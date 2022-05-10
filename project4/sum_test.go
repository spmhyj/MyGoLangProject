// sum_test.go
package main

import (
	"reflect"
	"testing"
)

func Test_sum(t *testing.T) {
	t.Helper()

	t.Run("using sumArrayEle", func(t *testing.T) {
		intArray := []int{1, 2, 3, 4, 5}
		got := sumArrayEle(intArray)
		want := 15
		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}
	})

	t.Run("using sumArrayEle2 with range", func(t *testing.T) {
		intArray := []int{2, 4, 6, 8}
		got := sumArrayEle2(intArray)
		want := 20
		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}
	})
}

func Test_sumall(t *testing.T) {
	t.Helper()
	t.Run("Return a new slice with each element is sum of a slice ",
		func(t *testing.T) {
			s1 := []int{1, 3, 5}
			s2 := []int{2, 4, 6}
			want := []int{9, 12}
			got := sumAll(s1, s2)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got:%+v, want: %+v", got, want)
			}
		})

	t.Run("Return a new slice with empty slick ",
		func(t *testing.T) {
			s1 := []int{1, 3, 5}
			s2 := []int{}
			want := []int{9, 0}
			got := sumAll(s1, s2)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got:%+v, want: %+v", got, want)
			}
		})
}
