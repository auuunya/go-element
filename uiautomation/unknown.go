package uiautomation

type IUnKnown struct {
	Vtbl *interface{}
}

type IUnKnownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}
