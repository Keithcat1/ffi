package ffi
import (
	"syscall"
)
type windowsDLLLoader struct {
	dll *syscall.DLL
}
func (self *windowsDLLLoader) Load(libname string) error {
	dll, err := syscall.LoadDLL(libname)
	self.dll = dll
	return err
}
func (self *windowsDLLLoader) Unload() error {
	return self.dll.Release()
}
func (self *windowsDLLLoader) GetProcAddress(name string) (uintptr, error) {
	proc, err := self.dll.FindProc(name)
	if err != nil {
		return 0, err
	} else {
		return proc.Addr(), nil
	}
}
func LoadLibrary(libname string) (*Library, error) {
	return loadLibrary(libname, new(windowsDLLLoader))
}