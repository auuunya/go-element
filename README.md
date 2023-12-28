## Uiautomation Library for Windows in Go
[![Go Reference](https://pkg.go.dev/badge/github.com/auuunya/go-element.svg)](https://pkg.go.dev/github.com/auuunya/go-element)

### Overview
This Go library provides a convenient interface for automating UI tasks on Windows using the UI Automation framework. It allows you to interact with and manipulate UI elements in applications through the corresponding Windows COM interfaces. The library maps Go methods to the methods defined in the COM interfaces, making it easy to work with UI Automation in a Go application.

### Features
- **COM Interface Mapping:** The library mirrors the structure of UI Automation COM interfaces, Corresponds to the Windows API structure
- **Element Structuring:** The structure of the application's UI elements can be visualized more intuitively
- **Automation:** Perform UI automation tasks programmatically, such as reading or manipulating the state of UI controls.

### Install
```shell
go get -u github.com/auuunya/go-element
```
```go
import (
    ...
    uiautomation "github.com/auuunya/go-element"
)
```

### Examples:
**Browser Element Structure Output**
```go
func main() {
	uiautomation.CoInitialize()
	defer uiautomation.CoUninitialize()
	findhwnd := uiautomation.GetWindowForString("Chrome_WidgetWin_1", "")
	instance, _ := uiautomation.CreateInstance(uiautomation.CLSID_CUIAutomation, uiautomation.IID_IUIAutomation, uiautomation.CLSCTX_INPROC_SERVER|uiautomation.CLSCTX_LOCAL_SERVER|uiautomation.CLSCTX_REMOTE_SERVER)
	unk := uiautomation.NewIUnKnown(instance)
	ppv := uiautomation.NewIUIAutomation(unk)
	root, _ := uiautomation.ElementFromHandle(ppv, findhwnd)
	elems := uiautomation.TraverseUIElementTree(ppv, root)
	uiautomation.TreeString(elems, 0)
}
```

**Find all button controls in Notepad**
```go
func main() {
	uiautomation.CoInitialize()
	defer uiautomation.CoUninitialize()
	findhwnd := uiautomation.GetWindowForString("Notepad", "")
	instance, _ := uiautomation.CreateInstance(uiautomation.CLSID_CUIAutomation, uiautomation.IID_IUIAutomation, uiautomation.CLSCTX_INPROC_SERVER|uiautomation.CLSCTX_LOCAL_SERVER|uiautomation.CLSCTX_REMOTE_SERVER)
	unk := uiautomation.NewIUnKnown(instance)
	ppv := uiautomation.NewIUIAutomation(unk)
	root := uiautomation.ElementFromHandle(ppv, findhwnd)
	elems := uiautomation.TraverseUIElementTree(ppv, root)
	fn := func(elem *uiautomation.Element) bool {
		return elem.CurrentName == "Button"
	}
	foundAll := uiautomation.FindElems(elems, fn)
	fmt.Printf("foundAll elems: %v\n", foundAll)
}
```
**Search for controls specified in the folder**
```go
func main() {
	uiautomation.CoInitialize()
	defer uiautomation.CoUninitialize()
	findhwnd := uiautomation.GetWindowForString("CabinetWClass", "")
	instance, _ := uiautomation.CreateInstance(uiautomation.CLSID_CUIAutomation, uiautomation.IID_IUIAutomation, uiautomation.CLSCTX_INPROC_SERVER|uiautomation.CLSCTX_LOCAL_SERVER|uiautomation.CLSCTX_REMOTE_SERVER)
	unk := uiautomation.NewIUnKnown(instance)
	ppv := uiautomation.NewIUIAutomation(unk)
	root, _ := uiautomation.ElementFromHandle(ppv, findhwnd)
	elems := uiautomation.TraverseUIElementTree(ppv, root)
	fn := func(elem *uiautomation.Element) bool {
		return elem.CurrentClassName == "SelectorButton"
	}
	foundElement := uiautomation.SearchElem(elems, fn)
	fmt.Printf("foundElement: %v\n", foundElement)
}
```

### Contribution
Contributions are welcome! If you find a bug or want to enhance the library, feel free to open an issue or submit a pull request.

### License
This library is distributed under the MIT License