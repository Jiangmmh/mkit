package main

import (
	"fmt"
	"mkit/list"
)

func main() {
	// 创建一个新的链表
	ll := list.NewLinkedList[int]()

	// 追加元素
	ll.Append(1, 2, 3)
	fmt.Println("After Append:", ll.AsSlice())

	// 在指定位置插入元素
	ll.Add(1, 99)
	fmt.Println("After Add(1, 99):", ll.AsSlice())

	// 获取指定位置的元素
	val, err := ll.Get(2)
	if err != nil {
		fmt.Println("Get error:", err)
	} else {
		fmt.Println("Get(2):", val)
	}

	// 设置指定位置的元素
	err = ll.Set(0, 42)
	if err != nil {
		fmt.Println("Set error:", err)
	} else {
		fmt.Println("After Set(0, 42):", ll.AsSlice())
	}

	// 删除指定位置的元素
	removed, err := ll.Remove(1)
	if err != nil {
		fmt.Println("Remove error:", err)
	} else {
		fmt.Println("Removed:", removed)
		fmt.Println("After Remove(1):", ll.AsSlice())
	}

	// 遍历链表
	fmt.Print("Range: ")
	ll.Range(func(index int, t int) error {
		fmt.Printf("[%d]=%d ", index, t)
		return nil
	})
	fmt.Println()

	// 打印长度和容量
	fmt.Println("Len:", ll.Len(), "Cap:", ll.Cap())
}
