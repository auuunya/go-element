package uiautomation

import (
	"errors"
	"syscall"
	"unsafe"
)

// https://learn.microsoft.com/zh-cn/windows/win32/winauto/uiauto-controlpattern-ids
type PatternId uintptr

const (
	UIA_AnnotationPatternId        PatternId = 10023
	UIA_CustomNavigationPatternId  PatternId = 10033
	UIA_DockPatternId              PatternId = 10011
	UIA_DragPatternId              PatternId = 10030
	UIA_DropTargetPatternId        PatternId = 10031
	UIA_ExpandCollapsePatternId    PatternId = 10005
	UIA_GridItemPatternId          PatternId = 10007
	UIA_GridPatternId              PatternId = 10006
	UIA_InvokePatternId            PatternId = 10000
	UIA_ItemContainerPatternId     PatternId = 10019
	UIA_LegacyIAccessiblePatternId PatternId = 10018
	UIA_MultipleViewPatternId      PatternId = 10008
	UIA_ObjectModelPatternId       PatternId = 10022
	UIA_RangeValuePatternId        PatternId = 10003
	UIA_ScrollItemPatternId        PatternId = 10017
	UIA_ScrollPatternId            PatternId = 10004
	UIA_SelectionItemPatternId     PatternId = 10010
	UIA_SelectionPatternId         PatternId = 10001
	UIA_SpreadsheetPatternId       PatternId = 10026
	UIA_SpreadsheetItemPatternId   PatternId = 10027
	UIA_StylesPatternId            PatternId = 10025
	UIA_SynchronizedInputPatternId PatternId = 10021
	UIA_TableItemPatternId         PatternId = 10013
	UIA_TablePatternId             PatternId = 10012
	UIA_TextChildPatternId         PatternId = 10029
	UIA_TextEditPatternId          PatternId = 10032
	UIA_TextPatternId              PatternId = 10014
	UIA_TextPattern2Id             PatternId = 10024
	UIA_TogglePatternId            PatternId = 10015
	UIA_TransformPatternId         PatternId = 10016
	UIA_TransformPattern2Id        PatternId = 10028
	UIA_ValuePatternId             PatternId = 10002
	UIA_VirtualizedItemPatternId   PatternId = 10020
	UIA_WindowPatternId            PatternId = 10009
)

var (
	ErrorBstrPointerNil = errors.New("BSTR pointer is nil")
)

// windows 8 start support
// UIA_AnnotationPatternId
type IAnnotationProvider struct {
	vtbl *IUnKnown
}
type IAnnotationProviderVtbl struct {
	IUnKnownVtbl
	Get_AnnotationTypeId   uintptr
	Get_AnnotationTypeName uintptr
	Get_Author             uintptr
	Get_DateTime           uintptr
	Get_Target             uintptr
}

type IRawElementProviderSimple struct {
	vtbl *IUnKnown
}

type IRawElementProviderSimpleVtbl struct {
	IUnKnownVtbl
	Get_HostRawElementProvider uintptr
	Get_ProviderOptions        uintptr
	GetPatternProvider         uintptr
	GetPropertyValue           uintptr
}

func NewIAnnotationProvider(obj uintptr) *IAnnotationProvider {
	return newIAnnotationProvider(obj)
}
func newIAnnotationProvider(obj uintptr) *IAnnotationProvider {
	return (*IAnnotationProvider)(unsafe.Pointer(obj))
}
func (v *IAnnotationProvider) Get_AnnotationTypeId() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IAnnotationProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_AnnotationTypeId,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IAnnotationProvider) Get_AnnotationTypeName() string {
	var bstr uintptr
	syscall.SyscallN(
		(*IAnnotationProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_AnnotationTypeName,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return ""
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(uintptr(bstr))
	return retVal
}
func (v *IAnnotationProvider) Get_Author() string {
	var bstr uintptr
	syscall.SyscallN(
		(*IAnnotationProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_Author,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return ""
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(uintptr(bstr))
	return retVal
}
func (v *IAnnotationProvider) Get_DateTime() string {
	var bstr uintptr
	syscall.SyscallN(
		(*IAnnotationProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_DateTime,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return ""
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(uintptr(bstr))
	return retVal
}
func (v *IAnnotationProvider) Get_Target() *IRawElementProviderSimple {
	var retVal *IRawElementProviderSimple
	syscall.SyscallN(
		(*IAnnotationProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_Target,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

// CustomNavigation: windows 10 start support
// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationclient/nn-uiautomationclient-iuiautomationelement4
// int automationelement4

type IDockProvider struct {
	vtbl *IUnKnown
}

type IDockProviderVtbl struct {
	IUnKnownVtbl
	Get_DockPosition uintptr
	SetDockPosition  uintptr
}

// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationcore/ne-uiautomationcore-dockposition
type DockPosition int

const (
	DockPosition_Top DockPosition = iota
	DockPosition_Left
	DockPosition_Bottom
	DockPosition_Right
	DockPosition_Fill
	DockPosition_None
)

func NewIDockProvider(obj uintptr) *IDockProvider {
	return newIDockProvider(obj)
}
func newIDockProvider(obj uintptr) *IDockProvider {
	return (*IDockProvider)(unsafe.Pointer(obj))
}
func (v *IDockProvider) Get_DockPosition() *DockPosition {
	var retVal *DockPosition
	syscall.SyscallN(
		(*IDockProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_DockPosition,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IDockProvider) SetDockPosition(position *DockPosition) error {
	ret, _, _ := syscall.SyscallN(
		(*IDockProviderVtbl)(unsafe.Pointer(v.vtbl)).SetDockPosition,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(position)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

type IDragProvider struct {
	vtbl *IUnKnown
}
type IDragProviderVtbl struct {
	IUnKnownVtbl
	Get_DropEffect  uintptr
	Get_DropEffects uintptr
	Get_IsGrabbed   uintptr
	GetGrabbedItems uintptr
}

func NewIDragProvider(obj uintptr) *IDragProvider {
	return newIDragProvider(obj)
}
func newIDragProvider(obj uintptr) *IDragProvider {
	return (*IDragProvider)(unsafe.Pointer(obj))
}
func (v *IDragProvider) Get_DropEffect() string {
	var bstr uintptr
	syscall.SyscallN(
		(*IDragProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_DropEffect,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return ""
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(uintptr(bstr))
	return retVal
}
func (v *IDragProvider) Get_DropEffects() *SafeArray {
	var retVal *SafeArray
	syscall.SyscallN(
		(*IDragProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_DropEffects,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IDragProvider) Get_IsGrabbed() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IDragProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_IsGrabbed,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IDragProvider) GetGrabbedItems() *SafeArray {
	var retVal *SafeArray
	syscall.SyscallN(
		(*IDragProviderVtbl)(unsafe.Pointer(v.vtbl)).GetGrabbedItems,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

type IDropTargetProvider struct {
	vtbl *IUnKnown
}
type IDropTargetProviderVtbl struct {
	IUnKnownVtbl
	Get_DropTargetEffect  uintptr
	Get_DropTargetEffects uintptr
}

func NewIDropTargetProvider(obj uintptr) *IDropTargetProvider {
	return newIDropTargetProvider(obj)
}
func newIDropTargetProvider(obj uintptr) *IDropTargetProvider {
	return (*IDropTargetProvider)(unsafe.Pointer(obj))
}
func (v *IDropTargetProvider) Get_DropTargetEffect() string {
	var bstr uintptr
	syscall.SyscallN(
		(*IDropTargetProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_DropTargetEffect,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return ""
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(uintptr(bstr))
	return retVal
}
func (v *IDropTargetProvider) Get_DropTargetEffects() *SafeArray {
	var retVal *SafeArray
	syscall.SyscallN(
		(*IDropTargetProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_DropTargetEffects,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

type IExpandCollapseProvider struct {
	vtbl *IUnKnown
}
type IExpandCollapseProviderVtbl struct {
	IUnKnownVtbl
	Collapse                uintptr
	Expand                  uintptr
	Get_ExpandCollapseState uintptr
}
type ExpandCollapseState int

const (
	ExpandCollapseState_Collapsed ExpandCollapseState = iota
	ExpandCollapseState_Expanded
	ExpandCollapseState_PartiallyExpanded
	ExpandCollapseState_LeafNode
)

func NewIExpandCollapseProvider(obj uintptr) *IExpandCollapseProvider {
	return newIExpandCollapseProvider(obj)
}
func newIExpandCollapseProvider(obj uintptr) *IExpandCollapseProvider {
	return (*IExpandCollapseProvider)(unsafe.Pointer(obj))
}
func (v *IExpandCollapseProvider) Collapse() (int32, error) {
	ret, _, _ := syscall.SyscallN(
		(*IExpandCollapseProviderVtbl)(unsafe.Pointer(v.vtbl)).Collapse,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return 0, HResult(ret)
	}
	return int32(ret), nil
}
func (v *IExpandCollapseProvider) Expand() (int32, error) {
	ret, _, _ := syscall.SyscallN(
		(*IExpandCollapseProviderVtbl)(unsafe.Pointer(v.vtbl)).Expand,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return 0, HResult(ret)
	}
	return int32(ret), nil
}
func (v *IExpandCollapseProvider) GetExpandCollapseState() *ExpandCollapseState {
	var retVal *ExpandCollapseState
	syscall.SyscallN(
		(*IExpandCollapseProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_ExpandCollapseState,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

type IGridItemProvider struct {
	vtbl *IUnKnown
}
type IGridItemProviderVtbl struct {
	IUnKnownVtbl
	Get_Column         uintptr
	Get_ColumnSpan     uintptr
	Get_ContainingGrid uintptr
	Get_Row            uintptr
	Get_RowSpan        uintptr
}

func NewIGridItemProvider(obj uintptr) *IGridItemProvider {
	return newIGridItemProvider(obj)
}
func newIGridItemProvider(obj uintptr) *IGridItemProvider {
	return (*IGridItemProvider)(unsafe.Pointer(obj))
}
func (v *IGridItemProvider) Get_Column() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IGridItemProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_Column,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IGridItemProvider) Get_ColumnSpan() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IGridItemProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_ColumnSpan,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IGridItemProvider) Get_ContainingGrid() *IRawElementProviderSimple {
	var retVal *IRawElementProviderSimple
	syscall.SyscallN(
		(*IGridItemProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_ColumnSpan,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IGridItemProvider) Get_Row() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IGridItemProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_Row,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IGridItemProvider) Get_RowSpan() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IGridItemProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_RowSpan,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

type IGridProvider struct {
	vtbl *IUnKnown
}
type IGridProviderVtbl struct {
	IUnKnownVtbl
	Get_ColumnCount uintptr
	Get_RowCount    uintptr
	GetItem         uintptr
}

func NewIGridProvider(obj uintptr) *IGridProvider {
	return newIGridProvider(obj)
}
func newIGridProvider(obj uintptr) *IGridProvider {
	return (*IGridProvider)(unsafe.Pointer(obj))
}
func (v *IGridProvider) Get_ColumnCount() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IGridProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_ColumnCount,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IGridProvider) Get_RowCount() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IGridProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_RowCount,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IGridProvider) GetItem(row, column int32) (*IRawElementProviderSimple, error) {
	var retVal *IRawElementProviderSimple
	ret, _, _ := syscall.SyscallN(
		(*IGridProviderVtbl)(unsafe.Pointer(v.vtbl)).GetItem,
		uintptr(unsafe.Pointer(v)),
		uintptr(row),
		uintptr(column),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

type IInvokeProvider struct {
	vtbl *IUnKnown
}
type IInvokeProviderVtbl struct {
	IUnKnownVtbl
	Invoke uintptr
}

func NewIInvokeProvider(obj uintptr) *IInvokeProvider {
	return newIInvokeProvider(obj)
}
func newIInvokeProvider(obj uintptr) *IInvokeProvider {
	return (*IInvokeProvider)(unsafe.Pointer(obj))
}
func (v *IInvokeProvider) Invoke() error {
	ret, _, _ := syscall.SyscallN(
		(*IInvokeProviderVtbl)(unsafe.Pointer(v.vtbl)).Invoke,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

type IItemContainerProvider struct {
	vtbl *IUnKnown
}
type IItemContainerProviderVtbl struct {
	IUnKnownVtbl
	FindItemByProperty uintptr
}

func NewIItemContainerProvider(obj uintptr) *IItemContainerProvider {
	return newIItemContainerProvider(obj)
}
func newIItemContainerProvider(obj uintptr) *IItemContainerProvider {
	return (*IItemContainerProvider)(unsafe.Pointer(obj))
}
func (v *IItemContainerProvider) FindItemByProperty(simple *IRawElementProviderSimple, id uintptr, val *VARIANT) (*IRawElementProviderSimple, error) {
	var retVal *IRawElementProviderSimple
	ret, _, _ := syscall.SyscallN(
		(*IItemContainerProviderVtbl)(unsafe.Pointer(v.vtbl)).FindItemByProperty,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(simple)),
		uintptr(id),
		uintptr(unsafe.Pointer(val)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

type ILegacyIAccessibleProvider struct {
	vtbl *IUnKnown
}
type ILegacyIAccessibleProviderVtbl struct {
	IUnKnownVtbl
	DoDefaultAction      uintptr
	Get_ChildId          uintptr
	Get_DefaultAction    uintptr
	Get_Description      uintptr
	Get_Help             uintptr
	Get_KeyboardShortcut uintptr
	Get_Name             uintptr
	Get_Role             uintptr
	Get_State            uintptr
	Get_Value            uintptr
	GetIAccessible       uintptr
	GetSelection         uintptr
	Select               uintptr
	SetValue             uintptr
}

func newILegacyIAccessibleProvider(obj uintptr) *ILegacyIAccessibleProvider {
	return (*ILegacyIAccessibleProvider)(unsafe.Pointer(obj))
}
func NewILegacyIAccessibleProvider(obj uintptr) *ILegacyIAccessibleProvider {
	return newILegacyIAccessibleProvider(obj)
}
func (v *ILegacyIAccessibleProvider) DoDefaultAction() error {
	ret, _, _ := syscall.SyscallN(
		(*ILegacyIAccessibleProviderVtbl)(unsafe.Pointer(v.vtbl)).DoDefaultAction,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *ILegacyIAccessibleProvider) Get_ChildId() int32 {
	var retVal int32
	syscall.SyscallN(
		(*ILegacyIAccessibleProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_ChildId,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *ILegacyIAccessibleProvider) Get_DefaultAction() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*ILegacyIAccessibleProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_DefaultAction,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return "", ErrorBstrPointerNil
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(uintptr(bstr))
	return retVal, nil
}
func (v *ILegacyIAccessibleProvider) Get_Description() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*ILegacyIAccessibleProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_Description,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return "", ErrorBstrPointerNil
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(uintptr(bstr))
	return retVal, nil
}
func (v *ILegacyIAccessibleProvider) Get_Help() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*ILegacyIAccessibleProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_Help,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return "", ErrorBstrPointerNil
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(uintptr(bstr))
	return retVal, nil
}
func (v *ILegacyIAccessibleProvider) Get_KeyboardShortcut() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*ILegacyIAccessibleProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_KeyboardShortcut,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return "", ErrorBstrPointerNil
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(uintptr(bstr))
	return retVal, nil
}
func (v *ILegacyIAccessibleProvider) Get_Name() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*ILegacyIAccessibleProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_Name,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return "", ErrorBstrPointerNil
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(uintptr(bstr))
	return retVal, nil
}
func (v *ILegacyIAccessibleProvider) Get_Role() uint32 {
	var retVal uint32
	syscall.SyscallN(
		(*ILegacyIAccessibleProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_Role,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *ILegacyIAccessibleProvider) Get_State() uint32 {
	var retVal uint32
	syscall.SyscallN(
		(*ILegacyIAccessibleProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_State,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *ILegacyIAccessibleProvider) Get_Value() uint32 {
	var retVal uint32
	syscall.SyscallN(
		(*ILegacyIAccessibleProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_Value,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *ILegacyIAccessibleProvider) GetIAccessible() (*IAccessible, error) {
	var retVal *IAccessible
	ret, _, _ := syscall.SyscallN(
		(*ILegacyIAccessibleProviderVtbl)(unsafe.Pointer(v.vtbl)).GetIAccessible,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *ILegacyIAccessibleProvider) GetSelection() (*SafeArray, error) {
	var retVal *SafeArray
	ret, _, _ := syscall.SyscallN(
		(*ILegacyIAccessibleProviderVtbl)(unsafe.Pointer(v.vtbl)).GetSelection,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

var (
	// https://learn.microsoft.com/zh-cn/windows/win32/winauto/selflag
	SELFLAG_NONE            = 0
	SELFLAG_TAKEFOCUS       = 0x1
	SELFLAG_TAKESELECTION   = 0x2
	SELFLAG_EXTENDSELECTION = 0x4
	SELFLAG_ADDSELECTION    = 0x8
	SELFLAG_REMOVESELECTION = 0x10
)

func (v *ILegacyIAccessibleProvider) Select(flag int32) error {
	ret, _, _ := syscall.SyscallN(
		(*ILegacyIAccessibleProviderVtbl)(unsafe.Pointer(v.vtbl)).Select,
		uintptr(unsafe.Pointer(v)),
		uintptr(flag),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *ILegacyIAccessibleProvider) SetValue(val string) error {
	retVal, err := syscall.UTF16PtrFromString(val)
	if err != nil {
		return err
	}
	ret, _, _ := syscall.SyscallN(
		(*ILegacyIAccessibleProviderVtbl)(unsafe.Pointer(v.vtbl)).SetValue,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(retVal)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

type IMultipleViewProvider struct {
	vtbl *IUnKnown
}
type IMultipleViewProviderVtbl struct {
	IUnKnownVtbl
	Get_CurrentView   uintptr
	GetSupportedViews uintptr
	GetViewName       uintptr
	SetCurrentView    uintptr
}

func newIMultipleViewProvider(obj uintptr) *IMultipleViewProvider {
	return (*IMultipleViewProvider)(unsafe.Pointer(obj))
}
func NewIMultipleViewProvider(obj uintptr) *IMultipleViewProvider {
	return newIMultipleViewProvider(obj)
}
func (v *IMultipleViewProvider) Get_CurrentView() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IMultipleViewProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentView,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IMultipleViewProvider) GetSupportedViews() (*SafeArray, error) {
	var retVal *SafeArray
	ret, _, _ := syscall.SyscallN(
		(*IMultipleViewProviderVtbl)(unsafe.Pointer(v.vtbl)).GetSupportedViews,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IMultipleViewProvider) GetViewName(viewId int32) (string, error) {
	var bstr uintptr
	ret, _, _ := syscall.SyscallN(
		(*IMultipleViewProviderVtbl)(unsafe.Pointer(v.vtbl)).GetViewName,
		uintptr(unsafe.Pointer(v)),
		uintptr(viewId),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if ret != 0 {
		return "", HResult(ret)
	}
	if bstr == 0 {
		return "", ErrorBstrPointerNil
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(uintptr(bstr))
	return retVal, nil
}
func (v *IMultipleViewProvider) SetCurrentView(viewId int32) error {
	ret, _, _ := syscall.SyscallN(
		(*IMultipleViewProviderVtbl)(unsafe.Pointer(v.vtbl)).SetCurrentView,
		uintptr(unsafe.Pointer(v)),
		uintptr(viewId),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

// windows 8 start support
type IObjectModelProvider struct {
	vtbl *IUnKnown
}
type IObjectModelProviderVtbl struct {
	IUnKnownVtbl
	GetUnderlyingObjectModel uintptr
}

func newIObjectModelProvider(obj uintptr) *IObjectModelProvider {
	return (*IObjectModelProvider)(unsafe.Pointer(obj))
}
func NewIObjectModelProvider(obj uintptr) *IObjectModelProvider {
	return newIObjectModelProvider(obj)
}
func (v *IObjectModelProvider) GetUnderlyingObjectModel() (*IUnKnown, error) {
	var retVal *IUnKnown
	ret, _, _ := syscall.SyscallN(
		(*IObjectModelProviderVtbl)(unsafe.Pointer(v.vtbl)).GetUnderlyingObjectModel,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

type IRangeValueProvider struct {
	vtbl *IUnKnown
}
type IRangeValueProviderVtbl struct {
	IUnKnownVtbl
	Get_IsReadOnly  uintptr
	Get_LargeChange uintptr
	Get_Maximum     uintptr
	Get_Minimum     uintptr
	Get_SmallChange uintptr
	Get_Value       uintptr
	SetValue        uintptr
}

func newIRangeValueProvider(obj uintptr) *IRangeValueProvider {
	return (*IRangeValueProvider)(unsafe.Pointer(obj))
}
func NewIRangeValueProvider(obj uintptr) *IRangeValueProvider {
	return newIRangeValueProvider(obj)
}
func (v *IRangeValueProvider) Get_IsReadOnly() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IRangeValueProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_IsReadOnly,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

type IScrollItemProvider struct {
	vtbl *IUnKnown
}
type IScrollItemProviderVtbl struct {
	IUnKnownVtbl
	ScrollIntoView uintptr
}

type IScrollProvider struct {
	vtbl *IUnKnown
}
type IScrollProviderVtbl struct {
	IUnKnownVtbl
	Get_HorizontallyScrollable  uintptr
	Get_HorizontalScrollPercent uintptr
	Get_HorizontalViewSize      uintptr
	Get_VerticallyScrollable    uintptr
	Get_VerticalScrollPercent   uintptr
	Get_VerticalViewSize        uintptr
	Scroll                      uintptr
	SetScrollPercent            uintptr
}

type ISelectionItemProvider struct {
	vtbl *IUnKnown
}
type ISelectionItemProviderVtbl struct {
	IUnKnownVtbl
	AddToSelection         uintptr
	Get_IsSelected         uintptr
	Get_SelectionContainer uintptr
	RemoveFromSelection    uintptr
	Select                 uintptr
}

type ISelectionProvider struct {
	vtbl *IUnKnown
}
type ISelectionProviderVtbl struct {
	IUnKnownVtbl
	Get_CanSelectMultiple   uintptr
	Get_IsSelectionRequired uintptr
	GetSelection            uintptr
}

type ISpreadsheetProvider struct {
	vtbl *IUnKnown
}
type ISpreadsheetProviderVtbl struct {
	IUnKnownVtbl
	GetItemByName uintptr
}

type ISpreadsheetItemProvider struct {
	vtbl *IUnKnown
}
type ISpreadsheetItemProviderVtbl struct {
	IUnKnownVtbl
	Get_Formula          uintptr
	GetAnnotationObjects uintptr
	GetAnnotationTypes   uintptr
}

// windows 8 start support
type IStylesProvider struct {
	vtbl *IUnKnown
}
type IStylesProviderVtbl struct {
	IUnKnownVtbl
	Get_ExtendedPropertvs uintptr
	Get_FillColor         uintptr
	Get_FillPatternColor  uintptr
	Get_FillPatternStyle  uintptr
	Get_Shape             uintptr
	Get_StyleId           uintptr
	Get_StyleName         uintptr
}
type ISynchronizedInputProvider struct {
	vtbl *IUnKnown
}

type ISynchronizedInputProviderVtbl struct {
	IUnKnownVtbl
	Cancel         uintptr
	StartListening uintptr
}
type ITableItemProvider struct {
	vtbl *IUnKnown
}
type ITableItemProviderVtbl struct {
	IUnKnownVtbl
	GetColumnHeaderItems uintptr
	GetRowHeaderItems    uintptr
}

type ITableProvider struct {
	vtbl *IUnKnown
}
type ITableProviderVtbl struct {
	IUnKnownVtbl
	Get_RowOrColumnMajor uintptr
	GetColumnHeaders     uintptr
	GetRowHeaders        uintptr
}
type ITextChildProvider struct {
	vtbl *IUnKnown
}
type ITextChildProviderVtbl struct {
	IUnKnownVtbl
	Get_TextContainer uintptr
	Get_TextRange     uintptr
}

type ITextEditProvider struct {
	vtbl *IUnKnown
}
type ITextEditProviderVtbl struct {
	IUnKnownVtbl
	GetActiveComposition uintptr
	GetConversionTarget  uintptr
}
type ITextProvider struct {
	vtbl *IUnKnown
}
type ITextProviderVtbl struct {
	IUnKnownVtbl
	Get_DocumentRange          uintptr
	Get_SupportedTextSelection uintptr
	GetSelection               uintptr
	GetVisibleRanges           uintptr
	RangeFromChild             uintptr
	RangeFromPoint             uintptr
}

// windows 8 start support
type ITextProvider2 struct {
	vtbl *IUnKnown
}
type ITextProvider2Vtbl struct {
	IUnKnownVtbl
	GetCaretRange       uintptr
	RangeFromAnnotation uintptr
}
type IToggleProvider struct {
	vtbl *IUnKnown
}
type IToggleProviderVtbl struct {
	IUnKnownVtbl
	Get_ToggleState uintptr
	Toggle          uintptr
}
type ITransformProvider struct {
	vtbl *IUnKnown
}
type ITransformProviderVtbl struct {
	IUnKnownVtbl
	Get_CanMove   uintptr
	Get_CanResize uintptr
	Get_CanRotate uintptr
	Move          uintptr
	Resize        uintptr
	Rotate        uintptr
}

// windows 8 start support
type ITransformProvider2 struct {
	vtbl *IUnKnown
}
type ITransformProvider2Vtbl struct {
	IUnKnownVtbl
	Get_CanZoom     uintptr
	Get_ZoomLevel   uintptr
	Get_ZoomMaximum uintptr
	Get_ZoomMinimum uintptr
	Zoom            uintptr
	ZoomByUnit      uintptr
}
type IValueProvider struct {
	vtbl *IUnKnown
}
type IValueProviderVtbl struct {
	IUnKnownVtbl
	Get_IsReadOnly uintptr
	Get_Value      uintptr
	SetValue       uintptr
}

type IVirtualizedItemProvider struct {
	vtbl *IUnKnown
}
type IVirtualizedItemProviderVtbl struct {
	IUnKnownVtbl
	Realize uintptr
}
type IWindowProvider struct {
	vtbl *IUnKnown
}
type IWindowProviderVtbl struct {
	IUnKnownVtbl
	Close                      uintptr
	Get_CanMaximize            uintptr
	Get_CanMinimize            uintptr
	Get_IsModal                uintptr
	Get_IsTopmost              uintptr
	Get_WindowInteractionState uintptr
	Get_WindowVisualState      uintptr
	SetVisualState             uintptr
	WaitForInputIdle           uintptr
}
