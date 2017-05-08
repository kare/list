package list_test

import (
	"testing"

	"kkn.fi/list"
)

func TestIntEquals(t *testing.T) {
	l := list.NewInt()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	same := list.NewIntWith(1, 2, 3)
	if !l.Equals(same) {
		t.Errorf("expected list %v and %v to be equal", *l, *same)
	}
	diff := list.NewIntWith(1, 5, 3)
	if diff.Equals(same) {
		t.Errorf("expected list %v and %v not to be equal", *diff, *same)
	}
	difflen := list.NewIntWith(1, 2, 3, 4)
	if same.Equals(difflen) {
		t.Error("lists of different len equal")
	}
	if diff.Equals(1) {
		t.Error("list equals to different type")
	}
}

func TestIntNew(t *testing.T) {
	l := list.NewIntWith(1, 2, 3)
	var expected = []int{1, 2, 3}
	if len(l.Slice()) != 3 {
		t.Errorf("expected len 3, but got %d", len(l.Slice()))
	}
	for i, v := range l.Slice() {
		if v != expected[i] {
			t.Errorf("expected list to contain '%v'", v)
		}
	}
}

func TestIntGet(t *testing.T) {
	l := list.NewInt()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	if l.Get(0) != 1 {
		t.Errorf("expected element %d to contain '%v', but got '%v'", 0, 1, l.Get(0))
	}
	if l.Get(1) != 2 {
		t.Errorf("expected element %d to contain '%v', but got '%v'", 1, 2, l.Get(1))
	}
	if l.Get(2) != 3 {
		t.Errorf("expected element %d to contain '%v', but got '%v'", 2, 3, l.Get(2))
	}
}

func TestIntLen(t *testing.T) {
	l := list.NewInt()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	if l.Len() != 3 {
		t.Error("expected len 3, but got", l.Len())
	}
}

func TestIntIsEmpty(t *testing.T) {
	l := list.NewInt()
	if !l.IsEmpty() {
		t.Error("expected empty list")
	}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	if l.IsEmpty() {
		t.Error("expected non empty list")
	}
}

func TestIntContains(t *testing.T) {
	l := list.NewInt()
	if l.Contains(1) {
		t.Error("expected list to not contain 1, but it did.")
	}
	l.Add(2)
	if !l.Contains(2) {
		t.Error("expected list to contain 2, but it did not.")
	}
	l.Add(1)
	l.Add(3)
	if l.IsEmpty() {
		t.Error("expected non empty list")
	}
}

func TestIntRemove(t *testing.T) {
	l := list.NewInt()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	if !l.Remove(2) {
		t.Error("expected set to contain 2")
	}
	if l.Len() != 2 {
		t.Error("expected list len to be 2, but got", l.Len())
	}
	if l.Contains(2) {
		t.Error("did not expect the list to contain 2")
	}
	if l.Remove(2) {
		t.Error("value was already removed")
	}
}

func TestIntSlice(t *testing.T) {
	l := list.NewInt()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	var expected = []int{1, 2, 3}
	if len(l.Slice()) != 3 {
		t.Errorf("expected len 3, but got %d", len(l.Slice()))
	}
	for i, v := range l.Slice() {
		if v != expected[i] {
			t.Errorf("expected list to contain '%v'", v)
		}
	}
}

func TestIntAddAll(t *testing.T) {
	l := list.NewInt()
	l.AddAll(1, 2, 3)
	var expected = []int{1, 2, 3}
	if len(l.Slice()) != 3 {
		t.Errorf("expected len 3, but got %d", len(l.Slice()))
	}
	for i, v := range l.Slice() {
		if v != expected[i] {
			t.Errorf("expected list to contain '%v'", v)
		}
	}
}
