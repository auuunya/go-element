package uiautomation

import (
	"errors"
	"runtime"
	"syscall"
	"unsafe"
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")
	ole32  = syscall.NewLazyDLL("Ole32.dll")

	procCoCreateInstance = ole32.NewProc("CoCreateInstance")
	procCoInitialize     = ole32.NewProc("CoInitialize")
	procCoUninitialize   = ole32.NewProc("CoUninitialize")
	procFindWindowW      = user32.NewProc("FindWindowW")
	procFindWindowExW    = user32.NewProc("FindWindowExW")

	ErrorNotFoundWindow = errors.New("not found window")
)

func GetWindowForString(classname, windowname string) (uintptr, error) {
	find := findWindowW(classname, windowname)
	if find == 0 {
		return 0, ErrorNotFoundWindow
	}
	return find, nil
}

func CoInitialize() error {
	runtime.LockOSThread()
	ret, _, _ := procCoInitialize.Call(
		uintptr(0),
	)
	// CoInitialize returns S_FALSE (1) when COM is already initialized
	// on this thread, which is still a successful result.
	if ret != 0 && ret != 1 {
		runtime.UnlockOSThread()
		return HResult(ret)
	}
	return nil
}

func CoUninitialize() {
	procCoUninitialize.Call()
	runtime.UnlockOSThread()
}

func CreateInstance(clsid, riid *syscall.GUID, clsctx CLSCTX) (unsafe.Pointer, error) {
	return createInstance(clsid, riid, clsctx)
}

func FindWindowW(lpclass, lpwindow string) uintptr {
	return findWindowW(lpclass, lpwindow)
}

func FindWindowExW(phwdn, chwdn uintptr, lpclass, lpwindow string) uintptr {
	return findWindowExW(phwdn, chwdn, lpclass, lpwindow)
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

func findWindowW(lpclass, lpwindow string) uintptr {
	var lpclassname, lpwindowname *uint16
	var err error
	if lpclass != "" {
		lpclassname, err = syscall.UTF16PtrFromString(lpclass)
		if err != nil {
			return 0
		}
	}
	if lpwindow != "" {
		lpwindowname, err = syscall.UTF16PtrFromString(lpwindow)
		if err != nil {
			return 0
		}
	}
	ret, _, _ := procFindWindowW.Call(
		uintptr(unsafe.Pointer(lpclassname)),
		uintptr(unsafe.Pointer(lpwindowname)),
	)
	return ret
}

func findWindowExW(phwdn, chwdn uintptr, lpclass, lpwindow string) uintptr {
	var lpclassname, lpwindowname *uint16
	var err error
	if lpclass != "" {
		lpclassname, err = syscall.UTF16PtrFromString(lpclass)
		if err != nil {
			return 0
		}
	}
	if lpwindow != "" {
		lpwindowname, err = syscall.UTF16PtrFromString(lpwindow)
		if err != nil {
			return 0
		}
	}
	ret, _, _ := procFindWindowExW.Call(
		phwdn,
		chwdn,
		uintptr(unsafe.Pointer(lpclassname)),
		uintptr(unsafe.Pointer(lpwindowname)),
	)
	return ret
}
