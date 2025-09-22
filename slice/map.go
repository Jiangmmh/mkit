package slice

/*
Map-Reduce 是一种编程模型，用于处理和生成大数据集。它包含两个主要操作：

Map 操作
功能：对集合中的每个元素应用一个函数，生成新的元素
特点：一对一映射，输入一个元素，输出一个元素
用途：数据转换、提取、计算等

Reduce 操作
功能：对集合中的元素进行归约，生成一个单一的值
特点：聚合操作，输入多个元素，输出一个元素
用途：求和、求平均、统计等
*/

// Map 对切片中的每个元素应用函数 f，返回新的切片
// 原切片不会被修改
func Map[T, U any](src []T, f func(T) U) []U { // 经过map后数据类型可能从T变成U
	if len(src) == 0 {
		return make([]U, 0)
	}

	result := make([]U, len(src))
	for i, v := range src {
		result[i] = f(v)
	}
	return result
}

// MapWithIndex 对切片中的每个元素应用函数 f，函数接收索引和值
// 原切片不会被修改
func MapWithIndex[T, U any](src []T, f func(int, T) U) []U {
	if len(src) == 0 {
		return make([]U, 0)
	}

	result := make([]U, len(src))
	for i, v := range src {
		result[i] = f(i, v)
	}
	return result
}
