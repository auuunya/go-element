package uiautomation

type IAccessible struct {
	vtbl *IDispatch
}

type IAccessibleVtbl struct {
	// https://learn.microsoft.com/zh-cn/windows/win32/api/oleacc/nn-oleacc-iaccessible
	IDispatchVtbl
	AccDoDefaultAction      uintptr
	AccHitTest              uintptr
	AccLocation             uintptr
	AccNavigate             uintptr
	AccSelect               uintptr
	Get_accChild            uintptr
	Get_accChildCount       uintptr
	Get_accDefaultAction    uintptr
	Get_accDescription      uintptr
	Get_accFocus            uintptr
	Get_accHelp             uintptr
	Get_accHelpTopic        uintptr
	Get_accKeyboardShortcut uintptr
	Get_accName             uintptr
	Get_accParent           uintptr
	Get_accRole             uintptr
	Get_accSelection        uintptr
	Get_accState            uintptr
	Get_accValue            uintptr
	Put_accName             uintptr
	Put_accValue            uintptr
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
