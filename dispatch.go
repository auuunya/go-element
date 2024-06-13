package uiautomation

import (
	"syscall"
	"unsafe"
)

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

func GetIDsOfNames(v *IDispatch, in *syscall.GUID, in2 uint16, in3, in4 uint32) (int32, error) {
	var retVal int32
	ret, _, _ := syscall.SyscallN(
		(*IDispatchVtbl)(unsafe.Pointer(v.vtbl)).GetIDsOfNames,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(in2),
		uintptr(in3),
		uintptr(in4),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return -1, HResult(ret)
	}
	return retVal, nil
}

func GetTypeInfo(v *IDispatch, in, in2 uint32) (*ITypeInfo, error) {
	var retVal *ITypeInfo
	ret, _, _ := syscall.SyscallN(
		(*IDispatchVtbl)(unsafe.Pointer(v.vtbl)).GetTypeInfo,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(in2),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

func GetTypeInfoCount(v *IDispatch) (uint32, error) {
	var retVal uint32
	ret, _, _ := syscall.SyscallN(
		(*IDispatchVtbl)(unsafe.Pointer(v.vtbl)).GetTypeInfoCount,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return 0, HResult(ret)
	}
	return retVal, nil
}

type InvokeReq struct {
	DispIdMember int32
	Riid         syscall.GUID
	LcId         uint32
	WFlags       uint16
	PDispParams  *TagDispParams
}
type InvokeResp struct {
	PDispParams *TagDispParams
	PVarResult  *VARIANT
	PExcepInfo  *TagExcepInfo
	PuArgErr    uint32
}

func Invoke(v *IDispatch, opt *InvokeReq) (*InvokeResp, error) {
	var retVal *InvokeResp
	ret, _, _ := syscall.SyscallN(
		(*IDispatchVtbl)(unsafe.Pointer(v.vtbl)).Invoke,
		uintptr(unsafe.Pointer(v)),
		uintptr(opt.DispIdMember),
		uintptr(unsafe.Pointer(&opt.Riid)),
		uintptr(opt.LcId),
		uintptr(opt.WFlags),
		uintptr(unsafe.Pointer(opt.PDispParams)),
		uintptr(unsafe.Pointer(&retVal.PDispParams)),
		uintptr(unsafe.Pointer(&retVal.PVarResult)),
		uintptr(unsafe.Pointer(&retVal.PExcepInfo)),
		uintptr(unsafe.Pointer(&retVal.PuArgErr)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}

type ITypeInfo struct {
	vtbl *IUnKnown
}

type ITypeInfoVtbl struct {
	IUnKnownVtbl
	AddressOfMember      uintptr
	CreateInstance       uintptr
	GetContainingTypeLib uintptr
	GetDllEntry          uintptr
	GetDocumentation     uintptr
	GetFuncDesc          uintptr
	GetIDsOfNames        uintptr
	GetImplTypeFlags     uintptr
	GetMops              uintptr
	GetNames             uintptr
	GetRefTypeInfo       uintptr
	GetRefTypeOfImplType uintptr
	GetTypeAttr          uintptr
	GetTypeComp          uintptr
	GetVarDesc           uintptr
	Invoke               uintptr
	ReleaseFuncDesc      uintptr
	ReleaseTypeAttr      uintptr
	ReleaseVarDesc       uintptr
}

// TODO:: ITypeInfo method
func newITypeInfo(unk *IUnKnown) *ITypeInfo {
	return (*ITypeInfo)(unsafe.Pointer(unk))
}
func NewITypeInfo(unk *IUnKnown) *ITypeInfo {
	return newITypeInfo(unk)
}
func (v *ITypeInfo) AddressOfMember(in int32, in2 TagInvokeKind) (unsafe.Pointer, error) {
	var retVal unsafe.Pointer
	ret, _, _ := syscall.SyscallN(
		(*ITypeInfoVtbl)(unsafe.Pointer(v.vtbl)).AddressOfMember,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(in2),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *ITypeInfo) CreateInstance(in *IUnKnown, in2 *syscall.GUID) (unsafe.Pointer, error) {
	var retVal unsafe.Pointer
	ret, _, _ := syscall.SyscallN(
		(*ITypeInfoVtbl)(unsafe.Pointer(v.vtbl)).CreateInstance,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(in)),
		uintptr(unsafe.Pointer(in2)),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *ITypeInfo) GetContainingTypeLib() (*ITypeLib, int32, error) {
	var retVal *ITypeLib
	var retVal2 int32
	ret, _, _ := syscall.SyscallN(
		(*ITypeInfoVtbl)(unsafe.Pointer(v.vtbl)).GetContainingTypeLib,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retVal)),
		uintptr(unsafe.Pointer(&retVal2)),
	)
	if ret != 0 {
		return nil, -1, HResult(ret)
	}
	return retVal, retVal2, nil
}
func (v *ITypeInfo) GetDllEntry(in int32, in2 TagInvokeKind) (string, string, uint16, error) {
	var bstr uintptr
	var bstr2 uintptr
	var retVal3 uint16
	ret, _, _ := syscall.SyscallN(
		(*ITypeInfoVtbl)(unsafe.Pointer(v.vtbl)).GetDllEntry,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(in2),
		uintptr(unsafe.Pointer(&bstr)),
		uintptr(unsafe.Pointer(&bstr2)),
		uintptr(unsafe.Pointer(&retVal3)),
	)
	if ret != 0 {
		return "", "", 0, HResult(ret)
	}
	var retVal string
	var retVal2 string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	if bstr2 != 0 {
		retVal2 = bstr2str(bstr2)
		procSysAllocString.Call(uintptr(bstr2))
	}
	return retVal, retVal2, retVal3, nil
}
func (v *ITypeInfo) GetDocumentation(in int32) (string, string, uint32, string, error) {
	var bstr uintptr
	var bstr2 uintptr
	var retVal3 uint32
	var bstr3 uintptr
	ret, _, _ := syscall.SyscallN(
		(*ITypeInfoVtbl)(unsafe.Pointer(v.vtbl)).GetDocumentation,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(unsafe.Pointer(&bstr)),
		uintptr(unsafe.Pointer(&bstr2)),
		uintptr(unsafe.Pointer(&retVal3)),
		uintptr(unsafe.Pointer(&bstr3)),
	)
	if ret != 0 {
		return "", "", 0, "", HResult(ret)
	}
	var retVal string
	var retVal2 string
	var retVal4 string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	if bstr2 != 0 {
		retVal2 = bstr2str(bstr2)
		procSysAllocString.Call(uintptr(bstr2))
	}
	if bstr3 != 0 {
		retVal4 = bstr2str(bstr2)
		procSysAllocString.Call(uintptr(bstr3))
	}
	return retVal, retVal2, retVal3, retVal4, nil
}
func (v *ITypeInfo) GetFuncDesc(in uint32) (*TagFuncDesc, error) {
	var retVal *TagFuncDesc
	ret, _, _ := syscall.SyscallN(
		(*ITypeInfoVtbl)(unsafe.Pointer(v.vtbl)).GetFuncDesc,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return nil, HResult(ret)
	}
	return retVal, nil
}
func (v *ITypeInfo) GetIDsOfNames() error {
	return nil
}
func (v *ITypeInfo) GetImplTypeFlags(in uint32) (int32, error) {
	var retVal int32
	ret, _, _ := syscall.SyscallN(
		(*ITypeInfoVtbl)(unsafe.Pointer(v.vtbl)).GetImplTypeFlags,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(unsafe.Pointer(&retVal)),
	)
	if ret != 0 {
		return -1, HResult(ret)
	}
	return retVal, nil
}
func (v *ITypeInfo) GetMops(in int32) (string, error) {
	var bstr uintptr
	ret, _, _ := syscall.SyscallN(
		(*ITypeInfoVtbl)(unsafe.Pointer(v.vtbl)).GetMops,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(unsafe.Pointer(&bstr)),
	)
	if ret != 0 {
		return "", HResult(ret)
	}
	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, nil
}
func (v *ITypeInfo) GetNames(in int32, in2 uint32) (string, uint32, error) {
	var bstr uintptr
	var retVal2 uint32
	ret, _, _ := syscall.SyscallN(
		(*ITypeInfoVtbl)(unsafe.Pointer(v.vtbl)).GetNames,
		uintptr(unsafe.Pointer(v)),
		uintptr(in),
		uintptr(unsafe.Pointer(&bstr)),
		uintptr(in2),
		uintptr(unsafe.Pointer(&retVal2)),
	)
	if ret != 0 {
		return "", 0, HResult(ret)
	}
	var retVal string
	if bstr != 0 {
		retVal = bstr2str(bstr)
		procSysAllocString.Call(uintptr(bstr))
	}
	return retVal, retVal2, nil
}

func (v *ITypeInfo) GetRefTypeInfo() error {
	return nil
}
func (v *ITypeInfo) GetRefTypeOfImplType() error {
	return nil
}
func (v *ITypeInfo) GetTypeAttr() error {
	return nil
}
func (v *ITypeInfo) GetTypeComp() error {
	return nil
}
func (v *ITypeInfo) GetVarDesc() error {
	return nil
}
func (v *ITypeInfo) Invoke() error {
	return nil
}
func (v *ITypeInfo) ReleaseFuncDesc() error {
	return nil
}
func (v *ITypeInfo) ReleaseTypeAttr() error {
	return nil
}
func (v *ITypeInfo) ReleaseVarDesc() error {
	return nil
}

type ITypeLib struct {
	vtbl *IUnKnown
}
type ITypeLibVtbl struct {
	IUnKnownVtbl
}

// TODO:: ITypeLib method
func newITypeLib(unk *IUnKnown) *ITypeLib {
	return (*ITypeLib)(unsafe.Pointer(unk))
}
func NewITypeLib(unk *IUnKnown) *ITypeLib {
	return newITypeLib(unk)
}
func (v *ITypeLib) FindName() error {
	return nil
}
func (v *ITypeLib) GetDocumentation() error {
	return nil
}
func (v *ITypeLib) GetLibAttr() error {
	return nil
}
func (v *ITypeLib) GetTypeComp() error {
	return nil
}
func (v *ITypeLib) GetTypeInfo() error {
	return nil
}
func (v *ITypeLib) GetTypeInfoCount() error {
	return nil
}
func (v *ITypeLib) GetTypeInfoOfGuid() error {
	return nil
}
func (v *ITypeLib) GetTypeInfoType() error {
	return nil
}
func (v *ITypeLib) IsName() error {
	return nil
}
func (v *ITypeLib) ReleaseTLibAttr() error {
	return nil
}
