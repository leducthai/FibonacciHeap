package main

func RemoveEle[T comparable](t []T, el T) []T {
	for i, v := range t {
		if v == el {
			t = append(t[:i], t[i+1:]...)
			break
		}
	}
	return t
}

func InsertEle[T comparable](sl []T, el T) []T {
	sl = append(sl, el)
	return sl
}

func setPoiter[A any](a A) *A {
	return &a
}
