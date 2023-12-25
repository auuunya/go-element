package com

import (
	"errors"
	"syscall"
	"unsafe"
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")
	ole32  = syscall.NewLazyDLL("Ole32.dll")

	procCoCreateInstance = ole32.NewProc("CoCreateInstance")
	procCoInitialize     = ole32.NewProc("CoInitialize")
	procCoUninitialize   = ole32.NewProc("CoUninitialize")
	procFindWindowA      = user32.NewProc("FindWindowA")

	ErrorNotFoundWindow = errors.New("not found window")
)

var (
	CLSID_CUIAutomation = &syscall.GUID{0xff48dba4, 0x60ef, 0x4201, [8]byte{0xaa, 0x87, 0x54, 0x10, 0x3e, 0xef, 0x59, 0x4e}}
	IID_IUIAutomation   = &syscall.GUID{0x30cbe57d, 0xd9d0, 0x452a, [8]byte{0xab, 0x13, 0x7a, 0xc5, 0xac, 0x48, 0x25, 0xee}}
)

// CLSCTX
var (
	CLSCTX_INPROC_SERVER                   = 0x1
	CLSCTX_INPROC_HANDLER                  = 0x2
	CLSCTX_LOCAL_SERVER                    = 0x4
	CLSCTX_INPROC_SERVER16                 = 0x8
	CLSCTX_REMOTE_SERVER                   = 0x10
	CLSCTX_INPROC_HANDLER16                = 0x20
	CLSCTX_RESERVED1                       = 0x40
	CLSCTX_RESERVED2                       = 0x80
	CLSCTX_RESERVED3                       = 0x100
	CLSCTX_RESERVED4                       = 0x200
	CLSCTX_NO_CODE_DOWNLOAD                = 0x400
	CLSCTX_RESERVED5                       = 0x800
	CLSCTX_NO_CUSTOM_MARSHAL               = 0x1000
	CLSCTX_ENABLE_CODE_DOWNLOAD            = 0x2000
	CLSCTX_NO_FAILURE_LOG                  = 0x4000
	CLSCTX_DISABLE_AAA                     = 0x8000
	CLSCTX_ENABLE_AAA                      = 0x10000
	CLSCTX_FROM_DEFAULT_CONTEXT            = 0x20000
	CLSCTX_ACTIVATE_X86_SERVER             = 0x40000
	_CLSCTX_ACTIVATE_32_BIT_SERVER         = 0
	CLSCTX_ACTIVATE_64_BIT_SERVER          = 0x80000
	CLSCTX_ENABLE_CLOAKING                 = 0x100000
	CLSCTX_APPCONTAINER                    = 0x400000
	CLSCTX_ACTIVATE_AAA_AS_IU              = 0x800000
	CLSCTX_RESERVED6                       = 0x1000000
	CLSCTX_ACTIVATE_ARM32_SERVER           = 0x2000000
	_CLSCTX_ALLOW_LOWER_TRUST_REGISTRATION = 0
	CLSCTX_PS_DLL                          = 0x80000000
)

func StringToCharPtr(str string) *uint8 {
	return stringToCharPtr(str)
}

func GetWindowForString(classname, windowname string) uintptr {
	find := findWindowA(classname, windowname)
	if find == 0 {
		panic(ErrorNotFoundWindow)
	}
	return find
}

func CoInitialize() {
	procCoInitialize.Call(
		uintptr(0),
	)
}
func CoUninitialize() {
	procCoUninitialize.Call()
}

func CreateInstance() (instanceVal *interface{}) {
	return createInstance()
}

func createInstance() (instanceVal *interface{}) {
	procCoCreateInstance.Call(
		uintptr(unsafe.Pointer(CLSID_CUIAutomation)),
		uintptr(0),
		uintptr(CLSCTX_INPROC_SERVER|CLSCTX_LOCAL_SERVER|CLSCTX_REMOTE_SERVER),
		uintptr(unsafe.Pointer(IID_IUIAutomation)),
		uintptr(unsafe.Pointer(&instanceVal)),
	)
	return
}
func stringToCharPtr(str string) *uint8 {
	chars := append([]byte(str), 0) // null terminated
	return &chars[0]
}

func findWindowA(lpclass, lpwindow string) uintptr {
	lpclassname := uintptr(0)
	lpwindowname := uintptr(0)
	if lpclass != "" {
		lpclassname = uintptr(unsafe.Pointer(stringToCharPtr(lpclass)))
	}
	if lpwindow != "" {
		lpwindowname = uintptr(unsafe.Pointer(stringToCharPtr(lpwindow)))
	}
	ret, _, _ := procFindWindowA.Call(
		lpclassname,
		lpwindowname,
	)
	return ret
}
