package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFilterFunc 针对 filterFunc 泛型函数的单元测试，覆盖不同类型和过滤条件
func TestFilterFunc(t *testing.T) {
	type testCase[T any] struct {
		name    string
		src     []T
		filter  func(T) bool
		want    []T
		wantErr bool
	}

	// int 类型测试
	testsInt := []testCase[int]{
		{
			name:   "过滤偶数",
			src:    []int{1, 2, 3, 4, 5, 6},
			filter: func(i int) bool { return i%2 == 0 },
			want:   []int{2, 4, 6},
		},
		{
			name:   "全部过滤掉",
			src:    []int{1, 3, 5},
			filter: func(i int) bool { return i%2 == 0 },
			want:   []int{},
		},
		{
			name:   "全部保留",
			src:    []int{2, 4, 6},
			filter: func(i int) bool { return i%2 == 0 },
			want:   []int{2, 4, 6},
		},
		{
			name:   "空切片",
			src:    []int{},
			filter: func(i int) bool { return true },
			want:   []int{},
		},
	}

	for _, tc := range testsInt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := FilterFunc(tc.src, tc.filter)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}

	// string 类型测试
	testsString := []testCase[string]{
		{
			name:   "过滤长度大于3的字符串",
			src:    []string{"go", "java", "python", "c"},
			filter: func(s string) bool { return len(s) > 3 },
			want:   []string{"java", "python"},
		},
		{
			name:   "全部保留",
			src:    []string{"a", "b"},
			filter: func(s string) bool { return true },
			want:   []string{"a", "b"},
		},
		{
			name:   "全部过滤掉",
			src:    []string{"a", "b"},
			filter: func(s string) bool { return false },
			want:   []string{},
		},
	}

	for _, tc := range testsString {
		t.Run(tc.name, func(t *testing.T) {
			got, err := FilterFunc(tc.src, tc.filter)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}
