// Code generated by "gends"; DO NOT EDIT.

package list // import "kkn.fi/list"

// Int64List holds an ordered collection of int64's.
type Int64List struct {
	values []int64
}

// NewInt64 creates an empty list of Int64List's.
func NewInt64() *Int64List {
	return &Int64List{}
}

// NewInt64With creates a list of Int64List's filled with given
// values.
func NewInt64With(values ...int64) *Int64List {
	return &Int64List{
		values: values,
	}
}

// Equals checks list equality given other list. In order to be equal both lists
// must be of same type.
func (l *Int64List) Equals(other interface{}) bool {
	list, ok := other.(*Int64List)
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
func (l *Int64List) Add(value int64) {
	l.values = append(l.values, value)
}

// AddAll appends all given values to the end of the list.
func (l *Int64List) AddAll(values ...int64) {
	l.values = append(l.values, values...)
}

// Get returns the int64 element at given index.
func (l *Int64List) Get(index int) int64 {
	return l.values[index]
}

// Len returns the size of the list.
func (l *Int64List) Len() int {
	return len(l.values)
}

// IsEmpty return true if list is empty, false otherwise.
func (l *Int64List) IsEmpty() bool {
	return len(l.values) == 0
}

// Contains returns true if given value is contained by the list, false
// otherwise.
func (l *Int64List) Contains(value int64) bool {
	for _, v := range l.values {
		if v == value {
			return true
		}
	}
	return false
}

// Slice returns int64 slice of the elements contained within the list.
func (l *Int64List) Slice() []int64 {
	return l.values
}

// Remove deletes the given value from the list. Returns true if the list
// contained specified value, false otherwise.
func (l *Int64List) Remove(value int64) bool {
	for i, v := range l.values {
		if v == value {
			l.values = append(l.values[:i], l.values[i+1:]...)
			return true
		}
	}
	return false
}
