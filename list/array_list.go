package list

import (
	"mkit/internal/errs"
	"mkit/internal/slice"
)

type ArrayList[T any] struct {
	elems []T
}

// 这里使用 make([]T, 0) 而不是 []T{}，主要是为了明确地分配一个新的 slice 底层数组，
// 保证每次创建的 ArrayList 都有独立的底层存储，避免潜在的内存重用问题。
// 在高性能场景下，make 也可以方便地指定初始容量（如 make([]T, 0, cap)），便于后续扩容优化
func NewArrayList[T any]() *ArrayList[T] {
	return &ArrayList[T]{
		elems: make([]T, 0),
	}
}

// Get 返回对应下标的元素，在下标超出范围的情况下，返回错误
func (l *ArrayList[T]) Get(index int) (T, error) {
	if index < 0 || index >= l.Len() {
		var zeroValue T
		return zeroValue, errs.NewErrIndexOutOfRange(l.Len(), index)
	}
	return l.elems[index], nil
}

// Append 在末尾追加元素
func (l *ArrayList[T]) Append(ts ...T) error {
	l.elems = append(l.elems, ts...)
	return nil
}

// Add 在特定下标处增加一个新元素
// 如果下标不在[0, Len()]范围之内
// 应该返回错误
// 如果index == Len()则表示往List末端增加一个值
func (l *ArrayList[T]) Add(index int, t T) error {
	newElems, err := slice.Add(l.elems, t, index)
	if err != nil {
		return err
	}
	l.elems = newElems
	return nil
}

// Set 重置 index 位置的值
// 如果下标超出范围，应该返回错误
func (l *ArrayList[T]) Set(index int, t T) error {
	if index < 0 || index >= l.Len() {
		return errs.NewErrIndexOutOfRange(l.Len(), index)
	}
	l.elems[index] = t
	return nil
}

// Delete 删除目标元素的位置，并且返回该位置的值,如有必要会进行缩容
// - 如果容量 > 2048，并且长度小于容量一半，那么就会缩容为原本的 5/8
// - 如果容量 (64, 2048]，如果长度是容量的 1/4，那么就会缩容为原本的一半
// - 如果此时容量 <= 64，那么我们将不会执行缩容。在容量很小的情况下，浪费的内存很少，所以没必要消耗 CPU去执行缩容
func (l *ArrayList[T]) Delete(index int) (T, error) {
	res, t, err := slice.Delete(l.elems, index)
	if err != nil {
		var zeroValue T
		return zeroValue, err
	}
	l.elems = res
	l.shrink()
	return t, nil
}

func (l *ArrayList[T]) shrink() {
	l.elems = slice.Shrink(l.elems)
}

// Len 返回长度
func (l *ArrayList[T]) Len() int {
	return len(l.elems)
}

// Cap 返回容量
func (l *ArrayList[T]) Cap() int {
	return cap(l.elems)
}

// Range 遍历 List 的所有元素
func (l *ArrayList[T]) Range(fn func(index int, t T) error) error {
	for i, v := range l.elems {
		if err := fn(i, v); err != nil {
			return err
		}
	}
	return nil
}

// AsSlice 将 List 转化为一个切片
// 不允许返回nil，在没有元素的情况下，
// 必须返回一个长度和容量都为 0 的切片
// AsSlice 每次调用都必须返回一个全新的切片
func (l *ArrayList[T]) AsSlice() []T {
	if len(l.elems) == 0 {
		return make([]T, 0)
	}
	result := make([]T, len(l.elems))
	copy(result, l.elems)
	return result
}
