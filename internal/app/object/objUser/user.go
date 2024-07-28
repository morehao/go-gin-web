package objUser

type UserBaseInfo struct {
    CompanyId uint64 `json:"companyId" form:"companyId"` // 公司id
    DepartmentId uint64 `json:"departmentId" form:"departmentId"` // 部门id
    Name string `json:"name" form:"name"` // 姓名
}
