package list

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestArrayList_AppendAndGet(t *testing.T) {
	list := NewArrayList[int]()
	err := list.Append(1, 2, 3)
	if err != nil {
		t.Fatalf("Append 失败: %v", err)
	}
	if list.Len() != 3 {
		t.Fatalf("期望长度3，实际为%d", list.Len())
	}
	for i := 0; i < 3; i++ {
		val, err := list.Get(i)
		if err != nil {
			t.Fatalf("Get(%d) 失败: %v", i, err)
		}
		if val != i+1 {
			t.Fatalf("期望值%d，实际为%d", i+1, val)
		}
	}
}

func TestArrayList_Get_OutOfRange(t *testing.T) {
	list := NewArrayList[int]()
	list.Append(1)
	_, err := list.Get(1)
	if err == nil {
		t.Fatal("期望越界错误，但未发生")
	}

	// 检查错误是否包含下标超出范围的提示
	if err == nil || !strings.Contains(err.Error(), "下标超出范围") {
		t.Fatalf("期望下标超出范围错误，实际为%v", err)
	}
}

func TestArrayList_Add(t *testing.T) {
	list := NewArrayList[int]()
	list.Append(1, 2, 3)
	err := list.Add(1, 99)
	if err != nil {
		t.Fatalf("Add 失败: %v", err)
	}
	expected := []int{1, 99, 2, 3}
	actual := list.AsSlice()
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Add 后切片不一致，期望%v，实际%v", expected, actual)
	}
}

func TestArrayList_Add_OutOfRange(t *testing.T) {
	list := NewArrayList[int]()
	err := list.Add(2, 10)
	if err == nil {
		t.Fatal("期望越界错误，但未发生")
	}
}

func TestArrayList_Set(t *testing.T) {
	list := NewArrayList[int]()
	list.Append(1, 2, 3)
	err := list.Set(1, 100)
	if err != nil {
		t.Fatalf("Set 失败: %v", err)
	}
	val, _ := list.Get(1)
	if val != 100 {
		t.Fatalf("Set 后值不正确，期望100，实际%d", val)
	}
}

func TestArrayList_Set_OutOfRange(t *testing.T) {
	list := NewArrayList[int]()
	err := list.Set(0, 10)
	if err == nil {
		t.Fatal("期望越界错误，但未发生")
	}
}

func TestArrayList_Delete(t *testing.T) {
	list := NewArrayList[int]()
	list.Append(1, 2, 3)
	val, err := list.Delete(1)
	if err != nil {
		t.Fatalf("Delete 失败: %v", err)
	}
	if val != 2 {
		t.Fatalf("Delete 返回值错误，期望2，实际%d", val)
	}
	expected := []int{1, 3}
	actual := list.AsSlice()
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Delete 后切片不一致，期望%v，实际%v", expected, actual)
	}
}

func TestArrayList_Delete_OutOfRange(t *testing.T) {
	list := NewArrayList[int]()
	_, err := list.Delete(0)
	if err == nil {
		t.Fatal("期望越界错误，但未发生")
	}
}

func TestArrayList_Range(t *testing.T) {
	list := NewArrayList[int]()
	list.Append(1, 2, 3)
	sum := 0
	err := list.Range(func(index int, t int) error {
		sum += t
		return nil
	})
	if err != nil {
		t.Fatalf("Range 失败: %v", err)
	}
	if sum != 6 {
		t.Fatalf("Range 累加错误，期望6，实际%d", sum)
	}
}

func TestArrayList_Range_EarlyExit(t *testing.T) {
	list := NewArrayList[int]()
	list.Append(1, 2, 3)
	errTest := errors.New("test")
	err := list.Range(func(index int, t int) error {
		if index == 1 {
			return errTest
		}
		return nil
	})
	if !errors.Is(err, errTest) {
		t.Fatalf("Range 未能提前退出，期望%v，实际%v", errTest, err)
	}
}

func TestArrayList_AsSlice(t *testing.T) {
	list := NewArrayList[int]()
	slice1 := list.AsSlice()
	if slice1 == nil || len(slice1) != 0 {
		t.Fatalf("空列表AsSlice应返回空切片，实际为%v", slice1)
	}
	list.Append(1, 2)
	slice2 := list.AsSlice()
	if !reflect.DeepEqual(slice2, []int{1, 2}) {
		t.Fatalf("AsSlice 返回值错误，期望[1 2]，实际%v", slice2)
	}
	// 检查返回的是新切片
	slice2[0] = 100
	val, _ := list.Get(0)
	if val == 100 {
		t.Fatal("AsSlice 返回的切片应为新切片，修改不应影响原数据")
	}
}

func TestArrayList_LenAndCap(t *testing.T) {
	list := NewArrayList[int]()
	if list.Len() != 0 {
		t.Fatalf("新建列表长度应为0，实际为%d", list.Len())
	}
	list.Append(1, 2, 3)
	if list.Len() != 3 {
		t.Fatalf("追加后长度应为3，实际为%d", list.Len())
	}
	if list.Cap() < 3 {
		t.Fatalf("容量应大于等于3，实际为%d", list.Cap())
	}
}
