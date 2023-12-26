package uiautomation

type IDispatch struct {
	vtbl *IUnKnown
}

type IDispatchVtbl struct {
	IUnKnownVtbl
	GetIDsOfNames    uintptr
	GetTypeInfo      uintptr
	GetTypeInfoCount uintptr
	Invoke           uintptr
}
