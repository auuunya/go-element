package main

import (
	"log"
	"os"
	"time"

	uia "github.com/auuunya/go-element"
)

var logFile string = "notepad_debug.log"

func main() {
	// 配置日志输出到文件
	fp, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic("无法配置日志文件: " + err.Error())
	}
	defer fp.Close()
	log.SetOutput(fp)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("--- 启动记事本自动化测试 ---")

	// 初始化 COM
	err = uia.CoInitialize()
	if err != nil {
		log.Fatalf("COM 初始化失败: %v", err)
	}
	defer uia.CoUninitialize()
	log.Println("COM 初始化成功")

	log.Println("正在查找记事本窗口...")
	hwnd, err := uia.GetWindowForString("Notepad", "")
	if err != nil {
		log.Fatalf("查找窗口失败: %v (请先打开记事本)", err)
	}
	log.Printf("找到记事本窗口，句柄: 0x%X", hwnd)

	// 创建 UIAutomation 实例
	log.Println("创建 UIAutomation 实例...")
	// 使用显式定义的常量组合
	clsctx := uia.CLSCTX_INPROC_SERVER | uia.CLSCTX_LOCAL_SERVER | uia.CLSCTX_REMOTE_SERVER
	instance, err := uia.CreateInstance(uia.CLSID_CUIAutomation, uia.IID_IUIAutomation, clsctx)
	if err != nil {
		log.Fatalf("无法创建 UIAutomation 实例: %v", err)
	}
	log.Printf("实例创建成功: %p", instance)

	unk := uia.NewIUnKnown(instance)
	ppv := uia.NewIUIAutomation(unk)
	defer ppv.Release()

	// 从句柄获取根元素
	log.Println("获取窗口 Root 元素...")
	root, err := uia.ElementFromHandle(ppv, hwnd)
	if err != nil {
		log.Fatalf("无法获取窗口元素: %v", err)
	}
	defer root.Release()

	rootName, _ := root.Get_CurrentName()
	log.Printf("Root 元素获取成功: %s", rootName)

	log.Println("正在构建 UI 树（使用缓存模式）...")
	start := time.Now()
	tree := uia.TraverseUIElementTree(ppv, root)
	log.Printf("UI 树构建完成，耗时: %v", time.Since(start))
	log.Printf("构建的树根节点: %s", tree.String())

	// 查找编辑器
	log.Println("正在定位文本编辑器...")
	// 尝试通过名称查找 (中文系统通常是 "文本编辑器")
	editor := tree.FindByName("文本编辑器")
	if editor == nil {
		log.Println("通过名称 '文本编辑器' 未找到，尝试通过控件类型查找...")
		// 备选方案：查找第一个 Document 类型的控件
		fnSearch := func(e *uia.Element) bool {
			return e.CurrentControlType == uia.UIA_DocumentControlTypeId ||
				e.CurrentClassName == "Edit" ||
				e.CurrentClassName == "NotepadTextBox" ||
				e.CurrentName == "文本编辑器" ||
				e.CurrentName == "Text Editor"
		}
		found := uia.FindElems(tree, fnSearch)
		if len(found) > 0 {
			editor = found[0]
		}
	}

	if editor != nil {
		log.Printf("找到编辑器元素: %s", editor.String())

		// 获取 ValuePattern
		log.Println("尝试获取 ValuePattern...")
		vp, err := editor.GetValuePattern()
		if err != nil {
			log.Printf("获取 ValuePattern 失败: %v", err)
		} else if vp != nil {
			defer vp.Release()
			log.Println("获取 ValuePattern 成功")

			msg := "来自 go-element 的自动输入！\n当前时间: " + time.Now().Format("2006-01-02 15:04:05")
			log.Printf("正在设置文本内容: %q", msg)

			err = vp.SetValue(msg)
			if err != nil {
				log.Printf("SetValue 操作失败: %v", err)
			} else {
				log.Println("文本设置成功！")
			}
		} else {
			log.Println("该元素虽然被找到，但不提供 ValuePattern 接口")
		}
	} else {
		log.Println("【错误】未能在 UI 树中定位到编辑器控件。")
		log.Println("请检查记事本版本（Win11 的记事本 UI 结构与旧版不同）。")
	}

	log.Println("--- 测试结束 ---")
}
