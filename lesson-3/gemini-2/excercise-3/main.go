package main

func Filter[T any](in []T, keep func(T) bool) []T {
	result := make([]T, 0, len(in))

	for _, v := range in {

		if keep(v) {
			result = append(result, v)
		}
	}
	return result

}
func GroupBy[T any, K comparable](in []T, key func(T) K) map[K][]T {

	result := make(map[K][]T)

	for _, v := range in {
		k := key(v)
		result[k] = append(result[k], v)
	}

	return result
}
