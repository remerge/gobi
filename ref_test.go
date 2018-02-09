package gobi

import (
	"bytes"
	"testing"
)

type M1 struct {
	SomeInt int
}

type M0 struct {
	A *M1
	B *M1
	X *int
	Y *int
}

func TestPointerStruct(t *testing.T) {
	var t0 M0
	i := 777
	t0.A = &M1{SomeInt: 9999}
	t0.B = t0.A
	t0.X = &i
	t0.Y = &i
	b := new(bytes.Buffer)
	NewEncoder(b).Encode(t0)
	dec := NewDecoder(b)
	var t1 M0
	err := dec.Decode(&t1)
	if err != nil {
		t.Error("error: ", err)
	}
	if t0.X != t0.Y {
		t.Error("should be equal")
	}
	if t0.A != t0.B {
		t.Error("should be equal")
	}
	t0.A.SomeInt = 20
	if t0.B.SomeInt != 20 {
		t.Error("ref error")
	}
}

type M0D struct {
	A *M1 `gobi:"deprecated"`
	B *M1
	X *int
	Y *int
}

type M0D0 struct {
	B *M1
	X *int
	Y *int
}

func TestPointerStructDeprecatedField(t *testing.T) {
	var t0 M0D
	i := 777
	t0.A = &M1{SomeInt: 9999}
	t0.B = t0.A
	t0.X = &i
	t0.Y = &i
	b := new(bytes.Buffer)
	NewEncoder(b).Encode(t0)
	var t1 M0D
	err := NewDecoder(b).Decode(&t1)
	if err != nil {
		t.Fatal("error: ", err)
	}
	if t1.X != t1.Y {
		t.Fatal("should be equal")
	}
	if t1.A != nil {
		t.Fatal("should be nil")
	}
	if t1.B == nil {
		t.Fatal("should not be nil")
	}
	if t1.B.SomeInt != 9999 {
		t.Fatal("ref error")
	}
	b = new(bytes.Buffer)
	NewEncoder(b).Encode(t1)
	var t2 M0D0
	err = NewDecoder(b).Decode(&t2)
	if err != nil {
		t.Fatal("error: ", err)
	}
	if t2.X != t2.Y {
		t.Fatal("should be equal")
	}
	if t2.B == nil {
		t.Fatal("should not be nil")
	}
	if t2.B.SomeInt != 9999 {
		t.Fatal("ref error")
	}
}

type PointerArray struct {
	Array []*M1
}

func TestPointerArray(t *testing.T) {
	m := &M1{SomeInt: 9999}

	pa := PointerArray{Array: []*M1{m, m, m, m}}

	b := new(bytes.Buffer)
	NewEncoder(b).Encode(pa)
	dec := NewDecoder(b)
	var r PointerArray
	err := dec.Decode(&r)
	if err != nil {
		t.Error("error: ", err)
	}
	for i := 0; i < len(r.Array); i++ {
		if r.Array[i] != r.Array[0] {
			t.Error("array entry pointer should be equal")
		}
	}
}

type PointerMap struct {
	M map[int]*M1
}

func TestPointerMap(t *testing.T) {
	m := &M1{SomeInt: 9999}
	pm := PointerMap{M: map[int]*M1{1: m, 2: m}}

	b := new(bytes.Buffer)
	NewEncoder(b).Encode(pm)
	dec := NewDecoder(b)
	var r PointerMap
	err := dec.Decode(&r)
	if err != nil {
		t.Error("error: ", err)
	}
	if r.M[1] != r.M[2] {
		t.Error("map entry pointer should be equal")
	}
}
