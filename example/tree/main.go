package main

import (
	"unsafe"

	"go-element/uiautomation"

	"go-element/com"
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
