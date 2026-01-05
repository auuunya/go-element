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

func NewIUnKnown(v unsafe.Pointer) *IUnKnown {
	return (*IUnKnown)(v)
}

func (v *IUnKnown) AddRef() uint32 {
	ret, _, _ := syscall.SyscallN(
		v.Vtbl.AddRef,
		uintptr(unsafe.Pointer(v)),
	)
	return uint32(ret)
}

func (v *IUnKnown) Release() uint32 {
	if v == nil || v.Vtbl == nil {
		return 0
	}
	ret, _, _ := syscall.SyscallN(
		v.Vtbl.Release,
		uintptr(unsafe.Pointer(v)),
	)
	return uint32(ret)
}

func AddRef(v *IUnKnown) uint32 {
	return v.AddRef()
}
func Release(v *IUnKnown) uint32 {
	return v.Release()
}

func QueryInterface(v *IUnKnown, riid *syscall.GUID) (unsafe.Pointer, error) {
	var retVal unsafe.Pointer
	ret, _, _ := syscall.SyscallN(
		(*IUnKnownVtbl)(unsafe.Pointer(v.Vtbl)).QueryInterface,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

func HResult(ret uintptr) error {
	if ret == 0 {
		return nil
	}
	// error info https://pkg.go.dev/golang.org/x/sys/windows
	return fmt.Errorf("COM Error: 0x%08X", uint32(ret))
}

func UnKnownToUintptr(obj interface{}) uintptr {
	// 获取接口内部的指针
	type iface struct {
		typ uintptr
		ptr uintptr
	}
	return (*iface)(unsafe.Pointer(&obj)).ptr
}
