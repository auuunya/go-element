package uiautomation

import (
	"syscall"
	"unsafe"
)

// IUIAutomationInvokePattern
type IUIAutomationInvokePattern struct {
	vtbl *IUnKnown
}
type IUIAutomationInvokePatternVtbl struct {
	IUnKnownVtbl
	Invoke uintptr
}

func (v *IUIAutomationInvokePattern) Invoke() error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationInvokePatternVtbl)(unsafe.Pointer(v.vtbl)).Invoke,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

func (v *IUIAutomationInvokePattern) Release() uint32 {
	return v.vtbl.Release()
}

// IUIAutomationValuePattern
type IUIAutomationValuePattern struct {
	vtbl *IUnKnown
}
type IUIAutomationValuePatternVtbl struct {
	IUnKnownVtbl
	SetValue              uintptr
	Get_CurrentValue      uintptr
	Get_CurrentIsReadOnly uintptr
	Get_CachedValue       uintptr
	Get_CachedIsReadOnly  uintptr
}

func (v *IUIAutomationValuePattern) SetValue(val string) error {
	bstr, err := string2Bstr(val)
	if err != nil {
		return err
	}
	defer procSysFreeString.Call(bstr)
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationValuePatternVtbl)(unsafe.Pointer(v.vtbl)).SetValue,
		uintptr(unsafe.Pointer(v)),
		uintptr(bstr),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

func (v *IUIAutomationValuePattern) Get_CurrentValue() (string, error) {
	var bstr uintptr
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationValuePatternVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentValue,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if ret != 0 {
		return "", HResult(ret)
	}
	defer procSysFreeString.Call(bstr)
	return bstr2str(bstr), nil
}

func (v *IUIAutomationValuePattern) Release() uint32 {
	return v.vtbl.Release()
}

// IUIAutomationSelectionItemPattern
type IUIAutomationSelectionItemPattern struct {
	vtbl *IUnKnown
}
type IUIAutomationSelectionItemPatternVtbl struct {
	IUnKnownVtbl
	Select                        uintptr
	AddToSelection                uintptr
	RemoveFromSelection           uintptr
	Get_CurrentIsSelected         uintptr
	Get_CurrentSelectionContainer uintptr
	Get_CachedIsSelected          uintptr
	Get_CachedSelectionContainer  uintptr
}

func (v *IUIAutomationSelectionItemPattern) Select() error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationSelectionItemPatternVtbl)(unsafe.Pointer(v.vtbl)).Select,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

func (v *IUIAutomationSelectionItemPattern) Release() uint32 {
	return v.vtbl.Release()
}

// IUIAutomationTogglePattern
type IUIAutomationTogglePattern struct {
	vtbl *IUnKnown
}
type IUIAutomationTogglePatternVtbl struct {
	IUnKnownVtbl
	Toggle                 uintptr
	Get_CurrentToggleState uintptr
	Get_CachedToggleState  uintptr
}

func (v *IUIAutomationTogglePattern) Toggle() error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationTogglePatternVtbl)(unsafe.Pointer(v.vtbl)).Toggle,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

func (v *IUIAutomationTogglePattern) Get_CurrentToggleState() (ToggleState, error) {
	var state ToggleState
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationTogglePatternVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentToggleState,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&state)),
	)
	if ret != 0 {
		return 0, HResult(ret)
	}
	return state, nil
}

func (v *IUIAutomationTogglePattern) Release() uint32 {
	return v.vtbl.Release()
}

// IUIAutomationExpandCollapsePattern
type IUIAutomationExpandCollapsePattern struct {
	vtbl *IUnKnown
}
type IUIAutomationExpandCollapsePatternVtbl struct {
	IUnKnownVtbl
	Expand                         uintptr
	Collapse                       uintptr
	Get_CurrentExpandCollapseState uintptr
	Get_CachedExpandCollapseState  uintptr
}

func (v *IUIAutomationExpandCollapsePattern) Expand() error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationExpandCollapsePatternVtbl)(unsafe.Pointer(v.vtbl)).Expand,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

func (v *IUIAutomationExpandCollapsePattern) Collapse() error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationExpandCollapsePatternVtbl)(unsafe.Pointer(v.vtbl)).Collapse,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

func (v *IUIAutomationExpandCollapsePattern) Release() uint32 {
	return v.vtbl.Release()
}
