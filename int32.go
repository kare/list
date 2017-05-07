// Code generated by "gends"; DO NOT EDIT.

package list // import "kkn.fi/list"

// Int32List holds an ordered collection of int32's.
type Int32List struct {
	values []int32
}

// NewInt32 creates an empty list of Int32List's.
func NewInt32() *Int32List {
	return &Int32List{}
}

// NewInt32With creates a list of Int32List's filled with given
// values.
func NewInt32With(values ...int32) *Int32List {
	return &Int32List{
		values: values,
	}
}

// Equals checks list equality given other list. In order to be equal both lists
// must be of same type.
func (l Int32List) Equals(other interface{}) bool {
	list, ok := other.(*Int32List)
	if !ok {
		return false
	}
	if len(l.values) != len(list.values) {
		return false
	}
	for i, v := range list.Slice() {
		if v != l.values[i] {
			return false
		}
	}
	return true

}

// Add appends a value to the end of the list.
func (l *Int32List) Add(value int32) {
	l.values = append(l.values, value)
}

// AddAll appends all given values to the end of the list.
func (l *Int32List) AddAll(values ...int32) {
	l.values = append(l.values, values...)
}

// Get returns the int32 element at given index.
func (l Int32List) Get(index int) int32 {
	return l.values[index]
}

// Len returns the size of the list.
func (l Int32List) Len() int {
	return len(l.values)
}

// IsEmpty return true if list is empty, false otherwise.
func (l Int32List) IsEmpty() bool {
	return len(l.values) == 0
}

// Contains returns true if given value is contained by the list, false
// otherwise.
func (l Int32List) Contains(value int32) bool {
	for _, v := range l.values {
		if v == value {
			return true
		}
	}
	return false
}

// Slice returns int32 slice of the elements contained within the list.
func (l Int32List) Slice() []int32 {
	result := make([]int32, 0, len(l.values))
	result = append(result, l.values...)
	return result
}

// Remove deletes the given value from the list. Returns true if the list
// contained specified value, false otherwise.
func (l *Int32List) Remove(value int32) bool {
	for i, v := range l.values {
		if v == value {
			l.values = append(l.values[:i], l.values[i+1:]...)
			return true
		}
	}
	return false
}
