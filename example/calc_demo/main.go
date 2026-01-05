package main

import (
	"fmt"
	"log"
	"time"

	uia "github.com/auuunya/go-element"
)

func main() {
	uia.CoInitialize()
	defer uia.CoUninitialize()

	log.Println("正在查找计算器窗口...")
	// 计算器类名通常是 ApplicationFrameWindow (Win10+) 或 CalcFrame
	hwnd, err := uia.GetWindowForString("ApplicationFrameWindow", "计算器")
	if err != nil {
		hwnd, err = uia.GetWindowForString("CalcFrame", "计算器")
	}
	if err != nil {
		log.Fatal("请先打开计算器 (calc.exe)")
	}

	instance, _ := uia.CreateInstance(uia.CLSID_CUIAutomation, uia.IID_IUIAutomation, uia.CLSCTX_INPROC_SERVER|uia.CLSCTX_LOCAL_SERVER|uia.CLSCTX_REMOTE_SERVER)
	ppv := uia.NewIUIAutomation(uia.NewIUnKnown(instance))
	defer ppv.Release()

	root, _ := uia.ElementFromHandle(ppv, hwnd)
	defer root.Release()

	log.Println("快速遍历计算器 UI...")
	tree := uia.TraverseUIElementTree(ppv, root)

	// 定义点击按钮的逻辑
	clickButton := func(name string) {
		btn := tree.FindByName(name)
		if btn != nil {
			fmt.Printf("点击按钮: %s\n", name)
			ip, err := btn.GetInvokePattern()
			if err == nil && ip != nil {
				defer ip.Release()
				ip.Invoke()
				time.Sleep(200 * time.Millisecond)
			}
		} else {
			fmt.Printf("未找到按钮: %s\n", name)
		}
	}

	// 执行 1 + 2 =
	clickButton("一")
	clickButton("加")
	clickButton("二")
	clickButton("等于")

	log.Println("演示完成")
}
