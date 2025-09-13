package slice

func filterFunc[T any](src []T, filter func(T) bool) ([]T, error) {
	dst := make([]T, 0, len(src))
	for _, v := range src {
		if filter(v) {
			dst = append(dst, v)
		}
	}
	return dst, nil
}
