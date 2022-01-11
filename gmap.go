package ggnr

type Gmap[K comparable, V any] map[K]V

func (m * Gmap[K,V]) Keys() []K {
	return Keys(*m)
}

func (m * Gmap[K,V]) Values() []V {
	return Values(*m)
}