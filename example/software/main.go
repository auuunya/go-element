package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	msi                    = syscall.NewLazyDLL("Msi.dll")
	procMsiEnumProductsW   = msi.NewProc("MsiEnumProductsW")
	procMsiGetProductInfoW = msi.NewProc("MsiGetProductInfoW")
)

func main() {
	var productIndex uint32 = 0
	ptr, err := windows.UTF16PtrFromString("ProductName")
	if err != nil {
		return
	}
	for {
		var lpProductBuf [39]uint16
		ret, _, _ := procMsiEnumProductsW.Call(
			uintptr(productIndex),
			uintptr(unsafe.Pointer(&lpProductBuf[0])),
		)
		if ret == 259 { // ERROR_NO_MORE_ITEMS
			break
		} else if ret != 0 { // other errors
			productIndex++
			continue
		}
		productIndex++
		var valueSize uint32 = 256
		var valueUTF16 = make([]uint16, valueSize)

		ret2, _, _ := procMsiGetProductInfoW.Call(
			uintptr(unsafe.Pointer(&lpProductBuf)),
			uintptr(unsafe.Pointer(ptr)),
			uintptr(unsafe.Pointer(&valueUTF16[0])),
			uintptr(unsafe.Pointer(&valueSize)),
		)
		if ret2 != 0 {
			return
		}
		valueUTF16 = valueUTF16[:valueSize]
		value := syscall.UTF16ToString(valueUTF16)
		fmt.Printf("product name: %#v\n", value)
	}
}
