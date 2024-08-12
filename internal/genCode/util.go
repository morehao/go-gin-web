package genCode

const (
	TplFuncIsSysField = "isSysField"
)

func IsSysField(name string) bool {
	sysFieldMap := map[string]struct{}{
		"Id":        {},
		"CreatedAt": {},
		"CreatedBy": {},
		"UpdatedAt": {},
		"UpdatedBy": {},
		"DeletedAt": {},
		"DeletedBy": {},
	}
	_, ok := sysFieldMap[name]
	return ok
}
