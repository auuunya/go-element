package uiautomation

type VARIANT struct {
	VT         TagVarenum
	wReserved1 uint16
	wReserved2 uint16
	wReserved3 uint16
	Val        int64
}

func NewVariant(vt TagVarenum, val int64) VARIANT {
	return VARIANT{
		VT:  vt,
		Val: val,
	}
}

type TagVarenum int

const (
	VT_EMPTY            TagVarenum = 0
	VT_NULL             TagVarenum = 1
	VT_I2               TagVarenum = 2
	VT_I4               TagVarenum = 3
	VT_R4               TagVarenum = 4
	VT_R8               TagVarenum = 5
	VT_CY               TagVarenum = 6
	VT_DATE             TagVarenum = 7
	VT_BSTR             TagVarenum = 8
	VT_DISPATCH         TagVarenum = 9
	VT_ERROR            TagVarenum = 10
	VT_BOOL             TagVarenum = 11
	VT_VARIANT          TagVarenum = 12
	VT_UNKNOWN          TagVarenum = 13
	VT_DECIMAL          TagVarenum = 14
	VT_I1               TagVarenum = 16
	VT_UI1              TagVarenum = 17
	VT_UI2              TagVarenum = 18
	VT_UI4              TagVarenum = 19
	VT_I8               TagVarenum = 20
	VT_UI8              TagVarenum = 21
	VT_INT              TagVarenum = 22
	VT_UINT             TagVarenum = 23
	VT_VOID             TagVarenum = 24
	VT_HRESULT          TagVarenum = 25
	VT_PTR              TagVarenum = 26
	VT_SAFEARRAY        TagVarenum = 27
	VT_CARRAY           TagVarenum = 28
	VT_USERDEFINED      TagVarenum = 29
	VT_LPSTR            TagVarenum = 30
	VT_LPWSTR           TagVarenum = 31
	VT_RECORD           TagVarenum = 36
	VT_INT_PTR          TagVarenum = 37
	VT_UINT_PTR         TagVarenum = 38
	VT_FILETIME         TagVarenum = 64
	VT_BLOB             TagVarenum = 65
	VT_STREAM           TagVarenum = 66
	VT_STORAGE          TagVarenum = 67
	VT_STREAMED_OBJECT  TagVarenum = 68
	VT_STORED_OBJECT    TagVarenum = 69
	VT_BLOB_OBJECT      TagVarenum = 70
	VT_CF               TagVarenum = 71
	VT_CLSID            TagVarenum = 72
	VT_VERSIONED_STREAM TagVarenum = 73
	VT_BSTR_BLOB        TagVarenum = 0xfff
	VT_VECTOR           TagVarenum = 0x1000
	VT_ARRAY            TagVarenum = 0x2000
	VT_BYREF            TagVarenum = 0x4000
	VT_RESERVED         TagVarenum = 0x8000
	VT_ILLEGAL          TagVarenum = 0xffff
	VT_ILLEGALMASKED    TagVarenum = 0xfff
	VT_TYPEMASK         TagVarenum = 0xfff
)
