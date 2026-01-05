## 用于 Windows 窗体自动化的 Go 语言 UI Automation 库
[![Go Reference](https://pkg.go.dev/badge/github.com/auuunya/go-element.svg)](https://pkg.go.dev/github.com/auuunya/go-element)

### 概述
本库为使用 Windows UI Automation 框架进行 UI 自动化提供了一个高性能且易用的 Go 语言接口。它允许开发者通过 Windows COM 接口与应用程序中的 UI 元素进行高效交互和操作。

### 核心改进 
- **极致性能**: 深度集成 `IUIAutomationCacheRequest`，实现属性和模式的批量获取，大幅减少跨进程通信 (IPC) 开销。
- **零反射设计**: 彻底移除反射逻辑，改为显式属性填充，提升运行时效率并增强类型安全。
- **完善的模式支持**: 内置常用自动化模式 (Patterns) 的包装，包括：`Value`（值）、`Invoke`（调用）、`Toggle`（切换）、`ExpandCollapse`（展开折叠）及 `SelectionItem`（选中项）。
- **资源安全**: 严格的 COM 对象生命周期管理，通过显式的 `Release()` 调用确保无内存泄漏和句柄耗尽。
- **便捷查询 API**: `Element` 结构体集成 `FindByName` 和 `FindByAutomationId` 等快捷方法，定位元素更简单。

### 安装
```shell
go get -u github.com/auuunya/go-element
```

### 编译说明
为了确保在不同架构的 Windows 系统（如 XP/7/10/11 的 32 位或 64 位版本）上正常运行，请参考以下编译命令：

- **64 位系统 (默认)**:
  ```shell
  GOOS=windows GOARCH=amd64 go build .
  ```
- **32 位系统 (高兼容性)**:
  ```shell
  GOOS=windows GOARCH=386 go build .
  ```
> **注意**: 如果在运行 64 位编译的程序时提示 `is not valid win32 application`，请尝试使用 `386` 架构重新编译。

### 快速开始：

#### 1. 输出 UI 结构树 (以浏览器为例)
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

	// 使用 Cache Request 进行高性能遍历
	elems := uia.TraverseUIElementTree(ppv, root)
	uia.TreeString(elems, 0)
}
```

#### 2. 自动化记事本 (模式调用示例)
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

	// 通过名称快速查找元素并使用 ValuePattern 设置内容
	if editor := tree.FindByName("文本编辑器"); editor != nil {
		if vp, err := editor.GetValuePattern(); err == nil {
			defer vp.Release()
			vp.SetValue("你好，来自 go-element 的自动输入！")
		}
	}
}
```

### 任务列表
- [x] 高性能属性缓存 (Cache Request)
- [x] UI 结构 JSON 序列化导出
- [x] 常用 Pattern 包装 (Invoke, Value, Toggle 等)
- [ ] UI 事件监听支持 (Event Handlers)
- [ ] 更丰富的 UI 交互操作

### 贡献
欢迎贡献！如果您发现 Bug 或想添加新功能，请随时提交 Issue 或 Pull Request。

### 许可证
本库采用 MIT 许可证。