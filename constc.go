package uiautomation

type CLSCTX int

var (
	// https://learn.microsoft.com/zh-cn/windows/win32/api/wtypesbase/ne-wtypesbase-clsctx
	CLSCTX_INPROC_SERVER                   CLSCTX = 0x1
	CLSCTX_INPROC_HANDLER                  CLSCTX = 0x2
	CLSCTX_LOCAL_SERVER                    CLSCTX = 0x4
	CLSCTX_INPROC_SERVER16                 CLSCTX = 0x8
	CLSCTX_REMOTE_SERVER                   CLSCTX = 0x10
	CLSCTX_INPROC_HANDLER16                CLSCTX = 0x20
	CLSCTX_RESERVED1                       CLSCTX = 0x40
	CLSCTX_RESERVED2                       CLSCTX = 0x80
	CLSCTX_RESERVED3                       CLSCTX = 0x100
	CLSCTX_RESERVED4                       CLSCTX = 0x200
	CLSCTX_NO_CODE_DOWNLOAD                CLSCTX = 0x400
	CLSCTX_RESERVED5                       CLSCTX = 0x800
	CLSCTX_NO_CUSTOM_MARSHAL               CLSCTX = 0x1000
	CLSCTX_ENABLE_CODE_DOWNLOAD            CLSCTX = 0x2000
	CLSCTX_NO_FAILURE_LOG                  CLSCTX = 0x4000
	CLSCTX_DISABLE_AAA                     CLSCTX = 0x8000
	CLSCTX_ENABLE_AAA                      CLSCTX = 0x10000
	CLSCTX_FROM_DEFAULT_CONTEXT            CLSCTX = 0x20000
	CLSCTX_ACTIVATE_X86_SERVER             CLSCTX = 0x40000
	_CLSCTX_ACTIVATE_32_BIT_SERVER         CLSCTX = 0
	CLSCTX_ACTIVATE_64_BIT_SERVER          CLSCTX = 0x80000
	CLSCTX_ENABLE_CLOAKING                 CLSCTX = 0x100000
	CLSCTX_APPCONTAINER                    CLSCTX = 0x400000
	CLSCTX_ACTIVATE_AAA_AS_IU              CLSCTX = 0x800000
	CLSCTX_RESERVED6                       CLSCTX = 0x1000000
	CLSCTX_ACTIVATE_ARM32_SERVER           CLSCTX = 0x2000000
	_CLSCTX_ALLOW_LOWER_TRUST_REGISTRATION CLSCTX = 0
	CLSCTX_PS_DLL                          CLSCTX = 0x80000000
)

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

type PropertyId int

const (
	// https://learn.microsoft.com/zh-cn/windows/win32/winauto/uiauto-entry-propids
	UIA_AcceleratorKeyPropertyId           PropertyId = 30006
	UIA_AccessKeyPropertyId                PropertyId = 30007
	UIA_AnnotationObjectsPropertyId        PropertyId = 30156
	UIA_AnnotationTypesPropertyId          PropertyId = 30155
	UIA_AriaPropertiesPropertyId           PropertyId = 30102
	UIA_AriaRolePropertyId                 PropertyId = 30101
	UIA_AutomationIdPropertyId             PropertyId = 30011
	UIA_BoundingRectanglePropertyId        PropertyId = 30001
	UIA_CenterPointPropertyId              PropertyId = 30165
	UIA_ClassNamePropertyId                PropertyId = 30012
	UIA_ClickablePointPropertyId           PropertyId = 30014
	UIA_ControllerForPropertyId            PropertyId = 30104
	UIA_ControlTypePropertyId              PropertyId = 30003
	UIA_CulturePropertyId                  PropertyId = 30015
	UIA_DescribedByPropertyId              PropertyId = 30105
	UIA_FillColorPropertyId                PropertyId = 30160
	UIA_FillTypePropertyId                 PropertyId = 30162
	UIA_FlowsFromPropertyId                PropertyId = 30148
	UIA_FlowsToPropertyId                  PropertyId = 30106
	UIA_FrameworkIdPropertyId              PropertyId = 30024
	UIA_FullDescriptionPropertyId          PropertyId = 30159
	UIA_HasKeyboardFocusPropertyId         PropertyId = 30008
	UIA_HeadingLevelPropertyId             PropertyId = 30173
	UIA_HelpTextPropertyId                 PropertyId = 30013
	UIA_IsContentElementPropertyId         PropertyId = 30017
	UIA_IsControlElementPropertyId         PropertyId = 30016
	UIA_IsDataValidForFormPropertyId       PropertyId = 30103
	UIA_IsDialogPropertyId                 PropertyId = 30174
	UIA_IsEnabledPropertyId                PropertyId = 30010
	UIA_IsKeyboardFocusablePropertyId      PropertyId = 30009
	UIA_IsOffscreenPropertyId              PropertyId = 30022
	UIA_IsPasswordPropertyId               PropertyId = 30019
	UIA_IsPeripheralPropertyId             PropertyId = 30150
	UIA_IsRequiredForFormPropertyId        PropertyId = 30025
	UIA_ItemStatusPropertyId               PropertyId = 30026
	UIA_ItemTypePropertyId                 PropertyId = 300021
	UIA_LabeledByPropertyId                PropertyId = 30018
	UIA_LandmarkTypePropertyId             PropertyId = 30157
	UIA_LevelPropertyId                    PropertyId = 30154
	UIA_LiveSettingPropertyId              PropertyId = 30135
	UIA_LocalizedControlTypePropertyId     PropertyId = 30004
	UIA_LocalizedLandmarkTypePropertyId    PropertyId = 30158
	UIA_NamePropertyId                     PropertyId = 30005
	UIA_NativeWindowHandlePropertyId       PropertyId = 30020
	UIA_OptimizeForVisualContentPropertyId PropertyId = 30111
	UIA_OrientationPropertyId              PropertyId = 300023
	UIA_OutlineColorPropertyId             PropertyId = 30161
	UIA_OutlineThicknessPropertyId         PropertyId = 30164
	UIA_PositionInSetPropertyId            PropertyId = 30152
	UIA_ProcessIdPropertyId                PropertyId = 30002
	UIA_ProviderDescriptionPropertyId      PropertyId = 30107
	UIA_RotationPropertyId                 PropertyId = 30166
	UIA_RuntimeIdPropertyId                PropertyId = 30000
	UIA_SizePropertyId                     PropertyId = 30167
	UIA_SizeOfSetPropertyId                PropertyId = 30153
	UIA_VisualEffectsPropertyId            PropertyId = 30163
)

type PatternId int

const (
	// https://learn.microsoft.com/zh-cn/windows/win32/winauto/uiauto-controlpattern-ids
	// https://learn.microsoft.com/zh-cn/windows/win32/winauto/uiauto-controlsupport
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

type ControlTypeId int

const (
	// https://learn.microsoft.com/zh-cn/windows/win32/winauto/uiauto-controlpatternmapping
	// https://learn.microsoft.com/zh-cn/windows/win32/winauto/uiauto-supportinguiautocontroltypes
	UIA_AppBarControlTypeId       ControlTypeId = 50040
	UIA_ButtonControlTypeId       ControlTypeId = 50000
	UIA_CalendarControlTypeId     ControlTypeId = 50001
	UIA_CheckBoxControlTypeId     ControlTypeId = 50002
	UIA_ComboBoxControlTypeId     ControlTypeId = 50003
	UIA_CustomControlTypeId       ControlTypeId = 50025
	UIA_DataGridControlTypeId     ControlTypeId = 50028
	UIA_DataItemControlTypeId     ControlTypeId = 50029
	UIA_DocumentControlTypeId     ControlTypeId = 50030
	UIA_EditControlTypeId         ControlTypeId = 50004
	UIA_GroupControlTypeId        ControlTypeId = 50026
	UIA_HeaderControlTypeId       ControlTypeId = 50034
	UIA_HeaderItemControlTypeId   ControlTypeId = 50035
	UIA_HyperlinkControlTypeId    ControlTypeId = 50005
	UIA_ImageControlTypeId        ControlTypeId = 50006
	UIA_ListControlTypeId         ControlTypeId = 50008
	UIA_ListItemControlTypeId     ControlTypeId = 50007
	UIA_MenuBarControlTypeId      ControlTypeId = 50010
	UIA_MenuControlTypeId         ControlTypeId = 50009
	UIA_MenuItemControlTypeId     ControlTypeId = 50011
	UIA_PaneControlTypeId         ControlTypeId = 50033
	UIA_ProgressBarControlTypeId  ControlTypeId = 50012
	UIA_RadioButtonControlTypeId  ControlTypeId = 50013
	UIA_ScrollBarControlTypeId    ControlTypeId = 50014
	UIA_SemanticZoomControlTypeId ControlTypeId = 50039
	UIA_SeparatorControlTypeId    ControlTypeId = 50038
	UIA_SliderControlTypeId       ControlTypeId = 50015
	UIA_SpinnerControlTypeId      ControlTypeId = 50016
	UIA_SplitButtonControlTypeId  ControlTypeId = 50031
	UIA_StatusBarControlTypeId    ControlTypeId = 50017
	UIA_TabControlTypeId          ControlTypeId = 50018
	UIA_TabItemControlTypeId      ControlTypeId = 50019
	UIA_TableControlTypeId        ControlTypeId = 50036
	UIA_TextControlTypeId         ControlTypeId = 50020
	UIA_ThumbControlTypeId        ControlTypeId = 50027
	UIA_TitleBarControlTypeId     ControlTypeId = 50037
	UIA_ToolBarControlTypeId      ControlTypeId = 50021
	UIA_ToolTipControlTypeId      ControlTypeId = 50022
	UIA_TreeControlTypeId         ControlTypeId = 50023
	UIA_TreeItemControlTypeId     ControlTypeId = 50024
	UIA_WindowControlTypeId       ControlTypeId = 50032
)

type TagDvTargetDevice struct {
	TdSize             uint32
	TdDriverNameOffset uint16
	TdDeviceNameOffset uint16
	TdPortNameOffset   uint16
	TdExtDevmodeOffset uint16
	TdData             [1]byte
}
type TagDvAspect int

const (
	DVASPECT_CONTENT   TagDvAspect = 1
	DVASPECT_THUMBNAIL TagDvAspect = 2
	DVASPECT_ICON      TagDvAspect = 4
	DVASPECT_DOCPRINT  TagDvAspect = 8
)

type TxtHitResult int

const (
	TXTHITRESULT_NOHIT TxtHitResult = iota
	TXTHITRESULT_TRANSPARENT
	TXTHITRESULT_CLOSE
	TXTHITRESULT_HIT
)

type UiaRect struct {
	Left   float64
	Top    float64
	Width  float64
	Height float64
}
