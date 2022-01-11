package ggnr

import (
	"testing"
)

var (
	m1 = map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}
	m2 = map[int64]bool{
		1: true,
		2: true,
		3: false,
		4: true,
	}
)

func Test_Keys(t *testing.T) {
	keys1 := Keys(m1)
	if !ArrValuesEqual(keys1, []string{"a", "b", "c"}) {
		t.Error(keys1)
		t.Fail()
	}
	keys2 := Keys(m2)
	if !ArrValuesEqual(keys2, []int64{1, 2, 3, 4}) {
		t.Error(keys2)
		t.Fail()
	}
}

func Test_Values(t *testing.T) {
	values1 := Values(m1)
	if !ArrValuesEqual(values1, []string{"1", "2", "3"}) {
		t.Error(values1)
		t.Fail()
	}

	values2 := Values(m2)
	if !ArrValuesEqual(values2, []bool{true, true, false, true}) {
		t.Error(values2)
		t.Fail()
	}
}
