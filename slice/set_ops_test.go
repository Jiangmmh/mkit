package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 针对 Intersection、Union、Difference、SymmetricDifference 的单元测试
func TestSetOps_Int(t *testing.T) {
	type testCase struct {
		name              string
		a, b              []int
		wantIntersection  []int
		wantUnion         []int
		wantDifference    []int
		wantSymmetricDiff []int
	}

	tests := []testCase{
		{
			name:              "无交集",
			a:                 []int{1, 2, 3},
			b:                 []int{4, 5, 6},
			wantIntersection:  []int{},
			wantUnion:         []int{1, 2, 3, 4, 5, 6},
			wantDifference:    []int{1, 2, 3},
			wantSymmetricDiff: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:              "有交集",
			a:                 []int{1, 2, 3, 4},
			b:                 []int{3, 4, 5, 6},
			wantIntersection:  []int{3, 4},
			wantUnion:         []int{1, 2, 3, 4, 5, 6},
			wantDifference:    []int{1, 2},
			wantSymmetricDiff: []int{1, 2, 5, 6},
		},
		{
			name:              "a为空",
			a:                 []int{},
			b:                 []int{1, 2},
			wantIntersection:  []int{},
			wantUnion:         []int{1, 2},
			wantDifference:    []int{},
			wantSymmetricDiff: []int{1, 2},
		},
		{
			name:              "b为空",
			a:                 []int{1, 2},
			b:                 []int{},
			wantIntersection:  []int{},
			wantUnion:         []int{1, 2},
			wantDifference:    []int{1, 2},
			wantSymmetricDiff: []int{1, 2},
		},
		{
			name:              "完全相同",
			a:                 []int{1, 2, 3},
			b:                 []int{1, 2, 3},
			wantIntersection:  []int{1, 2, 3},
			wantUnion:         []int{1, 2, 3},
			wantDifference:    []int{},
			wantSymmetricDiff: []int{},
		},
	}

	// 因为集合无序，需用assert.ElementsMatch
	for _, tc := range tests {
		t.Run(tc.name+"-Intersection", func(t *testing.T) {
			got := Intersection(tc.a, tc.b)
			assert.ElementsMatch(t, tc.wantIntersection, got)
		})
		t.Run(tc.name+"-Union", func(t *testing.T) {
			got := Union(tc.a, tc.b)
			assert.ElementsMatch(t, tc.wantUnion, got)
		})
		t.Run(tc.name+"-Difference", func(t *testing.T) {
			got := Difference(tc.a, tc.b)
			assert.ElementsMatch(t, tc.wantDifference, got)
		})
		t.Run(tc.name+"-SymmetricDifference", func(t *testing.T) {
			got := SymmetricDifference(tc.a, tc.b)
			assert.ElementsMatch(t, tc.wantSymmetricDiff, got)
		})
	}
}

func TestSetOps_String(t *testing.T) {
	type testCase struct {
		name              string
		a, b              []string
		wantIntersection  []string
		wantUnion         []string
		wantDifference    []string
		wantSymmetricDiff []string
	}

	tests := []testCase{
		{
			name:              "有交集",
			a:                 []string{"go", "java", "python"},
			b:                 []string{"java", "c++", "go"},
			wantIntersection:  []string{"go", "java"},
			wantUnion:         []string{"go", "java", "python", "c++"},
			wantDifference:    []string{"python"},
			wantSymmetricDiff: []string{"python", "c++"},
		},
		{
			name:              "无交集",
			a:                 []string{"a"},
			b:                 []string{"b"},
			wantIntersection:  []string{},
			wantUnion:         []string{"a", "b"},
			wantDifference:    []string{"a"},
			wantSymmetricDiff: []string{"a", "b"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name+"-Intersection", func(t *testing.T) {
			got := Intersection(tc.a, tc.b)
			assert.ElementsMatch(t, tc.wantIntersection, got)
		})
		t.Run(tc.name+"-Union", func(t *testing.T) {
			got := Union(tc.a, tc.b)
			assert.ElementsMatch(t, tc.wantUnion, got)
		})
		t.Run(tc.name+"-Difference", func(t *testing.T) {
			got := Difference(tc.a, tc.b)
			assert.ElementsMatch(t, tc.wantDifference, got)
		})
		t.Run(tc.name+"-SymmetricDifference", func(t *testing.T) {
			got := SymmetricDifference(tc.a, tc.b)
			assert.ElementsMatch(t, tc.wantSymmetricDiff, got)
		})
	}
}
