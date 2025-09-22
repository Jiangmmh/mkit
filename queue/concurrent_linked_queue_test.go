package queue

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestConcurrentLinkedQueue_Basic(t *testing.T) {
	q := NewConcurrentLinkedQueue[int]()

	// 测试空队列出队
	_, err := q.Dequeue()
	if !errors.Is(err, ErrOutOfCapacity) {
		t.Errorf("Dequeue 空队列应返回 ErrOutOfCapacity, 实际: %v", err)
	}

	// 入队
	if err := q.Enqueue(1); err != nil {
		t.Errorf("Enqueue 失败: %v", err)
	}
	if err := q.Enqueue(2); err != nil {
		t.Errorf("Enqueue 失败: %v", err)
	}
	if err := q.Enqueue(3); err != nil {
		t.Errorf("Enqueue 失败: %v", err)
	}

	// 出队
	v, err := q.Dequeue()
	if err != nil || v != 1 {
		t.Errorf("Dequeue 期望1, 实际%v, err=%v", v, err)
	}
	v, err = q.Dequeue()
	if err != nil || v != 2 {
		t.Errorf("Dequeue 期望2, 实际%v, err=%v", v, err)
	}
	v, err = q.Dequeue()
	if err != nil || v != 3 {
		t.Errorf("Dequeue 期望3, 实际%v, err=%v", v, err)
	}

	// 再次出队应为空
	_, err = q.Dequeue()
	if !errors.Is(err, ErrOutOfCapacity) {
		t.Errorf("Dequeue 空队列应返回 ErrOutOfCapacity, 实际: %v", err)
	}
}

func TestConcurrentLinkedQueue_Concurrent(t *testing.T) {
	q := NewConcurrentLinkedQueue[int]()
	const (
		producerCount = 10
		consumerCount = 10
		perProducer   = 1000
	)
	var wg sync.WaitGroup
	var produced, consumed sync.Map

	// 生产者
	for p := 0; p < producerCount; p++ {
		wg.Add(1)
		go func(pid int) {
			defer wg.Done()
			for i := 0; i < perProducer; i++ {
				val := pid*perProducer + i
				if err := q.Enqueue(val); err != nil {
					t.Errorf("Enqueue 失败: %v", err)
				} else {
					produced.Store(val, true)
				}
			}
		}(p)
	}

	// 消费者
	var consumeWg sync.WaitGroup
	total := producerCount * perProducer
	consumeWg.Add(consumerCount)
	for c := 0; c < consumerCount; c++ {
		go func() {
			defer consumeWg.Done()
			for {
				v, err := q.Dequeue()
				if err != nil {
					// 可能队列暂时为空，稍等重试
					time.Sleep(time.Millisecond)
					continue
				}
				consumed.Store(v, true)
				if lenMap(&consumed) >= total {
					return
				}
			}
		}()
	}

	wg.Wait()
	// 等待所有元素被消费
	waitTimeout(&consumeWg, 5*time.Second)

	// 检查所有生产的都被消费
	producedCount := lenMap(&produced)
	consumedCount := lenMap(&consumed)
	if producedCount != total {
		t.Errorf("生产数量不对: %d, 期望: %d", producedCount, total)
	}
	if consumedCount != total {
		t.Errorf("消费数量不对: %d, 期望: %d", consumedCount, total)
	}

	// 检查每个生产的都被消费
	produced.Range(func(key, value any) bool {
		if _, ok := consumed.Load(key); !ok {
			t.Errorf("元素 %v 被生产但未被消费", key)
		}
		return true
	})
}

// 辅助函数：统计 sync.Map 长度
func lenMap(m *sync.Map) int {
	cnt := 0
	m.Range(func(_, _ any) bool {
		cnt++
		return true
	})
	return cnt
}

// 辅助函数：带超时的 Wait
func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return true // 正常完成
	case <-time.After(timeout):
		return false // 超时
	}
}
