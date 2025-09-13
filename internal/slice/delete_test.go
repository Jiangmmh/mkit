package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDelete 针对 Delete 泛型函数的单元测试，覆盖头部、中间、尾部、越界、空切片等情况
func TestDelete(t *testing.T) {
	type testCase[T any] struct {
		name      string
		slice     []T
		index     int
		wantSlice []T
		wantVal   T
		wantErr   bool
	}

	// int 类型测试
	testsInt := []testCase[int]{
		{
			name:      "删除头部元素",
			slice:     []int{1, 2, 3},
			index:     0,
			wantSlice: []int{2, 3},
			wantVal:   1,
			wantErr:   false,
		},
		{
			name:      "删除中间元素",
			slice:     []int{1, 2, 3},
			index:     1,
			wantSlice: []int{1, 3},
			wantVal:   2,
			wantErr:   false,
		},
		{
			name:      "删除尾部元素",
			slice:     []int{1, 2, 3},
			index:     2,
			wantSlice: []int{1, 2},
			wantVal:   3,
			wantErr:   false,
		},
		{
			name:      "下标越界(负数)",
			slice:     []int{1, 2, 3},
			index:     -1,
			wantSlice: []int{1, 2, 3},
			wantVal:   0,
			wantErr:   true,
		},
		{
			name:      "下标越界(大于等于len)",
			slice:     []int{1, 2, 3},
			index:     3,
			wantSlice: []int{1, 2, 3},
			wantVal:   0,
			wantErr:   true,
		},
		{
			name:      "空切片删除",
			slice:     []int{},
			index:     0,
			wantSlice: []int{},
			wantVal:   0,
			wantErr:   true,
		},
	}

	for _, tc := range testsInt {
		t.Run("int-"+tc.name, func(t *testing.T) {
			res, val, err := Delete(tc.slice, tc.index)
			if tc.wantErr {
				assert.Error(t, err, "期望出现错误，但未出现")
				assert.Equal(t, tc.slice, res, "错误情况下切片应保持不变")
			} else {
				assert.NoError(t, err, "期望无错误，但出现了错误")
				assert.Equal(t, tc.wantSlice, res, "删除结果与预期不符")
				assert.Equal(t, tc.wantVal, val, "删除值与预期不符")
			}
		})
	}

	// string 类型测试
	testsString := []testCase[string]{
		{
			name:      "删除string中间元素",
			slice:     []string{"a", "b", "c"},
			index:     1,
			wantSlice: []string{"a", "c"},
			wantVal:   "b",
			wantErr:   false,
		},
		{
			name:      "string空切片删除",
			slice:     []string{},
			index:     0,
			wantSlice: []string{},
			wantVal:   "",
			wantErr:   true,
		},
	}

	for _, tc := range testsString {
		t.Run("string-"+tc.name, func(t *testing.T) {
			res, val, err := Delete(tc.slice, tc.index)
			if tc.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tc.slice, res)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.wantSlice, res)
				assert.Equal(t, tc.wantVal, val)
			}
		})
	}

	// float64 类型测试
	testsFloat := []testCase[float64]{
		{
			name:      "删除float64元素",
			slice:     []float64{1.1, 2.2, 3.3},
			index:     2,
			wantSlice: []float64{1.1, 2.2},
			wantVal:   3.3,
			wantErr:   false,
		},
		{
			name:      "float64空切片删除",
			slice:     []float64{},
			index:     0,
			wantSlice: []float64{},
			wantVal:   0,
			wantErr:   true,
		},
	}

	for _, tc := range testsFloat {
		t.Run("float64-"+tc.name, func(t *testing.T) {
			res, val, err := Delete(tc.slice, tc.index)
			if tc.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tc.slice, res)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.wantSlice, res)
				assert.Equal(t, tc.wantVal, val)
			}
		})
	}
}
