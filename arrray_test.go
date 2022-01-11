package ggnr

import (
	"strconv"
	"testing"
)

func Test_Equal(t *testing.T) {
	if ArrEqual([]string{"1", "2", "3"}, []string{"1", "2"}) {
		t.Fail()
	}

	if ArrEqual([]string{"1", "2", "3"}, []string{"1", "3", "2"}) {
		t.Fail()
	}

	if !ArrEqual([]string{"1", "2", "3"}, []string{"1", "2", "3"}) {
		t.Fail()
	}

	if !ArrEqual([]int{1, 2, 3}, []int{1, 2, 3}) {
		t.Fail()
	}

	if !ArrEqual([]float64{1.1, 2.2, 3.3}, []float64{1.1, 2.2, 3.3}) {
		t.Fail()
	}
}

func Test_ArrGen(t *testing.T) {
	if !ArrEqual(ArrGen("a", 3), []string{"a", "a", "a"}) {
		t.Fail()
	}
}

func Test_Map(t *testing.T) {
	if !ArrEqual(Map([]string{"1", "2", "3"}, func(v string, _ int) int {
		n, _ := strconv.Atoi(v)
		return n
	}), []int{1, 2, 3}) {
		t.Fail()
	}
}

func Test_Sort(t *testing.T) {
	a := []int{1, 5, 3}
	SortAsc(a)
	if !ArrEqual(a, []int{1, 3, 5}) {
		t.Fail()
	}

	SortDsc(a)
	if !ArrEqual(a, []int{5, 3, 1}) {
		t.Fail()
	}

	Sort(a, func(a, b int) bool { return a < b })
	if !ArrEqual(a, []int{1, 3, 5}) {
		t.Fail()
	}
}
