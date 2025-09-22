package slice

// 缩容标准：
//  1. cap小于等于64，不缩容
//  2. cap小于等于2048且空闲超过3/4，缩容一半
//  3. cap大于2048且空闲超过一半，缩容到0.625
func calCapacity(capacity, length int) (int, bool) {
	if capacity <= 64 {
		return capacity, false
	}

	if capacity > 2048 && (capacity/length >= 2) {
		factor := 0.625
		return int(float32(capacity) * float32(factor)), true
	}

	if capacity <= 2048 && (capacity/length >= 4) {
		return capacity / 2, true
	}

	return capacity, false
}

// 切片缩容
func Shrink[T any](src []T) []T {
	c, l := cap(src), len(src)
	if l == 0 {
		return src
	}

	n, isShrink := calCapacity(c, l)
	if isShrink { // 需要缩容
		s := make([]T, l, n)
		copy(s, src)
		return s
	}
	return src
}
