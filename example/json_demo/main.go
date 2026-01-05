package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	uiautomation "github.com/auuunya/go-element"
)

func main() {
	uiautomation.CoInitialize()
	defer uiautomation.CoUninitialize()

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

	fmt.Println("正在构建 UI 树并缓存属性...")
	elems := uiautomation.TraverseUIElementTree(ppv, root)

	// 序列化为 JSON
	data, err := json.MarshalIndent(elems, "", "  ")
	if err != nil {
		log.Fatalf("JSON 序列化失败: %v", err)
	}

	fileName := "ui_tree.json"
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		log.Fatalf("保存文件失败: %v", err)
	}

	fmt.Printf("UI 树已保存至: %s\n", fileName)
}
