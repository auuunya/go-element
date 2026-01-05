package main

import (
	"fmt"
	"log"
	"os"
	"time"

	uia "github.com/auuunya/go-element"
)

func main() {
	// 配置日志
	f, err := os.OpenFile("calc_demo.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("无法打开日志文件: %v\n", err)
	} else {
		defer f.Close()
		log.SetOutput(f)
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("--- 启动计算器演示 (极简/极稳版) ---")

	err = uia.CoInitialize()
	if err != nil {
		log.Fatalf("COM 初始化失败: %v", err)
	}
	defer uia.CoUninitialize()

	log.Println("正在寻找计算器主窗口...")
	hwnd, err := uia.GetWindowForString("", "计算器")
	if err != nil {
		hwnd, _ = uia.GetWindowForString("", "Calculator")
	}
	if hwnd == 0 {
		log.Fatal("未找到计算器窗口，请确保它已打开且标题匹配。")
	}
	log.Printf("找到句柄: 0x%X", hwnd)

	instance, err := uia.CreateInstance(uia.CLSID_CUIAutomation, uia.IID_IUIAutomation, uia.CLSCTX_ALL)
	if err != nil {
		log.Fatalf("无法创建 UIAutomation: %v", err)
	}
	ppv := uia.NewIUIAutomation(uia.NewIUnKnown(instance))
	defer ppv.Release()

	// 定位按钮的逻辑
	findButton := func(rootHwnd uintptr, id string, nameFallback string) *uia.IUIAutomationElement {
		root, _ := uia.ElementFromHandle(ppv, rootHwnd)
		if root == nil {
			return nil
		}
		defer root.Release()

		// 1. 尝试通过 ID 查找
		variantID, _ := uia.VariantFromString(id)
		condID, _ := ppv.CreatePropertyCondition(uia.UIA_AutomationIdPropertyId, variantID)
		btn, _ := root.FindFirst(uia.TreeScope_Descendants, condID)
		if btn != nil {
			log.Printf("通过 ID [%s] 找到按钮", id)
			return btn
		}

		// 2. 尝试通过名称查找
		variantName, _ := uia.VariantFromString(nameFallback)
		condName, _ := ppv.CreatePropertyCondition(uia.UIA_NamePropertyId, variantName)
		btn, _ = root.FindFirst(uia.TreeScope_Descendants, condName)
		if btn != nil {
			log.Printf("通过名称 [%s] 找到按钮", nameFallback)
			return btn
		}

		return nil
	}

	click := func(btn *uia.IUIAutomationElement) {
		if btn == nil {
			log.Println("跳过点击：元素为空")
			return
		}
		defer btn.Release()

		unk, err := btn.GetCurrentPattern(uia.UIA_InvokePatternId)
		if err == nil && unk != nil {
			// 正确的类型转换
			ip := uia.NewIUIAutomationInvokePattern(unk)
			defer ip.Release()
			ip.Invoke()
			log.Println("点击成功")
			time.Sleep(300 * time.Millisecond)
		} else {
			log.Println("该按钮不支持 Invoke")
		}
	}

	log.Println("开始执行计算操作...")

	click(findButton(hwnd, "num1Button", "1"))
	click(findButton(hwnd, "num1Button", "一"))

	click(findButton(hwnd, "plusButton", "+"))
	click(findButton(hwnd, "plusButton", "加"))

	click(findButton(hwnd, "num2Button", "2"))
	click(findButton(hwnd, "num2Button", "二"))

	click(findButton(hwnd, "equalButton", "="))
	click(findButton(hwnd, "equalButton", "等于"))

	log.Println("演示结束")
}
