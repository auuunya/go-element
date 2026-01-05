package uiautomation

import (
	"fmt"
	"reflect"
	"strings"
	"syscall"
	"unsafe"
)

var (
	oleAut             = syscall.NewLazyDLL("OleAut32.dll")
	procSysFreeString  = oleAut.NewProc("SysFreeString")
	procSysAllocString = oleAut.NewProc("SysAllocString")
)

type Element struct {
	UIAutoElement               *IUIAutomationElement
	CurrentAcceleratorKey       string
	CurrentAccessKey            string
	CurrentAriaProperties       string
	CurrentAriaRole             string
	CurrentAutomationId         string
	CurrentBoundingRectangle    *TagRect
	CurrentClassName            string
	CurrentControllerFor        *IUIAutomationElementArray
	CurrentControlType          ControlTypeId
	CurrentCulture              int32
	CurrentDescribedBy          *IUIAutomationElementArray
	CurrentFlowsTo              *IUIAutomationElementArray
	CurrentFrameworkId          string
	CurrentHasKeyboardFocus     int32
	CurrentHelpText             string
	CurrentIsContentElement     int32
	CurrentIsControlElement     int32
	CurrentIsDataValidForForm   int32
	CurrentIsEnabled            int32
	CurrentIsKeyboardFocusable  int32
	CurrentIsOffscreen          int32
	CurrentIsPassword           int32
	CurrentIsRequiredForForm    int32
	CurrentItemStatus           string
	CurrentItemType             string
	CurrentLabeledBy            *IUIAutomationElement
	CurrentLocalizedControlType string
	CurrentName                 string
	CurrentNativeWindowHandle   uintptr
	CurrentOrientation          OrientationType
	CurrentProcessId            int32
	CurrentProviderDescription  string

	SupportedPatterns []PatternId `json:"supported_patterns,omitempty"`
	Child             []*Element  `json:"child,omitempty"`
}

func NewElement(uiaumation *IUIAutomationElement) *Element {
	return &Element{
		UIAutoElement: uiaumation,
	}
}
func (e *Element) FormatString() string {
	if e == nil {
		return ""
	}
	elemType := reflect.TypeOf(e)
	if elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
	}
	elemValue := reflect.ValueOf(e)
	if elemValue.Kind() == reflect.Ptr {
		elemValue = elemValue.Elem()
	}
	buf := ""
	for i := 0; i < elemType.NumField(); i++ {
		field := elemType.Field(i)
		fieldName := field.Name
		fieldValue := elemValue.Field(i)
		if fieldValue.Kind() == reflect.Ptr && !fieldValue.IsNil() {
			if fieldValue.CanAddr() {
				v := fieldValue.Addr().Interface()
				buf += fmt.Sprintf("[%s]:[%v],", fieldName, v)
			}
		} else {
			buf += fmt.Sprintf("[%s]:[%v],", fieldName, fieldValue.Interface())
		}
	}
	return buf
}
func (e *Element) SetUIAutomation(uiaumation *IUIAutomationElement) {
	e.UIAutoElement = uiaumation
}

func (e *Element) AcceleratorKey() error {
	val, err := e.UIAutoElement.Get_CurrentAcceleratorKey()
	e.CurrentAcceleratorKey = val
	return err
}

func (e *Element) AccessKey() error {
	val, err := e.UIAutoElement.Get_CurrentAccessKey()
	e.CurrentAccessKey = val
	return err
}

func (e *Element) AriaProperties() error {
	val, err := e.UIAutoElement.Get_CurrentAriaProperties()
	e.CurrentAriaProperties = val
	return err
}

func (e *Element) AriaRole() error {
	val, err := e.UIAutoElement.Get_CurrentAriaRole()
	e.CurrentAriaRole = val
	return err
}

func (e *Element) AutomationId() error {
	val, err := e.UIAutoElement.Get_CurrentAutomationId()
	e.CurrentAutomationId = val
	return err
}

func (e *Element) BoundingRectangle() {
	val := e.UIAutoElement.Get_CurrentBoundingRectangle()
	e.CurrentBoundingRectangle = val
}

func (e *Element) ClassName() error {
	val, err := e.UIAutoElement.Get_CurrentClassName()
	e.CurrentClassName = val
	return err
}

func (e *Element) ControllerFor() {
	val := e.UIAutoElement.Get_CurrentControllerFor()
	e.CurrentControllerFor = val
}

func (e *Element) ControlType() {
	val := e.UIAutoElement.Get_CurrentControlType()
	e.CurrentControlType = val
}

func (e *Element) Culture() {
	val := e.UIAutoElement.Get_CurrentCulture()
	e.CurrentCulture = val
}

func (e *Element) DescribedBy() {
	val := e.UIAutoElement.Get_CurrentDescribedBy()
	e.CurrentDescribedBy = val
}

func (e *Element) FlowsTo() {
	val := e.UIAutoElement.Get_CurrentFlowsTo()
	e.CurrentFlowsTo = val
}

func (e *Element) FrameworkId() error {
	val, err := e.UIAutoElement.Get_CurrentFrameworkId()
	e.CurrentFrameworkId = val
	return err
}

func (e *Element) HasKeyboardFocus() {
	val := e.UIAutoElement.Get_CurrentHasKeyboardFocus()
	e.CurrentHasKeyboardFocus = val
}

func (e *Element) HelpText() error {
	val, err := e.UIAutoElement.Get_CurrentHelpText()
	e.CurrentHelpText = val
	return err
}

func (e *Element) IsControlElement() {
	val := e.UIAutoElement.Get_CurrentIsControlElement()
	e.CurrentIsControlElement = val
}

func (e *Element) IsContentElement() {
	val := e.UIAutoElement.Get_CurrentIsContentElement()
	e.CurrentIsContentElement = val
}

func (e *Element) IsDataValidForForm() {
	val := e.UIAutoElement.Get_CurrentIsDataValidForForm()
	e.CurrentIsDataValidForForm = val
}

func (e *Element) IsEnabled() {
	val := e.UIAutoElement.Get_CurrentIsEnabled()
	e.CurrentIsEnabled = val
}

func (e *Element) IsKeyboardFocusable() {
	val := e.UIAutoElement.Get_CurrentIsKeyboardFocusable()
	e.CurrentIsKeyboardFocusable = val
}

func (e *Element) IsOffscreen() {
	val := e.UIAutoElement.Get_CurrentIsOffscreen()
	e.CurrentIsOffscreen = val
}

func (e *Element) IsPassword() {
	val := e.UIAutoElement.Get_CurrentIsPassword()
	e.CurrentIsPassword = val
}

func (e *Element) IsRequiredForForm() {
	val := e.UIAutoElement.Get_CurrentIsRequiredForForm()
	e.CurrentIsRequiredForForm = val
}

func (e *Element) ItemStatus() error {
	val, err := e.UIAutoElement.Get_CurrentItemStatus()
	e.CurrentItemStatus = val
	return err
}

func (e *Element) ItemType() error {
	val, err := e.UIAutoElement.Get_CurrentItemType()
	e.CurrentItemType = val
	return err
}

func (e *Element) LabeledBy() {
	val := e.UIAutoElement.Get_CurrentLabeledBy()
	e.CurrentLabeledBy = val
}

func (e *Element) LocalizedControlType() error {
	val, err := e.UIAutoElement.Get_CurrentLocalizedControlType()
	e.CurrentLocalizedControlType = val
	return err
}

func (e *Element) Name() error {
	val, err := e.UIAutoElement.Get_CurrentName()
	e.CurrentName = val
	return err
}

func (e *Element) NativeWindowHandle() {
	val := e.UIAutoElement.Get_CurrentNativeWindowHandle()
	e.CurrentNativeWindowHandle = val
}

func (e *Element) Orientation() {
	val := e.UIAutoElement.Get_CurrentOrientation()
	e.CurrentOrientation = val
}

func (e *Element) ProcessId() {
	val := e.UIAutoElement.Get_CurrentProcessId()
	e.CurrentProcessId = val
}

func (e *Element) ProviderDescription() error {
	val, err := e.UIAutoElement.Get_CurrentProviderDescription()
	e.CurrentProviderDescription = val
	return err
}

type SearchFunc func(elem *Element) bool

func SearchElem(elem *Element, searchFunc SearchFunc) *Element {
	if elem == nil || searchFunc == nil {
		return nil
	}
	if searchFunc(elem) {
		return elem
	}
	for _, childElem := range elem.Child {
		if found := SearchElem(childElem, searchFunc); found != nil {
			return found
		}
	}
	return nil
}

func FindElems(elem *Element, searchFunc SearchFunc) (elems []*Element) {
	if searchFunc(elem) {
		elems = append(elems, elem)
	}
	for _, childElem := range elem.Child {
		if found := FindElems(childElem, searchFunc); found != nil {
			elems = append(elems, found...)
		}
	}
	return elems
}

func bstr2str(bstr uintptr) string {
	if bstr == 0 {
		return ""
	}
	// 在 Windows 上 BSTR 是以 null 结尾的 UTF-16 字符串
	// 指针直接指向字符数组。我们手动计算长度以避免创建巨型切片。
	p := (*uint16)(unsafe.Pointer(bstr))
	n := 0
	for ptr := p; *ptr != 0; ptr = (*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + 2)) {
		n++
		if n > 0x100000 { // 限制最大 1MB 字符，防止无限循环
			break
		}
	}
	// 使用 syscall.UTF16ToString 将其转换为 Go 字符串
	// 传入切片头，长度为我们计算出的 n
	var slice []uint16
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	header.Data = bstr
	header.Len = n
	header.Cap = n
	return syscall.UTF16ToString(slice)
}

func string2Bstr(str string) (uintptr, error) {
	// 将Go字符串转换为UTF-16编码
	utf16Str, err := syscall.UTF16PtrFromString(str)
	if err != nil {
		return 0, err
	}
	bstrPtr, _, _ := procSysAllocString.Call(
		uintptr(unsafe.Pointer(utf16Str)),
	)
	if bstrPtr == 0 {
		return 0, ErrorBstrPointerNil
	}
	return bstrPtr, nil
}

func (e *Element) Populate(cached bool) {
	if cached {
		e.CurrentName, _ = e.UIAutoElement.Get_CachedName()
		e.CurrentClassName, _ = e.UIAutoElement.Get_CachedClassName()
		e.CurrentControlType = e.UIAutoElement.Get_CachedControlType()
		e.CurrentAutomationId, _ = e.UIAutoElement.Get_CachedAutomationId()
		e.CurrentIsEnabled = e.UIAutoElement.Get_CachedIsEnabled()
		e.CurrentProcessId = e.UIAutoElement.Get_CachedProcessId()
		e.CurrentLocalizedControlType, _ = e.UIAutoElement.Get_CachedLocalizedControlType()
	} else {
		e.CurrentName, _ = e.UIAutoElement.Get_CurrentName()
		e.CurrentClassName, _ = e.UIAutoElement.Get_CurrentClassName()
		e.CurrentControlType = e.UIAutoElement.Get_CurrentControlType()
		e.CurrentAutomationId, _ = e.UIAutoElement.Get_CurrentAutomationId()
		e.CurrentIsEnabled = e.UIAutoElement.Get_CurrentIsEnabled()
		e.CurrentProcessId = e.UIAutoElement.Get_CurrentProcessId()
		e.CurrentLocalizedControlType, _ = e.UIAutoElement.Get_CurrentLocalizedControlType()
	}

	commonPatterns := []PatternId{
		UIA_ValuePatternId,
		UIA_InvokePatternId,
		UIA_SelectionPatternId,
		UIA_ExpandCollapsePatternId,
		UIA_TogglePatternId,
	}
	for _, pid := range commonPatterns {
		var pattern *IUnKnown
		var err error
		if cached {
			pattern, err = e.UIAutoElement.GetCachedPattern(pid)
		} else {
			pattern, err = e.UIAutoElement.GetCurrentPattern(pid)
		}
		if err == nil && pattern != nil {
			e.SupportedPatterns = append(e.SupportedPatterns, pid)
			pattern.Release()
		}
	}
}

func TraverseUIElementTree(ppv *IUIAutomation, root *IUIAutomationElement) *Element {
	cacheRequest, _ := ppv.CreateCacheRequest()
	if cacheRequest != nil {
		defer cacheRequest.Release()
		cacheRequest.AddProperty(UIA_NamePropertyId)
		cacheRequest.AddProperty(UIA_ClassNamePropertyId)
		cacheRequest.AddProperty(UIA_ControlTypePropertyId)
		cacheRequest.AddProperty(UIA_AutomationIdPropertyId)
		cacheRequest.AddProperty(UIA_IsEnabledPropertyId)
		cacheRequest.AddProperty(UIA_ProcessIdPropertyId)
		cacheRequest.AddProperty(UIA_LocalizedControlTypePropertyId)

		cacheRequest.AddPattern(UIA_ValuePatternId)
		cacheRequest.AddPattern(UIA_InvokePatternId)
		cacheRequest.AddPattern(UIA_SelectionPatternId)
		cacheRequest.AddPattern(UIA_ExpandCollapsePatternId)
		cacheRequest.AddPattern(UIA_TogglePatternId)

		cacheRequest.Put_TreeScope(TreeScope_Subtree)
	}

	return traverseUIElementTree(ppv, root, cacheRequest)
}

func traverseUIElementTree(ppv *IUIAutomation, root *IUIAutomationElement, cacheRequest *IUIAutomationCacheRequest) *Element {
	var elementArr *IUIAutomationElementArray
	var err error

	if cacheRequest != nil {
		elementArr, err = root.FindAllBuildCache(TreeScope_Children, CreateTrueCondition(ppv), cacheRequest)
	} else {
		elementArr, err = root.FindAll(TreeScope_Children, CreateTrueCondition(ppv))
	}

	newElement := NewElement(root)
	if cacheRequest != nil {
		// 如果支持缓存，由于 FindAllBuildCache 只处理子节点，根节点需要手动更新缓存
		cachedRoot, _ := root.BuildUpdatedCache(cacheRequest)
		if cachedRoot != nil {
			newElement.UIAutoElement = cachedRoot
			newElement.Populate(true)
		} else {
			newElement.Populate(false)
		}
	} else {
		newElement.Populate(false)
	}

	if err == nil && elementArr != nil {
		defer elementArr.Release()
		arrLen := elementArr.Get_Length()
		for i := 0; i < int(arrLen); i++ {
			elem, _ := elementArr.GetElement(int32(i))
			if elem != nil {
				childElem := traverseUIElementTree(ppv, elem, cacheRequest)
				newElement.Child = append(newElement.Child, childElem)
			}
		}
	}

	return newElement
}

func TreeString(root *Element, level int) {
	if root == nil {
		return
	}
	fmt.Printf("%s- %v\n", getIndentation(level), root)
	for _, child := range root.Child {
		TreeString(child, level+1)
	}
}

func (e *Element) String() string {
	return fmt.Sprintf("Element{Name: %q, Class: %q, Type: %d}", e.CurrentName, e.CurrentClassName, e.CurrentControlType)
}

func getIndentation(level int) string {
	return strings.Repeat("  ", level)
}

func (e *Element) FindByName(name string) *Element {
	if e.CurrentName == name {
		return e
	}
	for _, child := range e.Child {
		if found := child.FindByName(name); found != nil {
			return found
		}
	}
	return nil
}

func (e *Element) FindByAutomationId(id string) *Element {
	if e.CurrentAutomationId == id {
		return e
	}
	for _, child := range e.Child {
		if found := child.FindByAutomationId(id); found != nil {
			return found
		}
	}
	return nil
}

func (e *Element) GetValuePattern() (*IUIAutomationValuePattern, error) {
	unk, err := e.UIAutoElement.GetCurrentPattern(UIA_ValuePatternId)
	if err != nil {
		return nil, err
	}
	return (*IUIAutomationValuePattern)(unsafe.Pointer(unk)), nil
}

func (e *Element) GetInvokePattern() (*IUIAutomationInvokePattern, error) {
	unk, err := e.UIAutoElement.GetCurrentPattern(UIA_InvokePatternId)
	if err != nil {
		return nil, err
	}
	return (*IUIAutomationInvokePattern)(unsafe.Pointer(unk)), nil
}

func (e *Element) GetTogglePattern() (*IUIAutomationTogglePattern, error) {
	unk, err := e.UIAutoElement.GetCurrentPattern(UIA_TogglePatternId)
	if err != nil {
		return nil, err
	}
	return (*IUIAutomationTogglePattern)(unsafe.Pointer(unk)), nil
}

func (e *Element) GetExpandCollapsePattern() (*IUIAutomationExpandCollapsePattern, error) {
	unk, err := e.UIAutoElement.GetCurrentPattern(UIA_ExpandCollapsePatternId)
	if err != nil {
		return nil, err
	}
	return (*IUIAutomationExpandCollapsePattern)(unsafe.Pointer(unk)), nil
}

func (e *Element) GetSelectionItemPattern() (*IUIAutomationSelectionItemPattern, error) {
	unk, err := e.UIAutoElement.GetCurrentPattern(UIA_SelectionItemPatternId)
	if err != nil {
		return nil, err
	}
	return (*IUIAutomationSelectionItemPattern)(unsafe.Pointer(unk)), nil
}
