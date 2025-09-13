package slice

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		fn       func(int) int
		expected []int
	}{
		{
			name:     "空切片",
			input:    []int{},
			fn:       func(x int) int { return x * 2 },
			expected: []int{},
		},
		{
			name:     "单个元素",
			input:    []int{5},
			fn:       func(x int) int { return x * 2 },
			expected: []int{10},
		},
		{
			name:     "多个元素",
			input:    []int{1, 2, 3, 4, 5},
			fn:       func(x int) int { return x * 2 },
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "零值",
			input:    []int{0, 1, 0},
			fn:       func(x int) int { return x + 1 },
			expected: []int{1, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Map(tt.input, tt.fn)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Map() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMapWithIndex(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		fn       func(int, int) int
		expected []int
	}{
		{
			name:     "空切片",
			input:    []int{},
			fn:       func(i, x int) int { return i + x },
			expected: []int{},
		},
		{
			name:     "索引加值",
			input:    []int{10, 20, 30},
			fn:       func(i, x int) int { return i + x },
			expected: []int{10, 21, 32},
		},
		{
			name:     "索引乘值",
			input:    []int{2, 3, 4},
			fn:       func(i, x int) int { return i * x },
			expected: []int{0, 3, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapWithIndex(tt.input, tt.fn)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MapWithIndex() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMap_StringConversion(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []string{"1", "2", "3", "4", "5"}

	result := Map(input, func(x int) string {
		return string(rune('0' + x))
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Map() = %v, want %v", result, expected)
	}
}

func TestMap_OriginalUnchanged(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	input := make([]int, len(original))
	copy(input, original)

	Map(input, func(x int) int { return x * 2 })

	if !reflect.DeepEqual(input, original) {
		t.Errorf("原始切片被修改了，期望 %v，实际 %v", original, input)
	}
}
