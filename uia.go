package uiautomation

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

type UIA_EventId int

const (
	UIA_ActiveTextPositionChangedEventId                 UIA_EventId = 20036
	UIA_AsyncContentLoadedEventId                        UIA_EventId = 20006
	UIA_AutomationFocusChangedEventId                    UIA_EventId = 20005
	UIA_AutomationPropertyChangedEventId                 UIA_EventId = 20004
	UIA_ChangesEventId                                   UIA_EventId = 20034
	UIA_Drag_DragCancelEventId                           UIA_EventId = 20027
	UIA_Drag_DragCompleteEventId                         UIA_EventId = 20028
	UIA_Drag_DragStartEventId                            UIA_EventId = 20026
	UIA_DropTarget_DragEnterEventId                      UIA_EventId = 20029
	UIA_DropTarget_DragLeaveEventId                      UIA_EventId = 20030
	UIA_DropTarget_DroppedEventId                        UIA_EventId = 20031
	UIA_HostedFragmentRootsInvalidatedEventId            UIA_EventId = 20025
	UIA_InputDiscardedEventId                            UIA_EventId = 20022
	UIA_InputReachedOtherElementEventId                  UIA_EventId = 20021
	UIA_InputReachedTargetEventId                        UIA_EventId = 20020
	UIA_Invoke_InvokedEventId                            UIA_EventId = 20009
	UIA_LayoutInvalidatedEventId                         UIA_EventId = 20008
	UIA_LiveRegionChangedEventId                         UIA_EventId = 20024
	UIA_MenuClosedEventId                                UIA_EventId = 20007
	UIA_MenuModeEndEventId                               UIA_EventId = 20019
	UIA_MenuModeStartEventId                             UIA_EventId = 20018
	UIA_MenuOpenedEventId                                UIA_EventId = 20003
	UIA_NotificationEventId                              UIA_EventId = 20035
	UIA_Selection_InvalidatedEventId                     UIA_EventId = 20013
	UIA_SelectionItem_ElementAddedToSelectionEventId     UIA_EventId = 20010
	UIA_SelectionItem_ElementRemovedFromSelectionEventId UIA_EventId = 20011
	UIA_SelectionItem_ElementSelectedEventId             UIA_EventId = 20012
	UIA_StructureChangedEventId                          UIA_EventId = 20002
	UIA_SystemAlertEventId                               UIA_EventId = 20023
	UIA_Text_TextChangedEventId                          UIA_EventId = 20015
	UIA_Text_TextSelectionChangedEventId                 UIA_EventId = 20014
	UIA_TextEdit_ConversionTargetChangedEventId          UIA_EventId = 20033
	UIA_TextEdit_TextChangedEventId                      UIA_EventId = 20032
	UIA_ToolTipClosedEventId                             UIA_EventId = 20001
	UIA_ToolTipOpenedEventId                             UIA_EventId = 20000
	UIA_Window_WindowClosedEventId                       UIA_EventId = 20017
	UIA_Window_WindowOpenedEventId                       UIA_EventId = 20016
)
