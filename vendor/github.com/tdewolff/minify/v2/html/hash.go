package html

// generated by hasher -type=Hash -file=hash.go; DO NOT EDIT, except for adding more constants to the list and rerun go generate

// uses github.com/tdewolff/hasher
//go:generate hasher -type=Hash -file=hash.go

// Hash defines perfect hashes for a predefined list of strings
type Hash uint32

// Unique hash definitions to be used instead of strings
const (
	A               Hash = 0x1     // a
	Abbr            Hash = 0x3b804 // abbr
	About           Hash = 0x5     // about
	Accept          Hash = 0x1106  // accept
	Accept_Charset  Hash = 0x110e  // accept-charset
	Acronym         Hash = 0x4a07  // acronym
	Action          Hash = 0x21d06 // action
	Address         Hash = 0x7807  // address
	Align           Hash = 0x35b05 // align
	Alink           Hash = 0x3a405 // alink
	Allowfullscreen Hash = 0x2e10f // allowfullscreen
	Amp_Boilerplate Hash = 0x7f0f  // amp-boilerplate
	Applet          Hash = 0xd706  // applet
	Area            Hash = 0x2fd04 // area
	Article         Hash = 0x2707  // article
	Aside           Hash = 0x5b05  // aside
	Async           Hash = 0x8e05  // async
	Audio           Hash = 0x9605  // audio
	Autofocus       Hash = 0xcc09  // autofocus
	Autoplay        Hash = 0x10c08 // autoplay
	Axis            Hash = 0x11404 // axis
	B               Hash = 0x101   // b
	Background      Hash = 0x300a  // background
	Base            Hash = 0x17804 // base
	Basefont        Hash = 0x17808 // basefont
	Bb              Hash = 0x3b902 // bb
	Bdi             Hash = 0x18403 // bdi
	Bdo             Hash = 0x35303 // bdo
	Bgcolor         Hash = 0x12a07 // bgcolor
	Big             Hash = 0x13103 // big
	Blockquote      Hash = 0x1340a // blockquote
	Body            Hash = 0xd04   // body
	Br              Hash = 0x36102 // br
	Button          Hash = 0x13e06 // button
	Canvas          Hash = 0x5706  // canvas
	Caption         Hash = 0x1fe07 // caption
	Center          Hash = 0xb706  // center
	Charset         Hash = 0x1807  // charset
	Checked         Hash = 0x19707 // checked
	Cite            Hash = 0x9204  // cite
	Class           Hash = 0x15105 // class
	Classid         Hash = 0x15107 // classid
	Clear           Hash = 0x2b05  // clear
	Code            Hash = 0x17404 // code
	Codebase        Hash = 0x17408 // codebase
	Codetype        Hash = 0x18808 // codetype
	Col             Hash = 0x12c03 // col
	Colgroup        Hash = 0x1af08 // colgroup
	Color           Hash = 0x12c05 // color
	Cols            Hash = 0x1c904 // cols
	Colspan         Hash = 0x1c907 // colspan
	Compact         Hash = 0x1d707 // compact
	Content         Hash = 0x27b07 // content
	Controls        Hash = 0x1e708 // controls
	Data            Hash = 0x1f04  // data
	Datalist        Hash = 0x1f08  // datalist
	Datatype        Hash = 0xac08  // datatype
	Dd              Hash = 0x7902  // dd
	Declare         Hash = 0x5e07  // declare
	Default         Hash = 0xeb07  // default
	DefaultChecked  Hash = 0x2270e // defaultChecked
	DefaultMuted    Hash = 0xeb0c  // defaultMuted
	DefaultSelected Hash = 0xf60f  // defaultSelected
	Defer           Hash = 0x10405 // defer
	Del             Hash = 0x37903 // del
	Details         Hash = 0x15707 // details
	Dfn             Hash = 0x16403 // dfn
	Dialog          Hash = 0xc606  // dialog
	Dir             Hash = 0x18503 // dir
	Disabled        Hash = 0x19d08 // disabled
	Div             Hash = 0x1a403 // div
	Dl              Hash = 0x1e502 // dl
	Dt              Hash = 0x21702 // dt
	Em              Hash = 0x4302  // em
	Embed           Hash = 0x37505 // embed
	Enabled         Hash = 0x26307 // enabled
	Enctype         Hash = 0x2a207 // enctype
	Face            Hash = 0xb504  // face
	Fieldset        Hash = 0x1f308 // fieldset
	Figcaption      Hash = 0x1fb0a // figcaption
	Figure          Hash = 0x20c06 // figure
	Font            Hash = 0x17c04 // font
	Footer          Hash = 0xa006  // footer
	For             Hash = 0x21903 // for
	Form            Hash = 0x21904 // form
	Formaction      Hash = 0x2190a // formaction
	Formnovalidate  Hash = 0x2350e // formnovalidate
	Frame           Hash = 0x14505 // frame
	Frameborder     Hash = 0x2830b // frameborder
	Frameset        Hash = 0x14508 // frameset
	H1              Hash = 0x2d002 // h1
	H2              Hash = 0x24302 // h2
	H3              Hash = 0x24502 // h3
	H4              Hash = 0x24702 // h4
	H5              Hash = 0x24902 // h5
	H6              Hash = 0x24b02 // h6
	Head            Hash = 0x2c204 // head
	Header          Hash = 0x2c206 // header
	Hgroup          Hash = 0x24d06 // hgroup
	Hidden          Hash = 0x25f06 // hidden
	Hr              Hash = 0x16802 // hr
	Href            Hash = 0x16804 // href
	Hreflang        Hash = 0x16808 // hreflang
	Html            Hash = 0x26a04 // html
	Http_Equiv      Hash = 0x26e0a // http-equiv
	I               Hash = 0x2401  // i
	Icon            Hash = 0x27a04 // icon
	Id              Hash = 0x5d02  // id
	Iframe          Hash = 0x28206 // iframe
	Image           Hash = 0x28e05 // image
	Img             Hash = 0x29303 // img
	Inert           Hash = 0x5205  // inert
	Inlist          Hash = 0x29606 // inlist
	Input           Hash = 0x2a905 // input
	Ins             Hash = 0x2ae03 // ins
	Ismap           Hash = 0x11605 // ismap
	Itemscope       Hash = 0xe209  // itemscope
	Kbd             Hash = 0x18303 // kbd
	Keygen          Hash = 0x29e06 // keygen
	Label           Hash = 0x6505  // label
	Lang            Hash = 0x16c04 // lang
	Language        Hash = 0x16c08 // language
	Legend          Hash = 0x31706 // legend
	Li              Hash = 0x2302  // li
	Link            Hash = 0x3a504 // link
	Longdesc        Hash = 0x6908  // longdesc
	Main            Hash = 0x5004  // main
	Manifest        Hash = 0x11e08 // manifest
	Map             Hash = 0xd603  // map
	Mark            Hash = 0x2b404 // mark
	Marquee         Hash = 0x2b807 // marquee
	Math            Hash = 0x2bf04 // math
	Max             Hash = 0x2c803 // max
	Maxlength       Hash = 0x2c809 // maxlength
	Media           Hash = 0xc405  // media
	Menu            Hash = 0xde04  // menu
	Menuitem        Hash = 0xde08  // menuitem
	Meta            Hash = 0x2d204 // meta
	Meter           Hash = 0x30605 // meter
	Method          Hash = 0x30b06 // method
	Multiple        Hash = 0x31108 // multiple
	Muted           Hash = 0x31d05 // muted
	Name            Hash = 0xc204  // name
	Nav             Hash = 0x35803 // nav
	Nobr            Hash = 0x35f04 // nobr
	Noembed         Hash = 0x37307 // noembed
	Noframes        Hash = 0x14308 // noframes
	Nohref          Hash = 0x16606 // nohref
	Noresize        Hash = 0x1cf08 // noresize
	Noscript        Hash = 0x20408 // noscript
	Noshade         Hash = 0x22207 // noshade
	Novalidate      Hash = 0x2390a // novalidate
	Nowrap          Hash = 0x2ef06 // nowrap
	Object          Hash = 0x9a06  // object
	Ol              Hash = 0x7202  // ol
	Open            Hash = 0x35504 // open
	Optgroup        Hash = 0x39908 // optgroup
	Option          Hash = 0x32206 // option
	Output          Hash = 0x206   // output
	P               Hash = 0x501   // p
	Param           Hash = 0x11a05 // param
	Pauseonexit     Hash = 0x1b60b // pauseonexit
	Picture         Hash = 0x25207 // picture
	Plaintext       Hash = 0x2f409 // plaintext
	Portal          Hash = 0x3a006 // portal
	Poster          Hash = 0x38c06 // poster
	Pre             Hash = 0x38503 // pre
	Prefix          Hash = 0x38506 // prefix
	Profile         Hash = 0x32807 // profile
	Progress        Hash = 0x32f08 // progress
	Property        Hash = 0x33e08 // property
	Q               Hash = 0x13901 // q
	Rb              Hash = 0x2f02  // rb
	Readonly        Hash = 0x2fe08 // readonly
	Rel             Hash = 0x6303  // rel
	Required        Hash = 0x21008 // required
	Resource        Hash = 0x25708 // resource
	Rev             Hash = 0xa503  // rev
	Reversed        Hash = 0xa508  // reversed
	Rows            Hash = 0xbc04  // rows
	Rowspan         Hash = 0xbc07  // rowspan
	Rp              Hash = 0x8802  // rp
	Rt              Hash = 0x2802  // rt
	Rtc             Hash = 0x5503  // rtc
	Ruby            Hash = 0x10804 // ruby
	Rules           Hash = 0x36205 // rules
	S               Hash = 0x1c01  // s
	Samp            Hash = 0x7e04  // samp
	Scope           Hash = 0xe605  // scope
	Scoped          Hash = 0xe606  // scoped
	Script          Hash = 0x20606 // script
	Scrolling       Hash = 0x6f09  // scrolling
	Seamless        Hash = 0x36608 // seamless
	Section         Hash = 0x36d07 // section
	Select          Hash = 0x15d06 // select
	Selected        Hash = 0x15d08 // selected
	Shape           Hash = 0x1ee05 // shape
	Size            Hash = 0x1d304 // size
	Slot            Hash = 0x2b004 // slot
	Small           Hash = 0x2df05 // small
	Sortable        Hash = 0x33608 // sortable
	Source          Hash = 0x25906 // source
	Span            Hash = 0xbf04  // span
	Src             Hash = 0x34603 // src
	Srcset          Hash = 0x34606 // srcset
	Start           Hash = 0x2505  // start
	Strike          Hash = 0x29a06 // strike
	Strong          Hash = 0x12406 // strong
	Style           Hash = 0x34c05 // style
	Sub             Hash = 0x35103 // sub
	Summary         Hash = 0x37c07 // summary
	Sup             Hash = 0x38303 // sup
	Svg             Hash = 0x39203 // svg
	Tabindex        Hash = 0x2d408 // tabindex
	Table           Hash = 0x33905 // table
	Target          Hash = 0x706   // target
	Tbody           Hash = 0xc05   // tbody
	Td              Hash = 0x1e02  // td
	Template        Hash = 0x4208  // template
	Text            Hash = 0x2f904 // text
	Textarea        Hash = 0x2f908 // textarea
	Tfoot           Hash = 0x9f05  // tfoot
	Th              Hash = 0x2c102 // th
	Thead           Hash = 0x2c105 // thead
	Time            Hash = 0xdc04  // time
	Title           Hash = 0x14c05 // title
	Tr              Hash = 0x12502 // tr
	Track           Hash = 0x17f05 // track
	Translate       Hash = 0x1c009 // translate
	Truespeed       Hash = 0x1dd09 // truespeed
	Tt              Hash = 0x14002 // tt
	Type            Hash = 0xb004  // type
	Typemustmatch   Hash = 0x18c0d // typemustmatch
	Typeof          Hash = 0xb006  // typeof
	U               Hash = 0x301   // u
	Ul              Hash = 0xef02  // ul
	Undeterminate   Hash = 0x370d  // undeterminate
	Usemap          Hash = 0xd306  // usemap
	Valign          Hash = 0x35a06 // valign
	Value           Hash = 0x1a605 // value
	Valuetype       Hash = 0x1a609 // valuetype
	Var             Hash = 0x27703 // var
	Video           Hash = 0x39505 // video
	Visible         Hash = 0x3a907 // visible
	Vlink           Hash = 0x3b005 // vlink
	Vocab           Hash = 0x3b505 // vocab
	Wbr             Hash = 0x3bc03 // wbr
	Xmlns           Hash = 0x2db05 // xmlns
	Xmp             Hash = 0x38a03 // xmp
)

// String returns the hash' name.
func (i Hash) String() string {
	start := uint32(i >> 8)
	n := uint32(i & 0xff)
	if start+n > uint32(len(_Hash_text)) {
		return ""
	}
	return _Hash_text[start : start+n]
}

// ToHash returns the hash whose name is s. It returns zero if there is no
// such hash. It is case sensitive.
func ToHash(s []byte) Hash {
	if len(s) == 0 || len(s) > _Hash_maxLen {
		return 0
	}
	h := uint32(_Hash_hash0)
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= 16777619
	}
	if i := _Hash_table[h&uint32(len(_Hash_table)-1)]; int(i&0xff) == len(s) {
		t := _Hash_text[i>>8 : i>>8+i&0xff]
		for i := 0; i < len(s); i++ {
			if t[i] != s[i] {
				goto NEXT
			}
		}
		return i
	}
NEXT:
	if i := _Hash_table[(h>>16)&uint32(len(_Hash_table)-1)]; int(i&0xff) == len(s) {
		t := _Hash_text[i>>8 : i>>8+i&0xff]
		for i := 0; i < len(s); i++ {
			if t[i] != s[i] {
				return 0
			}
		}
		return i
	}
	return 0
}

const _Hash_hash0 = 0x67ac9bb5
const _Hash_maxLen = 15
const _Hash_text = "aboutputargetbodyaccept-charsetdatalistarticlearbackgroundet" +
	"erminatemplateacronymainertcanvasideclarelabelongdescrolling" +
	"addressamp-boilerplateasynciteaudiobjectfootereversedatatype" +
	"ofacenterowspanamedialogautofocusemappletimenuitemscopedefau" +
	"ltMutedefaultSelectedeferubyautoplayaxismaparamanifestrongbg" +
	"colorbigblockquotebuttonoframesetitleclassidetailselectedfno" +
	"hreflanguagecodebasefontrackbdircodetypemustmatcheckedisable" +
	"divaluetypecolgroupauseonexitranslatecolspanoresizecompactru" +
	"espeedlcontrolshapefieldsetfigcaptionoscriptfigurequiredtfor" +
	"mactionoshadefaultCheckedformnovalidateh2h3h4h5h6hgroupictur" +
	"esourcehiddenabledhtmlhttp-equivaricontentiframeborderimagei" +
	"mginlistrikeygenctypeinputinslotmarkmarqueematheadermaxlengt" +
	"h1metabindexmlnsmallowfullscreenowraplaintextareadonlymeterm" +
	"ethodmultiplegendmutedoptionprofileprogressortablepropertysr" +
	"csetstylesubdopenavalignobruleseamlessectionoembedelsummarys" +
	"uprefixmpostersvgvideoptgrouportalinkvisiblevlinkvocabbrwbr"

var _Hash_table = [1 << 9]Hash{
	0x1:   0x13e06, // button
	0x3:   0x2a207, // enctype
	0x4:   0x32206, // option
	0x5:   0x1fb0a, // figcaption
	0x7:   0x2ae03, // ins
	0x9:   0x9605,  // audio
	0xb:   0x2830b, // frameborder
	0xd:   0x2190a, // formaction
	0xe:   0x5,     // about
	0xf:   0x34606, // srcset
	0x10:  0x1dd09, // truespeed
	0x11:  0xeb0c,  // defaultMuted
	0x13:  0xa006,  // footer
	0x15:  0x19d08, // disabled
	0x16:  0x26e0a, // http-equiv
	0x19:  0x3a504, // link
	0x1a:  0x29606, // inlist
	0x1d:  0x10804, // ruby
	0x21:  0x2a905, // input
	0x22:  0x35803, // nav
	0x25:  0x7902,  // dd
	0x26:  0x2350e, // formnovalidate
	0x28:  0x16804, // href
	0x29:  0x24702, // h4
	0x2b:  0x10405, // defer
	0x2d:  0x1f308, // fieldset
	0x2e:  0xeb07,  // default
	0x34:  0x2fd04, // area
	0x36:  0xb006,  // typeof
	0x37:  0x37307, // noembed
	0x38:  0x5e07,  // declare
	0x3a:  0x4a07,  // acronym
	0x3b:  0xc05,   // tbody
	0x3e:  0x15107, // classid
	0x41:  0x9a06,  // object
	0x43:  0x16403, // dfn
	0x44:  0xef02,  // ul
	0x45:  0x16c04, // lang
	0x47:  0x16606, // nohref
	0x49:  0x2c803, // max
	0x4a:  0x6505,  // label
	0x4c:  0x1d304, // size
	0x4d:  0xe606,  // scoped
	0x4f:  0x15105, // class
	0x50:  0x11404, // axis
	0x54:  0xbf04,  // span
	0x56:  0x19707, // checked
	0x59:  0x38506, // prefix
	0x5b:  0x4208,  // template
	0x5c:  0x370d,  // undeterminate
	0x5d:  0xc606,  // dialog
	0x5e:  0x6908,  // longdesc
	0x60:  0x21903, // for
	0x61:  0x2c102, // th
	0x64:  0x15d08, // selected
	0x65:  0x35103, // sub
	0x6a:  0xd306,  // usemap
	0x6e:  0x24d06, // hgroup
	0x6f:  0x38303, // sup
	0x70:  0x2b404, // mark
	0x71:  0x28206, // iframe
	0x72:  0x30605, // meter
	0x74:  0x21008, // required
	0x75:  0x1f04,  // data
	0x78:  0x14308, // noframes
	0x83:  0x7807,  // address
	0x88:  0x10c08, // autoplay
	0x8a:  0x28e05, // image
	0x8b:  0x16c08, // language
	0x8e:  0x2f904, // text
	0x8f:  0x16802, // hr
	0x90:  0x5d02,  // id
	0x92:  0x31108, // multiple
	0x94:  0x16808, // hreflang
	0x95:  0x2db05, // xmlns
	0x96:  0x24902, // h5
	0x98:  0x25207, // picture
	0x99:  0x1106,  // accept
	0x9a:  0x1a609, // valuetype
	0x9b:  0x3a006, // portal
	0x9d:  0xac08,  // datatype
	0x9e:  0x18403, // bdi
	0xa0:  0x27a04, // icon
	0xa2:  0xa503,  // rev
	0xa5:  0x25708, // resource
	0xa8:  0x35504, // open
	0xac:  0x4302,  // em
	0xae:  0x1340a, // blockquote
	0xb0:  0x2f409, // plaintext
	0xb1:  0x2d204, // meta
	0xb2:  0x1c01,  // s
	0xb4:  0xdc04,  // time
	0xb5:  0x1fe07, // caption
	0xb8:  0x33e08, // property
	0xb9:  0x1,     // a
	0xbb:  0x2b807, // marquee
	0xbc:  0x3b505, // vocab
	0xbd:  0x1e502, // dl
	0xbf:  0xbc07,  // rowspan
	0xc4:  0x18503, // dir
	0xc5:  0x39908, // optgroup
	0xcc:  0x38c06, // poster
	0xcd:  0x24502, // h3
	0xce:  0x3b804, // abbr
	0xd1:  0x17408, // codebase
	0xd2:  0x27b07, // content
	0xd4:  0x7e04,  // samp
	0xd6:  0xc204,  // name
	0xd9:  0x14c05, // title
	0xda:  0x1a605, // value
	0xdd:  0xb004,  // type
	0xde:  0x35f04, // nobr
	0xe0:  0x17c04, // font
	0xe1:  0xd603,  // map
	0xe2:  0x2d002, // h1
	0xe3:  0x22207, // noshade
	0xe4:  0x6303,  // rel
	0xe5:  0x14002, // tt
	0xe7:  0xde04,  // menu
	0xeb:  0x2f908, // textarea
	0xee:  0x35b05, // align
	0xf1:  0x29303, // img
	0xf2:  0x35a06, // valign
	0xf3:  0x2c204, // head
	0xf4:  0x12a07, // bgcolor
	0xf5:  0x5004,  // main
	0xf6:  0x2302,  // li
	0xf7:  0x5205,  // inert
	0xfa:  0x5706,  // canvas
	0xfb:  0xe605,  // scope
	0xfc:  0x15d06, // select
	0x100: 0xa508,  // reversed
	0x101: 0x20408, // noscript
	0x102: 0x37c07, // summary
	0x103: 0x24b02, // h6
	0x106: 0x17404, // code
	0x107: 0x14508, // frameset
	0x10a: 0x12406, // strong
	0x10d: 0x300a,  // background
	0x10e: 0x18303, // kbd
	0x114: 0x31706, // legend
	0x116: 0x32f08, // progress
	0x118: 0x2d408, // tabindex
	0x119: 0x34603, // src
	0x11c: 0x39505, // video
	0x11f: 0x29a06, // strike
	0x121: 0xd706,  // applet
	0x123: 0x2802,  // rt
	0x125: 0x20606, // script
	0x128: 0xbc04,  // rows
	0x129: 0x2707,  // article
	0x12e: 0x9204,  // cite
	0x131: 0x18c0d, // typemustmatch
	0x133: 0x17f05, // track
	0x135: 0x3b902, // bb
	0x136: 0x1ee05, // shape
	0x137: 0x5b05,  // aside
	0x138: 0x1b60b, // pauseonexit
	0x13c: 0x38503, // pre
	0x140: 0x301,   // u
	0x149: 0x1a403, // div
	0x14c: 0x3a405, // alink
	0x14e: 0x27703, // var
	0x14f: 0x21d06, // action
	0x152: 0x2b05,  // clear
	0x154: 0x2401,  // i
	0x155: 0x21702, // dt
	0x156: 0x36608, // seamless
	0x157: 0x21904, // form
	0x15b: 0x15707, // details
	0x15f: 0x8e05,  // async
	0x160: 0x26a04, // html
	0x161: 0x33608, // sortable
	0x165: 0x2f02,  // rb
	0x167: 0x2e10f, // allowfullscreen
	0x168: 0x17804, // base
	0x169: 0x25f06, // hidden
	0x16e: 0x2ef06, // nowrap
	0x16f: 0x2505,  // start
	0x170: 0x14505, // frame
	0x171: 0x1f08,  // datalist
	0x173: 0x12502, // tr
	0x174: 0x30b06, // method
	0x175: 0x101,   // b
	0x176: 0x1c904, // cols
	0x178: 0x110e,  // accept-charset
	0x17a: 0x36205, // rules
	0x17b: 0x7f0f,  // amp-boilerplate
	0x17f: 0x2270e, // defaultChecked
	0x180: 0x32807, // profile
	0x181: 0x2b004, // slot
	0x182: 0x11a05, // param
	0x185: 0x1c907, // colspan
	0x186: 0x34c05, // style
	0x187: 0x1e02,  // td
	0x188: 0x12c05, // color
	0x18c: 0x13901, // q
	0x18d: 0x3b005, // vlink
	0x18e: 0x39203, // svg
	0x18f: 0x33905, // table
	0x190: 0x29e06, // keygen
	0x192: 0x20c06, // figure
	0x193: 0x3a907, // visible
	0x195: 0x17808, // basefont
	0x196: 0x8802,  // rp
	0x197: 0xf60f,  // defaultSelected
	0x198: 0x1af08, // colgroup
	0x19a: 0x3bc03, // wbr
	0x19c: 0x36d07, // section
	0x19d: 0x25906, // source
	0x19f: 0x2bf04, // math
	0x1a1: 0x2fe08, // readonly
	0x1a7: 0x1e708, // controls
	0x1a9: 0xde08,  // menuitem
	0x1ad: 0x206,   // output
	0x1b0: 0x2c809, // maxlength
	0x1b2: 0xe209,  // itemscope
	0x1b9: 0x501,   // p
	0x1bc: 0x2df05, // small
	0x1bd: 0x36102, // br
	0x1c0: 0x5503,  // rtc
	0x1c1: 0x1c009, // translate
	0x1c4: 0x35303, // bdo
	0x1c5: 0xd04,   // body
	0x1c8: 0xb706,  // center
	0x1c9: 0x2c105, // thead
	0x1ca: 0xcc09,  // autofocus
	0x1cc: 0xb504,  // face
	0x1cd: 0x24302, // h2
	0x1ce: 0x11e08, // manifest
	0x1d0: 0x706,   // target
	0x1d1: 0x11605, // ismap
	0x1d3: 0xc405,  // media
	0x1d7: 0x13103, // big
	0x1da: 0x37903, // del
	0x1dc: 0x6f09,  // scrolling
	0x1de: 0x37505, // embed
	0x1e0: 0x31d05, // muted
	0x1e4: 0x2390a, // novalidate
	0x1e6: 0x7202,  // ol
	0x1eb: 0x9f05,  // tfoot
	0x1ec: 0x18808, // codetype
	0x1ee: 0x26307, // enabled
	0x1f0: 0x2c206, // header
	0x1f1: 0x1cf08, // noresize
	0x1f6: 0x1d707, // compact
	0x1f9: 0x12c03, // col
	0x1fa: 0x38a03, // xmp
	0x1fb: 0x1807,  // charset
}