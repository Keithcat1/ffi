package ffi
import (
	"unsafe"
)

type Int8 int8
type Int16 int16
type Int32 int32
type Int64 int64
type Uint8 uint8
type Uint16 uint16
type Uint32 uint32
type Uint64 uint64
type Bool bool
type Uintptr uintptr
type Uint uint
type Int int

type Byte = uint8
type Rune = Uint32
func (self Uint8) ToUintptr() (uintptr, bool) { return uintptr(self), false }
func (self Uint16) ToUintptr() (uintptr, bool) { return uintptr(self), false }
func (self Uint32) ToUintptr() (uintptr, bool) { return uintptr(self), false }
func (self Uint64) ToUintptr() (uintptr, bool) { return uintptr(self), false }
func (self Uint) ToUintptr() (uintptr, bool) { return uintptr(self), false }
func (self Int8) ToUintptr() (uintptr, bool) { return uintptr(self), false }
func (self Int16) ToUintptr() (uintptr, bool) { return uintptr(self), false }
func (self Int32) ToUintptr() (uintptr, bool) { return uintptr(self), false }
func (self Int64) ToUintptr() (uintptr, bool) { return uintptr(self), false }
func (self Int) ToUintptr() (uintptr, bool) { return uintptr(self), false }
func (self Uintptr) ToUintptr() (uintptr, bool) { return uintptr(self), false }



func (self Int8) Free(value uintptr) {}
func (self Int16) Free(value uintptr) {}
func (self Int32) Free(value uintptr) {}
func (self Int64) Free(value uintptr) {}
func (self Int) Free(value uintptr) {}
func (self Uint8) Free(value uintptr) {}
func (self Uint16) Free(value uintptr) {}
func (self Uint32) Free(value uintptr) {}
func (self Uint64) Free(value uintptr) {}
func (self Uint) Free(value uintptr) {}
func (self Uintptr) Free(value uintptr) {}
func (self *Int8) FromUintptr(a, b uintptr) { *self = Int8(a) }
func (self *Int16) FromUintptr(a, b uintptr) { *self = Int16(a) }
func (self *Int32) FromUintptr(a, b uintptr) { *self = Int32(a) }
func (self *Int64) FromUintptr(a, b uintptr) { *self = Int64(a) }
func (self *Int) FromUintptr(a, b uintptr) { *self = Int(a) }
func (self *Uint8) FromUintptr(a, b uintptr) { *self = Uint8(a) }
func (self *Uint16) FromUintptr(a, b uintptr) { *self = Uint16(a) }
func (self *Uint32) FromUintptr(a, b uintptr) { *self = Uint32(a) }
func (self *Uint64) FromUintptr(a, b uintptr) { *self = Uint64(a) }
func (self *Uint) FromUintptr(a, b uintptr) { *self = Uint(a) }
func (self *Uintptr) FromUintptr(a, b uintptr) { *self = Uintptr(a) }
type UTF8Pointer uintptr
func UTF8StringToPointer(s string) UTF8Pointer {
	data := append([]byte(s), 0)
	return UTF8Pointer(unsafe.Pointer(&data[0]))
}
func (self UTF8Pointer) toUint8Pointer() *uint8 {
	return (*uint8)(unsafe.Pointer(self))
}
func (self UTF8Pointer) ToUintptr() (uintptr, bool) {
	return uintptr(self), false
}
func (self UTF8Pointer) Free(value uintptr) {}
func (self UTF8Pointer) String() string {
	var ptr unsafe.Pointer
	for ptr = unsafe.Pointer(self); *(*uint8)(ptr) !=0; ptr = unsafe.Add(ptr, 1) {}
	size := uintptr(ptr) - uintptr(self)
	println("String offset: %v")
	memory := unsafe.Slice(self.toUint8Pointer(), size - 1)
	return string(memory)
}
type BytesPointer uintptr
func PointerToBytes(data []byte) BytesPointer {
	return BytesPointer(unsafe.Pointer(&data[0]))
}
func (self BytesPointer) ToUintptr() (uintptr, bool) {
	return uintptr(self), false
}
func (self BytesPointer) Free(value uintptr) {}
func (self *BytesPointer) FromUintptr(a, b uintptr) { *self = BytesPointer(a) }
func (self BytesPointer) Bytes(length int) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(self)), length)
}
func (self BytesPointer) CopyBytes(length int) []byte {
	return append([]byte(nil), self.Bytes(length)...)
}