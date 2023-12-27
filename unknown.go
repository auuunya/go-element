package uiautomation

import (
	"fmt"
	"syscall"
	"unsafe"
)

type IUnKnown struct {
	Vtbl *IUnKnownVtbl
}

type IUnKnownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

func NewIUnKnown() *IUnKnown {
	return &IUnKnown{}
}

func AddRef(v *IUnKnown) uint32 {
	ret, _, _ := syscall.SyscallN(
		(*IUnKnownVtbl)(unsafe.Pointer(v.Vtbl)).AddRef,
		uintptr(unsafe.Pointer(v)),
	)
	return uint32(ret)
}
func Release(v *IUnKnown) uint32 {
	ret, _, _ := syscall.SyscallN(
		(*IUnKnownVtbl)(unsafe.Pointer(v.Vtbl)).Release,
		uintptr(unsafe.Pointer(v)),
	)
	return uint32(ret)
}

func QueryInterface(v *IUnKnown, riid syscall.GUID) (unsafe.Pointer, error) {
	var retVal unsafe.Pointer
	ret, _, _ := syscall.SyscallN(
		(*IUnKnownVtbl)(unsafe.Pointer(v.Vtbl)).QueryInterface,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&riid)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

func HResult(ret uintptr) error {
	// error info https://pkg.go.dev/golang.org/x/sys/windows
	return fmt.Errorf("COM Error: 0x%x", ret)
}

func UnKnownToUintptr(obj *interface{}) uintptr {
	return uintptr(unsafe.Pointer(obj))
}
