package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_Basic(t *testing.T) {
	ll := NewLinkedList[int]()
	assert.Equal(t, 0, ll.Len())
	assert.Equal(t, 0, ll.Cap())
	assert.Equal(t, []int{}, ll.AsSlice())

	// Append
	err := ll.Append(1, 2, 3)
	assert.NoError(t, err)
	assert.Equal(t, 3, ll.Len())
	assert.Equal(t, []int{1, 2, 3}, ll.AsSlice())

	// Add at head
	err = ll.Add(0, 10)
	assert.NoError(t, err)
	assert.Equal(t, []int{10, 1, 2, 3}, ll.AsSlice())

	// Add at tail
	err = ll.Add(ll.Len(), 20)
	assert.NoError(t, err)
	assert.Equal(t, []int{10, 1, 2, 3, 20}, ll.AsSlice())

	// Add at middle
	err = ll.Add(2, 99)
	assert.NoError(t, err)
	assert.Equal(t, []int{10, 1, 99, 2, 3, 20}, ll.AsSlice())

	// Get
	val, err := ll.Get(2)
	assert.NoError(t, err)
	assert.Equal(t, 99, val)

	// Set
	err = ll.Set(2, 100)
	assert.NoError(t, err)
	val, err = ll.Get(2)
	assert.NoError(t, err)
	assert.Equal(t, 100, val)

	// Remove head
	removed, err := ll.Remove(0)
	assert.NoError(t, err)
	assert.Equal(t, 10, removed)
	assert.Equal(t, []int{1, 100, 2, 3, 20}, ll.AsSlice())

	// Remove tail
	removed, err = ll.Remove(ll.Len() - 1)
	assert.NoError(t, err)
	assert.Equal(t, 20, removed)
	assert.Equal(t, []int{1, 100, 2, 3}, ll.AsSlice())

	// Remove middle
	removed, err = ll.Remove(1)
	assert.NoError(t, err)
	assert.Equal(t, 100, removed)
	assert.Equal(t, []int{1, 2, 3}, ll.AsSlice())
}

func TestLinkedList_Errors(t *testing.T) {
	ll := NewLinkedList[int]()
	_, err := ll.Get(0)
	assert.Error(t, err)

	err = ll.Set(0, 1)
	assert.Error(t, err)

	_, err = ll.Remove(0)
	assert.Error(t, err)

	err = ll.Add(-1, 1)
	assert.Error(t, err)

	err = ll.Add(1, 1)
	assert.Error(t, err)
}

func TestLinkedList_Range(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Append(1, 2, 3)
	var res []int
	err := ll.Range(func(index int, t int) error {
		res = append(res, t)
		return nil
	})
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, res)
}

func TestLinkedList_NewLinkedListOf(t *testing.T) {
	ll := NewLinkedListOf([]int{5, 6, 7})
	assert.Equal(t, 3, ll.Len())
	assert.Equal(t, []int{5, 6, 7}, ll.AsSlice())
}
