package ggnr

import (
	"constraints"
	"sort"
)

func ArrEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func ArrValuesEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	countMap := map[T]int{}
	for _, v := range a {
		countMap[v] += 1
	}
	for _, v := range b {
		countMap[v] -= 1
		if countMap[v] < 0 {
			return false
		}
	}
	return true
}

func ArrGen[T any, N constraints.Integer](v T, n N) []T {
	a := make([]T, int(n))
	for i := 0; i < int(n); i++ {
		a[i] = v
	}
	return a
}

func Map[T1 any, T2 any](a []T1, f (func(v T1, i int) T2)) []T2 {
	a2 := make([]T2, len(a))
	for i, v := range a {
		a2[i] = f(v, i)
	}
	return a2
}

func Sort[T any](a []T, less (func(v1, v2 T) bool)) {
	sort.Slice(a, func(i, j int) bool {
		return less(a[i], a[j])
	})
}

func SortAsc[T constraints.Ordered | string](a []T) {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
}

func SortDsc[T constraints.Ordered | string](a []T) {
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
}