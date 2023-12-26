package uiautomation

import (
	"fmt"
	"reflect"
	"strings"
	"syscall"
	"unsafe"
)

var (
	oleAut            = syscall.NewLazyDLL("OleAut32.dll")
	procSysFreeString = oleAut.NewProc("SysFreeString")
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
	CurrentControlType          syscall.GUID
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
	CurrentOrientation          *OrientationType
	CurrentProcessId            int32
	CurrentProviderDescription  string

	Child []*Element `json:"child,omitempty"`
}

func NewElement(uiaumation *IUIAutomationElement) *Element {
	return &Element{
		UIAutoElement: uiaumation,
	}
}
func (e *Element) FormatString() string {
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
				// fmt.Printf("fieldValue.Addr():%#v\n", fieldValue.Addr().Interface())
				v := fieldValue.Addr().Interface()
				// fmt.Printf("222: %#v\n", v)
				buf += fmt.Sprintf("[%s]:[%v],", fieldName, v)
			}
		} else {
			buf += fmt.Sprintf("[%s]:[%v],", fieldName, fieldValue.Interface())
		}
	}
	// buf += fmt.Sprintf("[%s]:[%v],", fieldName, fieldValue.Interface())
	return buf
}
func (e *Element) SetUIAutomation(uiaumation *IUIAutomationElement) {
	e.UIAutoElement = uiaumation
}

func (e *Element) AccleratorKey() error {
	val, err := e.UIAutoElement.Get_CurrentAcceleratorKey()
	e.CurrentAcceleratorKey = val
	return err
}

func (e *Element) AccessKey() error {
	val, err := e.UIAutoElement.Get_CurrentAccessKey()
	e.CurrentAcceleratorKey = val
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
	return syscall.UTF16ToString((*[1 << 30]uint16)(unsafe.Pointer(bstr))[:])
}

func TraverseUIElementTree(ppv *IUIAutomation, root *IUIAutomationElement) *Element {
	return traverseUIElementTree(ppv, root)
}

func traverseUIElementTree(ppv *IUIAutomation, root *IUIAutomationElement) *Element {
	condition := CreateTrueCondition(ppv)
	elementArr, _ := FindAll(root, condition)
	arrLen := Get_Length(elementArr)
	newElement := NewElement(root)
	newElement.AccleratorKey()
	newElement.AccessKey()
	newElement.AriaProperties()
	newElement.AriaRole()
	newElement.AutomationId()
	newElement.BoundingRectangle()
	newElement.ClassName()
	newElement.ControllerFor()
	newElement.ControlType()
	newElement.Culture()
	newElement.DescribedBy()
	newElement.FlowsTo()
	newElement.FrameworkId()
	newElement.HasKeyboardFocus()
	newElement.HelpText()
	newElement.IsContentElement()
	newElement.IsControlElement()
	newElement.IsDataValidForForm()
	newElement.IsEnabled()
	newElement.IsKeyboardFocusable()
	newElement.IsOffscreen()
	newElement.IsPassword()
	newElement.IsRequiredForForm()
	newElement.ItemStatus()
	newElement.ItemType()
	newElement.LabeledBy()
	newElement.LocalizedControlType()
	newElement.Name()
	newElement.NativeWindowHandle()
	newElement.Orientation()
	newElement.ProcessId()
	newElement.ProviderDescription()
	for i := 0; i < int(arrLen); i++ {
		elem, _ := GetElement(elementArr, int32(i))
		childElem := traverseUIElementTree(ppv, elem)
		newElement.Child = append(newElement.Child, childElem)
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

func getIndentation(level int) string {
	return strings.Repeat("  ", level)
}
