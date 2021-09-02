package ffi
import (
)
type Library struct {
	Loader DLLLoader
}
func loadLibrary(libname string, loader DLLLoader) (*Library, error) {
	self := new(Library)
	self.Loader = loader
	return self, self.Loader.Load(libname)
}
func MustLoadLibrary(libname string) *Library {
	lib, err := LoadLibrary(libname)
	if err != nil {
		panic(err)
	}
	return lib
}
func (self *Library) Unload() error {
	return self.Loader.Unload()
}
func (self *Library) MustUnload() {
	if err := self.Unload(); err != nil {
		panic(err)
	}
}
func (self *Library) GetProcAddress(name string) (uintptr, error) {
	return self.Loader.GetProcAddress(name)
}
func (self *Library) MustGetProcAddress(name string) uintptr {
	proc, err := self.GetProcAddress(name)
	if err != nil {
		panic(err)
	}
	return proc
}