package ordered_map

import (
	"sort"
	"reflect"
)

type CompareFunction func(interface{}, interface{}) int8

type Keys interface {
	sort.Interface
	Add(k interface{}) bool
	Remove(k interface{}) bool
	Clear()
	Get(index int) interface{}
	GetAll() []interface{}
	Search(e interface{}) (index int, contains bool)
	ElemType() reflect.Type
	CompareFunc() CompareFunction
}

type myKeys struct {
	container   []interface{}
	compareFunc CompareFunction
	elemType    reflect.Type
}

func (self *myKeys) Len() int {
	return len(self.container)
}
func (self *myKeys) Less(i, j int) bool {
	return self.compareFunc(self.container[i], self.container[j]) == -1
}
func (self *myKeys) Swap(i, j int) {
	self.container[i], self.container[j] = self.container[j], self.container[i]
}
func (self *myKeys) IsAcceptableElem(e interface{}) bool {
	if e == nil || reflect.TypeOf(e) != self.elemType {
		return false
	}
	return true
}

func (self *myKeys) Add(e interface{}) bool {
	if !self.IsAcceptableElem(e) {
		return false
	}
	self.container = append(self.container, e)
	sort.Sort(self)
	return true
}

//func Search(n int, f func(int) bool) int {
//	// Define f(-1) == false and f(n) == true.
//	// Invariant: f(i-1) == false, f(j) == true.
//	i, j := 0, n
//	for i < j {
//		h := int(uint(i+j) >> 1) // avoid overflow when computing h
//		// i ≤ h < j
//		if !f(h) {
//			i = h + 1 // preserves f(i-1) == false
//		} else {
//			j = h // preserves f(j) == true
//		}
//	}
//	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
//	return i
//}
func (self *myKeys) Search(e interface{}) (index int, contains bool) {
	if self.IsAcceptableElem(e) {
		_i := sort.Search(self.Len(), func(i int) bool {
			return self.compareFunc(self.container[i], e) > 0
		})
		if _i < self.Len() && self.container[_i] == e {
			index = _i
			contains = true
		} else {
			contains = false
		}
	} else {
		contains = false
	}
	return
}

func (self *myKeys) Remove(k interface{}) bool {
	var _i, contains = self.Search(k)
	if !contains {
		return false
	}
	self.container = append(self.container[0:_i], self.container[_i+1:]...)
	return true
}
func (self *myKeys) Clear() {
	self.container = make([]interface{}, 0)
}
func (self *myKeys) Get(index int) interface{} {
	if index >= self.Len() {
		return nil
	}
	return self.container[index]
}
func (self *myKeys) GetAll() []interface{} {
	initial_len := self.Len()
	snapshot := make([]interface{}, initial_len)
	actual_len := 0 // for concurrent-safe, but not totally
	for _, key := range self.container {
		if actual_len < initial_len {
			snapshot[actual_len] = key
		} else {
			snapshot = append(snapshot, key)
		}
		actual_len ++
	}
	if actual_len < initial_len {
		snapshot = snapshot[:actual_len]

	}
	return snapshot
}

func (self *myKeys) CompareFunc() CompareFunction {
	return self.compareFunc
}
func (self *myKeys) ElemType() reflect.Type {
	return self.elemType
}
func NewKeys(compareFunc CompareFunction, elemType reflect.Type) (Keys) {
	return &myKeys{
		make([]interface{}, 0),
		compareFunc,
		elemType,
	}
}

var (
	int64Keys = &myKeys{
		container: make([]interface{}, 0),
		compareFunc: func(e1 interface{}, e2 interface{}) int8 {
			k1 := e1.(int64)
			k2 := e2.(int64)
			if k1 < k2 {
				return -1
			} else if k1 > k2 {
				return 1
			} else {
				return 0
			}
		},
		elemType: reflect.TypeOf(int64(1))}
)
