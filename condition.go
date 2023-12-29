package uiautomation

import (
	"syscall"
	"unsafe"
)

type IUIAutomationCondition struct {
	vtbl *IUnKnown
}

type IUIAutomationConditionVtbl struct {
	IUnKnownVtbl
}

type IUIAutomationAndCondition struct {
	vtbl *IUIAutomationCondition
}

type IUIAutomationAndConditionVtbl struct {
	IUIAutomationConditionVtbl

	Get_ChildCount           uintptr
	GetChildren              uintptr
	GetChildrenAsNativeArray uintptr
}

func Get_ChildCount(v *IUIAutomationAndCondition) int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationAndConditionVtbl)(unsafe.Pointer(v.vtbl)).Get_ChildCount,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func GetChildren(v *IUIAutomationAndCondition) (*TagSafeArray, error) {
	var retVal *TagSafeArray
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationAndConditionVtbl)(unsafe.Pointer(v.vtbl)).GetChildren,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func GetChildrenAsNativeArray(v *IUIAutomationAndCondition) (*IUIAutomationCondition, int32, error) {
	var retVal *IUIAutomationCondition
	var retVal2 int32
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationAndConditionVtbl)(unsafe.Pointer(v.vtbl)).GetChildrenAsNativeArray,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
		uintptr(unsafe.Pointer(&retVal2)),
	)
	if ret != 0 {
		return nil, -1, HResult(ret)
	}
	return retVal, retVal2, nil
}

type IUIAutomationBoolCondition struct {
	vtbl *IUIAutomationCondition
}

type IUIAutomationBoolConditionVtbl struct {
	IUIAutomationConditionVtbl
	Get_BooleanValue uintptr
}

func Get_BooleanValue(v *IUIAutomationBoolCondition) int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationBoolConditionVtbl)(unsafe.Pointer(v.vtbl)).Get_BooleanValue,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

type IUIAutomationNotCondition struct {
	vtbl *IUIAutomationCondition
}

type IUIAutomationNotConditionVtbl struct {
	IUIAutomationConditionVtbl
	GetChild uintptr
}

func GetChild(v *IUIAutomationNotCondition) (*IUIAutomationCondition, error) {
	var retVal *IUIAutomationCondition
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationNotConditionVtbl)(unsafe.Pointer(v.vtbl)).GetChild,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

type IUIAutomationPropertyCondition struct {
	vtbl *IUIAutomationCondition
}

type IUIAutomationPropertyConditionVtbl struct {
	IUIAutomationConditionVtbl

	Get_PropertyConditionFlags uintptr
	Get_PropertyId             uintptr
	Get_PropertyValue          uintptr
}

func Get_PropertyConditionFlags(v *IUIAutomationPropertyCondition) *PropertyConditionFlags {
	var retVal *PropertyConditionFlags
	syscall.SyscallN(
		(*IUIAutomationPropertyConditionVtbl)(unsafe.Pointer(v.vtbl)).Get_PropertyConditionFlags,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func Get_PropertyId(v *IUIAutomationPropertyCondition) *PropertyId {
	var retVal *PropertyId
	syscall.SyscallN(
		(*IUIAutomationPropertyConditionVtbl)(unsafe.Pointer(v.vtbl)).Get_PropertyId,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func Get_PropertyValue(v *IUIAutomationPropertyCondition) *VARIANT {
	var retVal *VARIANT
	syscall.SyscallN(
		(*IUIAutomationPropertyConditionVtbl)(unsafe.Pointer(v.vtbl)).Get_PropertyValue,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
