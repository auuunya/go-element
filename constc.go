package uiautomation

type OrientationType int

const (
	OrientationType_None OrientationType = iota
	OrientationType_Horizontal
	OrientationType_Vertical
)

type TagPoint struct {
	X int32
	Y int32
}

type TagRect struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

// https://learn.microsoft.com/zh-cn/windows/win32/api/oaidl/ns-oaidl-safearray
type SafeArray struct {
	CbElement uint32
	CDims     uint16
	CLocks    uint32
	FFeatures uint16
	PvData    uintptr
	Rgsabound [1]byte
}

type TreeScope uintptr

var (
	TreeScope_None        TreeScope = 0x0
	TreeScope_Element     TreeScope = 0x1
	TreeScope_Children    TreeScope = 0x2
	TreeScope_Descendants TreeScope = 0x4
	TreeScope_Parent      TreeScope = 0x8
	TreeScope_Ancestors   TreeScope = 0x10
	TreeScope_Subtree     TreeScope = TreeScope_Element | TreeScope_Children | TreeScope_Descendants
)

type PROPERTYID uintptr
