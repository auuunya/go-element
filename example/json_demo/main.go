package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	uia "github.com/auuunya/go-element"
)

func main() {
	// 配置日志
	f, err := os.OpenFile("json_demo.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("无法打开日志文件: %v\n", err)
	} else {
		defer f.Close()
		log.SetOutput(f)
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("--- 启动 JSON 导出演示 ---")

	err = uia.CoInitialize()
	if err != nil {
		log.Fatalf("COM 初始化失败: %v", err)
	}
	defer uia.CoUninitialize()
	log.Println("COM 初始化成功")

	log.Println("正在查找资源管理器窗口...")
	findhwnd, err := uia.GetWindowForString("CabinetWClass", "")
	if err != nil {
		log.Fatal("未找到资源管理器窗口，请先打开一个资源管理器窗口 (Explorer)")
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

	log.Println("正在构建 UI 树并缓存属性...")
	elems := uia.TraverseUIElementTree(ppv, root)
	log.Println("UI 树构建完成，准备序列化...")

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

	log.Printf("UI 树已成功保存至: %s", fileName)
	fmt.Printf("已完成，请检查 %s 文件。\n", fileName)
}
