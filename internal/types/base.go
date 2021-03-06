package types

import "fmt"

const (
	pkg         = "types"
	typeNumMask = 0x1F
)

type Base byte

const (
	BaseEnum    Base = 0x00
	BaseSint8   Base = 0x01 // 2's complement format
	BaseUint8   Base = 0x02
	BaseSint16  Base = 0x03 // 2's complement format
	BaseUint16  Base = 0x04
	BaseSint32  Base = 0x05 // 2's complement format
	BaseUint32  Base = 0x06
	BaseString  Base = 0x07 // Null terminated string encoded in UTF-8
	BaseFloat32 Base = 0x08
	BaseFloat64 Base = 0x09
	BaseUint8z  Base = 0x0A
	BaseUint16z Base = 0x0B
	BaseUint32z Base = 0x0C
	BaseByte    Base = 0x0D // Array of bytes. Field is invalid if all bytes are invalid
	BaseSint64  Base = 0x0E // 2's complement format
	BaseUint64  Base = 0x0F
	BaseUint64z Base = 0x10
)

func DecodeBase(b byte) Base {
	return Base(b & typeNumMask)
}

func (t Base) Float() bool {
	return !t.Integer() && t.Signed()
}

func (t Base) GoInvalidValue() string {
	return binvalid[t]
}

func (t Base) GoType() string {
	return bgotype[t]
}

func (t Base) Integer() bool {
	return binteger[t]
}

func (t Base) Known() bool {
	return int(t) < len(bname)
}

func (t Base) PkgString() string {
	return pkg + "." + t.String()
}

func (t Base) Signed() bool {
	return bsigned[t]
}

func (t Base) Size() int {
	return bsize[t]
}

func (t Base) String() string {
	return bname[t]
}

var bsize = [...]int{
	1,
	1,
	1,
	2,
	2,
	4,
	4,
	1,
	4,
	8,
	1,
	2,
	4,
	1,
	8,
	8,
	8,
}

var bname = [...]string{
	"BaseEnum",
	"BaseSint8",
	"BaseUint8",
	"BaseSint16",
	"BaseUint16",
	"BaseSint32",
	"BaseUint32",
	"BaseString",
	"BaseFloat32",
	"BaseFloat64",
	"BaseUint8z",
	"BaseUint16z",
	"BaseUint32z",
	"BaseByte",
	"BaseSint64",
	"BaseUint64",
	"BaseUint64z",
}

var binteger = [...]bool{
	false,
	true,
	true,
	true,
	true,
	true,
	true,
	false,
	false,
	false,
	true,
	true,
	true,
	false,
	true,
	true,
	true,
}

var bsigned = [...]bool{
	false,
	true,
	false,
	true,
	false,
	true,
	false,
	false,
	true,
	true,
	false,
	false,
	false,
	false,
	true,
	false,
	false,
}

var bgotype = [...]string{
	"byte",
	"int8",
	"uint8",
	"int16",
	"uint16",
	"int32",
	"uint32",
	"string",
	"float32",
	"float64",
	"uint8",
	"uint16",
	"uint32",
	"byte",
	"int64",
	"uint64",
	"uint64",
}

var binvalid = [...]string{
	"0xFF",
	"0x7F",
	"0xFF",
	"0x7FFF",
	"0xFFFF",
	"0x7FFFFFFF",
	"0xFFFFFFFF",
	"\"\"",
	"0xFFFFFFFF",
	"0xFFFFFFFFFFFFFFFF",
	"0x00",
	"0x0000",
	"0x00000000",
	"0xFF",
	"0x7FFFFFFFFFFFFFFF",
	"0xFFFFFFFFFFFFFFFF",
	"0x0000000000000000",
}

func BaseFromString(s string) (Base, error) {
	t, found := baseStringToType[s]
	if !found {
		return 0xFF, fmt.Errorf("no base type found for string: %q", s)
	}
	return t, nil
}

var baseStringToType = map[string]Base{
	"enum":    BaseEnum,
	"sint8":   BaseSint8,
	"uint8":   BaseUint8,
	"sint16":  BaseSint16,
	"uint16":  BaseUint16,
	"sint32":  BaseSint32,
	"uint32":  BaseUint32,
	"string":  BaseString,
	"float32": BaseFloat32,
	"float64": BaseFloat64,
	"uint8z":  BaseUint8z,
	"uint16z": BaseUint16z,
	"uint32z": BaseUint32z,
	"byte":    BaseByte,
	"sint64":  BaseSint64,
	"uint64":  BaseUint64,
	"uint64z": BaseUint64z,

	// Typo in SDK 20.14:
	"unit8": BaseUint8,
}
