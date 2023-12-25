package uiautomation

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"syscall"
	"unsafe"
)

var (
	oleAut            = syscall.NewLazyDLL("OleAut32.dll")
	procSysFreeString = oleAut.NewProc("SysFreeString")
)

type OrientationType = int

const (
	OrientationType_None       OrientationType = 0
	OrientationType_Horizontal OrientationType = 1
	OrientationType_Vertical   OrientationType = 2
)

type Element struct {
	UIAutoElement               *IUIAutomationElement
	CurrentAcceleratorKey       string
	CurrentAccessKey            string
	CurrentAriaProperties       string
	CurrentAriaRole             string
	CurrentAutomationId         string
	CurrentBoundingRectangle    *RECT
	CurrentClassName            string
	CurrentControllerFor        *IUIAutomationElementArray
	CurrentControlType          int
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

	Child []*Element
}

func getElement(i int, elemArr *IUIAutomationElementArray) (elem *IUIAutomationElement) {
	syscall.SyscallN(
		(*IUIAutomationElementArrayVtbl)(unsafe.Pointer(elemArr.vtbl)).GetElement,
		uintptr(unsafe.Pointer(elemArr)),
		uintptr(i),
		uintptr(unsafe.Pointer(&elem)),
	)
	return
}

func getClickablePoint(ele *IUIAutomationElement) *TagPoint {
	var point TagPoint
	var gotClick int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(ele.vtbl)).GetClickablePoint,
		uintptr(unsafe.Pointer(ele)),
		uintptr(unsafe.Pointer(&point)),
		uintptr(unsafe.Pointer(&gotClick)),
	)
	if gotClick > 0 {
		return &point
	}
	return nil
}
func GetCurrentPattern(ele *IUIAutomationElement, patternid PatternId) uintptr {
	var patternObj uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(ele.vtbl)).GetCurrentPattern,
		uintptr(unsafe.Pointer(ele)),
		uintptr(patternid),
		uintptr(unsafe.Pointer(&patternObj)),
	)
	return patternObj
}
func getCurrentPatternAs(ele *IUIAutomationElement) {
}
func getCurrentPropertyValue(ele *IUIAutomationElement) {
}
func getCurrentPropertyValueEx(ele *IUIAutomationElement) {
}
func getRuntimeId(ele *IUIAutomationElement) {
	var safearray SafeArray
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(ele.vtbl)).GetRuntimeId,
		uintptr(unsafe.Pointer(ele)),
		uintptr(unsafe.Pointer(&safearray)),
	)
}
func setFocus(ele *IUIAutomationElement) {
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(ele.vtbl)).SetFocus,
		uintptr(unsafe.Pointer(ele)),
	)
}

func (e *Element) FormatString() string {
	val := reflect.ValueOf(e)
	val = reflect.Indirect(val)
	num := val.NumField()
	rType := val.Type()
	buf := ""
	for i := 0; i < num; i++ {
		v := rType.Field(i)
		buf += fmt.Sprintf("[%s]:[%v],", v.Name, val.Field(i))
	}
	return buf
}
func (e *Element) AccleratorKey() {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentAcceleratorKey,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr != 0 {
		e.CurrentAcceleratorKey = bstr2str(bstr)
	}
	procSysFreeString.Call(
		uintptr(bstr),
	)
}
func (e *Element) AccessKey() {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentAccessKey,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr != 0 {
		e.CurrentAccessKey = bstr2str(bstr)
	}
	procSysFreeString.Call(
		uintptr(bstr),
	)
}
func (e *Element) AriaProperties() {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentAriaProperties,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr != 0 {
		e.CurrentAriaProperties = bstr2str(bstr)
	}
	procSysFreeString.Call(
		uintptr(bstr),
	)
}
func (e *Element) AriaRole() {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentAriaRole,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	e.CurrentAriaRole = bstr2str(bstr)
	procSysFreeString.Call(
		uintptr(bstr),
	)
}
func (e *Element) AutomationId() {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentAutomationId,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr != 0 {
		e.CurrentAutomationId = bstr2str(bstr)
	}
	procSysFreeString.Call(
		uintptr(bstr),
	)
}

func (e *Element) BoundingRectangle() {
	var rect *RECT
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentBoundingRectangle,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(rect)),
	)
	e.CurrentBoundingRectangle = rect
}

func (e *Element) ClassName() {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentClassName,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr != 0 {
		e.CurrentClassName = bstr2str(bstr)
	}
	procSysFreeString.Call(
		uintptr(bstr),
	)
}
func (e *Element) ControllerFor() {
	var array *IUIAutomationElementArray
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentControllerFor,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&array)),
	)
	e.CurrentControllerFor = array
}

func (e *Element) ControlType() {
	var controlTypeId int
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentControlType,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&controlTypeId)),
	)
	e.CurrentControlType = controlTypeId
}

func (e *Element) Culture() {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentCulture,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	e.CurrentCulture = retVal
}

func (e *Element) DescribedBy() {
	var array *IUIAutomationElementArray
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentDescribedBy,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&array)),
	)
	e.CurrentDescribedBy = array
}

func (e *Element) FlowsTo() {
	var array *IUIAutomationElementArray
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentFlowsTo,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&array)),
	)
	e.CurrentFlowsTo = array
}

func (e *Element) FrameworkId() {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentFrameworkId,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr != 0 {
		e.CurrentFrameworkId = bstr2str(bstr)
	}
	procSysFreeString.Call(
		uintptr(bstr),
	)
}
func (e *Element) HasKeyboardFocus() {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentHasKeyboardFocus,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	e.CurrentHasKeyboardFocus = retVal
}

func (e *Element) HelpText() {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentHelpText,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr != 0 {
		e.CurrentHelpText = bstr2str(bstr)
	}
	procSysFreeString.Call(uintptr(bstr))
}
func (e *Element) IsControlElement() {
	var isControl int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentIsControlElement,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&isControl)),
	)
	e.CurrentIsControlElement = isControl
}
func (e *Element) IsContentElement() {
	var isContent int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentIsContentElement,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&isContent)),
	)
	e.CurrentIsContentElement = isContent
}
func (e *Element) IsDataValidForForm() {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentIsDataValidForForm,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	e.CurrentIsDataValidForForm = retVal
}
func (e *Element) IsEnabled() {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentIsEnabled,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	e.CurrentIsEnabled = retVal
}
func (e *Element) IsKeyboardFocusable() {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentIsKeyboardFocusable,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	e.CurrentIsKeyboardFocusable = retVal
}
func (e *Element) IsOffscreen() {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentIsOffscreen,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	e.CurrentIsOffscreen = retVal
}
func (e *Element) IsPassword() {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentIsPassword,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	e.CurrentIsPassword = retVal
}
func (e *Element) IsRequiredForForm() {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentIsRequiredForForm,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	e.CurrentIsRequiredForForm = retVal
}
func (e *Element) ItemStatus() {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentItemStatus,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr != 0 {
		e.CurrentItemStatus = bstr2str(bstr)
	}
	procSysFreeString.Call(uintptr(bstr))
}
func (e *Element) ItemType() {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentItemType,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr != 0 {
		e.CurrentItemType = bstr2str(bstr)
	}
	procSysFreeString.Call(uintptr(bstr))
}
func (e *Element) LabeledBy() {
	var elem *IUIAutomationElement
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentLabeledBy,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&elem)),
	)
	e.CurrentLabeledBy = elem
}
func (e *Element) LocalizedControlType() {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentLocalizedControlType,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr != 0 {
		e.CurrentLocalizedControlType = bstr2str(bstr)
	}
	procSysFreeString.Call(
		uintptr(bstr),
	)
}
func (e *Element) Name() {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentName,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr != 0 {
		e.CurrentName = bstr2str(bstr)
	}
	procSysFreeString.Call(
		uintptr(bstr),
	)
}
func (e *Element) NativeWindowHandle() {
	var retVal uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentNativeWindowHandle,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	e.CurrentNativeWindowHandle = retVal
}
func (e *Element) Orientation() {
	var retVal OrientationType
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentOrientation,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	e.CurrentOrientation = retVal
}
func (e *Element) ProcessId() {
	var retVal int32
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentProcessId,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	e.CurrentProcessId = retVal
}
func (e *Element) ProviderDescription() {
	var bstr uintptr
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(e.UIAutoElement.vtbl)).Get_CurrentProviderDescription,
		uintptr(unsafe.Pointer(e.UIAutoElement)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if bstr != 0 {
		e.CurrentProviderDescription = bstr2str(bstr)
	}
	procSysFreeString.Call(uintptr(bstr))
}

func (e *Element) SetUIAutomation(uiaumation *IUIAutomationElement) {
	e.UIAutoElement = uiaumation
}

type SearchFunc func(elem *Element) bool

func SearchElem(elem *Element, searchFunc SearchFunc) *Element {
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

func NewElement(uiaumation *IUIAutomationElement) *Element {
	return &Element{
		UIAutoElement: uiaumation,
	}
}

func bstr2str(bstr uintptr) string {
	return syscall.UTF16ToString((*[1 << 30]uint16)(unsafe.Pointer(bstr))[:])
}

func TraverseUIElementTree(ppv *IUIAutomation, root *IUIAutomationElement) *Element {
	return traverseUIElementTree(ppv, root)
}

func traverseUIElementTree(ppv *IUIAutomation, root *IUIAutomationElement) *Element {
	var arr *IUIAutomationElementArray
	var conditions *IUIAutomationCondition
	syscall.SyscallN(
		(*IUIAutomationVtbl)(unsafe.Pointer(ppv.vtbl)).CreateTrueCondition,
		uintptr(unsafe.Pointer(ppv)),
		uintptr(unsafe.Pointer(&conditions)),
	)
	syscall.SyscallN(
		(*IUIAutomationElementVtbl)(unsafe.Pointer(root.vtbl)).FindAll,
		uintptr(unsafe.Pointer(root)),
		uintptr(TreeScope_Children),
		uintptr(unsafe.Pointer(conditions)),
		uintptr(unsafe.Pointer(&arr)),
	)
	var arrLen int
	syscall.SyscallN(
		(*IUIAutomationElementArrayVtbl)(unsafe.Pointer(arr.vtbl)).Get_Length,
		uintptr(unsafe.Pointer(arr)),
		uintptr(unsafe.Pointer(&arrLen)),
	)
	main := NewElement(root)
	main.AccleratorKey()
	main.AccessKey()
	main.AriaProperties()
	main.AriaRole()
	main.AutomationId()
	main.BoundingRectangle()
	main.ClassName()
	main.ControllerFor()
	main.ControlType()
	main.Culture()
	main.DescribedBy()
	main.FlowsTo()
	main.FrameworkId()
	main.HasKeyboardFocus()
	main.HelpText()
	main.IsContentElement()
	main.IsControlElement()
	main.IsDataValidForForm()
	main.IsEnabled()
	main.IsKeyboardFocusable()
	main.IsOffscreen()
	main.IsPassword()
	main.IsRequiredForForm()
	main.ItemStatus()
	main.ItemType()
	main.LabeledBy()
	main.LocalizedControlType()
	main.Name()
	main.NativeWindowHandle()
	main.Orientation()
	main.ProcessId()
	main.ProviderDescription()
	for i := 0; i < arrLen; i++ {
		elem := getElement(i, arr)
		childElem := traverseUIElementTree(ppv, elem)
		main.Child = append(main.Child, childElem)
	}
	return main
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

func getIndentation(level int) string {
	return strings.Repeat("  ", level)
}

func WriteJSONToFile(data interface{}, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(data)
}
