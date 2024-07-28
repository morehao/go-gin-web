package genCode

const (
	TplFuncIsSysField = "isSysField"
)

func IsSysField(name string) bool {
	sysFieldMap := map[string]struct{}{
		"Id":          {},
		"CreatedTime": {},
		"CreatedBy":   {},
		"UpdatedTime": {},
		"UpdatedBy":   {},
		"DeletedTime": {},
		"DeletedBy":   {},
	}
	_, ok := sysFieldMap[name]
	return ok
}
