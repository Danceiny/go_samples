package ordered_map

import (
	_ "sort"
	"reflect"
)

type myOrderedMap struct {
	keys     Keys
	elemType reflect.Type
	m        map[interface{}]interface{}
}

type OrderedMap interface {
	Get(key interface{}) interface{}
	Put(key interface{}, elem interface{}) (interface{}, bool)
	Remove(key interface{}) interface{}
	Clear()
	Len() int
	Contains(key interface{}) bool
	FirstKey() interface{}
	LastKey() interface{}
	HeadMap(toKey interface{}) OrderedMap
	SubMap(fromKey, toKey interface{}) OrderedMap
	TailMap(fromKey interface{}) OrderedMap
	Keys() []interface{}
	Elems() []interface{}
	ToMap() map[interface{}]interface{}
	KeyType() reflect.Type
	ElemType() reflect.Type
}

//// sort.Interface
//// A type, typically a collection, that satisfies sort.Interface can be
//// sorted by the routines in this package. The methods require that the
//// elements of the collection be enumerated by an integer index.
//type Interface interface {
//	// Len is the number of elements in the collection.
//	Len() int
//	// Less reports whether the element with
//	// index i should sort before the element with index j.
//	Less(i, j int) bool
//	// Swap swaps the elements with indexes i and j.
//	Swap(i, j int)
//}

func (self *myOrderedMap) Len() int {
	return self.keys.Len()
}

func (self *myOrderedMap) Less(i, j int) bool {
	return self.keys.CompareFunc()(self.m[self.keys.Get(i)], self.m[self.keys.Get(j)]) < 0

}
func (self *myOrderedMap) Swap(i, j int) {
	self.m[i], self.m[j] = self.m[j], self.m[i]
}
func NewOrderedMap(elementType reflect.Type, ) {

}
