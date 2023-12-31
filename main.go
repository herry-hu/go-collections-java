package main

import (
	"fmt"
	"github.com/herry-hu/go-collections-java/collection/list/arraylist"
	"github.com/herry-hu/go-collections-java/collection/list/doublelinkedlist"
	"github.com/herry-hu/go-collections-java/collection/list/linkedlist"
	"github.com/herry-hu/go-collections-java/collection/set/hashset"
	"github.com/herry-hu/go-collections-java/collection/set/treeset"
	"github.com/herry-hu/go-collections-java/lang"
	"github.com/herry-hu/go-collections-java/map/concurrenthashmap"
	"github.com/herry-hu/go-collections-java/map/hashmap"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) CompareTo(other interface{}) int {
	if otherPerson, ok := other.(Person); ok {
		if p.Age < otherPerson.Age {
			return -1
		} else if p.Age > otherPerson.Age {
			return 1
		}
		return 0
	}
	panic("Cannot compare different types")
}

func main() {
	//arraylist
	list := arraylist.NewArrayList[lang.Int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	fmt.Println(list)
	list.Remove(1)
	fmt.Println(list)
	listc := arraylist.NewArrayList[int]()
	listc.Add(1)
	listc.Add(2)
	listc.Add(3)
	fmt.Println(listc)
	listc.Remove(1)
	fmt.Println(listc)

	ints := []lang.Int{1, 2, 3, 4}
	list.AddAll(ints...)
	fmt.Println(list)
	for index, value := range ints {
		fmt.Printf("index:value = %v:%v \n", index, value)
	}

	persons := arraylist.NewArrayList[Person]()
	persons.Add(Person{"张三", 20})
	persons.Add(Person{"李四", 25})
	persons.Add(Person{"王五", 13})
	fmt.Println(persons)

	//linkedlist
	linkedList := linkedlist.NewLinkedList[lang.Int]()
	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)
	fmt.Println(linkedList)
	linkedList.Remove(1)
	fmt.Println(linkedList)
	linkedListc := linkedlist.NewLinkedList[int]()
	linkedListc.Add(1)
	linkedListc.Add(2)
	linkedListc.Add(3)
	fmt.Println(linkedListc)
	linkedListc.Remove(1)
	fmt.Println(linkedListc)

	//doublelinkedlist
	do := doublelinkedlist.NewDoubleLinkedList[lang.Int]()
	do.Add(1)
	do.Add(2)
	do.Add(3)
	fmt.Println(do)
	do.Remove(2)
	fmt.Println(do)
	doc := doublelinkedlist.NewDoubleLinkedList[int]()
	doc.Add(1)
	doc.Add(2)
	doc.Add(3)
	fmt.Println(doc)
	doc.Remove(2)
	fmt.Println(doc)

	//hashmap
	hashMap := hashmap.NewHashMap[lang.String, lang.Int]()
	hashMap.Put("a", 123)
	hashMap.Put("b", 234)
	hashMap.Put("c", 345)
	fmt.Println(hashMap)
	hm, _ := hashMap.Get("a")
	fmt.Println(hm)
	hashMap.Delete("a")
	fmt.Println(hashMap)
	hashMapc := hashmap.NewHashMap[string, int]()
	hashMapc.Put("a", 123)
	hashMapc.Put("b", 234)
	hashMapc.Put("c", 345)
	fmt.Println(hashMapc)
	hmc, _ := hashMapc.Get("a")
	fmt.Println(hmc)
	hashMapc.Delete("a")
	fmt.Println(hashMapc)

	//并发安全hashmap
	//hashmap
	coHashmap := concurrenthashmap.NewConcurrentHashMap[lang.String, lang.Int]()
	coHashmap.Put("a", 123)
	coHashmap.Put("b", 234)
	coHashmap.Put("c", 345)
	fmt.Println(coHashmap)
	co, _ := coHashmap.Get("a")
	fmt.Println(co)
	coHashmap.Delete("a")
	fmt.Println(coHashmap)
	coHashmapc := concurrenthashmap.NewConcurrentHashMap[string, int]()
	coHashmapc.Put("a", 123)
	coHashmapc.Put("b", 234)
	coHashmapc.Put("c", 345)
	fmt.Println(coHashmapc)
	coc, _ := coHashmapc.Get("a")
	fmt.Println(coc)
	coHashmapc.Delete("a")
	fmt.Println(coHashmapc)

	//hashset
	set := hashset.NewHashSet[lang.String]()
	set.Add("apple")
	set.Add("banana")
	set.Add("orange")
	set.Add("apple")
	fmt.Println(set)
	setc := hashset.NewHashSet[string]()
	setc.Add("apple")
	setc.Add("banana")
	setc.Add("orange")
	setc.Add("apple")
	fmt.Println(setc)

	//treeset:只支持实现了compareTo的类型
	tree := treeset.NewTreeSet[lang.String]()
	tree.Add("apple")
	tree.Add("banana")
	tree.Add("orange")
	tree.Add("apple")
	fmt.Println(tree)

}
