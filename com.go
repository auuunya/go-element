package uiautomation

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

func CoInitialize() error {
	ret, _, _ := procCoInitialize.Call(
		uintptr(0),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func CoUninitialize() {
	procCoUninitialize.Call()
}

func CreateUiautomationInstance() (unsafe.Pointer, error) {
	return createInstance(CLSID_CUIAutomation, IID_IUIAutomation, CLSCTX_INPROC_SERVER|CLSCTX_LOCAL_SERVER|CLSCTX_REMOTE_SERVER)
}

func CreateInstance(clsid, riid *syscall.GUID, clsctx CLSCTX) (unsafe.Pointer, error) {
	return createInstance(clsid, riid, clsctx)
}

func createInstance(clsid, riid *syscall.GUID, clsctx CLSCTX) (unsafe.Pointer, error) {
	var retVal unsafe.Pointer
	ret, _, _ := procCoCreateInstance.Call(
		uintptr(unsafe.Pointer(clsid)),
		uintptr(0),
		uintptr(clsctx),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
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
