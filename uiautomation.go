package uiautomation

import (
	"syscall"
	"unsafe"
)

type IUIAutomation struct {
	vtbl *IUnKnown
}
type IUIAutomationVtbl struct {
	IUnKnownVtbl

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

func newIUIAutomation(obj uintptr) *IUIAutomation {
	return (*IUIAutomation)(unsafe.Pointer(obj))
}
func NewIUIAutomation(obj uintptr) *IUIAutomation {
	return newIUIAutomation(obj)
}

func ElementFromHandle(v *IUIAutomation, hwnd uintptr) (*IUIAutomationElement, error) {
	var retVal *IUIAutomationElement
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).ElementFromHandle,
		uintptr(unsafe.Pointer(v)),
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

func (v *IUIAutomation) CreateTrueCondition() *IUIAutomationCondition {
	var retVal *IUIAutomationCondition
	syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).CreateTrueCondition,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

func CreateTrueCondition(v *IUIAutomation) *IUIAutomationCondition {
	var retVal *IUIAutomationCondition
	syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).CreateTrueCondition,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

type IUIAutomationElement struct {
	vtbl *IUnKnown
}
type IUIAutomationElementVtbl struct {
	IUnKnownVtbl

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

func newIUIAutomationElement(obj uintptr) *IUIAutomationElement {
	return (*IUIAutomationElement)(unsafe.Pointer(obj))
}
func NewIUIAutomationElement(obj uintptr) *IUIAutomationElement {
	return newIUIAutomationElement(obj)
}

func (v *IUIAutomationElement) GetClickablePoint() (*TagPoint, int32, error) {
	var retVal *TagPoint
	var retVal2 int32
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).GetClickablePoint,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
		uintptr(unsafe.Pointer(&retVal2)),
	)
	if ret != 0 {
		return nil, -1, HResult(ret)
	}
	return retVal, retVal2, nil
}
func (v *IUIAutomationElement) GetCurrentPattern(patternId PatternId) (*IUnKnown, error) {
	var retVal *IUnKnown
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).GetCurrentPattern,
		uintptr(unsafe.Pointer(v)),
		uintptr(patternId),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomationElement) GetCurrentPatternAs(patternId PatternId, riid syscall.GUID) (unsafe.Pointer, error) {
	var retVal unsafe.Pointer
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).GetCurrentPatternAs,
		uintptr(unsafe.Pointer(v)),
		uintptr(patternId),
		uintptr(unsafe.Pointer(&riid)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomationElement) GetCurrentPropertyValue(id PROPERTYID) (*VARIANT, error) {
	var retVal *VARIANT
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).GetCurrentPropertyValue,
		uintptr(unsafe.Pointer(v)),
		uintptr(id),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomationElement) GetCurrentPropertyValueEx(id PROPERTYID, defaultVal int32) (*VARIANT, error) {
	var retVal *VARIANT
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).GetCurrentPropertyValue,
		uintptr(unsafe.Pointer(v)),
		uintptr(id),
		uintptr(defaultVal),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomationElement) GetRuntimeId() (*TagSafeArray, error) {
	var retVal *TagSafeArray
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).GetRuntimeId,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomationElement) SetFocus() error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).SetFocus,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *IUIAutomationElement) Get_CurrentAcceleratorKey() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentAcceleratorKey,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return "", ErrorBstrPointerNil
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(
		uintptr(bstr),
	)
	return retVal, nil
}
func (v *IUIAutomationElement) Get_CurrentAccessKey() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentAccessKey,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return "", ErrorBstrPointerNil
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(
		uintptr(bstr),
	)
	return retVal, nil
}
func (v *IUIAutomationElement) Get_CurrentAriaProperties() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentAriaProperties,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return "", ErrorBstrPointerNil
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(
		uintptr(bstr),
	)
	return retVal, nil
}
func (v *IUIAutomationElement) Get_CurrentAriaRole() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentAriaRole,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return "", ErrorBstrPointerNil
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(
		uintptr(bstr),
	)
	return retVal, nil
}
func (v *IUIAutomationElement) Get_CurrentAutomationId() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentAutomationId,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return "", ErrorBstrPointerNil
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(
		uintptr(bstr),
	)
	return retVal, nil
}
func (v *IUIAutomationElement) Get_CurrentBoundingRectangle() *TagRect {
	var retVal *TagRect
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentBoundingRectangle,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

func (v *IUIAutomationElement) Get_CurrentClassName() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentClassName,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return "", ErrorBstrPointerNil
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(
		uintptr(bstr),
	)
	return retVal, nil
}

func (v *IUIAutomationElement) Get_CurrentControllerFor() *IUIAutomationElementArray {
	var retVal *IUIAutomationElementArray
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentControllerFor,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

func (v *IUIAutomationElement) Get_CurrentControlType() syscall.GUID {
	var retVal syscall.GUID
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentControlType,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

func (v *IUIAutomationElement) Get_CurrentCulture() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentCulture,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

func (v *IUIAutomationElement) Get_CurrentDescribedBy() *IUIAutomationElementArray {
	var retVal *IUIAutomationElementArray
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentDescribedBy,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

func (v *IUIAutomationElement) Get_CurrentFlowsTo() *IUIAutomationElementArray {
	var retVal *IUIAutomationElementArray
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentFlowsTo,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

func (v *IUIAutomationElement) Get_CurrentFrameworkId() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentFrameworkId,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return "", ErrorBstrPointerNil
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(
		uintptr(bstr),
	)
	return retVal, nil
}

func (v *IUIAutomationElement) Get_CurrentHasKeyboardFocus() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentHasKeyboardFocus,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

func (v *IUIAutomationElement) Get_CurrentHelpText() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentHelpText,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return "", ErrorBstrPointerNil
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(
		uintptr(bstr),
	)
	return retVal, nil
}
func (v *IUIAutomationElement) Get_CurrentIsControlElement() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentIsControlElement,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

func (v *IUIAutomationElement) Get_CurrentIsContentElement() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentIsContentElement,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

func (v *IUIAutomationElement) Get_CurrentIsDataValidForForm() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentIsDataValidForForm,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

func (v *IUIAutomationElement) Get_CurrentIsEnabled() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentIsEnabled,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

func (v *IUIAutomationElement) Get_CurrentIsKeyboardFocusable() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentIsKeyboardFocusable,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

func (v *IUIAutomationElement) Get_CurrentIsOffscreen() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentIsOffscreen,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

func (v *IUIAutomationElement) Get_CurrentIsPassword() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentIsPassword,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

func (v *IUIAutomationElement) Get_CurrentIsRequiredForForm() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentIsRequiredForForm,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

func (v *IUIAutomationElement) Get_CurrentItemStatus() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentItemStatus,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return "", ErrorBstrPointerNil
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(uintptr(bstr))
	return retVal, nil
}

func (v *IUIAutomationElement) Get_CurrentItemType() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentItemType,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return "", ErrorBstrPointerNil
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(uintptr(bstr))
	return retVal, nil
}

func (v *IUIAutomationElement) Get_CurrentLabeledBy() *IUIAutomationElement {
	var retVal *IUIAutomationElement
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentLabeledBy,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

func (v *IUIAutomationElement) Get_CurrentLocalizedControlType() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentLocalizedControlType,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return "", ErrorBstrPointerNil
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(uintptr(bstr))
	return retVal, nil
}

func (v *IUIAutomationElement) Get_CurrentName() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentName,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return "", ErrorBstrPointerNil
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(uintptr(bstr))
	return retVal, nil
}
func (v *IUIAutomationElement) Get_CurrentNativeWindowHandle() uintptr {
	var retVal uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentNativeWindowHandle,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomationElement) Get_CurrentOrientation() *OrientationType {
	var retVal *OrientationType
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentOrientation,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomationElement) Get_CurrentProcessId() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentProcessId,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomationElement) Get_CurrentProviderDescription() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentProviderDescription,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr == 0 {
		return "", ErrorBstrPointerNil
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(uintptr(bstr))
	return retVal, nil
}

func (v *IUIAutomationElement) FindAll(condition *IUIAutomationCondition) (*IUIAutomationElementArray, error) {
	var retVal *IUIAutomationElementArray
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).FindAll,
		uintptr(unsafe.Pointer(v)),
		uintptr(TreeScope_Children),
		uintptr(unsafe.Pointer(condition)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

func FindAll(v *IUIAutomationElement, condition *IUIAutomationCondition) (*IUIAutomationElementArray, error) {
	var retVal *IUIAutomationElementArray
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).FindAll,
		uintptr(unsafe.Pointer(v)),
		uintptr(TreeScope_Children),
		uintptr(unsafe.Pointer(condition)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

type IUIAutomationTreeWalker struct {
	vtbl *IUnKnown
}

type IUIAutomationTreeWalkerVtbl struct {
	IUnKnownVtbl

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

func newIUIAutomationElementArray(obj uintptr) *IUIAutomationElementArray {
	return (*IUIAutomationElementArray)(unsafe.Pointer(obj))
}
func NewIUIAutomationElementArray(obj uintptr) *IUIAutomationElementArray {
	return newIUIAutomationElementArray(obj)
}

func (v *IUIAutomationElementArray) Get_Length() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementArrayVtbl)(unsafe.Pointer(v.vtbl)).Get_Length,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

func (v *IUIAutomationElementArray) GetElement(i int32) (*IUIAutomationElement, error) {
	var retVal *IUIAutomationElement
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationElementArrayVtbl)(unsafe.Pointer(v.vtbl)).GetElement,
		uintptr(unsafe.Pointer(v)),
		uintptr(i),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

func Get_Length(v *IUIAutomationElementArray) int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementArrayVtbl)(unsafe.Pointer(v.vtbl)).Get_Length,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}

func GetElement(v *IUIAutomationElementArray, i int32) (*IUIAutomationElement, error) {
	var retVal *IUIAutomationElement
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationElementArrayVtbl)(unsafe.Pointer(v.vtbl)).GetElement,
		uintptr(unsafe.Pointer(v)),
		uintptr(i),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
