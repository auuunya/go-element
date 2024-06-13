package uiautomation

import (
	"errors"
	"log"
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
	procFindWindowExA    = user32.NewProc("FindWindowExA")

	ErrorNotFoundWindow = errors.New("not found window")
)

func StringToCharPtr(str string) *uint8 {
	return stringToCharPtr(str)
}

func GetWindowForString(classname, windowname string) uintptr {
	find := findWindowA(classname, windowname)
	if find == 0 {
		log.Fatal(ErrorNotFoundWindow)
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

func CreateInstance(clsid, riid *syscall.GUID, clsctx CLSCTX) (unsafe.Pointer, error) {
	return createInstance(clsid, riid, clsctx)
}

func FindWindowA(lpclass, lpwindow string) uintptr {
	return findWindowA(lpclass, lpwindow)
}
func FindWindowExA(phwdn, chwdn uintptr, lpclass, lpwindow string) uintptr {
	return findWindowExA(phwdn, chwdn, lpclass, lpwindow)
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

func findWindowExA(phwdn, chwdn uintptr, lpclass, lpwindow string) uintptr {
	lpclassname := uintptr(0)
	lpwindowname := uintptr(0)
	if lpclass != "" {
		lpclassname = uintptr(unsafe.Pointer(stringToCharPtr(lpclass)))
	}
	if lpwindow != "" {
		lpwindowname = uintptr(unsafe.Pointer(stringToCharPtr(lpwindow)))
	}
	ret, _, _ := procFindWindowA.Call(
		phwdn,
		chwdn,
		lpclassname,
		lpwindowname,
	)
	return ret
}
