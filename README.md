# go-element

[![Go Reference](https://pkg.go.dev/badge/github.com/auuunya/go-element.svg)](https://pkg.go.dev/github.com/auuunya/go-element)

This is a windows window application for an element traversal, you can find the elements you want to operate the control, and then the next step to click or enter the operation, etc., can be traversed for the element tree, search for the specified element

### Examples:
**Output window element tree structure**
```go
func main() {
	com.CoInitialize()
	defer com.CoUninitialize()
	findhwnd := com.GetWindowForString("Notepad", "")
	instance := com.CreateInstance()
	ppv := (*uiautomation.IUIAutomation)(unsafe.Pointer(instance))
	root := uiautomation.ElementFromHandle(ppv, findhwnd)
	elems := uiautomation.TraverseUIElementTree(ppv, root)
	uiautomation.TreeString(elems, 0)
}
```
**Search window element**
```go
func main() {
	com.CoInitialize()
	defer com.CoUninitialize()
	findhwnd := com.GetWindowForString("Notepad", "")
	instance := com.CreateInstance()
	ppv := (*uiautomation.IUIAutomation)(unsafe.Pointer(instance))
	root := uiautomation.ElementFromHandle(ppv, findhwnd)
	elems := uiautomation.TraverseUIElementTree(ppv, root)
	fn := func(elem *uiautomation.Element) bool {
		return elem.CurrentLocalizedControlType == "菜单项目" && elem.CurrentName == "编辑"
	}
	foundElement := uiautomation.SearchElem(elems, fn)
	fmt.Printf("foundElement: %#v\n", foundElement)
}
```
**Output to Json file**
```go
func main() {
	com.CoInitialize()
	defer com.CoUninitialize()
	findhwnd := com.GetWindowForString("Notepad", "")
	instance := com.CreateInstance()
	ppv := (*uiautomation.IUIAutomation)(unsafe.Pointer(instance))
	root := uiautomation.ElementFromHandle(ppv, findhwnd)
	elems := uiautomation.TraverseUIElementTree(ppv, root)
	uiautomation.WriteJSONToFile(elems, "filename.json")
}
```
