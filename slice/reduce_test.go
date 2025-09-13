package slice

import (
	"testing"
)

func TestReduce(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		initial  int
		fn       func(int, int) int
		expected int
	}{
		{
			name:     "空切片",
			input:    []int{},
			initial:  0,
			fn:       func(acc, x int) int { return acc + x },
			expected: 0,
		},
		{
			name:     "求和",
			input:    []int{1, 2, 3, 4, 5},
			initial:  0,
			fn:       func(acc, x int) int { return acc + x },
			expected: 15,
		},
		{
			name:     "求积",
			input:    []int{1, 2, 3, 4},
			initial:  1,
			fn:       func(acc, x int) int { return acc * x },
			expected: 24,
		},
		{
			name:    "求最大值",
			input:   []int{3, 7, 2, 9, 1},
			initial: 0,
			fn: func(acc, x int) int {
				if x > acc {
					return x
				}
				return acc
			},
			expected: 9,
		},
		{
			name:    "求最小值",
			input:   []int{3, 7, 2, 9, 1},
			initial: 1000,
			fn: func(acc, x int) int {
				if x < acc {
					return x
				}
				return acc
			},
			expected: 1,
		},
		{
			name:     "计数",
			input:    []int{1, 2, 3, 4, 5},
			initial:  0,
			fn:       func(acc, x int) int { return acc + 1 },
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reduce(tt.input, tt.initial, tt.fn)
			if result != tt.expected {
				t.Errorf("Reduce() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestReduceWithIndex(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		initial  int
		fn       func(int, int, int) int
		expected int
	}{
		{
			name:     "空切片",
			input:    []int{},
			initial:  0,
			fn:       func(acc, i, x int) int { return acc + i + x },
			expected: 0,
		},
		{
			name:     "索引和值相加",
			input:    []int{10, 20, 30},
			initial:  0,
			fn:       func(acc, i, x int) int { return acc + i + x },
			expected: 63, // 0+10 + 1+20 + 2+30 = 63
		},
		{
			name:     "索引乘值",
			input:    []int{2, 3, 4},
			initial:  0,
			fn:       func(acc, i, x int) int { return acc + i*x },
			expected: 11, // 0*2 + 1*3 + 2*4 = 11
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReduceWithIndex(tt.input, tt.initial, tt.fn)
			if result != tt.expected {
				t.Errorf("ReduceWithIndex() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestReduce_StringConcatenation(t *testing.T) {
	input := []string{"Hello", " ", "World", "!"}
	expected := "Hello World!"

	result := Reduce(input, "", func(acc, x string) string {
		return acc + x
	})

	if result != expected {
		t.Errorf("Reduce() = %v, want %v", result, expected)
	}
}

func TestReduce_OriginalUnchanged(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	input := make([]int, len(original))
	copy(input, original)

	Reduce(input, 0, func(acc, x int) int { return acc + x })

	// 检查原始切片是否被修改
	if len(input) != len(original) {
		t.Errorf("原始切片长度被修改了")
	}
	for i, v := range input {
		if v != original[i] {
			t.Errorf("原始切片被修改了，索引 %d 期望 %v，实际 %v", i, original[i], v)
		}
	}
}
