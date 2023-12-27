# go-element

[![Go Reference](https://pkg.go.dev/badge/github.com/auuunya/go-element.svg)](https://pkg.go.dev/github.com/auuunya/go-element)

This is a windows window application for an element traversal, you can find the elements you want to operate the control, and then the next step to click or enter the operation, etc., can be traversed for the element tree, search for the specified element

### Examples:
**browser**
```go
func main() {
	uiautomation.CoInitialize()
	defer uiautomation.CoUninitialize()
	findhwnd := uiautomation.GetWindowForString("Chrome_WidgetWin_1", "")
	instance, _ := uiautomation.CreateUiautomationInstance()
	ppv := uiautomation.NewIUIAutomation(uintptr(instance))
	root, _ := uiautomation.ElementFromHandle(ppv, findhwnd)
	elems := uiautomation.TraverseUIElementTree(ppv, root)
	uiautomation.TreeString(elems, 0)
}
```

**Output Notepad application element tree structure**
```go
func main() {
	uiautomation.CoInitialize()
	defer uiautomation.CoUninitialize()
	findhwnd := uiautomation.GetWindowForString("Notepad", "")
	instance := uiautomation.CreateUiautomationInstance()
	ppv := (*uiautomation.IUIAutomation)(unsafe.Pointer(instance))
	root := uiautomation.ElementFromHandle(ppv, findhwnd)
	elems := uiautomation.TraverseUIElementTree(ppv, root)
	uiautomation.TreeString(elems, 0)
}
```
**Search directory window element**
```go
func main() {
	uiautomation.CoInitialize()
	defer uiautomation.CoUninitialize()
	findhwnd := uiautomation.GetWindowForString("CabinetWClass", "")
	instance, _ := uiautomation.CreateUiautomationInstance()
	ppv := uiautomation.NewIUIAutomation(uintptr(instance))
	root, _ := uiautomation.ElementFromHandle(ppv, findhwnd)
	elems := uiautomation.TraverseUIElementTree(ppv, root)
	fn := func(elem *uiautomation.Element) bool {
		return elem.CurrentClassName == "SelectorButton" && elem.CurrentName == "详细信息"
	}
	foundElements := uiautomation.SearchElem(elems, fn)
	fmt.Printf("foundElement: %#v\n", foundElements)

    fnAll := func(elem *uiautomation.Element) bool {
		return elem.CurrentClassName == "SelectorButton"
	}
	foundElements := uiautomation.FindElems(elems, fnAll)
	fmt.Printf("foundElements: %#v\n", foundElements)
}
```

