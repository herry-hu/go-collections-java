package arraylist

import (
	"bytes"
	"fmt"
	"go-collections-java/lang"
)

type ArrayList[T lang.Comparable] struct {
	size int
	data []T
}

func NewArrayList[T lang.Comparable]() *ArrayList[T] {
	return &ArrayList[T]{size: 0, data: make([]T, 0)}
}

func (list *ArrayList[T]) Add(item T) {
	list.data = append(list.data, item)
	list.size++
}

func (list *ArrayList[T]) AddAll(items ...T) {
	for _, item := range items {
		list.Add(item)
	}
}

func (list *ArrayList[T]) Get(index int) T {
	return list.data[index]
}

func (list *ArrayList[T]) Set(index int, item T) {
	list.data[index] = item
}

func (list *ArrayList[T]) Remove(index int) {
	copy(list.data[index:], list.data[index+1:])
	list.data = list.data[:len(list.data)-1]
	list.size--
}

func (list *ArrayList[T]) IndexOf(item T) int {
	for i, v := range list.data {
		if v.CompareTo(item) == 0 {
			return i
		}
	}
	return -1
}

func (list *ArrayList[T]) LastIndexOf(item T) int {
	for i := list.size - 1; i >= 0; i-- {
		if list.data[i].CompareTo(item) == 0 {
			return i
		}
	}
	return -1
}

func (list *ArrayList[T]) Size() int {
	return list.size
}

func (list *ArrayList[T]) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i, v := range list.data {
		if i > 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString(fmt.Sprintf("%v", v))
	}
	buffer.WriteString("]")
	return buffer.String()
}
