// Created by decgen --output dec_helpers.go; DO NOT EDIT

// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gobi

import (
	"math"
	"reflect"
)

var decArrayHelper = map[reflect.Kind]decHelper{
	reflect.Bool:       decBoolArray,
	reflect.Complex64:  decComplex64Array,
	reflect.Complex128: decComplex128Array,
	reflect.Float32:    decFloat32Array,
	reflect.Float64:    decFloat64Array,
	reflect.Int:        decIntArray,
	reflect.Int16:      decInt16Array,
	reflect.Int32:      decInt32Array,
	reflect.Int64:      decInt64Array,
	reflect.Int8:       decInt8Array,
	reflect.String:     decStringArray,
	reflect.Uint:       decUintArray,
	reflect.Uint16:     decUint16Array,
	reflect.Uint32:     decUint32Array,
	reflect.Uint64:     decUint64Array,
	reflect.Uintptr:    decUintptrArray,
}

var decSliceHelper = map[reflect.Kind]decHelper{
	reflect.Bool:       decBoolSlice,
	reflect.Complex64:  decComplex64Slice,
	reflect.Complex128: decComplex128Slice,
	reflect.Float32:    decFloat32Slice,
	reflect.Float64:    decFloat64Slice,
	reflect.Int:        decIntSlice,
	reflect.Int16:      decInt16Slice,
	reflect.Int32:      decInt32Slice,
	reflect.Int64:      decInt64Slice,
	reflect.Int8:       decInt8Slice,
	reflect.String:     decStringSlice,
	reflect.Uint:       decUintSlice,
	reflect.Uint16:     decUint16Slice,
	reflect.Uint32:     decUint32Slice,
	reflect.Uint64:     decUint64Slice,
	reflect.Uintptr:    decUintptrSlice,
}

func decBoolArray(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	// Can only slice if it is addressable.
	if !v.CanAddr() {
		return false
	}
	return decBoolSlice(state, v.Slice(0, v.Len()), length, ovfl)
}

func decBoolSlice(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	slice, ok := v.Interface().([]bool)
	if !ok {
		// It is kind bool but not type bool. TODO: We can handle this unsafely.
		return false
	}
	for i := 0; i < length; i++ {
		if state.b.Len() == 0 {
			errorf("decoding bool array or slice: length exceeds input size (%d elements)", length)
		}
		slice[i] = state.decodeUint() != 0
	}
	return true
}

func decComplex64Array(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	// Can only slice if it is addressable.
	if !v.CanAddr() {
		return false
	}
	return decComplex64Slice(state, v.Slice(0, v.Len()), length, ovfl)
}

func decComplex64Slice(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	slice, ok := v.Interface().([]complex64)
	if !ok {
		// It is kind complex64 but not type complex64. TODO: We can handle this unsafely.
		return false
	}
	for i := 0; i < length; i++ {
		if state.b.Len() == 0 {
			errorf("decoding complex64 array or slice: length exceeds input size (%d elements)", length)
		}
		real := float32FromBits(state.decodeUint(), ovfl)
		imag := float32FromBits(state.decodeUint(), ovfl)
		slice[i] = complex(float32(real), float32(imag))
	}
	return true
}

func decComplex128Array(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	// Can only slice if it is addressable.
	if !v.CanAddr() {
		return false
	}
	return decComplex128Slice(state, v.Slice(0, v.Len()), length, ovfl)
}

func decComplex128Slice(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	slice, ok := v.Interface().([]complex128)
	if !ok {
		// It is kind complex128 but not type complex128. TODO: We can handle this unsafely.
		return false
	}
	for i := 0; i < length; i++ {
		if state.b.Len() == 0 {
			errorf("decoding complex128 array or slice: length exceeds input size (%d elements)", length)
		}
		real := float64FromBits(state.decodeUint())
		imag := float64FromBits(state.decodeUint())
		slice[i] = complex(real, imag)
	}
	return true
}

func decFloat32Array(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	// Can only slice if it is addressable.
	if !v.CanAddr() {
		return false
	}
	return decFloat32Slice(state, v.Slice(0, v.Len()), length, ovfl)
}

func decFloat32Slice(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	slice, ok := v.Interface().([]float32)
	if !ok {
		// It is kind float32 but not type float32. TODO: We can handle this unsafely.
		return false
	}
	for i := 0; i < length; i++ {
		if state.b.Len() == 0 {
			errorf("decoding float32 array or slice: length exceeds input size (%d elements)", length)
		}
		slice[i] = float32(float32FromBits(state.decodeUint(), ovfl))
	}
	return true
}

func decFloat64Array(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	// Can only slice if it is addressable.
	if !v.CanAddr() {
		return false
	}
	return decFloat64Slice(state, v.Slice(0, v.Len()), length, ovfl)
}

func decFloat64Slice(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	slice, ok := v.Interface().([]float64)
	if !ok {
		// It is kind float64 but not type float64. TODO: We can handle this unsafely.
		return false
	}
	for i := 0; i < length; i++ {
		if state.b.Len() == 0 {
			errorf("decoding float64 array or slice: length exceeds input size (%d elements)", length)
		}
		slice[i] = float64FromBits(state.decodeUint())
	}
	return true
}

func decIntArray(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	// Can only slice if it is addressable.
	if !v.CanAddr() {
		return false
	}
	return decIntSlice(state, v.Slice(0, v.Len()), length, ovfl)
}

func decIntSlice(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	slice, ok := v.Interface().([]int)
	if !ok {
		// It is kind int but not type int. TODO: We can handle this unsafely.
		return false
	}
	for i := 0; i < length; i++ {
		if state.b.Len() == 0 {
			errorf("decoding int array or slice: length exceeds input size (%d elements)", length)
		}
		x := state.decodeInt()
		// MinInt and MaxInt
		if x < ^int64(^uint(0)>>1) || int64(^uint(0)>>1) < x {
			error_(ovfl)
		}
		slice[i] = int(x)
	}
	return true
}

func decInt16Array(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	// Can only slice if it is addressable.
	if !v.CanAddr() {
		return false
	}
	return decInt16Slice(state, v.Slice(0, v.Len()), length, ovfl)
}

func decInt16Slice(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	slice, ok := v.Interface().([]int16)
	if !ok {
		// It is kind int16 but not type int16. TODO: We can handle this unsafely.
		return false
	}
	for i := 0; i < length; i++ {
		if state.b.Len() == 0 {
			errorf("decoding int16 array or slice: length exceeds input size (%d elements)", length)
		}
		x := state.decodeInt()
		if x < math.MinInt16 || math.MaxInt16 < x {
			error_(ovfl)
		}
		slice[i] = int16(x)
	}
	return true
}

func decInt32Array(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	// Can only slice if it is addressable.
	if !v.CanAddr() {
		return false
	}
	return decInt32Slice(state, v.Slice(0, v.Len()), length, ovfl)
}

func decInt32Slice(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	slice, ok := v.Interface().([]int32)
	if !ok {
		// It is kind int32 but not type int32. TODO: We can handle this unsafely.
		return false
	}
	for i := 0; i < length; i++ {
		if state.b.Len() == 0 {
			errorf("decoding int32 array or slice: length exceeds input size (%d elements)", length)
		}
		x := state.decodeInt()
		if x < math.MinInt32 || math.MaxInt32 < x {
			error_(ovfl)
		}
		slice[i] = int32(x)
	}
	return true
}

func decInt64Array(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	// Can only slice if it is addressable.
	if !v.CanAddr() {
		return false
	}
	return decInt64Slice(state, v.Slice(0, v.Len()), length, ovfl)
}

func decInt64Slice(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	slice, ok := v.Interface().([]int64)
	if !ok {
		// It is kind int64 but not type int64. TODO: We can handle this unsafely.
		return false
	}
	for i := 0; i < length; i++ {
		if state.b.Len() == 0 {
			errorf("decoding int64 array or slice: length exceeds input size (%d elements)", length)
		}
		slice[i] = state.decodeInt()
	}
	return true
}

func decInt8Array(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	// Can only slice if it is addressable.
	if !v.CanAddr() {
		return false
	}
	return decInt8Slice(state, v.Slice(0, v.Len()), length, ovfl)
}

func decInt8Slice(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	slice, ok := v.Interface().([]int8)
	if !ok {
		// It is kind int8 but not type int8. TODO: We can handle this unsafely.
		return false
	}
	for i := 0; i < length; i++ {
		if state.b.Len() == 0 {
			errorf("decoding int8 array or slice: length exceeds input size (%d elements)", length)
		}
		x := state.decodeInt()
		if x < math.MinInt8 || math.MaxInt8 < x {
			error_(ovfl)
		}
		slice[i] = int8(x)
	}
	return true
}

func decStringArray(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	// Can only slice if it is addressable.
	if !v.CanAddr() {
		return false
	}
	return decStringSlice(state, v.Slice(0, v.Len()), length, ovfl)
}

func decStringSlice(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	slice, ok := v.Interface().([]string)
	if !ok {
		// It is kind string but not type string. TODO: We can handle this unsafely.
		return false
	}
	for i := 0; i < length; i++ {
		if state.b.Len() == 0 {
			errorf("decoding string array or slice: length exceeds input size (%d elements)", length)
		}
		u := state.decodeUint()
		n := int(u)
		if n < 0 || uint64(n) != u || n > state.b.Len() {
			errorf("length of string exceeds input size (%d bytes)", u)
		}
		if n > state.b.Len() {
			errorf("string data too long for buffer: %d", n)
		}
		// Read the data.
		data := make([]byte, n)
		if _, err := state.b.Read(data); err != nil {
			errorf("error decoding string: %s", err)
		}
		slice[i] = string(data)
	}
	return true
}

func decUintArray(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	// Can only slice if it is addressable.
	if !v.CanAddr() {
		return false
	}
	return decUintSlice(state, v.Slice(0, v.Len()), length, ovfl)
}

func decUintSlice(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	slice, ok := v.Interface().([]uint)
	if !ok {
		// It is kind uint but not type uint. TODO: We can handle this unsafely.
		return false
	}
	for i := 0; i < length; i++ {
		if state.b.Len() == 0 {
			errorf("decoding uint array or slice: length exceeds input size (%d elements)", length)
		}
		x := state.decodeUint()
		/*TODO if math.MaxUint32 < x {
			error_(ovfl)
		}*/
		slice[i] = uint(x)
	}
	return true
}

func decUint16Array(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	// Can only slice if it is addressable.
	if !v.CanAddr() {
		return false
	}
	return decUint16Slice(state, v.Slice(0, v.Len()), length, ovfl)
}

func decUint16Slice(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	slice, ok := v.Interface().([]uint16)
	if !ok {
		// It is kind uint16 but not type uint16. TODO: We can handle this unsafely.
		return false
	}
	for i := 0; i < length; i++ {
		if state.b.Len() == 0 {
			errorf("decoding uint16 array or slice: length exceeds input size (%d elements)", length)
		}
		x := state.decodeUint()
		if math.MaxUint16 < x {
			error_(ovfl)
		}
		slice[i] = uint16(x)
	}
	return true
}

func decUint32Array(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	// Can only slice if it is addressable.
	if !v.CanAddr() {
		return false
	}
	return decUint32Slice(state, v.Slice(0, v.Len()), length, ovfl)
}

func decUint32Slice(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	slice, ok := v.Interface().([]uint32)
	if !ok {
		// It is kind uint32 but not type uint32. TODO: We can handle this unsafely.
		return false
	}
	for i := 0; i < length; i++ {
		if state.b.Len() == 0 {
			errorf("decoding uint32 array or slice: length exceeds input size (%d elements)", length)
		}
		x := state.decodeUint()
		if math.MaxUint32 < x {
			error_(ovfl)
		}
		slice[i] = uint32(x)
	}
	return true
}

func decUint64Array(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	// Can only slice if it is addressable.
	if !v.CanAddr() {
		return false
	}
	return decUint64Slice(state, v.Slice(0, v.Len()), length, ovfl)
}

func decUint64Slice(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	slice, ok := v.Interface().([]uint64)
	if !ok {
		// It is kind uint64 but not type uint64. TODO: We can handle this unsafely.
		return false
	}
	for i := 0; i < length; i++ {
		if state.b.Len() == 0 {
			errorf("decoding uint64 array or slice: length exceeds input size (%d elements)", length)
		}
		slice[i] = state.decodeUint()
	}
	return true
}

func decUintptrArray(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	// Can only slice if it is addressable.
	if !v.CanAddr() {
		return false
	}
	return decUintptrSlice(state, v.Slice(0, v.Len()), length, ovfl)
}

func decUintptrSlice(state *decoderState, v reflect.Value, length int, ovfl error) bool {
	slice, ok := v.Interface().([]uintptr)
	if !ok {
		// It is kind uintptr but not type uintptr. TODO: We can handle this unsafely.
		return false
	}
	for i := 0; i < length; i++ {
		if state.b.Len() == 0 {
			errorf("decoding uintptr array or slice: length exceeds input size (%d elements)", length)
		}
		x := state.decodeUint()
		if uint64(^uintptr(0)) < x {
			error_(ovfl)
		}
		slice[i] = uintptr(x)
	}
	return true
}
