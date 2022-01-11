package ggnr

import (
	"testing"
)

func Test_Gmap(t* testing.T) {
	t.Helper()
	m := Gmap[int, int]{
		1: 3,
		2: 4,
	}
	
	keys := []int{1, 2}
	if !ArrValuesEqual(keys, m.Keys()) {
		t.Fail()
	}

	values := []int{3, 4}
	if !ArrValuesEqual(values, m.Values()) {
		t.Fail()
	}
}