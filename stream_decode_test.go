package gobi

import (
	"bytes"
	"reflect"
	"testing"
)

// streamSample exercises the value kinds the streaming decBuffer paths touch:
// varints (ReadByte), fixed/string bodies (Read), and skipped fields (Drop).
type streamSample struct {
	I    int64
	U    uint32
	F    float64
	S    string
	B    []byte
	M    map[string]int
	Sub  *streamSub
	List []streamSub
}

type streamSub struct {
	Name string
	Vals []int
}

// TestStreamDecodeMatchesBuffered round-trips the same payload through the
// default buffered decoder and the streaming decoder and asserts both yield
// an identical value.
func TestStreamDecodeMatchesBuffered(t *testing.T) {
	in := streamSample{
		I: -42, U: 4000000000, F: 3.14159,
		S: "the quick brown fox jumps over the lazy dog",
		B: bytes.Repeat([]byte{0xAB, 0xCD}, 1024),
		M: map[string]int{"a": 1, "b": 2, "c": 3},
		Sub:  &streamSub{Name: "sub", Vals: []int{1, 2, 3, 4, 5}},
		List: []streamSub{{Name: "x", Vals: []int{7}}, {Name: "y"}},
	}

	var buf bytes.Buffer
	if err := NewEncoder(&buf).Encode(in); err != nil {
		t.Fatalf("encode: %v", err)
	}
	encoded := buf.Bytes()

	var bufferedOut streamSample
	if err := NewDecoder(bytes.NewReader(encoded)).Decode(&bufferedOut); err != nil {
		t.Fatalf("buffered decode: %v", err)
	}

	var streamedOut streamSample
	dec := NewDecoder(bytes.NewReader(encoded))
	dec.StreamDecode = true
	if err := dec.Decode(&streamedOut); err != nil {
		t.Fatalf("streamed decode: %v", err)
	}

	if !reflect.DeepEqual(bufferedOut, streamedOut) {
		t.Fatalf("streamed result differs from buffered:\n buffered=%+v\n streamed=%+v", bufferedOut, streamedOut)
	}
	if !reflect.DeepEqual(in, streamedOut) {
		t.Fatalf("streamed result differs from input:\n input=%+v\n streamed=%+v", in, streamedOut)
	}
}

// TestStreamDecodeIgnoresFields decodes into a struct missing fields, forcing
// the decoder down the Drop (skip) path under streaming.
func TestStreamDecodeIgnoresFields(t *testing.T) {
	full := streamSample{
		I: 99, S: "drop me", B: bytes.Repeat([]byte{1}, 500),
		M: map[string]int{"k": 7}, List: []streamSub{{Name: "z", Vals: []int{9}}},
	}
	var buf bytes.Buffer
	if err := NewEncoder(&buf).Encode(full); err != nil {
		t.Fatalf("encode: %v", err)
	}

	// Decode into a subset type so the extra fields are skipped via Drop.
	type subset struct {
		I int64
		S string
	}
	var out subset
	dec := NewDecoder(bytes.NewReader(buf.Bytes()))
	dec.StreamDecode = true
	if err := dec.Decode(&out); err != nil {
		t.Fatalf("streamed decode: %v", err)
	}
	if out.I != full.I || out.S != full.S {
		t.Fatalf("got %+v, want I=%d S=%q", out, full.I, full.S)
	}
}
