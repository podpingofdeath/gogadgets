package set

import (
	"fmt"
	"testing"
)

type E struct {
	Ts  int64
	Typ int
	Val string
}

func (t *E) Hash() string {
	return fmt.Sprintf("%d|%s", t.Ts, t.Val)
}

func TestSet_Add(t *testing.T) {
	s := NewSet()
	es := []E{
		{1, 2, "F16944C8-5006-4887-A0B7-D2928C358009"},
		{1, 2, "F16944C8-5006-4887-A0B7-D2928C358009"},
		{2, 2, "0ee35c472c51377c3378d193d3406bb3"},
		{2, 2, "0ee35c472c51377c3378d193d3406bb3"},
	}

	for i := range es {
		s.Add(&es[i])
	}

	if s.Size() != 2 {
		t.Error("Set.Add error")
	}
	t.Logf("%v,%d\n", s.Elements, s.Size())

	s.Remove(&E{1, 2, "F16944C8-5006-4887-A0B7-D2928C358009"})
	if s.Size() != 1 {
		t.Error("Set.Remove error")
	}
	t.Logf("%v,%d\n", s.Elements, s.Size())

	s.Clean()
	if s.Size() != 0 {
		t.Error("Set.Clean error")
	}
	t.Logf("%v,%d\n", s.Elements, s.Size())
}

func TestSet_Union(t *testing.T) {
	s := NewSet()
	s.Add(&E{1, 2, "F512AD45-3F72-443D-BB15-63B923645AA1"})
	s.Add(&E{2, 2, "0ee35c472c51377c3378d193d3406bb3"})
	t.Logf("%v,%d\n", s.Elements, s.Size())

	s1 := NewSet()
	s1.Add(&E{1, 2, "F16944C8-5006-4887-A0B7-D2928C358009"})
	s.Add(&E{2, 2, "0ee35c472c51377c3378d193d3406bb3"})
	t.Logf("%v,%d\n", s1.Elements, s1.Size())

	s2 := NewSet()
	s2.Add(&E{2, 2, "0ee35c472c51377c3378d193d3406bb3"})
	t.Logf("%v,%d\n", s2.Elements, s2.Size())

	if s.Union(s1).Size() != 3 {
		t.Error("Set.Union error")
	}
	t.Logf("%v,%d\n", s.Elements, s.Size())

	if s.Join(s2).Size() != 1 {
		t.Error("Set.Join error")
	}
	t.Logf("%v,%d\n", s.Elements, s.Size())
}

func TestUnion(t *testing.T) {
	s := NewSet()
	s.Add(&E{1, 2, "F512AD45-3F72-443D-BB15-63B923645AA1"})
	s.Add(&E{2, 2, "0ee35c472c51377c3378d193d3406bb3"})
	t.Logf("%v,%d\n", s.Elements, s.Size())

	s1 := NewSet()
	s1.Add(&E{1, 2, "F16944C8-5006-4887-A0B7-D2928C358009"})
	s1.Add(&E{2, 2, "0ee35c472c51377c3378d193d3406bb3"})
	t.Logf("%v,%d\n", s1.Elements, s1.Size())

	s2 := Union(s, s1)
	if s2.Size() != 3 {
		t.Error("Union error")
	}
	t.Logf("%v,%d\n", s2.Elements, s2.Size())

	s3 := Join(s, s1)
	if s3.Size() != 1 {
		t.Error("Join error")
	}
	t.Logf("%v,%d\n", s3.Elements, s3.Size())
}
