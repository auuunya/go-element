package uiautomation

import (
	"syscall"
	"unsafe"
)

// IUIAutomationInvokePattern
type IUIAutomationInvokePattern struct {
	vtbl *IUIAutomationInvokePatternVtbl
}
type IUIAutomationInvokePatternVtbl struct {
	IUnKnownVtbl
	Invoke uintptr
}

func (v *IUIAutomationInvokePattern) Invoke() error {
	ret, _, _ := syscall.SyscallN(
		v.vtbl.Invoke,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

func (v *IUIAutomationInvokePattern) Release() uint32 {
	return (*IUnKnown)(unsafe.Pointer(v)).Release()
}

func newIUIAutomationInvokePattern(unk *IUnKnown) *IUIAutomationInvokePattern {
	return (*IUIAutomationInvokePattern)(unsafe.Pointer(unk))
}
func NewIUIAutomationInvokePattern(unk *IUnKnown) *IUIAutomationInvokePattern {
	return newIUIAutomationInvokePattern(unk)
}

// IUIAutomationValuePattern
type IUIAutomationValuePattern struct {
	vtbl *IUIAutomationValuePatternVtbl
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
		v.vtbl.SetValue,
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
		v.vtbl.Get_CurrentValue,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if ret != 0 {
		return "", HResult(ret)
	}
	defer procSysFreeString.Call(bstr)
	return bstr2str(bstr), nil
}

func newIUIAutomationValuePattern(unk *IUnKnown) *IUIAutomationValuePattern {
	return (*IUIAutomationValuePattern)(unsafe.Pointer(unk))
}
func NewIUIAutomationValuePattern(unk *IUnKnown) *IUIAutomationValuePattern {
	return newIUIAutomationValuePattern(unk)
}

func (v *IUIAutomationValuePattern) Release() uint32 {
	return (*IUnKnown)(unsafe.Pointer(v)).Release()
}

// IUIAutomationSelectionItemPattern
type IUIAutomationSelectionItemPattern struct {
	vtbl *IUIAutomationSelectionItemPatternVtbl
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
		v.vtbl.Select,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

func (v *IUIAutomationSelectionItemPattern) Release() uint32 {
	return (*IUnKnown)(unsafe.Pointer(v)).Release()
}

// IUIAutomationTogglePattern
type IUIAutomationTogglePattern struct {
	vtbl *IUIAutomationTogglePatternVtbl
}
type IUIAutomationTogglePatternVtbl struct {
	IUnKnownVtbl
	Toggle                 uintptr
	Get_CurrentToggleState uintptr
	Get_CachedToggleState  uintptr
}

func (v *IUIAutomationTogglePattern) Toggle() error {
	ret, _, _ := syscall.SyscallN(
		v.vtbl.Toggle,
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
		v.vtbl.Get_CurrentToggleState,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&state)),
	)
	if ret != 0 {
		return 0, HResult(ret)
	}
	return state, nil
}

func newIUIAutomationTogglePattern(unk *IUnKnown) *IUIAutomationTogglePattern {
	return (*IUIAutomationTogglePattern)(unsafe.Pointer(unk))
}
func NewIUIAutomationTogglePattern(unk *IUnKnown) *IUIAutomationTogglePattern {
	return newIUIAutomationTogglePattern(unk)
}

func (v *IUIAutomationTogglePattern) Release() uint32 {
	return (*IUnKnown)(unsafe.Pointer(v)).Release()
}

// IUIAutomationExpandCollapsePattern
type IUIAutomationExpandCollapsePattern struct {
	vtbl *IUIAutomationExpandCollapsePatternVtbl
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
		v.vtbl.Expand,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

func (v *IUIAutomationExpandCollapsePattern) Collapse() error {
	ret, _, _ := syscall.SyscallN(
		v.vtbl.Collapse,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

func (v *IUIAutomationExpandCollapsePattern) Release() uint32 {
	return (*IUnKnown)(unsafe.Pointer(v)).Release()
}
