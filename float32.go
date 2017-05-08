// Code generated by "gends"; DO NOT EDIT.

package list // import "kkn.fi/list"

// Float32List holds an ordered collection of float32's.
type Float32List struct {
	values []float32
}

// NewFloat32 creates an empty list of Float32List's.
func NewFloat32() *Float32List {
	return &Float32List{}
}

// NewFloat32With creates a list of Float32List's filled with given
// values.
func NewFloat32With(values ...float32) *Float32List {
	return &Float32List{
		values: values,
	}
}

// Equals checks list equality given other list. In order to be equal both lists
// must be of same type.
func (l *Float32List) Equals(other interface{}) bool {
	list, ok := other.(*Float32List)
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
func (l *Float32List) Add(value float32) {
	l.values = append(l.values, value)
}

// AddAll appends all given values to the end of the list.
func (l *Float32List) AddAll(values ...float32) {
	l.values = append(l.values, values...)
}

// Get returns the float32 element at given index.
func (l *Float32List) Get(index int) float32 {
	return l.values[index]
}

// Len returns the size of the list.
func (l *Float32List) Len() int {
	return len(l.values)
}

// IsEmpty return true if list is empty, false otherwise.
func (l *Float32List) IsEmpty() bool {
	return len(l.values) == 0
}

// Contains returns true if given value is contained by the list, false
// otherwise.
func (l *Float32List) Contains(value float32) bool {
	for _, v := range l.values {
		if v == value {
			return true
		}
	}
	return false
}

// Slice returns float32 slice of the elements contained within the list.
func (l *Float32List) Slice() []float32 {
	return l.values
}

// Remove deletes the given value from the list. Returns true if the list
// contained specified value, false otherwise.
func (l *Float32List) Remove(value float32) bool {
	for i, v := range l.values {
		if v == value {
			l.values = append(l.values[:i], l.values[i+1:]...)
			return true
		}
	}
	return false
}
