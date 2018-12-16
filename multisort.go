package multisort

import (
	"sort"
)

// MultiSorter implements the Sort interface,
// sorting the two dimensional string slices within.
type MultiSorter struct {
	list   []interface{}
	lesses []LessFunc
}

// LessFunc reports whether the element with
// index i should sort before the element with index j.
type LessFunc func(i, j interface{}) bool

// Len is part of sort.Interface.
func (ms *MultiSorter) Len() int {
	return len(ms.list)
}

// Swap is part of sort.Interface.
func (ms *MultiSorter) Swap(i, j int) {
	ms.list[i], ms.list[j] = ms.list[j], ms.list[i]
}

// Less is part of sort.Interface.
func (ms *MultiSorter) Less(i, j int) bool {
	var k int
	for k = 0; k < len(ms.lesses)-1; k++ {
		switch {
		case ms.lesses[k](ms.list[i], ms.list[j]):
			return true
		case ms.lesses[k](ms.list[j], ms.list[i]):
			return false
		}
	}
	return ms.lesses[k](ms.list[i], ms.list[j])
}

// Sort sorts the list according to lessFunc.
func (ms *MultiSorter) Sort(list []interface{}) {
	ms.list = list
	sort.Sort(ms)
}

// NewMultiSorter create the MultiSorter with given lessfunc.
func NewMultiSorter(lesses ...LessFunc) *MultiSorter {
	return &MultiSorter{lesses: lesses}
}
