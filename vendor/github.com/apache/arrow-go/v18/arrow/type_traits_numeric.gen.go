// Code generated by type_traits_numeric.gen.go.tmpl. DO NOT EDIT.

// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package arrow

import (
	"math"
	"unsafe"

	"github.com/apache/arrow-go/v18/arrow/endian"
)

var (
	Int64Traits    int64Traits
	Uint64Traits   uint64Traits
	Float64Traits  float64Traits
	Int32Traits    int32Traits
	Uint32Traits   uint32Traits
	Float32Traits  float32Traits
	Int16Traits    int16Traits
	Uint16Traits   uint16Traits
	Int8Traits     int8Traits
	Uint8Traits    uint8Traits
	Time32Traits   time32Traits
	Time64Traits   time64Traits
	Date32Traits   date32Traits
	Date64Traits   date64Traits
	DurationTraits durationTraits
)

// Int64 traits

const (
	// Int64SizeBytes specifies the number of bytes required to store a single int64 in memory
	Int64SizeBytes = int(unsafe.Sizeof(int64(0)))
)

type int64Traits struct{}

// BytesRequired returns the number of bytes required to store n elements in memory.
func (int64Traits) BytesRequired(n int) int { return Int64SizeBytes * n }

// PutValue
func (int64Traits) PutValue(b []byte, v int64) {
	endian.Native.PutUint64(b, uint64(v))
}

// CastFromBytes reinterprets the slice b to a slice of type int64.
//
// NOTE: len(b) must be a multiple of Int64SizeBytes.
func (int64Traits) CastFromBytes(b []byte) []int64 {
	return GetData[int64](b)
}

// CastToBytes reinterprets the slice b to a slice of bytes.
func (int64Traits) CastToBytes(b []int64) []byte {
	return GetBytes(b)
}

// Copy copies src to dst.
func (int64Traits) Copy(dst, src []int64) { copy(dst, src) }

// Uint64 traits

const (
	// Uint64SizeBytes specifies the number of bytes required to store a single uint64 in memory
	Uint64SizeBytes = int(unsafe.Sizeof(uint64(0)))
)

type uint64Traits struct{}

// BytesRequired returns the number of bytes required to store n elements in memory.
func (uint64Traits) BytesRequired(n int) int { return Uint64SizeBytes * n }

// PutValue
func (uint64Traits) PutValue(b []byte, v uint64) {
	endian.Native.PutUint64(b, uint64(v))
}

// CastFromBytes reinterprets the slice b to a slice of type uint64.
//
// NOTE: len(b) must be a multiple of Uint64SizeBytes.
func (uint64Traits) CastFromBytes(b []byte) []uint64 {
	return GetData[uint64](b)
}

// CastToBytes reinterprets the slice b to a slice of bytes.
func (uint64Traits) CastToBytes(b []uint64) []byte {
	return GetBytes(b)
}

// Copy copies src to dst.
func (uint64Traits) Copy(dst, src []uint64) { copy(dst, src) }

// Float64 traits

const (
	// Float64SizeBytes specifies the number of bytes required to store a single float64 in memory
	Float64SizeBytes = int(unsafe.Sizeof(float64(0)))
)

type float64Traits struct{}

// BytesRequired returns the number of bytes required to store n elements in memory.
func (float64Traits) BytesRequired(n int) int { return Float64SizeBytes * n }

// PutValue
func (float64Traits) PutValue(b []byte, v float64) {
	endian.Native.PutUint64(b, math.Float64bits(v))
}

// CastFromBytes reinterprets the slice b to a slice of type float64.
//
// NOTE: len(b) must be a multiple of Float64SizeBytes.
func (float64Traits) CastFromBytes(b []byte) []float64 {
	return GetData[float64](b)
}

// CastToBytes reinterprets the slice b to a slice of bytes.
func (float64Traits) CastToBytes(b []float64) []byte {
	return GetBytes(b)
}

// Copy copies src to dst.
func (float64Traits) Copy(dst, src []float64) { copy(dst, src) }

// Int32 traits

const (
	// Int32SizeBytes specifies the number of bytes required to store a single int32 in memory
	Int32SizeBytes = int(unsafe.Sizeof(int32(0)))
)

type int32Traits struct{}

// BytesRequired returns the number of bytes required to store n elements in memory.
func (int32Traits) BytesRequired(n int) int { return Int32SizeBytes * n }

// PutValue
func (int32Traits) PutValue(b []byte, v int32) {
	endian.Native.PutUint32(b, uint32(v))
}

// CastFromBytes reinterprets the slice b to a slice of type int32.
//
// NOTE: len(b) must be a multiple of Int32SizeBytes.
func (int32Traits) CastFromBytes(b []byte) []int32 {
	return GetData[int32](b)
}

// CastToBytes reinterprets the slice b to a slice of bytes.
func (int32Traits) CastToBytes(b []int32) []byte {
	return GetBytes(b)
}

// Copy copies src to dst.
func (int32Traits) Copy(dst, src []int32) { copy(dst, src) }

// Uint32 traits

const (
	// Uint32SizeBytes specifies the number of bytes required to store a single uint32 in memory
	Uint32SizeBytes = int(unsafe.Sizeof(uint32(0)))
)

type uint32Traits struct{}

// BytesRequired returns the number of bytes required to store n elements in memory.
func (uint32Traits) BytesRequired(n int) int { return Uint32SizeBytes * n }

// PutValue
func (uint32Traits) PutValue(b []byte, v uint32) {
	endian.Native.PutUint32(b, uint32(v))
}

// CastFromBytes reinterprets the slice b to a slice of type uint32.
//
// NOTE: len(b) must be a multiple of Uint32SizeBytes.
func (uint32Traits) CastFromBytes(b []byte) []uint32 {
	return GetData[uint32](b)
}

// CastToBytes reinterprets the slice b to a slice of bytes.
func (uint32Traits) CastToBytes(b []uint32) []byte {
	return GetBytes(b)
}

// Copy copies src to dst.
func (uint32Traits) Copy(dst, src []uint32) { copy(dst, src) }

// Float32 traits

const (
	// Float32SizeBytes specifies the number of bytes required to store a single float32 in memory
	Float32SizeBytes = int(unsafe.Sizeof(float32(0)))
)

type float32Traits struct{}

// BytesRequired returns the number of bytes required to store n elements in memory.
func (float32Traits) BytesRequired(n int) int { return Float32SizeBytes * n }

// PutValue
func (float32Traits) PutValue(b []byte, v float32) {
	endian.Native.PutUint32(b, math.Float32bits(v))
}

// CastFromBytes reinterprets the slice b to a slice of type float32.
//
// NOTE: len(b) must be a multiple of Float32SizeBytes.
func (float32Traits) CastFromBytes(b []byte) []float32 {
	return GetData[float32](b)
}

// CastToBytes reinterprets the slice b to a slice of bytes.
func (float32Traits) CastToBytes(b []float32) []byte {
	return GetBytes(b)
}

// Copy copies src to dst.
func (float32Traits) Copy(dst, src []float32) { copy(dst, src) }

// Int16 traits

const (
	// Int16SizeBytes specifies the number of bytes required to store a single int16 in memory
	Int16SizeBytes = int(unsafe.Sizeof(int16(0)))
)

type int16Traits struct{}

// BytesRequired returns the number of bytes required to store n elements in memory.
func (int16Traits) BytesRequired(n int) int { return Int16SizeBytes * n }

// PutValue
func (int16Traits) PutValue(b []byte, v int16) {
	endian.Native.PutUint16(b, uint16(v))
}

// CastFromBytes reinterprets the slice b to a slice of type int16.
//
// NOTE: len(b) must be a multiple of Int16SizeBytes.
func (int16Traits) CastFromBytes(b []byte) []int16 {
	return GetData[int16](b)
}

// CastToBytes reinterprets the slice b to a slice of bytes.
func (int16Traits) CastToBytes(b []int16) []byte {
	return GetBytes(b)
}

// Copy copies src to dst.
func (int16Traits) Copy(dst, src []int16) { copy(dst, src) }

// Uint16 traits

const (
	// Uint16SizeBytes specifies the number of bytes required to store a single uint16 in memory
	Uint16SizeBytes = int(unsafe.Sizeof(uint16(0)))
)

type uint16Traits struct{}

// BytesRequired returns the number of bytes required to store n elements in memory.
func (uint16Traits) BytesRequired(n int) int { return Uint16SizeBytes * n }

// PutValue
func (uint16Traits) PutValue(b []byte, v uint16) {
	endian.Native.PutUint16(b, uint16(v))
}

// CastFromBytes reinterprets the slice b to a slice of type uint16.
//
// NOTE: len(b) must be a multiple of Uint16SizeBytes.
func (uint16Traits) CastFromBytes(b []byte) []uint16 {
	return GetData[uint16](b)
}

// CastToBytes reinterprets the slice b to a slice of bytes.
func (uint16Traits) CastToBytes(b []uint16) []byte {
	return GetBytes(b)
}

// Copy copies src to dst.
func (uint16Traits) Copy(dst, src []uint16) { copy(dst, src) }

// Int8 traits

const (
	// Int8SizeBytes specifies the number of bytes required to store a single int8 in memory
	Int8SizeBytes = int(unsafe.Sizeof(int8(0)))
)

type int8Traits struct{}

// BytesRequired returns the number of bytes required to store n elements in memory.
func (int8Traits) BytesRequired(n int) int { return Int8SizeBytes * n }

// PutValue
func (int8Traits) PutValue(b []byte, v int8) {
	b[0] = byte(v)
}

// CastFromBytes reinterprets the slice b to a slice of type int8.
//
// NOTE: len(b) must be a multiple of Int8SizeBytes.
func (int8Traits) CastFromBytes(b []byte) []int8 {
	return GetData[int8](b)
}

// CastToBytes reinterprets the slice b to a slice of bytes.
func (int8Traits) CastToBytes(b []int8) []byte {
	return GetBytes(b)
}

// Copy copies src to dst.
func (int8Traits) Copy(dst, src []int8) { copy(dst, src) }

// Uint8 traits

const (
	// Uint8SizeBytes specifies the number of bytes required to store a single uint8 in memory
	Uint8SizeBytes = int(unsafe.Sizeof(uint8(0)))
)

type uint8Traits struct{}

// BytesRequired returns the number of bytes required to store n elements in memory.
func (uint8Traits) BytesRequired(n int) int { return Uint8SizeBytes * n }

// PutValue
func (uint8Traits) PutValue(b []byte, v uint8) {
	b[0] = byte(v)
}

// CastFromBytes reinterprets the slice b to a slice of type uint8.
//
// NOTE: len(b) must be a multiple of Uint8SizeBytes.
func (uint8Traits) CastFromBytes(b []byte) []uint8 {
	return GetData[uint8](b)
}

// CastToBytes reinterprets the slice b to a slice of bytes.
func (uint8Traits) CastToBytes(b []uint8) []byte {
	return GetBytes(b)
}

// Copy copies src to dst.
func (uint8Traits) Copy(dst, src []uint8) { copy(dst, src) }

// Time32 traits

const (
	// Time32SizeBytes specifies the number of bytes required to store a single Time32 in memory
	Time32SizeBytes = int(unsafe.Sizeof(Time32(0)))
)

type time32Traits struct{}

// BytesRequired returns the number of bytes required to store n elements in memory.
func (time32Traits) BytesRequired(n int) int { return Time32SizeBytes * n }

// PutValue
func (time32Traits) PutValue(b []byte, v Time32) {
	endian.Native.PutUint32(b, uint32(v))
}

// CastFromBytes reinterprets the slice b to a slice of type Time32.
//
// NOTE: len(b) must be a multiple of Time32SizeBytes.
func (time32Traits) CastFromBytes(b []byte) []Time32 {
	return GetData[Time32](b)
}

// CastToBytes reinterprets the slice b to a slice of bytes.
func (time32Traits) CastToBytes(b []Time32) []byte {
	return GetBytes(b)
}

// Copy copies src to dst.
func (time32Traits) Copy(dst, src []Time32) { copy(dst, src) }

// Time64 traits

const (
	// Time64SizeBytes specifies the number of bytes required to store a single Time64 in memory
	Time64SizeBytes = int(unsafe.Sizeof(Time64(0)))
)

type time64Traits struct{}

// BytesRequired returns the number of bytes required to store n elements in memory.
func (time64Traits) BytesRequired(n int) int { return Time64SizeBytes * n }

// PutValue
func (time64Traits) PutValue(b []byte, v Time64) {
	endian.Native.PutUint64(b, uint64(v))
}

// CastFromBytes reinterprets the slice b to a slice of type Time64.
//
// NOTE: len(b) must be a multiple of Time64SizeBytes.
func (time64Traits) CastFromBytes(b []byte) []Time64 {
	return GetData[Time64](b)
}

// CastToBytes reinterprets the slice b to a slice of bytes.
func (time64Traits) CastToBytes(b []Time64) []byte {
	return GetBytes(b)
}

// Copy copies src to dst.
func (time64Traits) Copy(dst, src []Time64) { copy(dst, src) }

// Date32 traits

const (
	// Date32SizeBytes specifies the number of bytes required to store a single Date32 in memory
	Date32SizeBytes = int(unsafe.Sizeof(Date32(0)))
)

type date32Traits struct{}

// BytesRequired returns the number of bytes required to store n elements in memory.
func (date32Traits) BytesRequired(n int) int { return Date32SizeBytes * n }

// PutValue
func (date32Traits) PutValue(b []byte, v Date32) {
	endian.Native.PutUint32(b, uint32(v))
}

// CastFromBytes reinterprets the slice b to a slice of type Date32.
//
// NOTE: len(b) must be a multiple of Date32SizeBytes.
func (date32Traits) CastFromBytes(b []byte) []Date32 {
	return GetData[Date32](b)
}

// CastToBytes reinterprets the slice b to a slice of bytes.
func (date32Traits) CastToBytes(b []Date32) []byte {
	return GetBytes(b)
}

// Copy copies src to dst.
func (date32Traits) Copy(dst, src []Date32) { copy(dst, src) }

// Date64 traits

const (
	// Date64SizeBytes specifies the number of bytes required to store a single Date64 in memory
	Date64SizeBytes = int(unsafe.Sizeof(Date64(0)))
)

type date64Traits struct{}

// BytesRequired returns the number of bytes required to store n elements in memory.
func (date64Traits) BytesRequired(n int) int { return Date64SizeBytes * n }

// PutValue
func (date64Traits) PutValue(b []byte, v Date64) {
	endian.Native.PutUint64(b, uint64(v))
}

// CastFromBytes reinterprets the slice b to a slice of type Date64.
//
// NOTE: len(b) must be a multiple of Date64SizeBytes.
func (date64Traits) CastFromBytes(b []byte) []Date64 {
	return GetData[Date64](b)
}

// CastToBytes reinterprets the slice b to a slice of bytes.
func (date64Traits) CastToBytes(b []Date64) []byte {
	return GetBytes(b)
}

// Copy copies src to dst.
func (date64Traits) Copy(dst, src []Date64) { copy(dst, src) }

// Duration traits

const (
	// DurationSizeBytes specifies the number of bytes required to store a single Duration in memory
	DurationSizeBytes = int(unsafe.Sizeof(Duration(0)))
)

type durationTraits struct{}

// BytesRequired returns the number of bytes required to store n elements in memory.
func (durationTraits) BytesRequired(n int) int { return DurationSizeBytes * n }

// PutValue
func (durationTraits) PutValue(b []byte, v Duration) {
	endian.Native.PutUint64(b, uint64(v))
}

// CastFromBytes reinterprets the slice b to a slice of type Duration.
//
// NOTE: len(b) must be a multiple of DurationSizeBytes.
func (durationTraits) CastFromBytes(b []byte) []Duration {
	return GetData[Duration](b)
}

// CastToBytes reinterprets the slice b to a slice of bytes.
func (durationTraits) CastToBytes(b []Duration) []byte {
	return GetBytes(b)
}

// Copy copies src to dst.
func (durationTraits) Copy(dst, src []Duration) { copy(dst, src) }
