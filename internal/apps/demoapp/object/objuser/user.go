package objuser

type UserBaseInfo struct {
	CompanyID    uint   `json:"companyId" form:"companyId"`       // 公司id
	DepartmentID uint   `json:"departmentId" form:"departmentId"` // 部门id
	Name         string `json:"name" form:"name"`                 // 姓名
}
