package utils

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func Max[T integer](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T integer](a, b T) T {
	if a > b {
		return b
	}
	return a
}
