package main

import (
	"unsafe"

	"github.com/auuunya/go-element/uiautomation"

	"github.com/auuunya/go-element/com"
)

func main() {
	com.CoInitialize()
	defer com.CoUninitialize()
	findhwnd := com.GetWindowForString("Notepad", "")
	instance := com.CreateInstance()
	ppv := (*uiautomation.IUIAutomation)(unsafe.Pointer(instance))
	// new
	root := uiautomation.ElementFromHandle(ppv, findhwnd)
	elems := uiautomation.TraverseUIElementTree(ppv, root)
	uiautomation.TreeString(elems, 0)
}
