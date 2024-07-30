package objUser

type UserBaseInfo struct {
    CompanyID uint64 `json:"companyId" form:"companyId"` // 公司id
    DepartmentID uint64 `json:"departmentId" form:"departmentId"` // 部门id
    Name string `json:"name" form:"name"` // 姓名
}
