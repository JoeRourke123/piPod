package util

func Map[A any, B any](vs []A, f func(A) B) []B {
	vsm := make([]B, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
