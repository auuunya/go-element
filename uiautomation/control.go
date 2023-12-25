package uiautomation

import (
	"syscall"
	"unsafe"
)

// https://learn.microsoft.com/zh-cn/windows/win32/winauto/uiauto-controlpattern-ids
type PatternId int

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

type PROPERTYID uintptr

// windows 8 start support
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

// CustomNavigation: windows 10 start support
// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationclient/nn-uiautomationclient-iuiautomationelement4
// int automationelement4
type iDockProvider struct {
	vtbl *IUnKnown
}

type iDockProviderVtbl struct {
	IUnKnownVtbl
	Get_DockPosition uintptr
	SetDockPosition  uintptr
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

type iDropTargetProvider struct {
	vtbl *IUnKnown
}
type iDropTargetProviderVtbl struct {
	IUnKnownVtbl
	Get_DropTargetEffect  uintptr
	Get_DropTargetEffects uintptr
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

func NewIExpandCollapse(obj uintptr) *IExpandCollapseProvider {
	return (*IExpandCollapseProvider)(unsafe.Pointer(obj))
}
func (ie *IExpandCollapseProvider) Collapse() {
	collapse(ie)
}
func (ie *IExpandCollapseProvider) Expand() {
	expand(ie)
}
func collapse(ie *IExpandCollapseProvider) uintptr {
	ret, _, _ := syscall.SyscallN(
		(*IExpandCollapseProviderVtbl)(unsafe.Pointer(ie.vtbl)).Collapse,
		uintptr(unsafe.Pointer(ie)),
	)
	return ret
}
func expand(ie *IExpandCollapseProvider) uintptr {
	ret, _, _ := syscall.SyscallN(
		(*IExpandCollapseProviderVtbl)(unsafe.Pointer(ie.vtbl)).Expand,
		uintptr(unsafe.Pointer(ie)),
	)
	return ret
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
type IGridProvider struct {
	vtbl *IUnKnown
}
type IGridProviderVtbl struct {
	IUnKnownVtbl
	Get_ColumnCount uintptr
	Get_RowCount    uintptr
	GetItem         uintptr
}
type IInvokeProvider struct {
	vtbl *IUnKnown
}
type IInvokeProviderVtbl struct {
	IUnKnownVtbl
	Invoke uintptr
}
type IItemContainerProvider struct {
	vtbl *IUnKnown
}
type IItemContainerProviderVtbl struct {
	IUnKnownVtbl
	FindItemByProperty uintptr
}
type iLegacyIAccessibleProvider struct {
	vtbl *IUnKnown
}
type iLegacyIAccessibleProviderVtbl struct {
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

// windows 8 start support
type IObjectModelProvider struct {
	vtbl *IUnKnown
}
type IObjectModelProviderVtbl struct {
	IUnKnownVtbl
	GetUnderlyingObjectModel uintptr
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
	Get_ExtendedProperties uintptr
	Get_FillColor          uintptr
	Get_FillPatternColor   uintptr
	Get_FillPatternStyle   uintptr
	Get_Shape              uintptr
	Get_StyleId            uintptr
	Get_StyleName          uintptr
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
