package set

import (
	"bytes"
	"fmt"
)

type HashSet struct {
	// value: bool类型,是否存在
	// key: 支持任意类型的元素
	m map[interface{}]bool
}

func NewHashSet() *HashSet {
	// to escape new(HashSet).m -> nil
	return &HashSet{m: make(map[interface{}]bool)}
}

func (self *HashSet) Add(e interface{}) bool {
	if !self.m[e] {
		self.m[e] = true
		return true
	}
	return false
}

func (self *HashSet) Remove(e interface{}) {
	delete(self.m, e)
}

func (self *HashSet) Clear() {
	// assign new value, instead of traversing which is not concurrent-safe
	self.m = make(map[interface{}]bool)
}

func (self *HashSet) Contains(e interface{}) bool {
	return self.m[e]
}

func (self *HashSet) Len() int {
	return len(self.m)
}
func (self *HashSet) Same(other *HashSet) bool {
	if other == nil {
		return false
	}
	if self.Len() != other.Len() {
		return false
	}
	for key := range self.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}
func (self *HashSet) Elements() []interface{} {
	initial_len := self.Len()
	snapshot := make([]interface{}, initial_len)
	actual_len := 0 // for concurrent-safe, but not totally
	for key := range self.m {
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

func (self *HashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("Set{")
	first := true
	for key := range self.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")
	return buf.String()
}

func (self *HashSet) IsSuperset(other *HashSet) bool {
	if other == nil {
		return false
	}
	one_len := self.Len()
	other_len := other.Len()

	if one_len <= other_len {
		// 超集当然要比子集大，包括了为0的情形

		return false
	}
	if one_len > 0 && other_len == 0 {
		return true
	}
	for _, v := range other.Elements() {
		if ! self.Contains(v) {
			return false
		}
	}
	return true
}

func (self *HashSet) Union(other *HashSet) *HashSet {
	//TOOD

	return nil
}
func (self *HashSet) Intersect(other *HashSet) *HashSet {
	//TODO
	return nil
}
func (self *HashSet) Difference(other *HashSet) *HashSet {
	//TODO
	return nil
}
func (self *HashSet) SymmetricDifference(other *HashSet) *HashSet {
	//TODO
	return nil
}

type Set interface {
	Add(e interface{}) bool
	Remove(e interface{})
	Clear()
	Contains(e interface{}) bool
	Len() int
	Same(other *Set) bool
	Elements() []interface{}
	String() string
}
