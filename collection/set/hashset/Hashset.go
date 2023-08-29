package hashset

import (
	"fmt"
	"gitee.com/herry-hu/go-collections-java/map/hashmap"
	"strings"
	"sync"
)

// HashSet 是一个线程安全的哈希集合
type HashSet[T comparable] struct {
	items *hashmap.HashMap[T, int] // 存储元素的哈希集合
	lock  sync.RWMutex             // 用于保护集合的读写锁
}

// NewHashSet 创建一个新的HashSet
func NewHashSet[T comparable]() *HashSet[T] {
	hashMap := hashmap.NewHashMap[T, int]()
	return &HashSet[T]{
		items: hashMap,
	}
}

// Add 将元素添加到HashSet
func (set *HashSet[T]) Add(key T) {
	set.lock.Lock()
	defer set.lock.Unlock()

	set.items.Put(key, 1)
}

// Contains 检查HashSet中是否包含指定的元素
func (set *HashSet[T]) Contains(key T) bool {
	set.lock.RLock()
	defer set.lock.RUnlock()

	_, found := set.items.Get(key)
	return found
}

// Remove 从HashSet中移除指定的元素
func (set *HashSet[T]) Remove(key T) {
	set.lock.Lock()
	defer set.lock.Unlock()

	set.items.Delete(key)
}

// Size 返回HashSet中的元素数量
func (set *HashSet[T]) Size() int {
	set.lock.RLock()
	defer set.lock.RUnlock()

	return set.items.Size()
}

// Clear 清空HashSet中的所有元素
func (set *HashSet[T]) Clear() {
	set.lock.Lock()
	defer set.lock.Unlock()

	set.items.Clear()
}

// IsEmpty 检查HashSet是否为空
func (set *HashSet[T]) IsEmpty() bool {
	set.lock.RLock()
	defer set.lock.RUnlock()

	return set.items.IsEmpty()
}

// String 返回HashSet的字符串表示形式
func (set *HashSet[T]) String() string {
	set.lock.RLock()
	defer set.lock.RUnlock()

	var items []string
	set.items.ForEach(func(key T, _ int) {
		items = append(items, fmt.Sprintf("%v", key))
	})
	return fmt.Sprintf("HashSet{%s}", strings.Join(items, ", "))
}

// Iterator 返回一个只读通道，用于遍历HashSet中的元素
func (set *HashSet[T]) Iterator() <-chan T {
	ch := make(chan T)

	go func() {
		set.lock.RLock()
		defer set.lock.RUnlock()

		set.items.ForEach(func(key T, _ int) {
			ch <- key
		})
		close(ch)
	}()

	return ch
}
