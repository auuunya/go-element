package uiautomation

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

type TagSafeArrayBound struct {
	// https://learn.microsoft.com/zh-cn/windows/win32/api/oaidl/ns-oaidl-safearraybound
	CElements uint32
	LLbound   int32
}

type TagSafeArray struct {
	// https://learn.microsoft.com/zh-cn/windows/win32/api/oaidl/ns-oaidl-safearray
	CDims     uint16
	FFeatures uint16
	CbElement uint32
	CLocks    uint32
	PvData    uintptr
	Rgsabound [1]TagSafeArrayBound // COM 中通常是尾随数组
}

type UiaPoint struct {
	// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationcore/ns-uiautomationcore-uiapoint
	X float64
	Y float64
}

type TagDvTargetDevice struct {
	TdSize             uint32
	TdDriverNameOffset uint16
	TdDeviceNameOffset uint16
	TdPortNameOffset   uint16
	TdExtDevmodeOffset uint16
	TdData             [1]byte
}

type UiaRect struct {
	Left   float64
	Top    float64
	Width  float64
	Height float64
}

type TagDispParams struct {
	Rgvarg            uintptr // 指向 VARIANT 数组
	RgdispidNamedArgs uintptr // 指向 DISPID 数组
	Cargs             uint32
	CNamedArgs        uint32
}

type TagExcepInfo struct {
	WCode             uint16
	WReserved         uint16
	BstrSource        uintptr
	BstrDescription   uintptr
	BstrHelpFile      uintptr
	DwHelpContext     uint32
	PvReserved        uintptr
	PFnDeferredFillIn uintptr
	Scode             int32
}

type TagSize struct {
	Cx int32
	Cy int32
}

type TYPEDESC struct {
}

type TagIdlDesc struct {
	DwReserved uintptr // ULONG_PTR
	WIdlFlags  uint16
}

type TagParamDesc struct {
	Pparamdescex uintptr // LPPARAMDESCEX
	WParamFlags  uint16  // PARAMFLAGS
}

type TagTypeDesc struct {
}

type TagElemDesc struct {
	Tdesc     TagTypeDesc
	Paramdesc TagParamDesc
}

type TagFuncDesc struct {
	Memid             int32   // MEMBERID
	LPrgsCode         uintptr // SCODE*
	LPrgelemdescParam uintptr // ELEMDESC*
	Funckind          int32   // FUNCKIND
	Invkind           int32   // INVOKEKIND
	Callconv          int32   // CALLCONV
	CParams           int16
	CParamsOpt        int16
	OVft              int16
	CScodes           int16
	ElemdescFunc      TagElemDesc
	WFuncFlags        uint16
}
