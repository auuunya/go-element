package uiautomation

type OrientationType int

const (
	OrientationType_None OrientationType = iota
	OrientationType_Horizontal
	OrientationType_Vertical
)

type TagPoint struct {
	X int32
	Y int32
}

type TagRect struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type TagSafeArrayBound struct {
	// https://learn.microsoft.com/zh-cn/windows/win32/api/oaidl/ns-oaidl-safearraybound
	CElements uint32
	LLbound   int32
}
type TagSafeArray struct {
	// https://learn.microsoft.com/zh-cn/windows/win32/api/oaidl/ns-oaidl-safearray
	CbElement uint32
	CDims     uint16
	CLocks    uint32
	FFeatures uint16
	PvData    uintptr
	Rgsabound []TagSafeArrayBound
}

type TreeScope int

var (
	// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationclient/ne-uiautomationclient-treescope
	TreeScope_None        TreeScope = 0x0
	TreeScope_Element     TreeScope = 0x1
	TreeScope_Children    TreeScope = 0x2
	TreeScope_Descendants TreeScope = 0x4
	TreeScope_Parent      TreeScope = 0x8
	TreeScope_Ancestors   TreeScope = 0x10
	TreeScope_Subtree     TreeScope = TreeScope_Element | TreeScope_Children | TreeScope_Descendants
)

type PROPERTYID uintptr

type ScrollAmount int

const (
	// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationcore/ne-uiautomationcore-scrollamount
	ScrollAmount_LargeDecrement ScrollAmount = iota
	ScrollAmount_SmallDecrement
	ScrollAmount_NoAmount
	ScrollAmount_LargeIncrement
	ScrollAmount_SmallIncrement
)

type SynchronizedInputType int

const (
	// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationcore/ne-uiautomationcore-synchronizedinputtype
	SynchronizedInputType_KeyUp          SynchronizedInputType = 0x1
	SynchronizedInputType_KeyDown        SynchronizedInputType = 0x2
	SynchronizedInputType_LeftMouseUp    SynchronizedInputType = 0x4
	SynchronizedInputType_LeftMouseDown  SynchronizedInputType = 0x8
	SynchronizedInputType_RightMouseUp   SynchronizedInputType = 0x10
	SynchronizedInputType_RightMouseDown SynchronizedInputType = 0x20
)

type RowOrColumnMajor int

const (
	// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationcore/ne-uiautomationcore-roworcolumnmajor
	RowOrColumnMajor_RowMajor RowOrColumnMajor = iota
	RowOrColumnMajor_ColumnMajor
	RowOrColumnMajor_Indeterminate
)

type TextPatternRangeEndpoint int

const (
	// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationcore/ne-uiautomationcore-textpatternrangeendpoint
	TextPatternRangeEndpoint_Start TextPatternRangeEndpoint = iota
	TextPatternRangeEndpoint_End
)

type TextUnit int

const (
	// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationcore/ne-uiautomationcore-textunit
	TextUnit_Character TextUnit = iota
	TextUnit_Format
	TextUnit_Word
	TextUnit_Line
	TextUnit_Paragraph
	TextUnit_Page
	TextUnit_Document
)

type TextArrtibuteId int

const (
// https://learn.microsoft.com/zh-cn/windows/win32/winauto/uiauto-textattribute-ids
)

type SupportedTextSelection int

const (
	// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationcore/ne-uiautomationcore-supportedtextselection
	SupportedTextSelection_None SupportedTextSelection = iota
	SupportedTextSelection_Single
	SupportedTextSelection_Multiple
)

type UiaPoint struct {
	// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationcore/ns-uiautomationcore-uiapoint
	X float64
	Y float64
}

type ToggleState int

const (
	// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationcore/ne-uiautomationcore-togglestate
	ToggleState_Off ToggleState = iota
	ToggleState_On
	ToggleState_Indeterminate
)

type ZoomUnit int

const (
	// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationcore/ne-uiautomationcore-zoomunit
	ZoomUnit_NoAmount ZoomUnit = iota
	ZoomUnit_LargeDecrement
	ZoomUnit_SmallDecrement
	ZoomUnit_LargeIncrement
	ZoomUnit_SmallIncrement
)

type WindowInteractionState int

const (
	// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationcore/ne-uiautomationcore-windowinteractionstate
	WindowInteractionState_Running WindowInteractionState = iota
	WindowInteractionState_Closing
	WindowInteractionState_ReadyForUserInteraction
	WindowInteractionState_BlockedByModalWindow
	WindowInteractionState_NotResponding
)

type WindowVisualState int

const (
	// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationcore/ne-uiautomationcore-windowvisualstate
	WindowVisualState_Normal WindowVisualState = iota
	WindowVisualState_Maximized
	WindowVisualState_Minimized
)
