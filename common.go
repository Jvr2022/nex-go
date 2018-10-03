/*
Package nex implements numerous protocols and related things used in the official Nintendo NEX servers
*/
package nex

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"reflect"
)

// Counter represents an incremental counter
type Counter struct {
	value uint16
}

// Value returns the counters current value
func (counter Counter) Value() uint16 {
	return counter.value
}

// Increment increments the counter by 1 and returns the value
func (counter *Counter) Increment() uint16 {
	counter.value++
	return counter.Value()
}

// Types represents the 5 NEX packet types
var Types = make(map[string]uint16, 5)

// Flags represents the 5 NEX packet flags
var Flags = make(map[string]uint16, 5)

// OptionsAll is used with OptionsSupport to support all methods
var OptionsAll = 0xFFFFFFFF

// OptionsSupport is the ID for the Supported Methods option in PRUDP v1 packets
var OptionsSupport = 0

// OptionsConnectionSignature is the ID for the Connection Signature option in PRUDP v1 packets
var OptionsConnectionSignature = 1

// OptionsFragment is the ID for the Fragment ID option in PRUDP v1 packets
var OptionsFragment = 2

// Options3 is unknown
var Options3 = 3 // Unknown

// Options4 is unknown
var Options4 = 4 // Unknown

func init() {
	Types["Syn"] = 0
	Types["Connect"] = 1
	Types["Data"] = 2
	Types["Disconnect"] = 3
	Types["Ping"] = 4

	Flags["Ack"] = 0x001
	Flags["Reliable"] = 0x002
	Flags["NeedAck"] = 0x004
	Flags["HasSize"] = 0x008
	Flags["MultiAck"] = 0x200
}

func readInt(data []byte, endianness binary.ByteOrder) (ret int) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, endianness, &ret)
	return
}

func readUInt16(data []byte, endianness binary.ByteOrder) (ret uint16) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, endianness, &ret)
	return
}

func readUInt32(data []byte, endianness binary.ByteOrder) (ret uint32) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, endianness, &ret)
	return
}

// Sum calculates the sum of the input
func sum(a interface{}) int {
	var (
		va = reflect.ValueOf(a)
		r  = float64(0)
		vb reflect.Value
	)

	if va.Kind() != reflect.Slice {
		panic(fmt.Sprintf("a %s is not a slice!", va.Kind().String()))
	}

	for i := 0; i < va.Len(); i++ {
		vb = va.Index(i)

		switch vb.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			r += float64(vb.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			r += float64(vb.Uint())
		case reflect.Float32, reflect.Float64:
			r += vb.Float()
		default:
			panic(fmt.Sprintf("a %s is not a summable type!", vb.Kind().String()))
		}
	}

	return int(r)
}

// MD5Hash returns the MD5 hash of the input
func MD5Hash(text []byte) []byte {
	hasher := md5.New()
	hasher.Write(text)
	return hasher.Sum(nil)
}
