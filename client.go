package uiautomation

import (
	"syscall"
	"unsafe"
)

type IUIAutomationValuePattern struct {
	vtbl *IUnKnown
}
type IUIAutomationValuePatternVtbl struct {
	*IUnKnownVtbl
	Get_CachedIsReadOnly  uintptr
	Get_CachedValue       uintptr
	Get_CurrentIsReadOnly uintptr
	Get_CurrentValue      uintptr
	SetValue              uintptr
}

// TODO:: clent IUIAutomationValuePattern method
func PatternSetValue(pointer unsafe.Pointer, in string) error {
	v := (*IUIAutomationValuePattern)(unsafe.Pointer(pointer))
	retVal, err := string2Bstr(in)
	if err != nil {
		return err
	}
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationValuePatternVtbl)(unsafe.Pointer(v.vtbl)).SetValue,
		uintptr(unsafe.Pointer(v)),
		uintptr(retVal),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
