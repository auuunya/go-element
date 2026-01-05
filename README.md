## UI Automation Library for Windows in Go
[![Go Reference](https://pkg.go.dev/badge/github.com/auuunya/go-element.svg)](https://pkg.go.dev/github.com/auuunya/go-element)

### Overview
This Go library provides a powerful and high-performance interface for automating UI tasks on Windows using the UI Automation framework. It allows you to interact with and manipulate UI elements in applications through Windows COM interfaces.

### Core Enhancements
- **High Performance**: Optimized with `IUIAutomationCacheRequest` to batch fetch properties and patterns, reducing IPC (Inter-Process Communication) overhead significantly.
- **Reflection-Free**: Removed reflection-based property population for improved runtime efficiency and better type safety.
- **Improved Pattern Support**: Ready-to-use wrappers for common Automation Patterns: `Value`, `Invoke`, `Toggle`, `ExpandCollapse`, and `SelectionItem`.
- **Resource Safety**: Strict COM object lifecycle management with explicit `Release()` calls to prevent memory leaks and handle exhaustion.
- **Better Search API**: Convenient methods like `FindByName` and `FindByAutomationId` integrated into the `Element` structure.

### Installation
```shell
go get -u github.com/auuunya/go-element
```

### Compilation Instructions
To ensure proper operation on various Windows architectures (such as 32-bit or 64-bit versions of XP/7/10/11), please refer to the following build commands:

- **64-bit System (Default)**:
  ```shell
  GOOS=windows GOARCH=amd64 go build .
  ```
- **32-bit System (High Compatibility)**:
  ```shell
  GOOS=windows GOARCH=386 go build .
  ```
> **Note**: If you receive an `is not valid win32 application` error when running a 64-bit compiled program, please try rebuilding with the `386` architecture.

### Quick Start:

#### 1. Output UI Tree (Browser Example)
```go
import uia "github.com/auuunya/go-element"

func main() {
	uia.CoInitialize()
	defer uia.CoUninitialize()
	
	hwnd, _ := uia.GetWindowForString("Chrome_WidgetWin_1", "")
	instance, _ := uia.CreateInstance(uia.CLSID_CUIAutomation, uia.IID_IUIAutomation, uia.CLSCTX_INPROC_SERVER)
	ppv := uia.NewIUIAutomation(uia.NewIUnKnown(instance))
	defer ppv.Release()

	root, _ := uia.ElementFromHandle(ppv, hwnd)
	defer root.Release()

	// High-performance traversal using Cache Request
	elems := uia.TraverseUIElementTree(ppv, root)
	uia.TreeString(elems, 0)
}
```

#### 2. Automobile Notepad (Pattern Usage)
```go
func main() {
	uia.CoInitialize()
	defer uia.CoUninitialize()
	
	hwnd, _ := uia.GetWindowForString("Notepad", "")
	instance, _ := uia.CreateInstance(uia.CLSID_CUIAutomation, uia.IID_IUIAutomation, uia.CLSCTX_ALL)
	ppv := uia.NewIUIAutomation(uia.NewIUnKnown(instance))
	defer ppv.Release()

	root, _ := uia.ElementFromHandle(ppv, hwnd)
	tree := uia.TraverseUIElementTree(ppv, root)

	// Find element by name and set value via Pattern
	if editor := tree.FindByName("文本编辑器"); editor != nil {
		if vp, err := editor.GetValuePattern(); err == nil {
			defer vp.Release()
			vp.SetValue("Hello from go-element!")
		}
	}
}
```

### Features Status
- [x] High-performance property caching
- [x] JSON serialization of UI structure
- [x] Common Pattern wrappers (Invoke, Value, Toggle, etc.)
- [ ] UI Event listener support
- [ ] More specialized UI operations

### Contribution
Contributions are welcome! If you find a bug or want to enhance the library, feel free to open an issue or submit a pull request.

### License
This library is distributed under the MIT License