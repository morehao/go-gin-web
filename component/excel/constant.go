package excel

const (
	TagExcel = "ex"

	SubTagHead     = "head"
	SubTagType     = "type"
	SubTagMax      = "max"
	SubTagRequired = "required"
)
const (
	typeDefault tagType = iota
	typeHead
	typeType
	typeRequired
	typeMax
)

const (
	tagSeparator    = ","
	tagKeySeparator = ":"
)
