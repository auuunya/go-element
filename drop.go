package uiautomation

import "unsafe"

type IDropTarget struct {
	vtbl *IUnKnown
}

// TODO:: IDropTarget method
type IDropTargetVtbl struct {
	IUnKnownVtbl
	DragEnter uintptr
	DragLeave uintptr
	DragOver  uintptr
	Drop      uintptr
}

func newIDropTarget(unk *IUnKnown) *IDropTarget {
	return (*IDropTarget)(unsafe.Pointer(unk))
}
func NewIDropTarget(unk *IUnKnown) *IDropTarget {
	return newIDropTarget(unk)
}
func (v *IDropTarget) DragEnter() error {
	return nil
}
func (v *IDropTarget) DragLeave() error {
	return nil
}
func (v *IDropTarget) DragOver() error {
	return nil
}
func (v *IDropTarget) Drop() error {
	return nil
}
