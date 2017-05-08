// Code generated by "gends"; DO NOT EDIT.

package list // import "kkn.fi/list"

// RuneList holds an ordered collection of rune's.
type RuneList struct {
	values []rune
}

// NewRune creates an empty list of RuneList's.
func NewRune() *RuneList {
	return &RuneList{}
}

// NewRuneWith creates a list of RuneList's filled with given
// values.
func NewRuneWith(values ...rune) *RuneList {
	return &RuneList{
		values: values,
	}
}

// Equals checks list equality given other list. In order to be equal both lists
// must be of same type.
func (l *RuneList) Equals(other interface{}) bool {
	list, ok := other.(*RuneList)
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
func (l *RuneList) Add(value rune) {
	l.values = append(l.values, value)
}

// AddAll appends all given values to the end of the list.
func (l *RuneList) AddAll(values ...rune) {
	l.values = append(l.values, values...)
}

// Get returns the rune element at given index.
func (l *RuneList) Get(index int) rune {
	return l.values[index]
}

// Len returns the size of the list.
func (l *RuneList) Len() int {
	return len(l.values)
}

// IsEmpty return true if list is empty, false otherwise.
func (l *RuneList) IsEmpty() bool {
	return len(l.values) == 0
}

// Contains returns true if given value is contained by the list, false
// otherwise.
func (l *RuneList) Contains(value rune) bool {
	for _, v := range l.values {
		if v == value {
			return true
		}
	}
	return false
}

// Slice returns rune slice of the elements contained within the list.
func (l *RuneList) Slice() []rune {
	return l.values
}

// Remove deletes the given value from the list. Returns true if the list
// contained specified value, false otherwise.
func (l *RuneList) Remove(value rune) bool {
	for i, v := range l.values {
		if v == value {
			l.values = append(l.values[:i], l.values[i+1:]...)
			return true
		}
	}
	return false
}
