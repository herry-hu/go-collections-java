package doublelinkedlist

import (
	"bytes"
	"fmt"
	"go-collections-java/lang"
)

type Node[T lang.Comparable] struct {
	value T        // 节点的值
	prev  *Node[T] // 上一个节点的指针
	next  *Node[T] // 下一个节点的指针
}

type LinkedList[T lang.Comparable] struct {
	head *Node[T] // 链表头部节点
	tail *Node[T] // 链表尾部节点
	size int      // 链表大小
}

func NewLinkedList[T lang.Comparable]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// 将指定元素添加到列表末尾。
func (list *LinkedList[T]) Add(item T) {
	node := &Node[T]{value: item, prev: list.tail, next: nil}
	if list.tail == nil {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		list.tail = node
	}
	list.size++
}

// 将指定元素插入到列表的指定位置。
func (list *LinkedList[T]) AddAt(index int, item T) {
	if index < 0 || index > list.size {
		panic("index out of bounds")
	}
	if index == 0 {
		node := &Node[T]{value: item, prev: nil, next: list.head}
		if list.head == nil {
			list.tail = node
		} else {
			list.head.prev = node
		}
		list.head = node
	} else if index == list.size {
		node := &Node[T]{value: item, prev: list.tail, next: nil}
		if list.tail == nil {
			list.head = node
		} else {
			list.tail.next = node
		}
		list.tail = node
	} else {
		current := list.getNode(index)
		prev := current.prev
		node := &Node[T]{value: item, prev: prev, next: current}
		prev.next = node
		current.prev = node
	}
	list.size++
}

// 返回列表中指定位置的元素。
func (list *LinkedList[T]) Get(index int) T {
	if index < 0 || index >= list.size {
		panic("index out of bounds")
	}
	return list.getNode(index).value
}

// 将列表中指定位置的元素替换为指定元素。
func (list *LinkedList[T]) Set(index int, item T) {
	if index < 0 || index >= list.size {
		panic("index out of bounds")
	}
	node := list.getNode(index)
	node.value = item
}

// 删除列表中指定位置的元素。
func (list *LinkedList[T]) Remove(index int) bool {
	if index < 0 || index >= list.size {
		panic("index out of bounds")
	}
	node := list.getNode(index)
	if node.prev == nil {
		list.head = node.next
	} else {
		node.prev.next = node.next
	}
	if node.next == nil {
		list.tail = node.prev
	} else {
		node.next.prev = node.prev
	}
	list.size--
	return true
}

// 返回列表的大小。
func (list *LinkedList[T]) Size() int {
	return list.size
}

// 返回列表中指定位置的节点。
func (list *LinkedList[T]) getNode(index int) *Node[T] {
	if index < 0 || index >= list.size {
		panic("index out of bounds")
	}
	var current *Node[T]
	if index < list.size/2 {
		current = list.head
		for i := 0; i < index; i++ {
			current = current.next
		}
	} else {
		current = list.tail
		for i := list.size - 1; i > index; i-- {
			current = current.prev
		}
	}
	return current
}

type Set[T lang.Comparable] struct {
	list *LinkedList[T] // 基于双向链表实现的 LinkedList
}

// 向集合中添加元素。
func (set *Set[T]) Add(item T) {
	if !set.Contains(item) {
		set.list.Add(item)
	}
}

// 从集合中删除元素。
func (set *Set[T]) Remove(item T) bool {
	for i := 0; i < set.list.Size(); i++ {
		if set.list.Get(i).CompareTo(item) == 0 {
			set.list.Remove(i)
			return true
		}
	}
	return false
}

// 检查集合中是否包含指定元素。
func (set *Set[T]) Contains(item T) bool {
	for i := 0; i < set.list.Size(); i++ {
		if set.list.Get(i).CompareTo(item) == 0 {
			return true
		}
	}
	return false
}

// 返回集合中的元素数量。
func (set *Set[T]) Size() int {
	return set.list.Size()
}

// 将集合转换为切片。
func (set *Set[T]) ToSlice() []T {
	slice := make([]T, set.Size())
	for i := 0; i < set.Size(); i++ {
		slice[i] = set.list.Get(i)
	}
	return slice
}

func (list *LinkedList[T]) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i := 0; i < list.size; i++ {
		if i > 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString(fmt.Sprintf("%v", list.Get(i)))
	}
	buffer.WriteString("]")
	return buffer.String()
}
