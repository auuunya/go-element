package uiautomation

import (
	"errors"
	"syscall"
	"unsafe"
)

type PatternId uintptr

const (
	// https://learn.microsoft.com/zh-cn/windows/win32/winauto/uiauto-controlpattern-ids
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

type IAnnotationProvider struct {
	// windows 8 start support
	// UIA_AnnotationPatternId
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

type DockPosition int

const (
	// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationcore/ne-uiautomationcore-dockposition
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
func (v *IDragProvider) Get_DropEffects() *TagSafeArray {
	var retVal *TagSafeArray
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
func (v *IDragProvider) GetGrabbedItems() *TagSafeArray {
	var retVal *TagSafeArray
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
func (v *IDropTargetProvider) Get_DropTargetEffects() *TagSafeArray {
	var retVal *TagSafeArray
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
		return -1, HResult(ret)
	}
	return int32(ret), nil
}
func (v *IExpandCollapseProvider) Expand() (int32, error) {
	ret, _, _ := syscall.SyscallN(
		(*IExpandCollapseProviderVtbl)(unsafe.Pointer(v.vtbl)).Expand,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return -1, HResult(ret)
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
func (v *ILegacyIAccessibleProvider) GetSelection() (*TagSafeArray, error) {
	var retVal *TagSafeArray
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
func (v *IMultipleViewProvider) GetSupportedViews() (*TagSafeArray, error) {
	var retVal *TagSafeArray
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

type IObjectModelProvider struct {
	vtbl *IUnKnown
}
type IObjectModelProviderVtbl struct {
	// windows 8 start support
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
func (v *IRangeValueProvider) Get_LargeChange() float64 {
	var retVal float64
	syscall.SyscallN(
		(*IRangeValueProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_LargeChange,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IRangeValueProvider) Get_Maximum() float64 {
	var retVal float64
	syscall.SyscallN(
		(*IRangeValueProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_Maximum,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IRangeValueProvider) Get_Minimum() float64 {
	var retVal float64
	syscall.SyscallN(
		(*IRangeValueProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_Minimum,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IRangeValueProvider) Get_SmallChange() float64 {
	var retVal float64
	syscall.SyscallN(
		(*IRangeValueProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_SmallChange,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IRangeValueProvider) Get_Value() float64 {
	var retVal float64
	syscall.SyscallN(
		(*IRangeValueProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_Value,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

func (v *IRangeValueProvider) SetValue(val float64) error {
	ret, _, _ := syscall.SyscallN(
		(*IRangeValueProviderVtbl)(unsafe.Pointer(v.vtbl)).SetValue,
		uintptr(unsafe.Pointer(v)),
		uintptr(val),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

type IScrollItemProvider struct {
	vtbl *IUnKnown
}
type IScrollItemProviderVtbl struct {
	IUnKnownVtbl
	ScrollIntoView uintptr
}

func newIScrollItemProvider(obj uintptr) *IScrollItemProvider {
	return (*IScrollItemProvider)(unsafe.Pointer(obj))
}
func NewIScrollItemProvider(obj uintptr) *IScrollItemProvider {
	return newIScrollItemProvider(obj)
}
func (v *IScrollItemProvider) ScrollIntoView() error {
	ret, _, _ := syscall.SyscallN(
		(*IScrollItemProviderVtbl)(unsafe.Pointer(v.vtbl)).ScrollIntoView,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
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

func newIScrollProvider(obj uintptr) *IScrollProvider {
	return (*IScrollProvider)(unsafe.Pointer(obj))
}
func NewIScrollProvider(obj uintptr) *IScrollProvider {
	return newIScrollProvider(obj)
}
func (v *IScrollProvider) Get_HorizontallyScrollable() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IScrollProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_HorizontallyScrollable,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IScrollProvider) Get_HorizontalScrollPercent() float64 {
	var retVal float64
	syscall.SyscallN(
		(*IScrollProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_HorizontalScrollPercent,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IScrollProvider) Get_HorizontalViewSize() float64 {
	var retVal float64
	syscall.SyscallN(
		(*IScrollProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_HorizontalViewSize,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IScrollProvider) Get_VerticallyScrollable() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IScrollProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_VerticallyScrollable,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IScrollProvider) Get_VerticalScrollPercent() float64 {
	var retVal float64
	syscall.SyscallN(
		(*IScrollProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_VerticalScrollPercent,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IScrollProvider) Get_VerticalViewSize() float64 {
	var retVal float64
	syscall.SyscallN(
		(*IScrollProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_VerticalViewSize,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IScrollProvider) Scroll(hr, vr ScrollAmount) error {
	ret, _, _ := syscall.SyscallN(
		(*IScrollProviderVtbl)(unsafe.Pointer(v.vtbl)).Scroll,
		uintptr(unsafe.Pointer(v)),
		uintptr(hr),
		uintptr(vr),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *IScrollProvider) SetScrollPercent(hr, vr float64) error {
	ret, _, _ := syscall.SyscallN(
		(*IScrollProviderVtbl)(unsafe.Pointer(v.vtbl)).SetScrollPercent,
		uintptr(unsafe.Pointer(v)),
		uintptr(hr),
		uintptr(vr),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
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

func newISelectionItemProvider(obj uintptr) *ISelectionItemProvider {
	return (*ISelectionItemProvider)(unsafe.Pointer(obj))
}
func NewISelectionItemProvider(obj uintptr) *ISelectionItemProvider {
	return newISelectionItemProvider(obj)
}
func (v *ISelectionItemProvider) AddToSelection() error {
	ret, _, _ := syscall.SyscallN(
		(*ISelectionItemProviderVtbl)(unsafe.Pointer(v.vtbl)).AddToSelection,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

func (v *ISelectionItemProvider) Get_IsSelected() int32 {
	var retVal int32
	syscall.SyscallN(
		(*ISelectionItemProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_IsSelected,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *ISelectionItemProvider) Get_SelectionContainer() *IRawElementProviderSimple {
	var retVal *IRawElementProviderSimple
	syscall.SyscallN(
		(*ISelectionItemProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_SelectionContainer,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *ISelectionItemProvider) RemoveFromSelection() error {
	ret, _, _ := syscall.SyscallN(
		(*ISelectionItemProviderVtbl)(unsafe.Pointer(v.vtbl)).RemoveFromSelection,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *ISelectionItemProvider) Select() error {
	ret, _, _ := syscall.SyscallN(
		(*ISelectionItemProviderVtbl)(unsafe.Pointer(v.vtbl)).Select,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
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

func newISelectionProvider(obj uintptr) *ISelectionProvider {
	return (*ISelectionProvider)(unsafe.Pointer(obj))
}
func NewISelectionProvider(obj uintptr) *ISelectionProvider {
	return newISelectionProvider(obj)
}
func (v *ISelectionProvider) Get_CanSelectMultiple() int32 {
	var retVal int32
	syscall.SyscallN(
		(*ISelectionProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_CanSelectMultiple,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *ISelectionProvider) Get_IsSelectionRequired() int32 {
	var retVal int32
	syscall.SyscallN(
		(*ISelectionProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_IsSelectionRequired,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *ISelectionProvider) GetSelection() (*TagSafeArray, error) {
	var retVal *TagSafeArray
	ret, _, _ := syscall.SyscallN(
		(*ISelectionProviderVtbl)(unsafe.Pointer(v.vtbl)).GetSelection,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

type ISpreadsheetProvider struct {
	vtbl *IUnKnown
}
type ISpreadsheetProviderVtbl struct {
	IUnKnownVtbl
	GetItemByName uintptr
}

func newISpreadsheetProvider(obj uintptr) *ISpreadsheetProvider {
	return (*ISpreadsheetProvider)(unsafe.Pointer(obj))
}
func NewISpreadsheetProvider(obj uintptr) *ISpreadsheetProvider {
	return newISpreadsheetProvider(obj)
}
func (v *ISpreadsheetProvider) GetItemByName(name string) (*IRawElementProviderSimple, error) {
	var retVal *IRawElementProviderSimple
	retVal2, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return nil, err
	}
	ret, _, _ := syscall.SyscallN(
		(*ISpreadsheetProviderVtbl)(unsafe.Pointer(v.vtbl)).GetItemByName,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(retVal2)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
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

func newISpreadsheetItemProvider(obj uintptr) *ISpreadsheetItemProvider {
	return (*ISpreadsheetItemProvider)(unsafe.Pointer(obj))
}
func NewISpreadsheetItemProvider(obj uintptr) *ISpreadsheetItemProvider {
	return newISpreadsheetItemProvider(obj)
}
func (v *ISpreadsheetItemProvider) Get_Formula() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*ISpreadsheetItemProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_Formula,
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
func (v *ISpreadsheetItemProvider) GetAnnotationObjects() (*TagSafeArray, error) {
	var retVal *TagSafeArray
	ret, _, _ := syscall.SyscallN(
		(*ISpreadsheetItemProviderVtbl)(unsafe.Pointer(v.vtbl)).GetAnnotationObjects,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *ISpreadsheetItemProvider) GetAnnotationTypes() (*TagSafeArray, error) {
	var retVal *TagSafeArray
	ret, _, _ := syscall.SyscallN(
		(*ISpreadsheetItemProviderVtbl)(unsafe.Pointer(v.vtbl)).GetAnnotationTypes,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

type IStylesProvider struct {
	vtbl *IUnKnown
}
type IStylesProviderVtbl struct {
	// windows 8 start support
	IUnKnownVtbl
	Get_ExtendedPropertvs uintptr
	Get_FillColor         uintptr
	Get_FillPatternColor  uintptr
	Get_FillPatternStyle  uintptr
	Get_Shape             uintptr
	Get_StyleId           uintptr
	Get_StyleName         uintptr
}

func newIStylesProvider(obj uintptr) *IStylesProvider {
	return (*IStylesProvider)(unsafe.Pointer(obj))
}
func NewIStylesProvider(obj uintptr) *IStylesProvider {
	return newIStylesProvider(obj)
}
func (v *IStylesProvider) Get_ExtendedPropertvs() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IStylesProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_ExtendedPropertvs,
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
func (v *IStylesProvider) Get_FillColor() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IStylesProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_FillColor,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IStylesProvider) Get_FillPatternColor() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IStylesProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_FillPatternColor,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IStylesProvider) Get_FillPatternStyle() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IStylesProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_FillPatternStyle,
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
func (v *IStylesProvider) Get_Shape() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IStylesProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_Shape,
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
func (v *IStylesProvider) Get_StyleId() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IStylesProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_StyleId,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IStylesProvider) Get_StyleName() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IStylesProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_StyleName,
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

type ISynchronizedInputProvider struct {
	vtbl *IUnKnown
}

type ISynchronizedInputProviderVtbl struct {
	IUnKnownVtbl
	Cancel         uintptr
	StartListening uintptr
}

func newISynchronizedInputProvider(obj uintptr) *ISynchronizedInputProvider {
	return (*ISynchronizedInputProvider)(unsafe.Pointer(obj))
}
func NewISynchronizedInputProvider(obj uintptr) *ISynchronizedInputProvider {
	return newISynchronizedInputProvider(obj)
}
func (v *ISynchronizedInputProvider) Cancel() error {
	ret, _, _ := syscall.SyscallN(
		(*ISynchronizedInputProviderVtbl)(unsafe.Pointer(v.vtbl)).Cancel,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *ISynchronizedInputProvider) StartListening(in SynchronizedInputType) error {
	ret, _, _ := syscall.SyscallN(
		(*ISynchronizedInputProviderVtbl)(unsafe.Pointer(v.vtbl)).StartListening,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

type ITableItemProvider struct {
	vtbl *IUnKnown
}
type ITableItemProviderVtbl struct {
	IUnKnownVtbl
	GetColumnHeaderItems uintptr
	GetRowHeaderItems    uintptr
}

func newITableItemProvider(obj uintptr) *ITableItemProvider {
	return (*ITableItemProvider)(unsafe.Pointer(obj))
}
func NewITableItemProvider(obj uintptr) *ITableItemProvider {
	return newITableItemProvider(obj)
}
func (v *ITableItemProvider) GetColumnHeaderItems() (*TagSafeArray, error) {
	var retVal *TagSafeArray
	ret, _, _ := syscall.SyscallN(
		(*ITableItemProviderVtbl)(unsafe.Pointer(v.vtbl)).GetColumnHeaderItems,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *ITableItemProvider) GetRowHeaderItems() (*TagSafeArray, error) {
	var retVal *TagSafeArray
	ret, _, _ := syscall.SyscallN(
		(*ITableItemProviderVtbl)(unsafe.Pointer(v.vtbl)).GetRowHeaderItems,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
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

func newITableProvider(obj uintptr) *ITableProvider {
	return (*ITableProvider)(unsafe.Pointer(obj))
}
func NewITableProvider(obj uintptr) *ITableProvider {
	return newITableProvider(obj)
}
func (v *ITableProvider) Get_RowOrColumnMajor() *RowOrColumnMajor {
	var retVal *RowOrColumnMajor
	syscall.SyscallN(
		(*ITableProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_RowOrColumnMajor,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *ITableProvider) GetColumnHeaders() (*TagSafeArray, error) {
	var retVal *TagSafeArray
	ret, _, _ := syscall.SyscallN(
		(*ITableProviderVtbl)(unsafe.Pointer(v.vtbl)).GetColumnHeaders,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *ITableProvider) GetRowHeaders() (*TagSafeArray, error) {
	var retVal *TagSafeArray
	ret, _, _ := syscall.SyscallN(
		(*ITableProviderVtbl)(unsafe.Pointer(v.vtbl)).GetRowHeaders,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

type ITextChildProvider struct {
	vtbl *IUnKnown
}
type ITextChildProviderVtbl struct {
	IUnKnownVtbl
	Get_TextContainer uintptr
	Get_TextRange     uintptr
}

func newITextChildProvider(obj uintptr) *ITextChildProvider {
	return (*ITextChildProvider)(unsafe.Pointer(obj))
}
func NewITextChildProvider(obj uintptr) *ITextChildProvider {
	return newITextChildProvider(obj)
}
func (v *ITextChildProvider) Get_TextContainer() *IRawElementProviderSimple {
	var retVal *IRawElementProviderSimple
	syscall.SyscallN(
		(*ITextChildProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_TextContainer,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *ITextChildProvider) Get_TextRange() *ITextRangeProvider {
	var retVal *ITextRangeProvider
	syscall.SyscallN(
		(*ITextChildProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_TextRange,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

type ITextRangeProvider struct {
	vtbl *IUnKnown
}
type ITextRangeProviderVtbl struct {
	IUnKnownVtbl
	AddToSelection        uintptr
	Clone                 uintptr
	Compare               uintptr
	CompareEndpoints      uintptr
	ExpandToEnclosingUnit uintptr
	FindAttribute         uintptr
	FindText              uintptr
	GetAttributeValue     uintptr
	GetBoundingRectangles uintptr
	GetChildren           uintptr
	GetEnclosingElement   uintptr
	GetText               uintptr
	Move                  uintptr
	MoveEndpointByRange   uintptr
	MoveEndpointByUnit    uintptr
	RemoveFromSelection   uintptr
	ScrollIntoView        uintptr
	Select                uintptr
}

func newITextRangeProvider(obj uintptr) *ITextRangeProvider {
	return (*ITextRangeProvider)(unsafe.Pointer(obj))
}
func NewITextRangeProvider(obj uintptr) *ITextRangeProvider {
	return newITextRangeProvider(obj)
}
func (v *ITextRangeProvider) AddToSelection() error {
	ret, _, _ := syscall.SyscallN(
		(*ITextRangeProviderVtbl)(unsafe.Pointer(v.vtbl)).AddToSelection,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *ITextRangeProvider) Clone() (*ITextRangeProvider, error) {
	var retVal *ITextRangeProvider
	ret, _, _ := syscall.SyscallN(
		(*ITextRangeProviderVtbl)(unsafe.Pointer(v.vtbl)).Clone,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *ITextRangeProvider) Compare(in *ITextRangeProvider) (int32, error) {
	var retVal int32
	ret, _, _ := syscall.SyscallN(
		(*ITextRangeProviderVtbl)(unsafe.Pointer(v.vtbl)).Compare,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return -1, HResult(ret)
	}
	return retVal, nil
}
func (v *ITextRangeProvider) CompareEndpoints(in TextPatternRangeEndpoint, in2 *ITextRangeProvider, in3 TextPatternRangeEndpoint) (int32, error) {
	var retVal int32
	ret, _, _ := syscall.SyscallN(
		(*ITextRangeProviderVtbl)(unsafe.Pointer(v.vtbl)).CompareEndpoints,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(unsafe.Pointer(in2)),
		uintptr(in3),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return -1, HResult(ret)
	}
	return retVal, nil
}
func (v *ITextRangeProvider) ExpandToEnclosingUnit(in TextUnit) error {
	ret, _, _ := syscall.SyscallN(
		(*ITextRangeProviderVtbl)(unsafe.Pointer(v.vtbl)).ExpandToEnclosingUnit,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *ITextRangeProvider) FindAttribute(in TextArrtibuteId, in2 *VARIANT, in3 int32) (*ITextRangeProvider, error) {
	var retVal *ITextRangeProvider
	ret, _, _ := syscall.SyscallN(
		(*ITextRangeProviderVtbl)(unsafe.Pointer(v.vtbl)).FindAttribute,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(unsafe.Pointer(in2)),
		uintptr(in3),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *ITextRangeProvider) FindText(in string, in2 int32, in3 int32) (*ITextRangeProvider, error) {
	var retVal *ITextRangeProvider
	retVal2, err := syscall.UTF16PtrFromString(in)
	if err != nil {
		return nil, err
	}
	ret, _, _ := syscall.SyscallN(
		(*ITextRangeProviderVtbl)(unsafe.Pointer(v.vtbl)).FindText,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(retVal2)),
		uintptr(in2),
		uintptr(in3),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *ITextRangeProvider) GetAttributeValue(in TextArrtibuteId) (*VARIANT, error) {
	var retVal *VARIANT
	ret, _, _ := syscall.SyscallN(
		(*ITextRangeProviderVtbl)(unsafe.Pointer(v.vtbl)).GetAttributeValue,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *ITextRangeProvider) GetBoundingRectangles() (*TagSafeArray, error) {
	var retVal *TagSafeArray
	ret, _, _ := syscall.SyscallN(
		(*ITextRangeProviderVtbl)(unsafe.Pointer(v.vtbl)).GetBoundingRectangles,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *ITextRangeProvider) GetChildren() (*TagSafeArray, error) {
	var retVal *TagSafeArray
	ret, _, _ := syscall.SyscallN(
		(*ITextRangeProviderVtbl)(unsafe.Pointer(v.vtbl)).GetChildren,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *ITextRangeProvider) GetEnclosingElement() (*IRawElementProviderSimple, error) {
	var retVal *IRawElementProviderSimple
	ret, _, _ := syscall.SyscallN(
		(*ITextRangeProviderVtbl)(unsafe.Pointer(v.vtbl)).GetEnclosingElement,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *ITextRangeProvider) GetText(in int32) (string, error) {
	var bstr uintptr
	ret, _, _ := syscall.SyscallN(
		(*ITextRangeProviderVtbl)(unsafe.Pointer(v.vtbl)).GetText,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
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
func (v *ITextRangeProvider) Move(in TextUnit, in2 int32) (int32, error) {
	var retVal int32
	ret, _, _ := syscall.SyscallN(
		(*ITextRangeProviderVtbl)(unsafe.Pointer(v.vtbl)).Move,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(in2),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return -1, HResult(ret)
	}
	return retVal, nil
}
func (v *ITextRangeProvider) MoveEndpointByRange(in TextPatternRangeEndpoint, in2 *ITextRangeProvider, in3 TextPatternRangeEndpoint) error {
	ret, _, _ := syscall.SyscallN(
		(*ITextRangeProviderVtbl)(unsafe.Pointer(v.vtbl)).MoveEndpointByRange,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(unsafe.Pointer(in2)),
		uintptr(in3),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *ITextRangeProvider) MoveEndpointByUnit(in TextPatternRangeEndpoint, in2 TextUnit, in3 int32) (int32, error) {
	var retVal int32
	ret, _, _ := syscall.SyscallN(
		(*ITextRangeProviderVtbl)(unsafe.Pointer(v.vtbl)).MoveEndpointByUnit,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(in2),
		uintptr(in3),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return -1, HResult(ret)
	}
	return retVal, nil
}
func (v *ITextRangeProvider) RemoveFromSelection() error {
	ret, _, _ := syscall.SyscallN(
		(*ITextRangeProviderVtbl)(unsafe.Pointer(v.vtbl)).RemoveFromSelection,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *ITextRangeProvider) ScrollIntoView(in int32) error {
	ret, _, _ := syscall.SyscallN(
		(*ITextRangeProviderVtbl)(unsafe.Pointer(v.vtbl)).ScrollIntoView,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *ITextRangeProvider) Select() error {
	ret, _, _ := syscall.SyscallN(
		(*ITextRangeProviderVtbl)(unsafe.Pointer(v.vtbl)).Select,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

type ITextEditProvider struct {
	vtbl *IUnKnown
}
type ITextEditProviderVtbl struct {
	IUnKnownVtbl
	GetActiveComposition uintptr
	GetConversionTarget  uintptr
}

func newITextEditProvider(obj uintptr) *ITextEditProvider {
	return (*ITextEditProvider)(unsafe.Pointer(obj))
}
func NewITextEditProvider(obj uintptr) *ITextEditProvider {
	return newITextEditProvider(obj)
}
func (v *ITextEditProvider) GetActiveComposition() (*ITextRangeProvider, error) {
	var retVal *ITextRangeProvider
	ret, _, _ := syscall.SyscallN(
		(*ITextEditProviderVtbl)(unsafe.Pointer(v.vtbl)).GetActiveComposition,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *ITextEditProvider) GetConversionTarget() (*ITextRangeProvider, error) {
	var retVal *ITextRangeProvider
	ret, _, _ := syscall.SyscallN(
		(*ITextEditProviderVtbl)(unsafe.Pointer(v.vtbl)).GetConversionTarget,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
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

func newITextProvider(obj uintptr) *ITextProvider {
	return (*ITextProvider)(unsafe.Pointer(obj))
}
func NewITextProvider(obj uintptr) *ITextProvider {
	return newITextProvider(obj)
}
func (v *ITextProvider) Get_DocumentRange() *ITextRangeProvider {
	var retVal *ITextRangeProvider
	syscall.SyscallN(
		(*ITextProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_DocumentRange,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *ITextProvider) Get_SupportedTextSelection() *SupportedTextSelection {
	var retVal *SupportedTextSelection
	syscall.SyscallN(
		(*ITextProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_SupportedTextSelection,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *ITextProvider) GetSelection() (*TagSafeArray, error) {
	var retVal *TagSafeArray
	ret, _, _ := syscall.SyscallN(
		(*ITextProviderVtbl)(unsafe.Pointer(v.vtbl)).GetSelection,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *ITextProvider) GetVisibleRanges() (*TagSafeArray, error) {
	var retVal *TagSafeArray
	ret, _, _ := syscall.SyscallN(
		(*ITextProviderVtbl)(unsafe.Pointer(v.vtbl)).GetVisibleRanges,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *ITextProvider) RangeFromChild(in *IRawElementProviderSimple) (*ITextRangeProvider, error) {
	var retVal *ITextRangeProvider
	ret, _, _ := syscall.SyscallN(
		(*ITextProviderVtbl)(unsafe.Pointer(v.vtbl)).RangeFromChild,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *ITextProvider) RangeFromPoint(in *UiaPoint) (*ITextRangeProvider, error) {
	var retVal *ITextRangeProvider
	ret, _, _ := syscall.SyscallN(
		(*ITextProviderVtbl)(unsafe.Pointer(v.vtbl)).RangeFromPoint,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

type ITextProvider2 struct {
	vtbl *IUnKnown
}
type ITextProvider2Vtbl struct {
	// windows 8 start support
	IUnKnownVtbl
	GetCaretRange       uintptr
	RangeFromAnnotation uintptr
}

func newITextProvider2(obj uintptr) *ITextProvider2 {
	return (*ITextProvider2)(unsafe.Pointer(obj))
}
func NewITextProvider2(obj uintptr) *ITextProvider2 {
	return newITextProvider2(obj)
}
func (v *ITextProvider2) GetCaretRange() (int32, *ITextRangeProvider, error) {
	var retVal int32
	var retVal2 *ITextRangeProvider
	ret, _, _ := syscall.SyscallN(
		(*ITextProvider2Vtbl)(unsafe.Pointer(v.vtbl)).GetCaretRange,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
		uintptr(unsafe.Pointer(&retVal2)),
	)
	if ret != 0 {
		return -1, nil, HResult(ret)
	}
	return retVal, retVal2, nil
}
func (v *ITextProvider2) RangeFromAnnotation(in *IRawElementProviderSimple) (*ITextRangeProvider, error) {
	var retVal *ITextRangeProvider
	ret, _, _ := syscall.SyscallN(
		(*ITextProvider2Vtbl)(unsafe.Pointer(v.vtbl)).RangeFromAnnotation,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

type IToggleProvider struct {
	vtbl *IUnKnown
}
type IToggleProviderVtbl struct {
	IUnKnownVtbl
	Get_ToggleState uintptr
	Toggle          uintptr
}

func newIToggleProvider(obj uintptr) *IToggleProvider {
	return (*IToggleProvider)(unsafe.Pointer(obj))
}
func NewIToggleProvider(obj uintptr) *IToggleProvider {
	return newIToggleProvider(obj)
}
func (v *IToggleProvider) Get_ToggleState() *ToggleState {
	var retVal *ToggleState
	syscall.SyscallN(
		(*IToggleProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_ToggleState,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IToggleProvider) Toggle() error {
	ret, _, _ := syscall.SyscallN(
		(*IToggleProviderVtbl)(unsafe.Pointer(v.vtbl)).Toggle,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
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

func newITransformProvider(obj uintptr) *ITransformProvider {
	return (*ITransformProvider)(unsafe.Pointer(obj))
}
func NewITransformProvider(obj uintptr) *ITransformProvider {
	return newITransformProvider(obj)
}
func (v *ITransformProvider) Get_CanMove() int32 {
	var retVal int32
	syscall.SyscallN(
		(*ITransformProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_CanMove,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *ITransformProvider) Get_CanResize() int32 {
	var retVal int32
	syscall.SyscallN(
		(*ITransformProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_CanResize,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *ITransformProvider) Get_CanRotate() int32 {
	var retVal int32
	syscall.SyscallN(
		(*ITransformProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_CanRotate,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *ITransformProvider) Move(x, y float64) error {
	ret, _, _ := syscall.SyscallN(
		(*ITransformProviderVtbl)(unsafe.Pointer(v.vtbl)).Move,
		uintptr(unsafe.Pointer(v)),
		uintptr(x),
		uintptr(y),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *ITransformProvider) Resize(w, h float64) error {
	ret, _, _ := syscall.SyscallN(
		(*ITransformProviderVtbl)(unsafe.Pointer(v.vtbl)).Resize,
		uintptr(unsafe.Pointer(v)),
		uintptr(w),
		uintptr(h),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *ITransformProvider) Rotate(in float64) error {
	ret, _, _ := syscall.SyscallN(
		(*ITransformProviderVtbl)(unsafe.Pointer(v.vtbl)).Rotate,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

type ITransformProvider2 struct {
	vtbl *IUnKnown
}
type ITransformProvider2Vtbl struct {
	// windows 8 start support
	IUnKnownVtbl
	Get_CanZoom     uintptr
	Get_ZoomLevel   uintptr
	Get_ZoomMaximum uintptr
	Get_ZoomMinimum uintptr
	Zoom            uintptr
	ZoomByUnit      uintptr
}

func newITransformProvider2(obj uintptr) *ITransformProvider2 {
	return (*ITransformProvider2)(unsafe.Pointer(obj))
}
func NewITransformProvider2(obj uintptr) *ITransformProvider2 {
	return newITransformProvider2(obj)
}
func (v *ITransformProvider2) Get_CanZoom() int32 {
	var retVal int32
	syscall.SyscallN(
		(*ITransformProvider2Vtbl)(unsafe.Pointer(v.vtbl)).Get_CanZoom,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *ITransformProvider2) Get_ZoomLevel() float64 {
	var retVal float64
	syscall.SyscallN(
		(*ITransformProvider2Vtbl)(unsafe.Pointer(v.vtbl)).Get_ZoomLevel,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *ITransformProvider2) Get_ZoomMaximum() float64 {
	var retVal float64
	syscall.SyscallN(
		(*ITransformProvider2Vtbl)(unsafe.Pointer(v.vtbl)).Get_ZoomMaximum,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *ITransformProvider2) Get_ZoomMinimum() float64 {
	var retVal float64
	syscall.SyscallN(
		(*ITransformProvider2Vtbl)(unsafe.Pointer(v.vtbl)).Get_ZoomMinimum,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *ITransformProvider2) Zoom(in float64) error {
	ret, _, _ := syscall.SyscallN(
		(*ITransformProvider2Vtbl)(unsafe.Pointer(v.vtbl)).Zoom,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *ITransformProvider2) ZoomByUnit(in ZoomUnit) error {
	ret, _, _ := syscall.SyscallN(
		(*ITransformProvider2Vtbl)(unsafe.Pointer(v.vtbl)).ZoomByUnit,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
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

func newIValueProvider(obj uintptr) *IValueProvider {
	return (*IValueProvider)(unsafe.Pointer(obj))
}
func NewIValueProvider(obj uintptr) *IValueProvider {
	return newIValueProvider(obj)
}
func (v *IValueProvider) Get_IsReadOnly() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IValueProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_IsReadOnly,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IValueProvider) Get_Value() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IValueProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_Value,
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
func (v *IValueProvider) SetValue(in string) error {
	retVal, err := syscall.UTF16PtrFromString(in)
	if err != nil {
		return nil
	}
	ret, _, _ := syscall.SyscallN(
		(*IValueProviderVtbl)(unsafe.Pointer(v.vtbl)).SetValue,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(retVal)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

type IVirtualizedItemProvider struct {
	vtbl *IUnKnown
}
type IVirtualizedItemProviderVtbl struct {
	IUnKnownVtbl
	Realize uintptr
}

func newIVirtualizedItemProvider(obj uintptr) *IVirtualizedItemProvider {
	return (*IVirtualizedItemProvider)(unsafe.Pointer(obj))
}
func NewIVirtualizedItemProvider(obj uintptr) *IVirtualizedItemProvider {
	return newIVirtualizedItemProvider(obj)
}
func (v *IVirtualizedItemProvider) Realize() error {
	ret, _, _ := syscall.SyscallN(
		(*IVirtualizedItemProviderVtbl)(unsafe.Pointer(v.vtbl)).Realize,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
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

func newIWindowProvider(obj uintptr) *IWindowProvider {
	return (*IWindowProvider)(unsafe.Pointer(obj))
}
func NewIWindowProvider(obj uintptr) *IWindowProvider {
	return newIWindowProvider(obj)
}
func (v *IWindowProvider) Close() error {
	ret, _, _ := syscall.SyscallN(
		(*IWindowProviderVtbl)(unsafe.Pointer(v.vtbl)).Close,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *IWindowProvider) Get_CanMaximize() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IWindowProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_CanMaximize,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IWindowProvider) Get_CanMinimize() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IWindowProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_CanMinimize,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IWindowProvider) Get_IsModal() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IWindowProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_IsModal,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IWindowProvider) Get_IsTopmost() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IWindowProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_IsTopmost,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IWindowProvider) Get_WindowInteractionState() *WindowInteractionState {
	var retVal *WindowInteractionState
	syscall.SyscallN(
		(*IWindowProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_WindowInteractionState,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IWindowProvider) Get_WindowVisualState() *WindowVisualState {
	var retVal *WindowVisualState
	syscall.SyscallN(
		(*IWindowProviderVtbl)(unsafe.Pointer(v.vtbl)).Get_WindowVisualState,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IWindowProvider) SetVisualState(in WindowVisualState) error {
	ret, _, _ := syscall.SyscallN(
		(*IWindowProviderVtbl)(unsafe.Pointer(v.vtbl)).SetVisualState,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *IWindowProvider) WaitForInputIdle(in int32) (int32, error) {
	var retVal int32
	ret, _, _ := syscall.SyscallN(
		(*IWindowProviderVtbl)(unsafe.Pointer(v.vtbl)).WaitForInputIdle,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return -1, HResult(ret)
	}
	return retVal, nil
}
