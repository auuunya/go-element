package uiautomation

type VARIANT struct {
	VT         uint16
	wReserved1 uint16
	wReserved2 uint16
	wReserved3 uint16
	Val        uintptr
}

func NewVariant(vt uint16, val uintptr) VARIANT {
	return VARIANT{
		VT:  vt,
		Val: val,
	}
}
