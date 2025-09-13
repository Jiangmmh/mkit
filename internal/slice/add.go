package slice

import "mkit/internal/errs"

func Add[T any](src []T, elem T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index > length {
		return src, errs.NewErrIndexOutOfRange(length, index)
	}

	// 先将src扩展一个元素
	var placeholder T
	src = append(src, placeholder)
	for i := len(src) - 1; i > index; i-- {
		if i-1 >= 0 {
			src[i] = src[i-1]
		}
	}

	src[index] = elem
	return src, nil
}
