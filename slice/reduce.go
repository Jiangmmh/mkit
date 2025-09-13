package slice

// Reduce 将切片中的元素通过函数 f 聚合为单个值
// initial 是初始值，f 是聚合函数，接收累积值和当前元素，返回新的累积值
func Reduce[T, U any](src []T, initial U, f func(U, T) U) U {
	result := initial
	for _, v := range src {
		result = f(result, v)
	}
	return result
}

// ReduceWithIndex 将切片中的元素通过函数 f 聚合为单个值，函数接收索引
// initial 是初始值，f 是聚合函数，接收累积值、索引和当前元素，返回新的累积值
func ReduceWithIndex[T, U any](src []T, initial U, f func(U, int, T) U) U {
	result := initial
	for i, v := range src {
		result = f(result, i, v)
	}
	return result
}
