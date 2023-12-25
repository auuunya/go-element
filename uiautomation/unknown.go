package uiautomation

type IUnKnown struct {
	Vtbl *IUnKnownVtbl
}

type IUnKnownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}
