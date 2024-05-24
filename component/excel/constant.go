package excel

const (
	TagExcel = "ex"

	SubTagHeader   = "header"
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
