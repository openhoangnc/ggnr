package ggnr

type Gset[K comparable] struct {
	m map[K]bool
}

func NewGset[K comparable](a ...K) Gset[K] {
	m := map[K]bool{}
	for _, k := range a {
		m[k] = true
	}
	return Gset[K]{m}
}

func(s * Gset[K])Has(v K) bool {
	return s.m[v]
}

func(s * Gset[K])Add(v ...K) {
	for _, k := range v {
		s.m[k] = true
	}
}

func(s * Gset[K])Delete(v ...K) {
	for _, k := range v {
		delete(s.m, k)
	}
}

func(s * Gset[K])Clear() {
	s.m = map[K]bool{}
}

func(s * Gset[K])Size() int {
	return len(s.m)
}

func(s * Gset[K])Values() []K {
	return Keys(s.m)
}