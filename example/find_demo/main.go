package main

import (
	"fmt"

	uiautomation "github.com/auuunya/go-element"
)

func main() {
	uiautomation.CoInitialize()
	defer uiautomation.CoUninitialize()

	// 查找资源管理器窗口
	findhwnd, err := uiautomation.GetWindowForString("CabinetWClass", "")
	if err != nil {
		fmt.Println("请先打开一个资源管理器窗口 (Explorer)")
		return
	}

	instance, _ := uiautomation.CreateInstance(uiautomation.CLSID_CUIAutomation, uiautomation.IID_IUIAutomation, uiautomation.CLSCTX_INPROC_SERVER|uiautomation.CLSCTX_LOCAL_SERVER|uiautomation.CLSCTX_REMOTE_SERVER)
	ppv := uiautomation.NewIUIAutomation(uiautomation.NewIUnKnown(instance))
	defer ppv.Release()

	root, _ := uiautomation.ElementFromHandle(ppv, findhwnd)
	defer root.Release()

	fmt.Println("遍历 UI 树...")
	tree := uiautomation.TraverseUIElementTree(ppv, root)

	// 使用快捷方法查找
	fmt.Println("查找名为 '详细信息' 的元素...")
	detailBtn := tree.FindByName("详细信息")
	if detailBtn != nil {
		fmt.Printf("找到元素: %s (Class: %s)\n", detailBtn.CurrentName, detailBtn.CurrentClassName)
	}

	// 使用传统搜索函数
	fmt.Println("搜索所有类型为 'SelectorButton' 的元素...")
	fnAll := func(elem *uiautomation.Element) bool {
		return elem.CurrentClassName == "SelectorButton"
	}
	foundElements := uiautomation.FindElems(tree, fnAll)
	fmt.Printf("找到 %d 个 SelectorButton 元素\n", len(foundElements))
}
