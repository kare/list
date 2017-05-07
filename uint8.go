// Code generated by "gends"; DO NOT EDIT.

package list // import "kkn.fi/list"

// Uint8List holds an ordered collection of uint8's.
type Uint8List struct {
	values []uint8
}

// NewUint8 creates an empty list of Uint8List's.
func NewUint8() *Uint8List {
	return &Uint8List{}
}

// NewUint8With creates a list of Uint8List's filled with given
// values.
func NewUint8With(values ...uint8) *Uint8List {
	return &Uint8List{
		values: values,
	}
}

// Equals checks list equality given other list. In order to be equal both lists
// must be of same type.
func (l Uint8List) Equals(other interface{}) bool {
	list, ok := other.(*Uint8List)
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
func (l *Uint8List) Add(value uint8) {
	l.values = append(l.values, value)
}

// AddAll appends all given values to the end of the list.
func (l *Uint8List) AddAll(values ...uint8) {
	l.values = append(l.values, values...)
}

// Get returns the uint8 element at given index.
func (l Uint8List) Get(index int) uint8 {
	return l.values[index]
}

// Len returns the size of the list.
func (l Uint8List) Len() int {
	return len(l.values)
}

// IsEmpty return true if list is empty, false otherwise.
func (l Uint8List) IsEmpty() bool {
	return len(l.values) == 0
}

// Contains returns true if given value is contained by the list, false
// otherwise.
func (l Uint8List) Contains(value uint8) bool {
	for _, v := range l.values {
		if v == value {
			return true
		}
	}
	return false
}

// Slice returns uint8 slice of the elements contained within the list.
func (l Uint8List) Slice() []uint8 {
	result := make([]uint8, 0, len(l.values))
	result = append(result, l.values...)
	return result
}

// Remove deletes the given value from the list. Returns true if the list
// contained specified value, false otherwise.
func (l *Uint8List) Remove(value uint8) bool {
	for i, v := range l.values {
		if v == value {
			l.values = append(l.values[:i], l.values[i+1:]...)
			return true
		}
	}
	return false
}
