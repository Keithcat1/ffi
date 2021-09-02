package ffi
import (
)

type FFIValue interface {
	ToUintptr() (uintptr, bool)
	Free(uintptr)
}
type FFIReturnable interface {
	FromUintptr(a, b uintptr)
}