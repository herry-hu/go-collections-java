package linkedlist

import (
	"bytes"
	"fmt"
)

type Node[T comparable] struct {
	value T        // 节点的值
	next  *Node[T] // 下一个节点的指针
}

type LinkedList[T comparable] struct {
	head *Node[T] // 链表头部节点
	size int      // 链表大小
}

type List[T comparable] struct {
	linkedList *LinkedList[T] // 基于单向链表实现的 LinkedList
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		size: 0,
	}
}

// 将指定元素添加到列表末尾。
func (list *LinkedList[T]) Add(item T) {
	node := &Node[T]{value: item, next: nil}
	if list.head == nil {
		list.head = node
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
	list.size++
}

// 将指定元素插入到列表的指定位置。
func (list *LinkedList[T]) AddAt(index int, item T) {
	if index < 0 || index > list.size {
		panic("index out of bounds")
	}
	if index == 0 {
		list.head = &Node[T]{value: item, next: list.head}
	} else {
		prev := list.getNode(index - 1)
		prev.next = &Node[T]{value: item, next: prev.next}
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
	if index == 0 {
		list.head = list.head.next
	} else {
		prev := list.getNode(index - 1)
		prev.next = prev.next.next
	}
	list.size--
	return true
}

// 返回列表中的元素数量。
func (list *LinkedList[T]) Size() int {
	return list.size
}

// 判断列表是否为空。
func (list *LinkedList[T]) IsEmpty() bool {
	return list.size == 0
}

// 获取指定位置的节点。
func (list *LinkedList[T]) getNode(index int) *Node[T] {
	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current
}

// 将指定元素添加到列表末尾。
func (list *List[T]) PushBack(item T) {
	list.linkedList.Add(item)
}

// 将指定元素插入到列表的指定位置。
func (list *List[T]) InsertAt(index int, item T) {
	list.linkedList.AddAt(index, item)
}

// 返回列表中指定位置的元素。
func (list *List[T]) Get(index int) T {
	return list.linkedList.Get(index)
}

// 将列表中指定位置的元素替换为指定元素。
func (list *List[T]) Set(index int, item T) {
	list.linkedList.Set(index, item)
}

// 删除列表中指定位置的元素。
func (list *List[T]) RemoveAt(index int) bool {
	return list.linkedList.Remove(index)
}

// 返回列表中的元素数量。
func (list *List[T]) Len() int {
	return list.linkedList.Size()
}

// 判断列表是否为空。
func (list *List[T]) IsEmpty() bool {
	return list.linkedList.IsEmpty()
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
