package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestShrink 针对 Shrink 泛型函数的单元测试，覆盖了不同容量和长度的切片情况
func TestShrink(t *testing.T) {
	type testCase[T any] struct {
		name      string
		slice     []T
		makeCap   int
		wantCap   int
		wantEqual bool // 缩容后内容是否应与原切片内容一致
	}

	testsInt := []testCase[int]{
		{
			name:      "cap大于2048且空闲超过一半，缩容到0.625",
			slice:     make([]int, 1000, 4000),
			makeCap:   4000,
			wantCap:   int(float32(4000) * 0.625),
			wantEqual: true,
		},
		{
			name:      "cap大于2048但空闲不超过一半，不缩容",
			slice:     make([]int, 3000, 4000),
			makeCap:   4000,
			wantCap:   4000,
			wantEqual: true,
		},
		{
			name:      "cap小于等于2048且空闲超过3/4，缩容一半",
			slice:     make([]int, 500, 2048),
			makeCap:   2048,
			wantCap:   1024,
			wantEqual: true,
		},
		{
			name:      "cap小于等于2048但空闲不超过3/4，不缩容",
			slice:     make([]int, 600, 2048),
			makeCap:   2048,
			wantCap:   2048,
			wantEqual: true,
		},
		{
			name:      "空切片",
			slice:     make([]int, 0, 100),
			makeCap:   100,
			wantCap:   100,
			wantEqual: true,
		},
		{
			name:      "长度等于容量，不缩容",
			slice:     make([]int, 10, 10),
			makeCap:   10,
			wantCap:   10,
			wantEqual: true,
		},
	}

	for _, tc := range testsInt {
		t.Run(tc.name, func(t *testing.T) {
			// 填充内容，便于后续校验
			for i := 0; i < len(tc.slice); i++ {
				tc.slice[i] = i
			}
			res := Shrink(tc.slice)
			assert.Equal(t, len(tc.slice), len(res), "缩容后长度应一致")
			assert.Equal(t, tc.wantCap, cap(res), "缩容后容量不符")
			if tc.wantEqual {
				assert.Equal(t, tc.slice, res, "缩容后内容应一致")
			}
		})
	}

	// string类型测试
	testsString := []testCase[string]{
		{
			name:      "string类型缩容",
			slice:     make([]string, 100, 400),
			makeCap:   400,
			wantCap:   200, // 400/2=200，因为400<=2048且400/100=4>=4
			wantEqual: true,
		},
		{
			name:      "string类型不缩容",
			slice:     make([]string, 300, 400),
			makeCap:   400,
			wantCap:   400,
			wantEqual: true,
		},
	}

	for _, tc := range testsString {
		t.Run(tc.name, func(t *testing.T) {
			for i := 0; i < len(tc.slice); i++ {
				tc.slice[i] = "a"
			}
			res := Shrink(tc.slice)
			assert.Equal(t, len(tc.slice), len(res))
			assert.Equal(t, tc.wantCap, cap(res))
			if tc.wantEqual {
				assert.Equal(t, tc.slice, res)
			}
		})
	}
}
