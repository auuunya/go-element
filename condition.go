package uiautomation

type IUIAutomationCondition struct {
	vtbl *IUnKnown
}

type IUIAutomationConditionVtbl struct {
	IUnKnownVtbl
}

type IUIAutomationAndCondition struct {
	vtbl *IUIAutomationCondition
}

type IUIAutomationAndConditionVtbl struct {
	IUIAutomationConditionVtbl

	Get_ChildCount           uintptr
	GetChildren              uintptr
	GetChildrenAsNativeArray uintptr
}

type IUIAutomationBoolCondition struct {
	vtbl *IUIAutomationCondition
}

type IUIAutomationBoolConditionVtbl struct {
	IUIAutomationConditionVtbl
	Get_BooleanValue uintptr
}

type IUIAutomationNotCondition struct {
	vtbl *IUIAutomationCondition
}

type IUIAutomationNotConditionVtbl struct {
	IUIAutomationConditionVtbl
	GetChild uintptr
}

type IUIAutomationPropertyCondition struct {
	vtbl *IUIAutomationCondition
}

type IUIAutomationPropertyConditionVtbl struct {
	IUIAutomationConditionVtbl

	Get_PropertyConditionFlags uintptr
	Get_PropertyId             uintptr
	Get_PropertyValue          uintptr
}
