package main

import (
	"fmt"
	"log"
	"os"

	uia "github.com/auuunya/go-element"
)

func main() {
	// 配置日志
	f, err := os.OpenFile("find_demo.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("无法打开日志文件: %v\n", err)
	} else {
		defer f.Close()
		log.SetOutput(f)
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("--- 启动元素查找演示 ---")

	err = uia.CoInitialize()
	if err != nil {
		log.Fatalf("COM 初始化失败: %v", err)
	}
	defer uia.CoUninitialize()
	log.Println("COM 初始化成功")

	log.Println("正在查找资源管理器窗口...")
	findhwnd, err := uia.GetWindowForString("CabinetWClass", "")
	if err != nil {
		log.Fatal("未找到资源管理器窗口，请先打开一个文件夹 (Explorer)")
	}
	log.Printf("找到资源管理器，句柄: 0x%X", findhwnd)

	clsctx := uia.CLSCTX_INPROC_SERVER | uia.CLSCTX_LOCAL_SERVER | uia.CLSCTX_REMOTE_SERVER
	instance, err := uia.CreateInstance(uia.CLSID_CUIAutomation, uia.IID_IUIAutomation, clsctx)
	if err != nil {
		log.Fatalf("无法创建 UIAutomation 实例: %v", err)
	}
	log.Printf("UIAutomation 实例创建成功: %p", instance)

	ppv := uia.NewIUIAutomation(uia.NewIUnKnown(instance))
	defer ppv.Release()

	root, err := uia.ElementFromHandle(ppv, findhwnd)
	if err != nil {
		log.Fatalf("无法获取 Root 元素: %v", err)
	}
	defer root.Release()
	rootName, _ := root.Get_CurrentName()
	log.Printf("Root 元素获取成功: %s", rootName)

	log.Println("正在遍历 UI 树并缓存属性...")
	tree := uia.TraverseUIElementTree(ppv, root)
	log.Println("UI 树构建完成")

	// 使用快捷方法查找
	log.Println("查找名为 '详细信息' 的元素...")
	detailBtn := tree.FindByName("详细信息")
	if detailBtn != nil {
		log.Printf("找到元素: %s (Class: %s)", detailBtn.CurrentName, detailBtn.CurrentClassName)
	} else {
		log.Println("未找到名为 '详细信息' 的元素")
	}

	// 使用传统搜索函数
	log.Println("搜索所有类名为 'SelectorButton' 的元素...")
	fnAll := func(elem *uia.Element) bool {
		return elem.CurrentClassName == "SelectorButton"
	}
	foundElements := uia.FindElems(tree, fnAll)
	log.Printf("找到 %d 个 SelectorButton 元素", len(foundElements))
	for i, elem := range foundElements {
		log.Printf("  [%d] %s", i, elem.String())
	}

	log.Println("演示结束")
}
