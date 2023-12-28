package uiautomation

type IDropTarget struct {
	vtbl *IUnKnown
}

type IDropTargetVtbl struct {
	IUnKnownVtbl
	DragEnter uintptr
	DragLeave uintptr
	DragOver  uintptr
	Drop      uintptr
}
