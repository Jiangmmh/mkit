package queue

import (
	"sync/atomic"
	"unsafe"
)

// node 节点结构体，存储值和下一个节点指针
type node[T any] struct {
	val  T
	next unsafe.Pointer // *node[T]，用于CAS操作
}

// ConcurrentLinkedQueue 无锁并发链表队列
type ConcurrentLinkedQueue[T any] struct {
	head unsafe.Pointer // *node[T]
	tail unsafe.Pointer // *node[T]
}

// NewConcurrentLinkedQueue 创建一个新的并发链表队列
func NewConcurrentLinkedQueue[T any]() *ConcurrentLinkedQueue[T] {
	head := &node[T]{}
	ptr := unsafe.Pointer(head)
	return &ConcurrentLinkedQueue[T]{
		head: ptr,
		tail: ptr,
	}
}

// Enqueue 入队操作（无锁，基于CAS）
// 参考ByteRhythm-main项目实现
func (c *ConcurrentLinkedQueue[T]) Enqueue(t T) error {
	newNode := &node[T]{val: t}
	newPtr := unsafe.Pointer(newNode)
	for {
		tailPtr := atomic.LoadPointer(&c.tail)
		tail := (*node[T])(tailPtr)
		tailNext := atomic.LoadPointer(&tail.next)
		if tailNext != nil {
			// 其他goroutine已插入新节点但未推进tail指针，尝试推进tail
			atomic.CompareAndSwapPointer(&c.tail, tailPtr, tailNext)
			continue
		}
		// 尝试将新节点插入到tail.next
		if atomic.CompareAndSwapPointer(&tail.next, nil, newPtr) {
			// 插入成功后，推进tail指针到新节点
			atomic.CompareAndSwapPointer(&c.tail, tailPtr, newPtr)
			return nil
		}
		// 插入失败，继续自旋
	}
}

// Dequeue 出队操作（无锁，基于CAS）
func (c *ConcurrentLinkedQueue[T]) Dequeue() (T, error) {
	for {
		headPtr := atomic.LoadPointer(&c.head)
		head := (*node[T])(headPtr)
		tailPtr := atomic.LoadPointer(&c.tail)
		tail := (*node[T])(tailPtr)
		nextPtr := atomic.LoadPointer(&head.next)
		if head == tail {
			// 队列为空
			var zero T
			return zero, ErrOutOfCapacity
		}
		if nextPtr == nil {
			// 理论上不会出现，保护性分支
			var zero T
			return zero, ErrOutOfCapacity
		}
		// 尝试推进head指针
		if atomic.CompareAndSwapPointer(&c.head, headPtr, nextPtr) {
			next := (*node[T])(nextPtr)
			return next.val, nil
		}
		// 推进失败，继续自旋
	}
}
