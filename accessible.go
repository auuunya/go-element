package uiautomation

import (
	"syscall"
	"unsafe"
)

type IAccessible struct {
	vtbl *IDispatch
}

type IAccessibleVtbl struct {
	// https://learn.microsoft.com/zh-cn/windows/win32/api/oleacc/nn-oleacc-iaccessible
	IDispatchVtbl
	AccDoDefaultAction      uintptr
	AccHitTest              uintptr
	AccLocation             uintptr
	AccNavigate             uintptr // unsupported
	AccSelect               uintptr
	Get_accChild            uintptr
	Get_accChildCount       uintptr
	Get_accDefaultAction    uintptr
	Get_accDescription      uintptr
	Get_accFocus            uintptr
	Get_accHelp             uintptr
	Get_accHelpTopic        uintptr // unsupported
	Get_accKeyboardShortcut uintptr
	Get_accName             uintptr
	Get_accParent           uintptr
	Get_accRole             uintptr
	Get_accSelection        uintptr
	Get_accState            uintptr
	Get_accValue            uintptr
	Put_accName             uintptr // unsupported
	Put_accValue            uintptr
}

func newIAccessible(unk *IDispatch) *IAccessible {
	return (*IAccessible)(unsafe.Pointer(unk))
}
func NewIAccessible(unk *IDispatch) *IAccessible {
	return newIAccessible(unk)
}
func (v *IAccessible) AccDoDefaultAction(in *VARIANT) error {
	ret, _, _ := syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).AccDoDefaultAction,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *IAccessible) AccHitTest(in, in2 int32) (*VARIANT, error) {
	var retVal *VARIANT
	ret, _, _ := syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).AccHitTest,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(in2),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IAccessible) AccLocation(in, in2, in3, in4 int32) (*VARIANT, error) {
	var retVal *VARIANT
	ret, _, _ := syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).AccLocation,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(in2),
		uintptr(in3),
		uintptr(in4),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IAccessible) AccNavigate(in int32, in2 *VARIANT) (*VARIANT, error) {
	var retVal *VARIANT
	ret, _, _ := syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).AccNavigate,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(unsafe.Pointer(in2)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IAccessible) AccSelect(in int32, in2 *VARIANT) error {
	ret, _, _ := syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).AccSelect,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(unsafe.Pointer(in2)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *IAccessible) Get_accChild(in *VARIANT) (*IDispatch, error) {
	var retVal *IDispatch
	ret, _, _ := syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).Get_accChild,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IAccessible) Get_accChildCount() (int32, error) {
	var retVal int32
	ret, _, _ := syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).Get_accChildCount,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return -1, HResult(ret)
	}
	return retVal, nil
}
func (v *IAccessible) Get_accDefaultAction(in *VARIANT) (string, error) {
	var bstr uintptr
	ret, _, _ := syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).Get_accDefaultAction,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if ret != 0 {
		return "", HResult(ret)
	}

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IAccessible) Get_accDescription(in *VARIANT) (string, error) {
	var bstr uintptr
	ret, _, _ := syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).Get_accDescription,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if ret != 0 {
		return "", HResult(ret)
	}

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IAccessible) Get_accFocus() (*VARIANT, error) {
	var retVal *VARIANT
	ret, _, _ := syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).Get_accFocus,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IAccessible) Get_accHelp(in *VARIANT) (string, error) {
	var bstr uintptr
	ret, _, _ := syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).Get_accHelp,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if ret != 0 {
		return "", HResult(ret)
	}

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IAccessible) Get_accHelpTopic(in *VARIANT, in2 int32) (string, error) {
	var bstr uintptr
	ret, _, _ := syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).Get_accHelpTopic,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(in2),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if ret != 0 {
		return "", HResult(ret)
	}

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IAccessible) Get_accKeyboardShortcut(in *VARIANT) (string, error) {
	var bstr uintptr
	ret, _, _ := syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).Get_accKeyboardShortcut,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if ret != 0 {
		return "", HResult(ret)
	}

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IAccessible) Get_accName(in *VARIANT) (string, error) {
	var bstr uintptr
	ret, _, _ := syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).Get_accName,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if ret != 0 {
		return "", HResult(ret)
	}

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IAccessible) Get_accParent() (*IDispatch, error) {
	var retVal *IDispatch
	ret, _, _ := syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).Get_accParent,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IAccessible) Get_accRole(in *VARIANT) (*VARIANT, error) {
	var retVal *VARIANT
	ret, _, _ := syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).Get_accRole,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IAccessible) Get_accSelection() (*VARIANT, error) {
	var retVal *VARIANT
	ret, _, _ := syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).Get_accSelection,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IAccessible) Get_accState(in *VARIANT) (*VARIANT, error) {
	var retVal *VARIANT
	ret, _, _ := syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).Get_accState,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IAccessible) Get_accValue(in *VARIANT) (string, error) {
	var bstr uintptr
	ret, _, _ := syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).Get_accValue,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if ret != 0 {
		return "", HResult(ret)
	}

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IAccessible) Put_accName(in *VARIANT, in2 string) error {
	retVal, err := syscall.UTF16PtrFromString(in2)
	if err != nil {
		return err
	}
	syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).Put_accName,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(retVal)),
	)
	return nil
}
func (v *IAccessible) Put_accValue(in *VARIANT, in2 string) error {
	retVal, err := syscall.UTF16PtrFromString(in2)
	if err != nil {
		return err
	}
	ret, _, _ := syscall.SyscallN(
		(*IAccessibleVtbl)(unsafe.Pointer(v.vtbl)).Put_accValue,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(retVal)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

type IUIAutomationLegacyIAccessiblePattern struct {
	vtbl *IUnKnown
}
type IUIAutomationLegacyIAccessiblePatternVtbl struct {
	IUnKnownVtbl
	DoDefaultAction             uintptr
	Get_CachedChildId           uintptr
	Get_CachedDefaultAction     uintptr
	Get_CachedDescription       uintptr
	Get_CachedHelp              uintptr
	Get_CachedKeyboardShortcut  uintptr
	Get_CachedName              uintptr
	Get_CachedRole              uintptr
	Get_CachedState             uintptr
	Get_CachedValue             uintptr
	Get_CurrentChildId          uintptr
	Get_CurrentDefaultAction    uintptr
	Get_CurrentDescription      uintptr
	Get_CurrentHelp             uintptr
	Get_CurrentKeyboardShortcut uintptr
	Get_CurrentName             uintptr
	Get_CurrentRole             uintptr
	Get_CurrentState            uintptr
	Get_CurrentValue            uintptr
	GetCachedSelection          uintptr
	GetCurrentSelection         uintptr
	GetIAccessible              uintptr
	Select                      uintptr
	SetValue                    uintptr
}

func newIUIAutomationLegacyIAccessiblePattern(unk *IDispatch) *IUIAutomationLegacyIAccessiblePattern {
	return (*IUIAutomationLegacyIAccessiblePattern)(unsafe.Pointer(unk))
}
func NewIUIAutomationLegacyIAccessiblePattern(unk *IDispatch) *IUIAutomationLegacyIAccessiblePattern {
	return newIUIAutomationLegacyIAccessiblePattern(unk)
}
func (v *IUIAutomationLegacyIAccessiblePattern) DoDefaultAction() error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).DoDefaultAction,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *IUIAutomationLegacyIAccessiblePattern) Get_CachedChildId() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).Get_CachedChildId,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomationLegacyIAccessiblePattern) Get_CachedDefaultAction() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).Get_CachedDefaultAction,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomationLegacyIAccessiblePattern) Get_CachedDescription() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).Get_CachedDescription,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomationLegacyIAccessiblePattern) Get_CachedHelp() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).Get_CachedHelp,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomationLegacyIAccessiblePattern) Get_CachedKeyboardShortcut() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).Get_CachedKeyboardShortcut,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomationLegacyIAccessiblePattern) Get_CachedName() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).Get_CachedName,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomationLegacyIAccessiblePattern) Get_CachedRole() uint32 {
	var retVal uint32
	syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).Get_CachedRole,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomationLegacyIAccessiblePattern) Get_CachedState() uint32 {
	var retVal uint32
	syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).Get_CachedState,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomationLegacyIAccessiblePattern) Get_CachedValue() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).Get_CachedValue,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomationLegacyIAccessiblePattern) Get_CurrentChildId() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentChildId,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomationLegacyIAccessiblePattern) Get_CurrentDefaultAction() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentDefaultAction,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomationLegacyIAccessiblePattern) Get_CurrentDescription() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentDescription,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomationLegacyIAccessiblePattern) Get_CurrentHelp() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentHelp,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomationLegacyIAccessiblePattern) Get_CurrentKeyboardShortcut() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentKeyboardShortcut,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomationLegacyIAccessiblePattern) Get_CurrentName() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentName,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomationLegacyIAccessiblePattern) Get_CurrentRole() uint32 {
	var retVal uint32
	syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentRole,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomationLegacyIAccessiblePattern) Get_CurrentState() uint32 {
	var retVal uint32
	syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentState,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomationLegacyIAccessiblePattern) Get_CurrentValue() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentValue,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomationLegacyIAccessiblePattern) GetCachedSelection() (*IUIAutomationElementArray, error) {
	var retVal *IUIAutomationElementArray
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).GetCachedSelection,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomationLegacyIAccessiblePattern) GetCurrentSelection() (*IUIAutomationElementArray, error) {
	var retVal *IUIAutomationElementArray
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).GetCurrentSelection,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomationLegacyIAccessiblePattern) GetIAccessible() (*IAccessible, error) {
	var retVal *IAccessible
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).GetIAccessible,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomationLegacyIAccessiblePattern) Select(in SelFlag) error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).Select,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *IUIAutomationLegacyIAccessiblePattern) SetValue(in string) error {
	retVal, err := syscall.UTF16PtrFromString(in)
	if err != nil {
		return err
	}
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.vtbl)).SetValue,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(retVal)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
