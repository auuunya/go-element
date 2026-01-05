package main

import (
	"fmt"
	"log"
	"time"

	uia "github.com/auuunya/go-element"
)

func main() {
	// 初始化 COM
	uia.CoInitialize()
	defer uia.CoUninitialize()

	log.Println("正在查找记事本窗口...")
	hwnd, err := uia.GetWindowForString("Notepad", "")
	if err != nil {
		log.Fatal("请先打开记事本 (Notepad.exe)")
	}

	// 创建 UIAutomation 实例
	instance, err := uia.CreateInstance(uia.CLSID_CUIAutomation, uia.IID_IUIAutomation, uia.CLSCTX_INPROC_SERVER|uia.CLSCTX_LOCAL_SERVER|uia.CLSCTX_REMOTE_SERVER)
	if err != nil {
		log.Fatalf("无法创建 UIAutomation 实例: %v", err)
	}
	ppv := uia.NewIUIAutomation(uia.NewIUnKnown(instance))
	defer ppv.Release()

	// 从句柄获取根元素
	root, err := uia.ElementFromHandle(ppv, hwnd)
	if err != nil {
		log.Fatalf("无法获取窗口元素: %v", err)
	}
	defer root.Release()

	log.Println("正在遍历 UI 树（使用缓存优化）...")
	start := time.Now()
	tree := uia.TraverseUIElementTree(ppv, root)
	fmt.Printf("遍历完成，耗时: %v\n", time.Since(start))

	// 查找编辑器
	log.Println("查找文本编辑器...")
	editor := tree.FindByName("文本编辑器") // 简体中文系统下的记事本编辑器名称
	if editor == nil {
		// 备选查找方式：通过类型查找
		for _, child := range tree.Child {
			if child.CurrentControlType == uia.UIA_DocumentControlTypeId || child.CurrentClassName == "Edit" {
				editor = child
				break
			}
		}
	}

	if editor != nil {
		fmt.Printf("找到编辑器: %s (Class: %s)\n", editor.CurrentName, editor.CurrentClassName)

		// 使用 ValuePattern 设置文本
		vp, err := editor.GetValuePattern()
		if err == nil && vp != nil {
			defer vp.Release()
			log.Println("通过 ValuePattern 设置文本...")
			err = vp.SetValue("来自 go-element 的自动输入！\n时间: " + time.Now().Format("2006-01-02 15:04:05"))
			if err != nil {
				log.Printf("设置文本失败: %v", err)
			}
		} else {
			log.Println("该元素不支持 ValuePattern")
		}
	} else {
		log.Println("未找到编辑器元素，请确保记事本处于打开状态且未最小化。")
	}
}
