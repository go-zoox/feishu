package department

type DepartmentEntity struct {
	// 部门名称
	Name string `json:"name"`

	// 国际化的部门名称
	I18nName struct {
		ZhCN string `json:"zh_cn"`
		EnUS string `json:"en_us"`
		JaJP string `json:"ja_jp"`
	} `json:"i18n_name"`

	// 父部门的ID
	ParentDepartmentID string `json:"parent_department_id"`

	// 本部门的自定义部门ID
	DeparntmentID string `json:"department_id"`

	// 部门的open_id
	OpenDepartmentID string `json:"open_department_id"`

	// 部门主管用户ID
	LeaderUserID string `json:"leader_user_id"`

	// 部门群ID
	ChatID string `json:"chat_id"`

	// 部门的排序，即部门在其同级部门的展示顺序
	Order string `json:"order"`

	// 部门单位自定义ID列表，当前只支持一个
	UnitIDS []string `json:"unit_ids"`

	// 部门下用户的个数
	MemberCount int `json:"member_count"`

	// 部门状态
	Status DepartmentEntityStatus `json:"status"`

	// 是否创建部门群，默认不创建
	CreateGroupChat bool `json:"create_group_chat"`
}

type DepartmentEntityStatus struct {
	// 是否被删除
	IsDeleted bool `json:"is_deleted"`
}
