package uiautomation

type TxtHitResult int

const (
	TXTHITRESULT_NOHIT TxtHitResult = iota
	TXTHITRESULT_TRANSPARENT
	TXTHITRESULT_CLOSE
	TXTHITRESULT_HIT
)

type StructureChangeType int

const (
	StructureChangeType_ChildAdded StructureChangeType = iota
	StructureChangeType_ChildRemoved
	StructureChangeType_ChildrenInvalidated
	StructureChangeType_ChildrenBulkAdded
	StructureChangeType_ChildrenBulkRemoved
	StructureChangeType_ChildrenReordered
)

type WindowVisualState int

const (
	// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationcore/ne-uiautomationcore-windowvisualstate
	WindowVisualState_Normal WindowVisualState = iota
	WindowVisualState_Maximized
	WindowVisualState_Minimized
)

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

type SupportedTextSelection int

const (
	// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationcore/ne-uiautomationcore-supportedtextselection
	SupportedTextSelection_None SupportedTextSelection = iota
	SupportedTextSelection_Single
	SupportedTextSelection_Multiple
)

type ScrollAmount int

const (
	// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationcore/ne-uiautomationcore-scrollamount
	ScrollAmount_LargeDecrement ScrollAmount = iota
	ScrollAmount_SmallDecrement
	ScrollAmount_NoAmount
	ScrollAmount_LargeIncrement
	ScrollAmount_SmallIncrement
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

type OrientationType int

const (
	OrientationType_None OrientationType = iota
	OrientationType_Horizontal
	OrientationType_Vertical
)

type TagFuncKind int

const (
	FUNC_VIRTUAL TagFuncKind = iota
	FUNC_PUREVIRTUAL
	FUNC_NONVIRTUAL
	FUNC_STATIC
	FUNC_DISPATCH
)

type TagCallConv int

const (
	CC_FASTCALL TagCallConv = iota
	CC_CDECL
	CC_MSCPASCAL
	CC_PASCAL
	CC_MACPASCAL
	CC_STDCALL
	CC_FPFASTCALL
	CC_SYSCALL
	CC_MPWCDECL
	CC_MPWPASCAL
	CC_MAX
)
