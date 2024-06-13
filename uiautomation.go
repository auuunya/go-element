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

func newIUIAutomation(unk *IUnKnown) *IUIAutomation {
	return (*IUIAutomation)(unsafe.Pointer(unk))
}
func NewIUIAutomation(unk *IUnKnown) *IUIAutomation {
	return newIUIAutomation(unk)
}

type EventHandler struct {
	EventId      UIA_EventId
	Element      *IUIAutomationElement
	Scope        TreeScope
	CacheRequest *IUIAutomationCacheRequest
	Handler      *IUIAutomationEventHandler
}

func (v *IUIAutomation) AddAutomationEventHandler(opt *EventHandler) error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).AddAutomationEventHandler,
		uintptr(unsafe.Pointer(v)),
		uintptr(opt.EventId),
		uintptr(unsafe.Pointer(opt.Element)),
		uintptr(opt.Scope),
		uintptr(unsafe.Pointer(opt.CacheRequest)),
		uintptr(unsafe.Pointer(opt.Handler)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

func (v *IUIAutomation) AddFocusChangedEventHandler(in *IUIAutomationCacheRequest, in2 *IUIAutomationFocusChangedEventHandler) error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).AddFocusChangedEventHandler,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(in2)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

type ChangeEventHandler struct {
	Element       *IUIAutomationElement
	Scope         TreeScope
	CacheRequest  *IUIAutomationCacheRequest
	Handler       *IUIAutomationPropertyChangedEventHandler
	PropertyArray *TagSafeArray
}

func (v *IUIAutomation) AddPropertyChangedEventHandler(opt *ChangeEventHandler) error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).AddPropertyChangedEventHandler,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(opt.Element)),
		uintptr(opt.Scope),
		uintptr(unsafe.Pointer(opt.CacheRequest)),
		uintptr(unsafe.Pointer(opt.Handler)),
		uintptr(unsafe.Pointer(opt.PropertyArray)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

type ChangeEventHandlerNativeArray struct {
	Element       *IUIAutomationElement
	Scope         TreeScope
	CacheRequest  *IUIAutomationCacheRequest
	Handler       *IUIAutomationPropertyChangedEventHandler
	PropertyArray *TagSafeArray
	PropertyCount int32
}

func (v *IUIAutomation) AddPropertyChangedEventHandlerNativeArray(opt *ChangeEventHandlerNativeArray) error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).AddPropertyChangedEventHandlerNativeArray,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(opt.Element)),
		uintptr(opt.Scope),
		uintptr(unsafe.Pointer(opt.CacheRequest)),
		uintptr(unsafe.Pointer(opt.Handler)),
		uintptr(unsafe.Pointer(opt.PropertyArray)),
		uintptr(opt.PropertyCount),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

type StructureChangedEventHandler struct {
	Element      *IUIAutomationElement
	Scope        TreeScope
	CacheRequest *IUIAutomationCacheRequest
	Handler      *IUIAutomationPropertyChangedEventHandler
}

func (v *IUIAutomation) AddStructureChangedEventHandler(opt *StructureChangedEventHandler) error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).AddStructureChangedEventHandler,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(opt.Element)),
		uintptr(opt.Scope),
		uintptr(unsafe.Pointer(opt.CacheRequest)),
		uintptr(unsafe.Pointer(opt.Handler)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *IUIAutomation) CheckNotSupported(in *VARIANT) (int32, error) {
	var retVal int32
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).CheckNotSupported,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return -1, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) CompareElements(in, in2 *IUIAutomationElement) (int32, error) {
	var retVal int32
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).CompareElements,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(in2)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return -1, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) CompareRuntimeIds(in, in2 *TagSafeArray) (int32, error) {
	var retVal int32
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).CompareRuntimeIds,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(in2)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return -1, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) CreateAndCondition(in, in2 *IUIAutomationCondition) (*IUIAutomationCondition, error) {
	var retVal *IUIAutomationCondition
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).CreateAndCondition,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(in2)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) CreateAndConditionFromArray(in *TagSafeArray) (*IUIAutomationCondition, error) {
	var retVal *IUIAutomationCondition
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).CreateAndConditionFromArray,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) CreateAndConditionFromNativeArray(in *IUIAutomationCondition, in2 int32) (*IUIAutomationCondition, error) {
	var retVal *IUIAutomationCondition
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).CreateAndConditionFromNativeArray,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(in2),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) CreateCacheRequest() (*IUIAutomationCacheRequest, error) {
	var retVal *IUIAutomationCacheRequest
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).CreateCacheRequest,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) CreateFalseCondition() (*IUIAutomationCondition, error) {
	var retVal *IUIAutomationCondition
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).CreateFalseCondition,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) CreateNotCondition(in *IUIAutomationCondition) (*IUIAutomationCondition, error) {
	var retVal *IUIAutomationCondition
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).CreateNotCondition,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) CreateOrCondition(in, in2 *IUIAutomationCondition) (*IUIAutomationCondition, error) {
	var retVal *IUIAutomationCondition
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).CreateOrCondition,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(in2)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) CreateOrConditionFromArray(in *TagSafeArray) (*IUIAutomationCondition, error) {
	var retVal *IUIAutomationCondition
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).CreateOrConditionFromArray,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) CreateOrConditionFromNativeArray(in *IUIAutomationCondition, in2 int32) (*IUIAutomationCondition, error) {
	var retVal *IUIAutomationCondition
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).CreateOrConditionFromNativeArray,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(in2),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) CreatePropertyCondition(in *PropertyId, in2 *VARIANT) (*IUIAutomationCondition, error) {
	var retVal *IUIAutomationCondition
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).CreatePropertyCondition,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(in2)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) CreatePropertyConditionEx(in *PropertyId, in2 *VARIANT, in3 PropertyConditionFlags) (*IUIAutomationCondition, error) {
	var retVal *IUIAutomationCondition
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).CreatePropertyConditionEx,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(in2)),
		uintptr(in3),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) CreateProxyFactoryEntry(in *IUIAutomationProxyFactory) (*IUIAutomationProxyFactoryEntry, error) {
	var retVal *IUIAutomationProxyFactoryEntry
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).CreateProxyFactoryEntry,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) CreateTreeWalker(in *IUIAutomationCondition) (*IUIAutomationTreeWalker, error) {
	var retVal *IUIAutomationTreeWalker
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).CreateTreeWalker,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
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
func (v *IUIAutomation) ElementFromHandle(in uintptr) (*IUIAutomationElement, error) {
	var retVal *IUIAutomationElement
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).ElementFromHandle,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) ElementFromHandleBuildCache(in uintptr, in2 *IUIAutomationCacheRequest) (*IUIAutomationElement, error) {
	var retVal *IUIAutomationElement
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).ElementFromHandleBuildCache,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(unsafe.Pointer(in2)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) ElementFromIAccessible(in *IAccessible, in2 int32) (*IUIAutomationElement, error) {
	var retVal *IUIAutomationElement
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).ElementFromIAccessible,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(in2),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) ElementFromIAccessibleBuildCache(in *IAccessible, in2 int32, in3 *IUIAutomationCacheRequest) (*IUIAutomationElement, error) {
	var retVal *IUIAutomationElement
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).ElementFromIAccessibleBuildCache,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(in2),
		uintptr(unsafe.Pointer(in3)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) ElementFromPoint(in *TagPoint) (*IUIAutomationElement, error) {
	var retVal *IUIAutomationElement
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).ElementFromPoint,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) ElementFromPointBuildCache(in *TagPoint, in2 *IUIAutomationCacheRequest) (*IUIAutomationElement, error) {
	var retVal *IUIAutomationElement
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).ElementFromPointBuildCache,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(in2)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) Get_ContentViewCondition() *IUIAutomationCondition {
	var retVal *IUIAutomationCondition
	syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).Get_ContentViewCondition,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomation) Get_ContentViewWalker() *IUIAutomationTreeWalker {
	var retVal *IUIAutomationTreeWalker
	syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).Get_ContentViewWalker,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomation) Get_ControlViewCondition() *IUIAutomationCondition {
	var retVal *IUIAutomationCondition
	syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).Get_ControlViewCondition,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomation) Get_ControlViewWalker() *IUIAutomationTreeWalker {
	var retVal *IUIAutomationTreeWalker
	syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).Get_ControlViewWalker,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomation) Get_ProxyFactoryMapping() *IUIAutomationProxyFactoryMapping {
	var retVal *IUIAutomationProxyFactoryMapping
	syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).Get_ProxyFactoryMapping,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomation) Get_RawViewCondition() *IUIAutomationCondition {
	var retVal *IUIAutomationCondition
	syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).Get_RawViewCondition,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomation) Get_RawViewWalker() *IUIAutomationTreeWalker {
	var retVal *IUIAutomationTreeWalker
	syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).Get_RawViewWalker,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomation) Get_ReservedMixedAttributeValue() *IUnKnown {
	var retVal *IUnKnown
	syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).Get_ReservedMixedAttributeValue,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomation) Get_ReservedNotSupportedValue() *IUnKnown {
	var retVal *IUnKnown
	syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).Get_ReservedNotSupportedValue,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomation) GetFocusedElement() (*IUIAutomationElement, error) {
	var retVal *IUIAutomationElement
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).GetFocusedElement,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) GetFocusedElementBuildCache(in *IUIAutomationCacheRequest) (*IUIAutomationElement, error) {
	var retVal *IUIAutomationElement
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).GetFocusedElementBuildCache,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) GetPatternProgrammaticName(in PatternId) (string, error) {
	var bstr uintptr
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).GetPatternProgrammaticName,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if ret != 0 {
		return "", HResult(ret)
	}

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomation) GetPropertyProgrammaticName(in PropertyId) (string, error) {
	var bstr uintptr
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).GetPropertyProgrammaticName,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if ret != 0 {
		return "", HResult(ret)
	}

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomation) GetRootElement() (*IUIAutomationElement, error) {
	var retVal *IUIAutomationElement
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).GetRootElement,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) GetRootElementBuildCache(in *IUIAutomationCacheRequest) (*IUIAutomationElement, error) {
	var retVal *IUIAutomationElement
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).GetRootElementBuildCache,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) IntNativeArrayToSafeArray(in, in2 int32) (*TagSafeArray, error) {
	var retVal *TagSafeArray
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).IntNativeArrayToSafeArray,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(in2),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) IntSafeArrayToNativeArray(in *TagSafeArray) (int32, int32, error) {
	var retVal int32
	var retVal2 int32
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).IntSafeArrayToNativeArray,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
		uintptr(unsafe.Pointer(&retVal2)),
	)
	if ret != 0 {
		return -1, -1, HResult(ret)
	}
	return retVal, retVal2, nil
}
func (v *IUIAutomation) PollForPotentialSupportedPatterns(in *IUIAutomationElement) (*TagSafeArray, *TagSafeArray, error) {
	var retVal *TagSafeArray
	var retVal2 *TagSafeArray
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).PollForPotentialSupportedPatterns,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
		uintptr(unsafe.Pointer(&retVal2)),
	)
	if ret != 0 {
		return nil, nil, HResult(ret)
	}
	return retVal, retVal2, nil
}
func (v *IUIAutomation) PollForPotentialSupportedProperties(in *IUIAutomationElement) (*TagSafeArray, *TagSafeArray, error) {
	var retVal *TagSafeArray
	var retVal2 *TagSafeArray
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).PollForPotentialSupportedProperties,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
		uintptr(unsafe.Pointer(&retVal2)),
	)
	if ret != 0 {
		return nil, nil, HResult(ret)
	}
	return retVal, retVal2, nil
}
func (v *IUIAutomation) RectToVariant(in *TagRect) (*VARIANT, error) {
	var retVal *VARIANT
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).RectToVariant,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomation) RemoveAllEventHandlers() error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).RemoveAllEventHandlers,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *IUIAutomation) RemoveAutomationEventHandler(in UIA_EventId, in2 *IUIAutomationElement, in3 *IUIAutomationEventHandler) error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).RemoveAutomationEventHandler,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(unsafe.Pointer(in2)),
		uintptr(unsafe.Pointer(in3)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *IUIAutomation) RemoveFocusChangedEventHandler(in *IUIAutomationFocusChangedEventHandler) error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).RemoveFocusChangedEventHandler,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *IUIAutomation) RemovePropertyChangedEventHandler(in *IUIAutomationElement, in2 *IUIAutomationPropertyChangedEventHandler) error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).RemovePropertyChangedEventHandler,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(in2)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *IUIAutomation) RemoveStructureChangedEventHandler(in *IUIAutomationElement, in2 *IUIAutomationStructureChangedEventHandler) error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).RemoveStructureChangedEventHandler,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(in2)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *IUIAutomation) SafeArrayToRectNativeArray(in *TagSafeArray) (*TagRect, int32, error) {
	var retVal *TagRect
	var retVal2 int32
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).SafeArrayToRectNativeArray,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
		uintptr(unsafe.Pointer(&retVal2)),
	)
	if ret != 0 {
		return nil, -1, HResult(ret)
	}
	return retVal, retVal2, nil
}
func (v *IUIAutomation) VariantToRect(in *VARIANT) (*TagRect, error) {
	var retVal *TagRect
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(v.vtbl)).VariantToRect,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
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

func newIUIAutomationElement(unk *IUnKnown) *IUIAutomationElement {
	return (*IUIAutomationElement)(unsafe.Pointer(unk))
}
func NewIUIAutomationElement(unk *IUnKnown) *IUIAutomationElement {
	return newIUIAutomationElement(unk)
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
func (v *IUIAutomationElement) GetCurrentPatternAs(patternId PatternId, riid *syscall.GUID) (unsafe.Pointer, error) {
	var retVal unsafe.Pointer
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).GetCurrentPatternAs,
		uintptr(unsafe.Pointer(v)),
		uintptr(patternId),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomationElement) GetCurrentPropertyValue(id PropertyId) (*VARIANT, error) {
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
func (v *IUIAutomationElement) GetCurrentPropertyValueEx(id PropertyId, defaultVal int32) (*VARIANT, error) {
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

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomationElement) Get_CurrentAccessKey() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentAccessKey,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomationElement) Get_CurrentAriaProperties() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentAriaProperties,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomationElement) Get_CurrentAriaRole() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentAriaRole,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomationElement) Get_CurrentAutomationId() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentAutomationId,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomationElement) Get_CurrentBoundingRectangle() *TagRect {
	var retVal TagRect
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentBoundingRectangle,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return &retVal
}

func (v *IUIAutomationElement) Get_CurrentClassName() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentClassName,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
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

func (v *IUIAutomationElement) Get_CurrentControlType() ControlTypeId {
	var retVal ControlTypeId
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

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
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

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
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

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}

func (v *IUIAutomationElement) Get_CurrentItemType() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentItemType,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
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

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}

func (v *IUIAutomationElement) Get_CurrentName() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(v.vtbl)).Get_CurrentName,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
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
func (v *IUIAutomationElement) Get_CurrentOrientation() OrientationType {
	var retVal OrientationType
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

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
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

func newIUIAutomationElementArray(unk *IUnKnown) *IUIAutomationElementArray {
	return (*IUIAutomationElementArray)(unsafe.Pointer(unk))
}
func NewIUIAutomationElementArray(unk *IUnKnown) *IUIAutomationElementArray {
	return newIUIAutomationElementArray(unk)
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

type IUIAutomationCacheRequest struct {
	vtbl *IUnKnown
}
type IUIAutomationCacheRequestVtbl struct {
	IUnKnownVtbl
	AddPattern                uintptr
	AddProperty               uintptr
	Clone                     uintptr
	Get_AutomationElementMode uintptr
	Get_TreeFilter            uintptr
	Get_TreeScope             uintptr
	Put_AutomationElementMode uintptr
	Put_TreeFilter            uintptr
	Put_TreeScope             uintptr
}

type IUIAutomationEventHandler struct {
	vtbl *IUnKnown
}
type IUIAutomationEventHandlerVtbl struct {
	IUnKnownVtbl
	HandleAutomationEvent uintptr
}

func newIUIAutomationEventHandler(unk *IUnKnown) *IUIAutomationEventHandler {
	return (*IUIAutomationEventHandler)(unsafe.Pointer(unk))
}
func NewIUIAutomationEventHandler(unk *IUnKnown) *IUIAutomationEventHandler {
	return newIUIAutomationEventHandler(unk)
}
func (v *IUIAutomationEventHandler) HandleAutomationEvent(in *IUIAutomationElement, in2 UIA_EventId) error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationEventHandlerVtbl)(unsafe.Pointer(v.vtbl)).HandleAutomationEvent,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(in2),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

type IUIAutomationFocusChangedEventHandler struct {
	vtbl *IUnKnown
}
type IUIAutomationFocusChangedEventHandlerVtbl struct {
	IUnKnownVtbl
	HandleFocusChangedEvent uintptr
}

func newIUIAutomationFocusChangedEventHandler(unk *IUnKnown) *IUIAutomationFocusChangedEventHandler {
	return (*IUIAutomationFocusChangedEventHandler)(unsafe.Pointer(unk))
}
func NewIUIAutomationFocusChangedEventHandler(unk *IUnKnown) *IUIAutomationFocusChangedEventHandler {
	return newIUIAutomationFocusChangedEventHandler(unk)
}
func (v *IUIAutomationFocusChangedEventHandler) HandleFocusChangedEvent(in *IUIAutomationElement) error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationFocusChangedEventHandlerVtbl)(unsafe.Pointer(v.vtbl)).HandleFocusChangedEvent,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

type IUIAutomationProxyFactory struct {
	vtbl *IUnKnown
}
type IUIAutomationProxyFactoryVtbl struct {
	IUnKnownVtbl
	CreateProvider     uintptr
	Get_ProxyFactoryId uintptr
}

func newIUIAutomationProxyFactory(unk *IUnKnown) *IUIAutomationProxyFactory {
	return (*IUIAutomationProxyFactory)(unsafe.Pointer(unk))
}
func NewIUIAutomationProxyFactory(unk *IUnKnown) *IUIAutomationProxyFactory {
	return newIUIAutomationProxyFactory(unk)
}
func (v *IUIAutomationProxyFactory) CreateProvider(in uintptr, in2, in3 int32) (*IRawElementProviderSimple, error) {
	var retVal *IRawElementProviderSimple
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationProxyFactoryVtbl)(unsafe.Pointer(v.vtbl)).CreateProvider,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(in2),
		uintptr(in3),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomationProxyFactory) Get_ProxyFactoryId() (string, error) {
	var bstr uintptr
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationProxyFactoryVtbl)(unsafe.Pointer(v.vtbl)).Get_ProxyFactoryId,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if ret != 0 {
		return "", HResult(ret)
	}

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}

type IUIAutomationProxyFactoryEntry struct {
	vtbl *IUnKnown
}
type IUIAutomationProxyFactoryEntryVtbl struct {
	IUnKnownVtbl
	Get_AllowSubstringMatch        uintptr
	Get_CanCheckBaseClass          uintptr
	Get_ClassName                  uintptr
	Get_ImageName                  uintptr
	Get_NeedsAdviseEvents          uintptr
	Get_ProxyFactory               uintptr
	GetWinEventsForAutomationEvent uintptr
	Put_AllowSubstringMatch        uintptr
	Put_CanCheckBaseClass          uintptr
	Put_ClassName                  uintptr
	Put_ImageName                  uintptr
	Put_NeedsAdviseEvents          uintptr
	SetWinEventsForAutomationEvent uintptr
}

func newIUIAutomationProxyFactoryEntry(unk *IUnKnown) *IUIAutomationProxyFactoryEntry {
	return (*IUIAutomationProxyFactoryEntry)(unsafe.Pointer(unk))
}
func NewIUIAutomationProxyFactoryEntry(unk *IUnKnown) *IUIAutomationProxyFactoryEntry {
	return newIUIAutomationProxyFactoryEntry(unk)
}
func (v *IUIAutomationProxyFactoryEntry) Get_AllowSubstringMatch() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationProxyFactoryEntryVtbl)(unsafe.Pointer(v.vtbl)).Get_AllowSubstringMatch,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomationProxyFactoryEntry) Get_CanCheckBaseClass() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationProxyFactoryEntryVtbl)(unsafe.Pointer(v.vtbl)).Get_CanCheckBaseClass,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomationProxyFactoryEntry) Get_ClassName() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationProxyFactoryEntryVtbl)(unsafe.Pointer(v.vtbl)).Get_ClassName,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomationProxyFactoryEntry) Get_ImageName() (string, error) {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationProxyFactoryEntryVtbl)(unsafe.Pointer(v.vtbl)).Get_ImageName,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *IUIAutomationProxyFactoryEntry) Get_NeedsAdviseEvents() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationProxyFactoryEntryVtbl)(unsafe.Pointer(v.vtbl)).Get_NeedsAdviseEvents,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomationProxyFactoryEntry) Get_ProxyFactory() *IUIAutomationProxyFactory {
	var retVal *IUIAutomationProxyFactory
	syscall.SyscallN(
		(*IUIAutomationProxyFactoryEntryVtbl)(unsafe.Pointer(v.vtbl)).Get_ProxyFactory,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomationProxyFactoryEntry) GetWinEventsForAutomationEvent(in UIA_EventId, in2 PropertyId) (*TagSafeArray, error) {
	var retVal *TagSafeArray
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationProxyFactoryEntryVtbl)(unsafe.Pointer(v.vtbl)).GetWinEventsForAutomationEvent,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(in2),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomationProxyFactoryEntry) Put_AllowSubstringMatch() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationProxyFactoryEntryVtbl)(unsafe.Pointer(v.vtbl)).Put_AllowSubstringMatch,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomationProxyFactoryEntry) Put_CanCheckBaseClass() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationProxyFactoryEntryVtbl)(unsafe.Pointer(v.vtbl)).Put_CanCheckBaseClass,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomationProxyFactoryEntry) Put_ClassName() string {
	var retVal []uint16
	syscall.SyscallN(
		(*IUIAutomationProxyFactoryEntryVtbl)(unsafe.Pointer(v.vtbl)).Put_ClassName,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	retValStr := syscall.UTF16ToString(retVal)
	return retValStr
}
func (v *IUIAutomationProxyFactoryEntry) Put_ImageName() string {
	var retVal []uint16
	syscall.SyscallN(
		(*IUIAutomationProxyFactoryEntryVtbl)(unsafe.Pointer(v.vtbl)).Put_ImageName,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	retValStr := syscall.UTF16ToString(retVal)
	return retValStr
}
func (v *IUIAutomationProxyFactoryEntry) Put_NeedsAdviseEvents() int32 {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationProxyFactoryEntryVtbl)(unsafe.Pointer(v.vtbl)).Put_NeedsAdviseEvents,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomationProxyFactoryEntry) SetWinEventsForAutomationEvent(in UIA_EventId, in2 PropertyId, in3 *TagSafeArray) error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationProxyFactoryEntryVtbl)(unsafe.Pointer(v.vtbl)).Put_NeedsAdviseEvents,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(in2),
		uintptr(unsafe.Pointer(in3)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

type IUIAutomationProxyFactoryMapping struct {
	vtbl *IUnKnown
}
type IUIAutomationProxyFactoryMappingVtbl struct {
	IUnKnownVtbl
	ClearTable          uintptr
	Get_Count           uintptr
	GetEntry            uintptr
	GetTable            uintptr
	InsertEntries       uintptr
	InsertEntry         uintptr
	RemoveEntry         uintptr
	RestoreDefaultTable uintptr
	SetTable            uintptr
}

func newIUIAutomationProxyFactoryMapping(unk *IUnKnown) *IUIAutomationProxyFactoryMapping {
	return (*IUIAutomationProxyFactoryMapping)(unsafe.Pointer(unk))
}
func NewIUIAutomationProxyFactoryMapping(unk *IUnKnown) *IUIAutomationProxyFactoryMapping {
	return newIUIAutomationProxyFactoryMapping(unk)
}
func (v *IUIAutomationProxyFactoryMapping) ClearTable() error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationProxyFactoryMappingVtbl)(unsafe.Pointer(v.vtbl)).ClearTable,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *IUIAutomationProxyFactoryMapping) Get_Count() uint32 {
	var retVal uint32
	syscall.SyscallN(
		(*IUIAutomationProxyFactoryMappingVtbl)(unsafe.Pointer(v.vtbl)).Get_Count,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	return retVal
}
func (v *IUIAutomationProxyFactoryMapping) GetEntry(in uint32) (*IUIAutomationProxyFactoryEntry, error) {
	var retVal *IUIAutomationProxyFactoryEntry
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationProxyFactoryMappingVtbl)(unsafe.Pointer(v.vtbl)).GetEntry,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomationProxyFactoryMapping) GetTable() (*TagSafeArray, error) {
	var retVal *TagSafeArray
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationProxyFactoryMappingVtbl)(unsafe.Pointer(v.vtbl)).GetTable,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IUIAutomationProxyFactoryMapping) InsertEntries(in uint32, in2 *TagSafeArray) error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationProxyFactoryMappingVtbl)(unsafe.Pointer(v.vtbl)).InsertEntries,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(unsafe.Pointer(in2)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *IUIAutomationProxyFactoryMapping) InsertEntry(in uint32, in2 *IUIAutomationProxyFactoryEntry) error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationProxyFactoryMappingVtbl)(unsafe.Pointer(v.vtbl)).InsertEntry,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(unsafe.Pointer(in2)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *IUIAutomationProxyFactoryMapping) RemoveEntry(in uint32) error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationProxyFactoryMappingVtbl)(unsafe.Pointer(v.vtbl)).RemoveEntry,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *IUIAutomationProxyFactoryMapping) RestoreDefaultTable() error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationProxyFactoryMappingVtbl)(unsafe.Pointer(v.vtbl)).RestoreDefaultTable,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *IUIAutomationProxyFactoryMapping) SetTable(in *TagSafeArray) error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationProxyFactoryMappingVtbl)(unsafe.Pointer(v.vtbl)).SetTable,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

type IUIAutomationStructureChangedEventHandler struct {
	vtbl *IUnKnown
}
type IUIAutomationStructureChangedEventHandlerVtbl struct {
	IUnKnownVtbl
	HandleStructureChangedEvent uintptr
}

func newIUIAutomationStructureChangedEventHandler(unk *IUnKnown) *IUIAutomationStructureChangedEventHandler {
	return (*IUIAutomationStructureChangedEventHandler)(unsafe.Pointer(unk))
}
func NewIUIAutomationStructureChangedEventHandler(unk *IUnKnown) *IUIAutomationStructureChangedEventHandler {
	return newIUIAutomationStructureChangedEventHandler(unk)
}

func (v *IUIAutomationStructureChangedEventHandler) HandleStructureChangedEvent(in *IUIAutomationElement, in2 StructureChangeType, in3 *TagSafeArray) error {
	ret, _, _ := syscall.SyscallN(
		(*IUIAutomationStructureChangedEventHandlerVtbl)(unsafe.Pointer(v.vtbl)).HandleStructureChangedEvent,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(in2),
		uintptr(unsafe.Pointer(in3)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
