package uiautomation

type PropertyConditionFlags int

const (
	PropertyConditionFlags_None           PropertyConditionFlags = 0
	PropertyConditionFlags_IgnoreCase     PropertyConditionFlags = 0x1
	PropertyConditionFlags_MatchSubstring PropertyConditionFlags = 0x2
)

type TreeScope int

var (
	// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationclient/ne-uiautomationclient-treescope
	TreeScope_None        TreeScope = 0x0
	TreeScope_Element     TreeScope = 0x1
	TreeScope_Children    TreeScope = 0x2
	TreeScope_Descendants TreeScope = 0x4
	TreeScope_Parent      TreeScope = 0x8
	TreeScope_Ancestors   TreeScope = 0x10
	TreeScope_Subtree     TreeScope = TreeScope_Element | TreeScope_Children | TreeScope_Descendants
)

type SynchronizedInputType int

const (
	// https://learn.microsoft.com/zh-cn/windows/win32/api/uiautomationcore/ne-uiautomationcore-synchronizedinputtype
	SynchronizedInputType_KeyUp          SynchronizedInputType = 0x1
	SynchronizedInputType_KeyDown        SynchronizedInputType = 0x2
	SynchronizedInputType_LeftMouseUp    SynchronizedInputType = 0x4
	SynchronizedInputType_LeftMouseDown  SynchronizedInputType = 0x8
	SynchronizedInputType_RightMouseUp   SynchronizedInputType = 0x10
	SynchronizedInputType_RightMouseDown SynchronizedInputType = 0x20
)

type TagDvAspect int

const (
	DVASPECT_CONTENT   TagDvAspect = 1
	DVASPECT_THUMBNAIL TagDvAspect = 2
	DVASPECT_ICON      TagDvAspect = 4
	DVASPECT_DOCPRINT  TagDvAspect = 8
)

type CLSCTX int

var (
	// https://learn.microsoft.com/zh-cn/windows/win32/api/wtypesbase/ne-wtypesbase-clsctx
	CLSCTX_INPROC_SERVER                   CLSCTX = 0x1
	CLSCTX_INPROC_HANDLER                  CLSCTX = 0x2
	CLSCTX_LOCAL_SERVER                    CLSCTX = 0x4
	CLSCTX_INPROC_SERVER16                 CLSCTX = 0x8
	CLSCTX_REMOTE_SERVER                   CLSCTX = 0x10
	CLSCTX_INPROC_HANDLER16                CLSCTX = 0x20
	CLSCTX_RESERVED1                       CLSCTX = 0x40
	CLSCTX_RESERVED2                       CLSCTX = 0x80
	CLSCTX_RESERVED3                       CLSCTX = 0x100
	CLSCTX_RESERVED4                       CLSCTX = 0x200
	CLSCTX_NO_CODE_DOWNLOAD                CLSCTX = 0x400
	CLSCTX_RESERVED5                       CLSCTX = 0x800
	CLSCTX_NO_CUSTOM_MARSHAL               CLSCTX = 0x1000
	CLSCTX_ENABLE_CODE_DOWNLOAD            CLSCTX = 0x2000
	CLSCTX_NO_FAILURE_LOG                  CLSCTX = 0x4000
	CLSCTX_DISABLE_AAA                     CLSCTX = 0x8000
	CLSCTX_ENABLE_AAA                      CLSCTX = 0x10000
	CLSCTX_FROM_DEFAULT_CONTEXT            CLSCTX = 0x20000
	CLSCTX_ACTIVATE_X86_SERVER             CLSCTX = 0x40000
	_CLSCTX_ACTIVATE_32_BIT_SERVER         CLSCTX = 0
	CLSCTX_ACTIVATE_64_BIT_SERVER          CLSCTX = 0x80000
	CLSCTX_ENABLE_CLOAKING                 CLSCTX = 0x100000
	CLSCTX_APPCONTAINER                    CLSCTX = 0x400000
	CLSCTX_ACTIVATE_AAA_AS_IU              CLSCTX = 0x800000
	CLSCTX_RESERVED6                       CLSCTX = 0x1000000
	CLSCTX_ACTIVATE_ARM32_SERVER           CLSCTX = 0x2000000
	_CLSCTX_ALLOW_LOWER_TRUST_REGISTRATION CLSCTX = 0
	CLSCTX_PS_DLL                          CLSCTX = 0x80000000
)

type SelFlag int

var (
	// https://learn.microsoft.com/zh-cn/windows/win32/winauto/selflag
	SELFLAG_NONE            SelFlag = 0
	SELFLAG_TAKEFOCUS       SelFlag = 0x1
	SELFLAG_TAKESELECTION   SelFlag = 0x2
	SELFLAG_EXTENDSELECTION SelFlag = 0x4
	SELFLAG_ADDSELECTION    SelFlag = 0x8
	SELFLAG_REMOVESELECTION SelFlag = 0x10
)

type TagInvokeKind int

const (
	INVOKE_FUNC           TagInvokeKind = 1
	INVOKE_PROPERTYGET    TagInvokeKind = 2
	INVOKE_PROPERTYPUT    TagInvokeKind = 4
	INVOKE_PROPERTYPUTREF TagInvokeKind = 8
)

type ParamFalg uint16

const (
	// https://learn.microsoft.com/zh-cn/previous-versions/windows/desktop/automat/paramflags
	PARAMFLAG_NONE         ParamFalg = 0
	PARAMFLAG_FIN          ParamFalg = 0x1
	PARAMFLAG_FOUT         ParamFalg = 0x2
	PARAMFLAG_FLCID        ParamFalg = 0x4
	PARAMFLAG_FRETVAL      ParamFalg = 0x8
	PARAMFLAG_FOPT         ParamFalg = 0x10
	PARAMFLAG_FHASDEFAULT  ParamFalg = 0x20
	PARAMFLAG_FHASCUSTDATA ParamFalg = 0x40
)
