// Code generated by "gends"; DO NOT EDIT.

package list // import "kkn.fi/list"

// ByteList holds an ordered collection of byte's.
type ByteList struct {
	values []byte
}

// NewByte creates an empty list of ByteList's.
func NewByte() *ByteList {
	return &ByteList{}
}

// NewByteWith creates a list of ByteList's filled with given
// values.
func NewByteWith(values ...byte) *ByteList {
	return &ByteList{
		values: values,
	}
}

// Equals checks list equality given other list. In order to be equal both lists
// must be of same type.
func (l ByteList) Equals(other interface{}) bool {
	list, ok := other.(*ByteList)
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
func (l *ByteList) Add(value byte) {
	l.values = append(l.values, value)
}

// AddAll appends all given values to the end of the list.
func (l *ByteList) AddAll(values ...byte) {
	l.values = append(l.values, values...)
}

// Get returns the byte element at given index.
func (l ByteList) Get(index int) byte {
	return l.values[index]
}

// Len returns the size of the list.
func (l ByteList) Len() int {
	return len(l.values)
}

// IsEmpty return true if list is empty, false otherwise.
func (l ByteList) IsEmpty() bool {
	return len(l.values) == 0
}

// Contains returns true if given value is contained by the list, false
// otherwise.
func (l ByteList) Contains(value byte) bool {
	for _, v := range l.values {
		if v == value {
			return true
		}
	}
	return false
}

// Slice returns byte slice of the elements contained within the list.
func (l ByteList) Slice() []byte {
	result := make([]byte, 0, len(l.values))
	result = append(result, l.values...)
	return result
}

// Remove deletes the given value from the list. Returns true if the list
// contained specified value, false otherwise.
func (l *ByteList) Remove(value byte) bool {
	for i, v := range l.values {
		if v == value {
			l.values = append(l.values[:i], l.values[i+1:]...)
			return true
		}
	}
	return false
}