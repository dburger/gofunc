package functional

func filter[T any](slice []T, pred func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if pred(v) {
			result = append(result, v)
		}
	}
	return result
}

func transform[T any, U any](slice []T, mapper func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = mapper(v)
	}
	return result
}

func flatmap[T any, U any](slice [][]T, mapper func(T) U) []U {
	size := 0
	for _, l := range slice {
		size += len(l)
	}
	result := make([]U, size)
	i := 0
	for _, l := range slice {
		for _, v := range l {
			result[i] = mapper(v)
			i++
		}
	}
	return result
}

func reduce[T any, U any](slice []T, comb func(a U, b T) U, first U) U {
	for _, v := range slice {
		first = comb(first, v)
	}
	return first
}
