/*
Package gofunc provides functional primitives.
*/
package gofunc

// Filter returns a new slice of only the items that pass predicate.
func Filter[T any](items []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range items {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map returns a new slice of items transformed by mapper.
func Map[T any, U any](items []T, mapper func(T) U) []U {
	result := make([]U, len(items))
	for i, v := range items {
		result[i] = mapper(v)
	}
	return result
}

// FlatMap returns a new slice of items mapped and flattened into a single dimensional slice.
func FlatMap[T any, U any](items [][]T, mapper func(T) U) []U {
	size := 0
	for _, l := range items {
		size += len(l)
	}
	result := make([]U, size)
	i := 0
	for _, l := range items {
		for _, v := range l {
			result[i] = mapper(v)
			i++
		}
	}
	return result
}

// Reduce combines items using a combiner function reducing them to a single value.
// Each item is combined with the result of the combiner function operating on the
// prior item. A first parameter allows the initial value for the combiner to be
// specified.
func Reduce[T any, U any](items []T, combiner func(a U, b T) U, first U) U {
	for _, v := range items {
		first = combiner(first, v)
	}
	return first
}
