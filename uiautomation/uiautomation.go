package uiautomation

import (
	"syscall"
	"unsafe"
)

type IUIAutomation struct {
	vtbl *IUIAutomationVtbl
}
type IUIAutomationVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	CompareElements                           uintptr
	CompareRuntimeIds                         uintptr
	GetRootElement                            uintptr
	ElementFromHandle                         uintptr
	ElementFromPoint                          uintptr
	GetFocusedElement                         uintptr
	GetRootElementBuildCache                  uintptr
	ElementFromHandleBuildCache               uintptr
	ElementFromPointBuildCache                uintptr
	GetFocusedElementBuildCache               uintptr
	CreateTreeWalker                          uintptr
	Get_ControlViewWalker                     uintptr
	Get_ContentViewWalker                     uintptr
	Get_RawViewWalker                         uintptr
	Get_RawViewCondition                      uintptr
	Get_ControlViewCondition                  uintptr
	Get_ContentViewCondition                  uintptr
	CreateCacheRequest                        uintptr
	CreateTrueCondition                       uintptr
	CreateFalseCondition                      uintptr
	CreatePropertyCondition                   uintptr
	CreatePropertyConditionEx                 uintptr
	CreateAndCondition                        uintptr
	CreateAndConditionFromArray               uintptr
	CreateAndConditionFromNativeArray         uintptr
	CreateOrCondition                         uintptr
	CreateOrConditionFromArray                uintptr
	CreateOrConditionFromNativeArray          uintptr
	CreateNotCondition                        uintptr
	AddAutomationEventHandler                 uintptr
	RemoveAutomationEventHandler              uintptr
	AddPropertyChangedEventHandlerNativeArray uintptr
	AddPropertyChangedEventHandler            uintptr
	RemovePropertyChangedEventHandler         uintptr
	AddStructureChangedEventHandler           uintptr
	RemoveStructureChangedEventHandler        uintptr
	AddFocusChangedEventHandler               uintptr
	RemoveFocusChangedEventHandler            uintptr
	RemoveAllEventHandlers                    uintptr
	IntNativeArrayToSafeArray                 uintptr
	IntSafeArrayToNativeArray                 uintptr
	RectToVariant                             uintptr
	VariantToRect                             uintptr
	SafeArrayToRectNativeArray                uintptr
	CreateProxyFactoryEntry                   uintptr
	Get_ProxyFactoryMapping                   uintptr
	GetPropertyProgrammaticName               uintptr
	GetPatternProgrammaticName                uintptr
	PollForPotentialSupportedPatterns         uintptr
	PollForPotentialSupportedProperties       uintptr
	CheckNotSupported                         uintptr
	Get_ReservedNotSupportedValue             uintptr
	Get_ReservedMixedAttributeValue           uintptr
	ElementFromIAccessible                    uintptr
	ElementFromIAccessibleBuildCache          uintptr
}

type IUIAutomationElement struct {
	vtbl *IUIAutomationElementVtbl
}
type IUIAutomationElementVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	SetFocus                        uintptr
	GetRuntimeId                    uintptr
	FindFirst                       uintptr
	FindAll                         uintptr
	FindFirstBuildCache             uintptr
	FindAllBuildCache               uintptr
	BuildUpdatedCache               uintptr
	GetCurrentPropertyValue         uintptr
	GetCurrentPropertyValueEx       uintptr
	GetCachedPropertyValue          uintptr
	GetCachedPropertyValueEx        uintptr
	GetCurrentPatternAs             uintptr
	GetCachedPatternAs              uintptr
	GetCurrentPattern               uintptr
	GetCachedPattern                uintptr
	GetCachedParent                 uintptr
	GetCachedChildren               uintptr
	Get_CurrentProcessId            uintptr
	Get_CurrentControlType          uintptr
	Get_CurrentLocalizedControlType uintptr
	Get_CurrentName                 uintptr
	Get_CurrentAcceleratorKey       uintptr
	Get_CurrentAccessKey            uintptr
	Get_CurrentHasKeyboardFocus     uintptr
	Get_CurrentIsKeyboardFocusable  uintptr
	Get_CurrentIsEnabled            uintptr
	Get_CurrentAutomationId         uintptr
	Get_CurrentClassName            uintptr
	Get_CurrentHelpText             uintptr
	Get_CurrentCulture              uintptr
	Get_CurrentIsControlElement     uintptr
	Get_CurrentIsContentElement     uintptr
	Get_CurrentIsPassword           uintptr
	Get_CurrentNativeWindowHandle   uintptr
	Get_CurrentItemType             uintptr
	Get_CurrentIsOffscreen          uintptr
	Get_CurrentOrientation          uintptr
	Get_CurrentFrameworkId          uintptr
	Get_CurrentIsRequiredForForm    uintptr
	Get_CurrentItemStatus           uintptr
	Get_CurrentBoundingRectangle    uintptr
	Get_CurrentLabeledBy            uintptr
	Get_CurrentAriaRole             uintptr
	Get_CurrentAriaProperties       uintptr
	Get_CurrentIsDataValidForForm   uintptr
	Get_CurrentControllerFor        uintptr
	Get_CurrentDescribedBy          uintptr
	Get_CurrentFlowsTo              uintptr
	Get_CurrentProviderDescription  uintptr
	Get_CachedProcessId             uintptr
	Get_CachedControlType           uintptr
	Get_CachedLocalizedControlType  uintptr
	Get_CachedName                  uintptr
	Get_CachedAcceleratorKey        uintptr
	Get_CachedAccessKey             uintptr
	Get_CachedHasKeyboardFocus      uintptr
	Get_CachedIsKeyboardFocusable   uintptr
	Get_CachedIsEnabled             uintptr
	Get_CachedAutomationId          uintptr
	Get_CachedClassName             uintptr
	Get_CachedHelpText              uintptr
	Get_CachedCulture               uintptr
	Get_CachedIsControlElement      uintptr
	Get_CachedIsContentElement      uintptr
	Get_CachedIsPassword            uintptr
	Get_CachedNativeWindowHandle    uintptr
	Get_CachedItemType              uintptr
	Get_CachedIsOffscreen           uintptr
	Get_CachedOrientation           uintptr
	Get_CachedFrameworkId           uintptr
	Get_CachedIsRequiredForForm     uintptr
	Get_CachedItemStatus            uintptr
	Get_CachedBoundingRectangle     uintptr
	Get_CachedLabeledBy             uintptr
	Get_CachedAriaRole              uintptr
	Get_CachedAriaProperties        uintptr
	Get_CachedIsDataValidForForm    uintptr
	Get_CachedControllerFor         uintptr
	Get_CachedDescribedBy           uintptr
	Get_CachedFlowsTo               uintptr
	Get_CachedProviderDescription   uintptr
	GetClickablePoint               uintptr
}

type IUIAutomationTreeWalker struct {
	vtbl *IUIAutomationTreeWalkerVtbl
}

type IUIAutomationTreeWalkerVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	Get_Condition                       uintptr
	GetFirstChildElement                uintptr
	GetFirstChildElementBuildCache      uintptr
	GetLastChildElement                 uintptr
	GetLastChildElementBuildCache       uintptr
	GetNextSiblingElement               uintptr
	GetNextSiblingElementBuildCache     uintptr
	GetParentElement                    uintptr
	GetParentElementBuildCache          uintptr
	GetPreviousSiblingElement           uintptr
	GetPreviousSiblingElementBuildCache uintptr
	NormalizeElement                    uintptr
	NormalizeElementBuildCache          uintptr
}

type IUIAutomationElementArray struct {
	vtbl *IUIAutomationElementArrayVtbl
}
type IUIAutomationElementArrayVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	Get_Length uintptr
	GetElement uintptr
}

type TagPoint struct {
	X int32
	Y int32
}
type RECT struct {
	Left, Top, Right, Bottom int32
}
type SafeArray struct {
	CbElement uint32
	CDims     uint16
	CLocks    uint32
	FFeatures uint16
	PvData    uintptr
	Rgsabound [1]byte
}

func AddRef(v *IUIAutomation) uintptr {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).AddRef,
		1,
		uintptr(unsafe.Pointer(v)),
	)
	return ret
}
func Release(v *IUIAutomation) uintptr {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).Release,
		1,
		uintptr(unsafe.Pointer(v)),
	)
	return ret
}

func ElementFromHandle(ppv *IUIAutomation, hwnd uintptr) (root *IUIAutomationElement) {
	syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(ppv.vtbl)).ElementFromHandle,
		uintptr(unsafe.Pointer(ppv)),
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&root)),
	)
	return
}
