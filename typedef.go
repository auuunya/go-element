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
	CbElement uint32
	CDims     uint16
	CLocks    uint32
	FFeatures uint16
	PvData    uintptr
	Rgsabound []TagSafeArrayBound
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
	Rgvarg            *VARIANT
	RgdispidNamedArgs int32
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
	DwReserved uint32
	WIdlFlags  uint16
}

type TagParamDesc struct {
	Pparamdescex *VARIANT
	WParamFlags  ParamFalg
}
type TagTypeDesc struct {
}

type TagElemDesc struct {
	Tdesc *TagTypeDesc
	Union struct {
		Idldesc   TagIdlDesc
		Paramdesc TagParamDesc
	}
}

type TagFuncDesc struct {
	Memid             int32
	LPrgsCode         int32
	LPrgelemdescParam *TagElemDesc
	Funckind          TagFuncKind
	Invkind           TagInvokeKind
	Callconv          TagCallConv
	CParams           int16
	CParamsOpt        int16
	OVft              int16
	CScodes           int16
	ElemdescFunc      TagElemDesc
	WFuncFlags        uint16
}
