package main

import (
	"fmt"
	"log"
	"os"

	uia "github.com/auuunya/go-element"
)

func main() {
	// 配置日志
	f, err := os.OpenFile("tree_demo.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("无法打开日志文件: %v\n", err)
	} else {
		defer f.Close()
		log.SetOutput(f)
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("--- 启动树遍历演示 ---")

	err = uia.CoInitialize()
	if err != nil {
		log.Fatalf("COM 初始化失败: %v", err)
	}
	defer uia.CoUninitialize()
	log.Println("COM 初始化成功")

	log.Println("正在查找目标窗口 (记事本或 Chrome)...")
	hwnd, err := uia.GetWindowForString("", "计算器")
	if err != nil {
		log.Println("尝试通过 CalcFrame 类名查找...")
		hwnd, err = uia.GetWindowForString("", "Calculator")
	}
	if err != nil {
		log.Fatal("未找到计算器窗口，请先打开计算器 (calc.exe)")
	}
	log.Printf("找到计算器窗口，句柄: 0x%X", hwnd)

	clsctx := uia.CLSCTX_INPROC_SERVER | uia.CLSCTX_LOCAL_SERVER | uia.CLSCTX_REMOTE_SERVER
	instance, err := uia.CreateInstance(uia.CLSID_CUIAutomation, uia.IID_IUIAutomation, clsctx)
	if err != nil {
		log.Fatalf("无法创建 UIAutomation 实例: %v", err)
	}
	log.Printf("UIAutomation 实例创建成功: %p", instance)

	ppv := uia.NewIUIAutomation(uia.NewIUnKnown(instance))
	defer ppv.Release()

	root, err := uia.ElementFromHandle(ppv, hwnd)
	if err != nil {
		log.Fatalf("无法获取 Root 元素: %v", err)
	}
	defer root.Release()
	rootName, _ := root.Get_CurrentName()
	log.Printf("Root 元素获取成功: %s", rootName)

	log.Println("正在构建 UI 树（缓存加速模式）...")
	elems := uia.TraverseUIElementTree(ppv, root)
	log.Println("UI 树构建完成")

	fmt.Println("\n--- UI Element Tree ---")
	uia.TreeString(elems, 0)
	log.Println("树结构已打印到控制台")
}
