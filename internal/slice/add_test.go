package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 使用 go test -v . 进行测试

// TestAdd 针对 Add 泛型函数的单元测试，覆盖了头部、中间、尾部插入，越界和空切片等情况。
// Add 是一个泛型函数，支持多种类型。
func TestAdd(t *testing.T) {
	type testCase[T any] struct {
		name      string
		slice     []T
		index     int
		val       T
		wantSlice []T
		wantErr   bool
	}

	// int 类型测试
	testsInt := []testCase[int]{
		{
			name:      "在头部插入",
			slice:     []int{1, 2, 3},
			index:     0,
			val:       100,
			wantSlice: []int{100, 1, 2, 3},
			wantErr:   false,
		},
		{
			name:      "在中间插入",
			slice:     []int{1, 2, 3},
			index:     1,
			val:       200,
			wantSlice: []int{1, 200, 2, 3},
			wantErr:   false,
		},
		{
			name:      "在尾部插入",
			slice:     []int{1, 2, 3},
			index:     3,
			val:       300,
			wantSlice: []int{1, 2, 3, 300},
			wantErr:   false,
		},
		{
			name:      "下标越界(负数)",
			slice:     []int{1, 2, 3},
			index:     -1,
			val:       400,
			wantSlice: []int{1, 2, 3},
			wantErr:   true,
		},
		{
			name:      "下标越界(大于len)",
			slice:     []int{1, 2, 3},
			index:     4,
			val:       500,
			wantSlice: []int{1, 2, 3},
			wantErr:   true,
		},
		{
			name:      "空切片插入",
			slice:     []int{},
			index:     0,
			val:       600,
			wantSlice: []int{600},
			wantErr:   false,
		},
	}

	for _, tc := range testsInt {
		t.Run("int-"+tc.name, func(t *testing.T) {
			res, err := Add(tc.slice, tc.val, tc.index)
			if tc.wantErr {
				assert.Error(t, err, "期望出现错误，但未出现")
				assert.Equal(t, tc.slice, res, "错误情况下切片应保持不变")
			} else {
				assert.NoError(t, err, "期望无错误，但出现了错误")
				assert.Equal(t, tc.wantSlice, res, "插入结果与预期不符")
			}
		})
	}

	// string 类型测试
	testsString := []testCase[string]{
		{
			name:      "string类型插入",
			slice:     []string{"a", "b", "c"},
			index:     2,
			val:       "x",
			wantSlice: []string{"a", "b", "x", "c"},
			wantErr:   false,
		},
		{
			name:      "string空切片插入",
			slice:     []string{},
			index:     0,
			val:       "z",
			wantSlice: []string{"z"},
			wantErr:   false,
		},
	}

	for _, tc := range testsString {
		t.Run("string-"+tc.name, func(t *testing.T) {
			res, err := Add(tc.slice, tc.val, tc.index)
			if tc.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tc.slice, res)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.wantSlice, res)
			}
		})
	}

	// float64 类型测试
	testsFloat := []testCase[float64]{
		{
			name:      "float64类型插入",
			slice:     []float64{1.1, 2.2},
			index:     1,
			val:       3.3,
			wantSlice: []float64{1.1, 3.3, 2.2},
			wantErr:   false,
		},
		{
			name:      "float64空切片插入",
			slice:     []float64{},
			index:     0,
			val:       4.4,
			wantSlice: []float64{4.4},
			wantErr:   false,
		},
	}

	for _, tc := range testsFloat {
		t.Run("float64-"+tc.name, func(t *testing.T) {
			res, err := Add(tc.slice, tc.val, tc.index)
			if tc.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tc.slice, res)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.wantSlice, res)
			}
		})
	}
}
