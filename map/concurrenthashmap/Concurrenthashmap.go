package concurrenthashmap

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"reflect"
	"sync/atomic"
	"unsafe"
)

type entry[T comparable, V comparable] struct {
	key   T              // 键
	value V              // 值
	next  unsafe.Pointer // 指向下一个节点的指针
}

type ConcurrentHashMap[T comparable, V comparable] struct {
	data           []*entry[T, V] // 存储数据的切片
	capacity       int            // 哈希表容量
	size           int            // 哈希表中元素数量
	loadFactor     float64        // 负载因子
	threshold      int            // 下一次扩容的阈值
	resizeCapacity int            // 扩容后的容量
}

// 创建一个新的并发安全的哈希表
func NewConcurrentHashMap[T comparable, V comparable]() *ConcurrentHashMap[T, V] {
	return &ConcurrentHashMap[T, V]{
		data:           make([]*entry[T, V], 16),
		capacity:       16,
		size:           0,
		loadFactor:     0.75,
		threshold:      12,
		resizeCapacity: 32,
	}
}

// 将键值对添加到并发安全的哈希表中
func (h *ConcurrentHashMap[T, V]) Put(key T, value V) {
	// 如果负载因子大于阈值，扩容哈希表
	if float64(h.size)/float64(h.capacity) > h.loadFactor {
		h.resize()
	}

	// 计算键的哈希值，并得到对应的索引
	hash := h.hash(key)
	index := int(hash & uint32(h.capacity-1))

	// 创建一个新的节点
	newEntry := &entry[T, V]{key: key, value: value}

	// 循环直到成功插入节点或找到相同的键
	for {
		oldEntry := (*entry[T, V])(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&h.data[index]))))

		// 如果该索引对应的节点为nil，直接将新节点插入
		if oldEntry == nil {
			if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&h.data[index])), unsafe.Pointer(oldEntry), unsafe.Pointer(newEntry)) {
				atomic.AddInt32((*int32)(unsafe.Pointer(&h.size)), 1)
				return
			}
		}

		// 遍历该索引对应的链表，查找是否存在相同的键
		for e := oldEntry; e != nil; e = (*entry[T, V])(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&e.next)))) {
			if e.key == key {
				// 如果存在相同的键，更新其对应的值
				atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&e.value)), unsafe.Pointer(&value))
				return
			}
		}

		// 将新节点插入链表的头部
		newEntry.next = unsafe.Pointer(oldEntry)
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&h.data[index])), unsafe.Pointer(oldEntry), unsafe.Pointer(newEntry)) {
			atomic.AddInt32((*int32)(unsafe.Pointer(&h.size)), 1)
			return
		}
	}
}

// 根据键获取并发安全的哈希表中对应的值
func (h *ConcurrentHashMap[T, V]) Get(key T) (V, bool) {
	// 计算键的哈希值，并得到对应的索引
	hash := h.hash(key)
	index := int(hash & uint32(h.capacity-1))

	// 遍历该索引对应的链表，查找是否存在相同的键
	for e := (*entry[T, V])(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&h.data[index])))); e != nil; e = (*entry[T, V])(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&e.next)))) {
		if e.key == key {
			return e.value, true // 如果存在相同的键，返回其对应的值和true
		}
	}

	var zeroValue V         // 很关键
	return zeroValue, false // 如果不存在相同的键，返回值的零值和false
}

// 删除并发安全的哈希表中指定键的键值对
func (h *ConcurrentHashMap[T, V]) Delete(key T) bool {
	// 计算键的哈希值，并得到对应的索引
	hash := h.hash(key)
	index := int(hash & uint32(h.capacity-1))

	// 循环直到成功删除节点或找不到相同的键
	for {
		oldEntry := (*entry[T, V])(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&h.data[index]))))

		// 如果该索引对应的节点为nil，直接返回false
		if oldEntry == nil {
			return false
		}

		var prev *entry[T, V] = nil
		var found bool = false

		// 遍历该索引对应的链表，查找是否存在相同的键，并删除其对应的节点
		for e := oldEntry; e != nil; e = (*entry[T, V])(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&e.next)))) {
			if e.key == key {
				found = true
				if prev == nil {
					// 如果要删除的节点是链表的头节点
					if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&h.data[index])), unsafe.Pointer(oldEntry), unsafe.Pointer(e.next)) {
						atomic.AddInt32((*int32)(unsafe.Pointer(&h.size)), -1)
						return true
					}
				} else {
					// 如果要删除的节点不是链表的头节点
					if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&prev.next)), unsafe.Pointer(e), unsafe.Pointer(e.next)) {
						atomic.AddInt32((*int32)(unsafe.Pointer(&h.size)), -1)
						return true
					}
				}
			}
			prev = e
		}

		if !found {
			return false
		}
	}
}

// 计算键的哈希值
func (h *ConcurrentHashMap[T, V]) hash(key T) uint32 {
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

// 扩容并发安全的哈希表
func (h *ConcurrentHashMap[T, V]) resize() {
	h.capacity = h.resizeCapacity
	h.resizeCapacity *= 2

	newData := make([]*entry[T, V], h.capacity)

	for i := 0; i < len(h.data); i++ {
		for e := h.data[i]; e != nil; e = (*entry[T, V])(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&e.next)))) {
			index := h.getIndex(e.key)
			if newData[index] == nil {
				newData[index] = &entry[T, V]{e.key, e.value, nil}
			} else {
				newData[index].next = unsafe.Pointer(e)
			}
		}
	}

	h.data = newData
}

// 计算键在并发安全的哈希表中对应的索引
func (h *ConcurrentHashMap[T, V]) getIndex(key T) int {
	hash := h.hash(key)
	return int(hash & uint32(h.capacity-1))
}

// 实现fmt.Stringer接口，将并发安全的哈希表转换为字符串表示形式
func (h *ConcurrentHashMap[T, V]) String() string {
	var buf bytes.Buffer
	buf.WriteString("{")
	for i := 0; i < h.capacity; i++ {
		for e := (*entry[T, V])(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&h.data[i])))); e != nil; e = (*entry[T, V])(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&e.next)))) {
			if buf.Len() > 1 {
				buf.WriteString(", ")
			}
			buf.WriteString(fmt.Sprintf("%v: %v", e.key, e.value))
		}
	}
	buf.WriteString("}")
	return buf.String()
}

// 返回并发安全的哈希表中元素的数量
func (h *ConcurrentHashMap[T, V]) Size() int {
	return int(atomic.LoadInt32((*int32)(unsafe.Pointer(&h.size))))
}
