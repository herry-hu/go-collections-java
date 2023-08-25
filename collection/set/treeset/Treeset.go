package treeset

import (
	"container/list"
	"fmt"
	"go-collections-java/lang"
	"reflect"
	"strings"
)

// TreeSet 是一个基于红黑树实现的有序集合。
type TreeSet[T lang.Comparable] struct {
	set *list.List   // 使用双向链表存储元素
	typ reflect.Type // 元素类型
}

// NewTreeSet 创建一个新的 TreeSet 实例。
func NewTreeSet[T lang.Comparable]() *TreeSet[T] {
	return &TreeSet[T]{
		set: list.New(),
	}
}

// Add 向集合中添加元素。
func (set *TreeSet[T]) Add(value T) {
	// 向集合中添加指定元素
	for e := set.set.Front(); e != nil; e = e.Next() {
		cmp := e.Value.(lang.Comparable).CompareTo(value)
		if cmp == 0 {
			return // 元素已存在，不重复添加
		} else if cmp > 0 {
			set.set.InsertBefore(value, e)
			return
		}
	}
	set.set.PushBack(value)

	// 更新元素类型
	if set.typ == nil {
		set.typ = reflect.TypeOf(value)
	}
}

// Clear 清空集合中的所有元素。
func (set *TreeSet[T]) Clear() {
	set.set = list.New()
}

// Contains 检查集合中是否包含指定元素。
func (set *TreeSet[T]) Contains(item T) bool {
	for e := set.set.Front(); e != nil; e = e.Next() {
		cmp := e.Value.(lang.Comparable).CompareTo(item)
		if cmp == 0 {
			return true
		} else if cmp > 0 {
			return false
		}
	}
	return false
}

// First 返回集合中的第一个元素。
func (set *TreeSet[T]) First() T {
	var t T
	if set.set.Len() == 0 {
		return t
	}
	return set.set.Front().Value.(T)
}

// IsEmpty 检查集合是否为空。
func (set *TreeSet[T]) IsEmpty() bool {
	return set.set.Len() == 0
}

// Set 用指定的元素替换集合中的所有元素。
func (set *TreeSet[T]) Set(items ...T) {
	set.Clear() // 清空集合

	for _, item := range items {
		set.Add(item) // 添加新的元素
	}
}

// Iter 返回一个通道，用于迭代集合中的元素。
func (set *TreeSet[T]) Iter() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for e := set.set.Front(); e != nil; e = e.Next() {
			ch <- e.Value.(T)
		}
	}()
	return ch
}

// String 返回集合的字符串表示形式。
func (set *TreeSet[T]) String() string {
	var values []string
	for item := range set.Iter() {
		values = append(values, fmt.Sprintf("%v", item))
	}
	return "treeSet{" + strings.Join(values, ", ") + "}"
}

// Size 返回集合中的元素数量。
func (set *TreeSet[T]) Size() int {
	return set.set.Len()
}

// Remove 从集合中移除指定的元素。
func (set *TreeSet[T]) Remove(item T) {
	for e := set.set.Front(); e != nil; e = e.Next() {
		if e.Value.(lang.Comparable).CompareTo(item) == 0 {
			set.set.Remove(e)
			return
		}
	}
}

// Union 返回当前集合与另一个集合的并集。
func (set *TreeSet[T]) Union(other *TreeSet[T]) *TreeSet[T] {
	unionSet := NewTreeSet[T]()

	for item := range set.Iter() {
		unionSet.Add(item)
	}

	for item := range other.Iter() {
		unionSet.Add(item)
	}

	return unionSet
}

// Intersection 返回当前集合与另一个集合的交集。
func (set *TreeSet[T]) Intersection(other *TreeSet[T]) *TreeSet[T] {
	intersectionSet := NewTreeSet[T]()

	for item := range set.Iter() {
		if other.Contains(item) {
			intersectionSet.Add(item)
		}
	}

	return intersectionSet
}
