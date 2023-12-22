package uiautomation

type TreeScope uintptr

var (
	TreeScope_None        TreeScope = 0x0
	TreeScope_Element     TreeScope = 0x1
	TreeScope_Children    TreeScope = 0x2
	TreeScope_Descendants TreeScope = 0x4
	TreeScope_Parent      TreeScope = 0x8
	TreeScope_Ancestors   TreeScope = 0x10
	TreeScope_Subtree     TreeScope = TreeScope_Element | TreeScope_Children | TreeScope_Descendants
)

type IUIAutomationCondition struct {
	vtbl *IUIAutomationConditionVtbl
}

type IUIAutomationConditionVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

type IUIAutomationAndCondition struct {
	// *IUIAutomationCondition
	vtbl *IUIAutomationAndConditionVtbl
}

type IUIAutomationAndConditionVtbl struct {
	// *IUIAutomationConditionVtbl
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	Get_ChildCount           uintptr
	GetChildren              uintptr
	GetChildrenAsNativeArray uintptr
}

type IUIAutomationBoolCondition struct {
	*IUIAutomationCondition
}

type IUIAutomationBoolConditionVtbl struct {
	*IUIAutomationConditionVtbl
	Get_BooleanValue uintptr
}

type IUIAutomationNotCondition struct {
	*IUIAutomationCondition
}

type IUIAutomationNotConditionVtbl struct {
	*IUIAutomationConditionVtbl
	GetChild uintptr
}

type IUIAutomationPropertyCondition struct {
	*IUIAutomationCondition
}

type IUIAutomationPropertyConditionVtbl struct {
	*IUIAutomationConditionVtbl

	Get_PropertyConditionFlags uintptr
	Get_PropertyId             uintptr
	Get_PropertyValue          uintptr
}
