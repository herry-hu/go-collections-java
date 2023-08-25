package hashmap

import (
	"bytes"
	"fmt"
	"go-collections-java/lang"
	"hash/fnv"
	"reflect"
)

type entry[T lang.Comparable, V lang.Comparable] struct {
	key   T            // 键
	value V            // 值
	next  *entry[T, V] // 指向下一个节点的指针
}

type HashMap[T lang.Comparable, V lang.Comparable] struct {
	data           []*entry[T, V] // 存储数据的切片
	capacity       int            // 哈希表容量
	size           int            // 哈希表中元素数量
	loadFactor     float64        // 负载因子
	threshold      int            // 下一次扩容的阈值
	resizeCapacity int            // 扩容后的容量
}

// 创建一个新的哈希表
func NewHashMap[T lang.Comparable, V lang.Comparable]() *HashMap[T, V] {
	return &HashMap[T, V]{
		data:           make([]*entry[T, V], 16),
		capacity:       16,
		size:           0,
		loadFactor:     0.75,
		threshold:      12,
		resizeCapacity: 32,
	}
}

// 将键值对添加到哈希表中
func (h *HashMap[T, V]) Put(key T, value V) {
	// 如果负载因子大于阈值，扩容哈希表
	if float64(h.size)/float64(h.capacity) > h.loadFactor {
		h.resize()
	}

	// 计算键的哈希值，并得到对应的索引
	hash := h.hash(key)
	index := int(hash & uint32(h.capacity-1))

	// 如果该索引对应的节点为nil，创建一个新的节点
	if h.data[index] == nil {
		h.data[index] = &entry[T, V]{key, value, nil}
		h.size++
		return
	}

	// 遍历该索引对应的链表或红黑树，查找是否存在相同的键
	for e := h.data[index]; e != nil; e = e.next {
		if e.key.CompareTo(key) == 0 {
			e.value = value // 如果存在相同的键，更新其对应的值
			return
		}
	}

	// 如果不存在相同的键，将键值对添加到该索引对应的链表或红黑树中
	h.data[index].next = &entry[T, V]{key, value, nil}
	h.size++
}

// 根据键获取哈希表中对应的值
func (h *HashMap[T, V]) Get(key T) (V, bool) {
	// 计算键的哈希值，并得到对应的索引
	hash := h.hash(key)
	index := int(hash & uint32(h.capacity-1))

	// 遍历该索引对应的链表或红黑树，查找是否存在相同的键
	for e := h.data[index]; e != nil; e = e.next {
		if e.key.CompareTo(key) == 0 {
			return e.value, true // 如果存在相同的键，返回其对应的值和true
		}
	}

	var zeroValue V         // 很关键
	return zeroValue, false // 如果不存在相同的键，返回值的零值和false
}

// 删除哈希表中指定键的键值对
func (h *HashMap[T, V]) Delete(key T) bool {
	// 计算键的哈希值，并得到对应的索引
	hash := h.hash(key)
	index := int(hash & uint32(h.capacity-1))

	// 遍历该索引对应的链表或红黑树，查找是否存在相同的键，并删除其对应的节点
	prev := h.data[index]
	for e := h.data[index]; e != nil; e = e.next {
		if e.key.CompareTo(key) == 0 {
			if prev == e {
				h.data[index] = e.next
			} else {
				prev.next = e.next
			}
			h.size--
			return true // 如果存在相同的键，删除其对应的节点并返回true
		}
		prev = e
	}

	return false // 如果不存在相同的键，返回false
}

// 计算键的哈希值
func (h *HashMap[T, V]) hash(key T) uint32 {
	switch reflect.TypeOf(key).Kind() {
	case reflect.Int:
		return uint32(reflect.ValueOf(key).Int())
	case reflect.String:
		var hash uint32 = 5381
		str := reflect.ValueOf(key).String()
		for i := 0; i < len(str); i++ {
			hash = ((hash << 5) + hash) + uint32(str[i])
		}
		return hash
	default:
		hash := fnv.New32()
		_, _ = hash.Write([]byte(fmt.Sprintf("%v", key)))
		return hash.Sum32()
	}
}

// 计算键在哈希表中对应的索引
func (h *HashMap[T, V]) getIndex(key T) int {
	hash := h.hash(key)
	return int(hash & uint32(h.capacity-1))
}

// 扩容哈希表
func (h *HashMap[T, V]) resize() {
	h.capacity = h.resizeCapacity
	h.resizeCapacity *= 2

	newData := make([]*entry[T, V], h.capacity)
	for i := 0; i < len(h.data); i++ {
		for e := h.data[i]; e != nil; e = e.next {
			index := h.getIndex(e.key)
			if newData[index] == nil {
				newData[index] = &entry[T, V]{e.key, e.value, nil}
			} else {
				newData[index].next = &entry[T, V]{e.key, e.value, newData[index].next}
			}
		}
	}

	h.data = newData
}

// 实现fmt.Stringer接口，将哈希表转换为字符串表示形式
func (h *HashMap[T, V]) String() string {
	var buf bytes.Buffer
	buf.WriteString("{")
	for i := 0; i < h.capacity; i++ {
		for e := h.data[i]; e != nil; e = e.next {
			if buf.Len() > 1 {
				buf.WriteString(", ")
			}
			buf.WriteString(fmt.Sprintf("%v: %v", e.key, e.value))
		}
	}
	buf.WriteString("}")
	return buf.String()
}

// 返回哈希表中元素的数量
func (h *HashMap[T, V]) Size() int {
	return h.size
}

// 清空哈希表中的所有元素
func (h *HashMap[T, V]) Clear() {
	h.data = make([]*entry[T, V], h.capacity)
	h.size = 0
}

// 检查哈希表是否为空
func (h *HashMap[T, V]) IsEmpty() bool {
	return h.size == 0
}

// 遍历哈希表中的所有元素，并对每个元素执行指定的操作
func (h *HashMap[T, V]) ForEach(fn func(key T, value V)) {
	for i := 0; i < h.capacity; i++ {
		for e := h.data[i]; e != nil; e = e.next {
			fn(e.key, e.value)
		}
	}
}
