package uiautomation

// https://learn.microsoft.com/zh-cn/windows/win32/api/oleacc/nn-oleacc-iaccessible
type IAccessible struct {
	vtbl *IDispatch
}

type IAccessibleVtbl struct {
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
