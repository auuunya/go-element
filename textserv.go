package uiautomation

import (
	"syscall"
	"unsafe"
)

var (
	msftedit                 = syscall.NewLazyDLL("Msftedit.dll")
	procCreateTextServices   = msftedit.NewProc("CreateTextServices")
	procShutdownTextServices = msftedit.NewProc("ShutdownTextServices")
)

func CreateTextServices(unk *IUnKnown, thost *ITextHost) (*IUnKnown, error) {
	return createTextServices(unk, thost)
}

func ShutdownTextServices(unk *IUnKnown) error {
	return shutdownTextServices(unk)
}

func createTextServices(in *IUnKnown, in2 *ITextHost) (*IUnKnown, error) {
	var retVal *IUnKnown
	ret, _, _ := procCreateTextServices.Call(
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(in2)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

func shutdownTextServices(in *IUnKnown) error {
	ret, _, _ := procShutdownTextServices.Call(uintptr(unsafe.Pointer(in)))
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

type IRichEditUiaInformation struct {
	vtbl *IUnKnown
}
type IRichEditUiaInformationVtbl struct {
	IUnKnownVtbl
	GetBoundaryRectangle uintptr
	IsVisible            uintptr
}

func newIRichEditUiaInformation(unk *IUnKnown) *IRichEditUiaInformation {
	return (*IRichEditUiaInformation)(unsafe.Pointer(unk))
}
func NewIRichEditUiaInformation(unk *IUnKnown) *IRichEditUiaInformation {
	return newIRichEditUiaInformation(unk)
}
func (v *IRichEditUiaInformation) GetBoundaryRectangle() (*UiaRect, error) {
	var retVal *UiaRect
	ret, _, _ := syscall.SyscallN(
		(*IRichEditUiaInformationVtbl)(unsafe.Pointer(v.vtbl)).GetBoundaryRectangle,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *IRichEditUiaInformation) IsVisible() bool {
	ret, _, _ := syscall.SyscallN(
		(*IRichEditUiaInformationVtbl)(unsafe.Pointer(v.vtbl)).IsVisible,
		uintptr(unsafe.Pointer(v)),
	)
	return ret != 0
}

type IRicheditUiaOverrides struct {
	vtbl *IUnKnown
}
type IRicheditUiaOverridesVtbl struct {
	IUnKnownVtbl
	GetPropertyOverrideValue uintptr
}
type IRicheditWindowlessAccessibility struct {
	vtbl *IUnKnown
}
type IRicheditWindowlessAccessibilityVtbl struct {
	IUnKnownVtbl
	CreateProvider uintptr
}

type ITextHost struct {
	vtbl *IUnKnown
}
type ITextHostVtbl struct {
	IUnKnownVtbl
	OnTxCharFormatChange   uintptr
	OnTxParaFormatChange   uintptr
	TxActivate             uintptr
	TxClientToScreen       uintptr
	TxCreateCaret          uintptr
	TxDeactivate           uintptr
	TxEnableScrollBar      uintptr
	TxGetAcceleratorPos    uintptr
	TxGetBackStyle         uintptr
	TxGetCharFormat        uintptr
	TxGetClientRect        uintptr
	TxGetDC                uintptr
	TxGetExtent            uintptr
	TxGetMaxLength         uintptr
	TxGetParaFormat        uintptr
	TxGetPasswordChar      uintptr
	TxGetPropertyBits      uintptr
	TxGetScrollBars        uintptr
	TxGetSelectionBarWidth uintptr
	TxGetSysColor          uintptr
	TxGetViewInset         uintptr
	TxImmGetContext        uintptr
	TxImmReleaseContext    uintptr
	TxInvalidateRect       uintptr
	TxKillTimer            uintptr
	TxNotify               uintptr
	TxReleaseDC            uintptr
	TxScreenToClient       uintptr
	TxScrollWindowEx       uintptr
	TxSetCapture           uintptr
	TxSetCaretPos          uintptr
	TxSetCursor            uintptr
	TxSetFocus             uintptr
	TxSetScrollPos         uintptr
	TxSetScrollRange       uintptr
	TxSetTimer             uintptr
	TxShowCaret            uintptr
	TxShowScrollBar        uintptr
	TxViewChange           uintptr
}

func newITextHost(unk *IUnKnown) *ITextHost {
	return (*ITextHost)(unsafe.Pointer(unk))
}
func NewITextHost(unk *IUnKnown) *ITextHost {
	return newITextHost(unk)
}
func (v *ITextHost) TxSetFocus() {
	syscall.SyscallN(
		(*ITextHostVtbl)(unsafe.Pointer(v.vtbl)).TxSetFocus,
		uintptr(unsafe.Pointer(v)),
	)
}
func (v *ITextHost) TxViewChange(in int32) {
	syscall.SyscallN(
		(*ITextHostVtbl)(unsafe.Pointer(v.vtbl)).TxViewChange,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
	)
}

type ITextHost2 struct {
	vtbl *ITextHost
}
type ITextHost2Vtbl struct {
	ITextHostVtbl
	TxDestroyCaret                 uintptr
	TxFreeTextServicesNotification uintptr
	TxGetEastAsianFlags            uintptr
	TxGetEditStyle                 uintptr
	TxGetHorzExtent                uintptr
	TxGetPalette                   uintptr
	TxGetWindow                    uintptr
	TxGetWindowStyles              uintptr
	TxIsDoubleClickPending         uintptr
	TxSetCursor2                   uintptr
	TxSetForegroundWindow          uintptr
	TxShowDropCaret                uintptr
}

type ITextServices struct {
	vtbl *IUnKnown
}

type ITextServicesVtbl struct {
	IUnKnownVtbl
	OnTxInPlaceActivate    uintptr
	OnTxInPlaceDeactivate  uintptr
	OnTxPropertyBitsChange uintptr
	OnTxSetCursor          uintptr
	OnTxUIActivate         uintptr
	OnTxUIDeactivate       uintptr
	TxDraw                 uintptr
	TxGetBaseLinePos       uintptr
	TxGetCachedSize        uintptr
	TxGetCurTargetX        uintptr
	TxGetDropTarget        uintptr
	TxGetHScroll           uintptr
	TxGetNaturalSize       uintptr
	TxGetText              uintptr
	TxGetVScroll           uintptr
	TxQueryHitPoint        uintptr
	TxSendMessage          uintptr
	TxSetText              uintptr
}

func newITextServices(unk *IUnKnown) *ITextServices {
	return (*ITextServices)(unsafe.Pointer(unk))
}
func NewITextServices(unk *IUnKnown) *ITextServices {
	return newITextServices(unk)
}
func (v *ITextServices) OnTxInPlaceActivate(in *TagRect) error {
	ret, _, _ := syscall.SyscallN(
		(*ITextServicesVtbl)(unsafe.Pointer(v.vtbl)).OnTxInPlaceActivate,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *ITextServices) OnTxInPlaceDeactivate(in *TagRect) error {
	ret, _, _ := syscall.SyscallN(
		(*ITextServicesVtbl)(unsafe.Pointer(v.vtbl)).OnTxInPlaceDeactivate,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *ITextServices) OnTxPropertyBitsChange(in, in2 uint32) error {
	ret, _, _ := syscall.SyscallN(
		(*ITextServicesVtbl)(unsafe.Pointer(v.vtbl)).OnTxPropertyBitsChange,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(in2),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

type TextCursor struct {
	DwDrawAspect TagDvAspect
	Lindex       int32
	PvAspect     unsafe.Pointer
	Ptd          *TagDvTargetDevice
	HdcDraw      uintptr
	HicTargetDev uintptr
	LPrcClient   *TagRect
	X            int32
	Y            int32
}

func (v *ITextServices) OnTxSetCursor(opt *TextCursor) error {
	ret, _, _ := syscall.SyscallN(
		(*ITextServicesVtbl)(unsafe.Pointer(v.vtbl)).OnTxSetCursor,
		uintptr(unsafe.Pointer(v)),
		uintptr(opt.DwDrawAspect),
		uintptr(opt.Lindex),
		uintptr(opt.PvAspect),
		uintptr(unsafe.Pointer(opt.Ptd)),
		uintptr(opt.HdcDraw),
		uintptr(opt.HicTargetDev),
		uintptr(unsafe.Pointer(opt.LPrcClient)),
		uintptr(opt.X),
		uintptr(opt.Y),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *ITextServices) OnTxUIActivate() error {
	ret, _, _ := syscall.SyscallN(
		(*ITextServicesVtbl)(unsafe.Pointer(v.vtbl)).OnTxUIActivate,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *ITextServices) OnTxUIDeactivate() error {
	ret, _, _ := syscall.SyscallN(
		(*ITextServicesVtbl)(unsafe.Pointer(v.vtbl)).OnTxUIDeactivate,
		uintptr(unsafe.Pointer(v)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

type TextDraw struct {
	DwDrawAspect TagDvAspect
	Lindex       int32
	PvAspect     unsafe.Pointer
	Ptd          *TagDvTargetDevice
	HdcDraw      uintptr
	HicTargetDev uintptr
	LPrcBounds   *TagRect
	LPrcWBounds  *TagRect
	LPrcUpdate   *TagRect
	PFnContinue  int32
	DwContinue   uint32
	LViewId      int32
}

func (v *ITextServices) TxDraw(opt *TextDraw) error {
	ret, _, _ := syscall.SyscallN(
		(*ITextServicesVtbl)(unsafe.Pointer(v.vtbl)).TxDraw,
		uintptr(unsafe.Pointer(v)),
		uintptr(opt.DwDrawAspect),
		uintptr(opt.Lindex),
		uintptr(opt.PvAspect),
		uintptr(unsafe.Pointer(opt.Ptd)),
		uintptr(opt.HdcDraw),
		uintptr(opt.HicTargetDev),
		uintptr(unsafe.Pointer(opt.LPrcBounds)),
		uintptr(unsafe.Pointer(opt.LPrcWBounds)),
		uintptr(unsafe.Pointer(opt.LPrcUpdate)),
		uintptr(opt.PFnContinue),
		uintptr(opt.DwContinue),
		uintptr(opt.LViewId),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *ITextServices) TxGetBaseLinePos() (int32, error) {
	var retVal int32
	ret, _, _ := syscall.SyscallN(
		(*ITextServicesVtbl)(unsafe.Pointer(v.vtbl)).TxGetBaseLinePos,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return -1, HResult(ret)
	}
	return retVal, nil
}
func (v *ITextServices) TxGetCachedSize(in, in2 uint32) error {
	ret, _, _ := syscall.SyscallN(
		(*ITextServicesVtbl)(unsafe.Pointer(v.vtbl)).TxGetCachedSize,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(in2),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}
func (v *ITextServices) TxGetCurTargetX() (int32, error) {
	var retVal int32
	ret, _, _ := syscall.SyscallN(
		(*ITextServicesVtbl)(unsafe.Pointer(v.vtbl)).TxGetCurTargetX,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return -1, HResult(ret)
	}
	return retVal, nil
}
func (v *ITextServices) TxGetDropTarget() (*IDropTarget, error) {
	var retVal *IDropTarget
	ret, _, _ := syscall.SyscallN(
		(*ITextServicesVtbl)(unsafe.Pointer(v.vtbl)).TxGetDropTarget,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

type TextHScroll struct {
	PlMin     int32
	PlMax     int32
	PlPos     int32
	PlPage    int32
	PfEnabled int32
}

func (v *ITextServices) TxGetHScroll() (*TextHScroll, error) {
	var retVal *TextHScroll
	ret, _, _ := syscall.SyscallN(
		(*ITextServicesVtbl)(unsafe.Pointer(v.vtbl)).TxGetHScroll,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal.PlMin)),
		uintptr(unsafe.Pointer(&retVal.PlMax)),
		uintptr(unsafe.Pointer(&retVal.PlPos)),
		uintptr(unsafe.Pointer(&retVal.PlPage)),
		uintptr(unsafe.Pointer(&retVal.PfEnabled)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

type TextNaturalSize struct {
	DwAspect     TagDvAspect
	HdcDraw      uintptr
	HicTargetDev uintptr
	Ptd          *TagDvTargetDevice
	DwMode       uint32
	// PSizelExtent uintptr
	PWidth  int32
	PHeight int32
}

func (v *ITextServices) TxGetNaturalSize(opt *TextNaturalSize) (int32, int32, error) {
	ret, _, _ := syscall.SyscallN(
		(*ITextServicesVtbl)(unsafe.Pointer(v.vtbl)).TxGetNaturalSize,
		uintptr(unsafe.Pointer(v)),
		uintptr(opt.DwAspect),
		uintptr(opt.HdcDraw),
		uintptr(opt.HicTargetDev),
		uintptr(unsafe.Pointer(opt.Ptd)),
		uintptr(opt.DwMode),
		uintptr(opt.PWidth),
		uintptr(opt.PHeight),
	)
	if ret != 0 {
		return -1, -1, HResult(ret)
	}
	return opt.PWidth, opt.PHeight, nil
}
func (v *ITextServices) TxGetText() (string, error) {
	var bstr uintptr
	ret, _, _ := syscall.SyscallN(
		(*ITextServicesVtbl)(unsafe.Pointer(v.vtbl)).TxGetText,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if ret != 0 {
		return "", HResult(ret)
	}
	if bstr == 0 {
		return "", ErrorBstrPointerNil
	}
	retVal := bstr2str(bstr)
	procSysFreeString.Call(uintptr(bstr))
	return retVal, nil
}

type TextVScroll struct {
	PlMin     int32
	PlMax     int32
	PlPos     int32
	PlPage    int32
	PfEnabled int32
}

func (v *ITextServices) TxGetVScroll() (*TextHScroll, error) {
	var retVal *TextHScroll
	ret, _, _ := syscall.SyscallN(
		(*ITextServicesVtbl)(unsafe.Pointer(v.vtbl)).TxGetVScroll,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal.PlMin)),
		uintptr(unsafe.Pointer(&retVal.PlMax)),
		uintptr(unsafe.Pointer(&retVal.PlPos)),
		uintptr(unsafe.Pointer(&retVal.PlPage)),
		uintptr(unsafe.Pointer(&retVal.PfEnabled)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

type QueryHitPoint struct {
	DwDrawAspect TagDvAspect
	Lindex       int32
	PvAspect     unsafe.Pointer
	Ptd          *TagDvTargetDevice
	HdcDraw      uintptr
	HicTargetDev uintptr
	LPrcClient   *TagRect
	X            int32
	Y            int32
}

func (v *ITextServices) TxQueryHitPoint(opt *QueryHitPoint) (TxtHitResult, error) {
	var retVal TxtHitResult
	ret, _, _ := syscall.SyscallN(
		(*ITextServicesVtbl)(unsafe.Pointer(v.vtbl)).TxQueryHitPoint,
		uintptr(unsafe.Pointer(v)),
		uintptr(opt.DwDrawAspect),
		uintptr(opt.Lindex),
		uintptr(opt.PvAspect),
		uintptr(unsafe.Pointer(opt.Ptd)),
		uintptr(opt.HdcDraw),
		uintptr(opt.HicTargetDev),
		uintptr(unsafe.Pointer(opt.LPrcClient)),
		uintptr(opt.X),
		uintptr(opt.Y),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return -1, HResult(ret)
	}
	return retVal, nil
}

type TextSendMessage struct {
	Msg      uint32
	WParam   uintptr
	LParam   uintptr
	PLResult uintptr
}

func (v *ITextServices) TxSendMessage(opt *TextSendMessage) error {
	ret, _, _ := syscall.SyscallN(
		(*ITextServicesVtbl)(unsafe.Pointer(v.vtbl)).TxSendMessage,
		uintptr(unsafe.Pointer(v)),
		uintptr(opt.Msg),
		uintptr(opt.WParam),
		uintptr(opt.LParam),
		uintptr(opt.PLResult),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

func (v *ITextServices) TxSetText(in string) error {
	retVal, err := syscall.UTF16PtrFromString(in)
	if err != nil {
		return err
	}
	ret, _, _ := syscall.SyscallN(
		(*ITextServicesVtbl)(unsafe.Pointer(v.vtbl)).TxSetText,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(retVal)),
	)
	if ret != 0 {
		return HResult(ret)
	}
	return nil
}

type ITextServices2 struct {
	vtbl *ITextServices
}
type ITextServices2Vtbl struct {
	ITextServicesVtbl
	TxDrawD2D         uintptr
	TxGetNaturalSize2 uintptr
}
