package list

import (
	"fmt"
	"mkit/internal/errs"
)

var (
	_ List[any] = &LinkedList[any]{}
)

// 这里定义了一个泛型链表节点类型 node[T]。
// 其中 prev *node[T] 和 next *node[T] 的语法表示：
// prev 和 next 字段都是指向 node[T] 类型的指针。
// 也就是说，每个节点都可以通过 prev 指针访问前一个节点，通过 next 指针访问下一个节点，形成双向链表结构。
// 例如：*node[int] 表示“指向 int 类型节点的指针”。
type node[T any] struct {
	val  T        // 节点存储的值
	prev *node[T] // 指向前一个节点的指针
	next *node[T] // 指向下一个节点的指针
}

// 双链表结构体
type LinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

// 创建一个新的链表
func NewLinkedList[T any]() *LinkedList[T] {
	// 空链表包含头尾两个节点
	head := &node[T]{}
	tail := &node[T]{next: head, prev: head}
	head.next, head.prev = tail, tail
	return &LinkedList[T]{
		head:   head,
		tail:   tail,
		length: 0,
	}
}

func NewLinkedListOf[T any](ts []T) *LinkedList[T] {
	list := NewLinkedList[T]()
	if err := list.Append(ts...); err != nil {
		panic(err)
	}
	return list
}

func (l *LinkedList[T]) findNode(index int) *node[T] {
	var cur *node[T]
	if index <= l.Len()/2 {
		cur = l.head
		for i := -1; i < index; i++ {
			cur = cur.next
		}
	} else {
		cur = l.tail
		for i := l.Len(); i > index; i-- {
			cur = cur.prev
		}
	}
	return cur
}

func (l *LinkedList[T]) checkIndex(index int) bool {
	return 0 <= index && index < l.Len()
}

func (l *LinkedList[T]) Get(index int) (T, error) {
	if !l.checkIndex(index) {
		var zeroVal T
		return zeroVal, errs.NewErrIndexOutOfRange(l.length, index)
	}

	node := l.findNode(index)
	return node.val, nil
}

func (l *LinkedList[T]) Append(ts ...T) error {
	fmt.Println(ts)
	for _, t := range ts {
		node := &node[T]{prev: l.tail.prev, next: l.tail, val: t}
		node.prev.next, node.next.prev = node, node
		l.length++
	}
	return nil
}

func (l *LinkedList[T]) Add(index int, t T) error {
	if index < 0 || index > l.Len() {
		return errs.NewErrIndexOutOfRange(l.length, index)
	}

	if index == l.Len() {
		return l.Append(t)
	}

	// 找到前驱节点
	nextNode := l.findNode(index)

	// 创建新节点，然后插入链表对应位置
	node := &node[T]{prev: nextNode.prev, next: nextNode, val: t}
	node.prev.next, node.next.prev = node, node
	l.length++
	return nil
}

func (l *LinkedList[T]) Set(index int, t T) error {
	if !l.checkIndex(index) {
		return errs.NewErrIndexOutOfRange(l.Len(), index)
	}
	node := l.findNode(index)
	node.val = t
	return nil
}

func (l *LinkedList[T]) Remove(index int) (T, error) {
	if !l.checkIndex(index) {
		var zeroVal T
		return zeroVal, errs.NewErrIndexOutOfRange(l.Len(), index)
	}

	node := l.findNode(index)
	node.prev.next, node.next.prev = node.next, node.prev
	l.length--
	return node.val, nil
}

func (l *LinkedList[T]) Len() int {
	return l.length
}

func (l *LinkedList[T]) Cap() int {
	return l.length
}

// 对链表中全部节点使用fn函数
func (l *LinkedList[T]) Range(fn func(index int, t T) error) error {
	i := 0
	for cur := l.head.next; cur != l.tail; cur = cur.next {
		err := fn(i, cur.val)
		if err != nil {
			return err
		}
		i++
	}
	return nil
}

func (l *LinkedList[T]) AsSlice() []T {
	ans := []T{}
	for cur := l.head.next; cur != l.tail; cur = cur.next {
		ans = append(ans, cur.val)
	}
	return ans
}
