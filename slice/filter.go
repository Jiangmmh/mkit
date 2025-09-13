package slice

func FilterFunc[T any](src []T, filter func(T) bool) ([]T, error) {
	dst := make([]T, 0, len(src))
	for _, v := range src {
		if filter(v) {
			dst = append(dst, v)
		}
	}
	return dst, nil
}

// func Filter[T any](src []T, predicate func(T) bool) []T {
// 	if len(src) == 0 {
// 		return make([]T, 0)
// 	}

// 	// 预分配容量，避免多次扩容
// 	result := make([]T, 0, len(src))
// 	for _, v := range src {
// 		if predicate(v) {
// 			result = append(result, v)
// 		}
// 	}
// 	return result
// }

// // FilterWithIndex 根据条件函数过滤切片中的元素，条件函数接收索引和值
// // 原切片不会被修改
// func FilterWithIndex[T any](src []T, predicate func(int, T) bool) []T {
// 	if len(src) == 0 {
// 		return make([]T, 0)
// 	}

// 	// 预分配容量，避免多次扩容
// 	result := make([]T, 0, len(src))
// 	for i, v := range src {
// 		if predicate(i, v) {
// 			result = append(result, v)
// 		}
// 	}
// 	return result
// }
