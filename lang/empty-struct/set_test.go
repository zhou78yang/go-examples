package empty_struct

import "testing"

func TestSet(t *testing.T) {
	s := NewSet(8)
	s.Add("hello")
	s.Add("world")
	if !s.Has("hello") {
		t.Errorf("has hello: exepct true, got false\n")
	}
	if s.Has("Hello") {
		t.Errorf("has Hello: expect false, got true\n")
	}
	l := s.Len()
	if l != 2 {
		t.Errorf("set Len: expect 2, got %d\n", l)
	}
	s.Remote("hello")
	if s.Has("hello") {
		t.Errorf("has hello: expect false, got true\n")
	}
	l = s.Len()
	if l != 1 {
		t.Errorf("set len: expect 1, got %d\n", l)
	}
}
