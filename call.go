package ffi
import (
	"syscall"
)


func call0(resultPtr FFIReturnable, fp uintptr) error {
	result1, result2, err := syscall.Syscall(fp, 0, 0, 0, 0)
	if resultPtr != nil {
		resultPtr.FromUintptr(result1, result2)
	}
	return err
}
func call1(resultPtr FFIReturnable, fp uintptr, a1 FFIValue) error {
	u1, free := a1.ToUintptr()
	if free {
		defer a1.Free(u1)
	}
	result1, result2, err := syscall.Syscall(fp, 1, u1, 0, 0)
	if resultPtr != nil {
		resultPtr.FromUintptr(result1, result2)
	}
	return err
}
func call2(resultPtr FFIReturnable, fp uintptr, a1, a2 FFIValue) error {
	u1, free := a1.ToUintptr()
	if free {
		defer a1.Free(u1)
	}
	u2, free := a2.ToUintptr()
	if free {
		defer a2.Free(u2)
	}
	result1, result2, err := syscall.Syscall(fp, 2, u1, u2, 0)
	if resultPtr != nil {
		resultPtr.FromUintptr(result1, result2)
	}
	return err
}
func call3(resultPtr FFIReturnable, fp uintptr, a1, a2, a3 FFIValue) error {
	u1, free := a1.ToUintptr()
	if free {
		defer a1.Free(u1)
	}
	u2, free := a2.ToUintptr()
	if free {
		defer a2.Free(u2)
	}
	u3, free := a3.ToUintptr()
	if free {
		defer a3.Free(u3)
	}
	result1, result2, err := syscall.Syscall(fp, 3, u1, u2, u3)
	if resultPtr != nil {
		resultPtr.FromUintptr(result1, result2)
	}
	return err
}
func call4(resultPtr FFIReturnable, fp uintptr, a1, a2, a3, a4 FFIValue) error {
	u1, free := a1.ToUintptr()
	if free {
		defer a1.Free(u1)
	}
	u2, free := a2.ToUintptr()
	if free {
		defer a2.Free(u2)
	}
	u3, free := a3.ToUintptr()
	if free {
		defer a3.Free(u3)
	}
	u4, free := a4.ToUintptr()
	if free {
		defer a4.Free(u4)
	}
	result1, result2, err := syscall.Syscall6(fp, 4, u1, u2, u3, u4, 0, 0)
	if resultPtr != nil {
		resultPtr.FromUintptr(result1, result2)
	}
	return err
}
func call5(resultPtr FFIReturnable, fp uintptr, a1, a2, a3, a4, a5 FFIValue) error {
	u1, free := a1.ToUintptr()
	if free {
		defer a1.Free(u1)
	}
	u2, free := a2.ToUintptr()
	if free {
		defer a2.Free(u2)
	}
	u3, free := a3.ToUintptr()
	if free {
		defer a3.Free(u3)
	}
	u4, free := a4.ToUintptr()
	if free {
		defer a4.Free(u4)
	}
	u5, free := a5.ToUintptr()
	if free {
		defer a5.Free(u5)
	}
	result1, result2, err := syscall.Syscall6(fp, 5, u1, u2, u3, u4, u5, 0)
	if resultPtr != nil {
		resultPtr.FromUintptr(result1, result2)
	}
	return err
}
func call6(resultPtr FFIReturnable, fp uintptr, a1, a2, a3, a4, a5, a6 FFIValue) error {
	u1, free := a1.ToUintptr()
	if free {
		defer a1.Free(u1)
	}
	u2, free := a2.ToUintptr()
	if free {
		defer a2.Free(u2)
	}
	u3, free := a3.ToUintptr()
	if free {
		defer a3.Free(u3)
	}
	u4, free := a4.ToUintptr()
	if free {
		defer a4.Free(u4)
	}
	u5, free := a5.ToUintptr()
	if free {
		defer a5.Free(u5)
	}
	u6, free := a6.ToUintptr()
	if free {
		defer a6.Free(u6)
	}
	result1, result2, err := syscall.Syscall6(fp, 6, u1, u2, u3, u4, u5, u6)
	if resultPtr != nil {
		resultPtr.FromUintptr(result1, result2)
	}
	return err
}