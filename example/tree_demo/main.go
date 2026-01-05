package main

import (
	"fmt"

	uiautomation "github.com/auuunya/go-element"
)

// chrome browser
func main() {
	uiautomation.CoInitialize()
	defer uiautomation.CoUninitialize()

	// 查找当前活动窗口（例如记事本或浏览器）
	findhwnd, err := uiautomation.GetWindowForString("Notepad", "")
	if err != nil {
		fmt.Println("未找到记事本，尝试查找 Chrome...")
		findhwnd, _ = uiautomation.GetWindowForString("Chrome_WidgetWin_1", "")
	}

	if findhwnd == 0 {
		fmt.Println("未找到目标窗口，请确保记事本或浏览器已打开。")
		return
	}

	instance, _ := uiautomation.CreateInstance(uiautomation.CLSID_CUIAutomation, uiautomation.IID_IUIAutomation, uiautomation.CLSCTX_INPROC_SERVER|uiautomation.CLSCTX_LOCAL_SERVER|uiautomation.CLSCTX_REMOTE_SERVER)
	ppv := uiautomation.NewIUIAutomation(uiautomation.NewIUnKnown(instance))
	defer ppv.Release()

	root, _ := uiautomation.ElementFromHandle(ppv, findhwnd)
	defer root.Release()

	fmt.Println("正在构建 UI 树（缓存加速模式）...")
	elems := uiautomation.TraverseUIElementTree(ppv, root)

	fmt.Println("\n--- UI Element Tree ---")
	uiautomation.TreeString(elems, 0)
}
