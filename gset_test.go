package ggnr

import (
	"testing"
)

func Test_Gset(t *testing.T) {
	t.Helper()
	s := NewGset(1, 1, 2, 3, 2)
	if !ArrValuesEqual(s.Values(), []int{1, 2, 3}) {
		t.Fail()
	}
	if !s.Has(1) {
		t.Fail()
	}
	if s.Size() != 3 {
		t.Fail()
	}
	s.Add(4)
	if !s.Has(4) {
		t.Fail()
	}
	if s.Size() != 4 {
		t.Fail()
	}
	s.Delete(1)
	if s.Has(1) {
		t.Fail()
	}
	if s.Size() != 3 {
		t.Fail()
	}
	s.Clear()
	if s.Size() != 0 {
		t.Fail()
	}
}
