package ffi
import (
)
// knows how to load and unload a DLL and retrieve an exported function
type DLLLoader interface {
	Load(libname string) error
	Unload() error
	GetProcAddress(name string) (uintptr, error)
}