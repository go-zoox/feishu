package user

type UserEntity struct {
	// 用户的union_id
	// 不同ID的说明参见 用户相关的 ID 概念: https://open.feishu.cn/document/home/user-identity-introduction/introduction
	UnionID string `json:"union_id"`

	// 租户内用户的唯一标识，用户的user_id
	// 不同ID的说明参见 用户相关的 ID 概念: https://open.feishu.cn/document/home/user-identity-introduction/introduction
	UserID string `json:"user_id"`

	// 用户的open_id
	// 不同ID的说明参见 用户相关的 ID 概念: https://open.feishu.cn/document/home/user-identity-introduction/introduction
	OpenID string `json:"open_id"`

	// 用户名
	Name string `json:"name"`

	// 英文名
	EnName string `json:"en_name"`

	// 昵称
	Nickname string `json:"nickname"`

	// 邮箱
	// 注意非 +86 手机号成员必须同时添加邮箱
	Email string `json:"email"`

	// 手机号，在本企业内不可重复；未认证企业仅支持添加中国大陆手机号，通过飞书认证的企业允许添加海外手机号，注意国际电话区号前缀中必须包含加号 +
	Mobile string `json:"mobile"`

	// 手机号码可见性，true 为可见，false 为不可见，目前默认为 true。不可见时，组织员工将无法查看该员工的手机号码
	MobileVisible bool `json:"mobile_visible"`

	// 性别
	// 可选值：0 为保密，1 为男性，2 为女性
	Gender int `json:"gender"`

	// 用户头像信息
	Avatar UserEntityAvatar `json:"avatar"`

	// 部门状态
	Status UserEntityStatus `json:"status"`

	// 用户所属部门的ID列表，一个用户可属于多个部门。
	// ID值与查询参数中的department_id_type 对应
	// 不同 ID 的说明参见：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/field-overview#23857fe0
	DepartmentIDs []string `json:"department_ids"`

	// 用户的直接主管的用户ID，ID值与查询参数中的user_id_type 对应
	// 不同 ID 的说明参见：https://open.feishu.cn/document/home/user-identity-introduction/introduction
	LeaderUserID string `json:"leader_user_id"`

	// 城市
	City string `json:"city"`

	// 国家或地区Code缩写
	// 具体写入格式请参考：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/country-code-description
	Country string `json:"country"`

	// 工位
	WorkStation string `json:"work_station"`

	// 入职时间
	JoinTime int `json:"join_time"`

	// 是否是租户超级管理员
	IsTenantManager bool `json:"is_tenant_manager"`

	// 工号
	EmployeeNo string `json:"employee_no"`

	// 员工类型
	// 可选值
	// 1: 正式员工
	// 2: 实习生
	// 3: 外包
	// 4: 劳务
	// 5: 顾问
	EmployeeType int `json:"employee_type"`

	// 用户排序信息
	Orders []UserEntityOrder `json:"orders"`

	// 自定义字段，请确保你的组织管理员已在管理后台/组织架构/成员字段管理/自定义字段管理/全局设置中开启了“允许开放平台 API 调用“，否则该字段不会生效/返回
	CustomAttrs []UserEntityCustomAttr `json:"custom_attrs"`

	// 企业邮箱，请先确保已在管理后台启用飞书邮箱服务
	EnterpriseEmail string `json:"enterprise_email"`

	// 职务
	JobTitle string `json:"job_title"`

	// 是否暂停用户
	IsFronzen bool `json:"is_fronzen"`
}

type UserEntityStatus struct {
	// 是否暂停
	IsFronzen bool `json:"is_fronzen"`

	// 是否离职
	IsResigned bool `json:"is_resigned"`

	// 是否激活
	IsActivatd bool `json:"is_ctivatd"`

	// 是否主动退出，主动退出一段时间后用户会自动转为已离职
	IsExisted bool `json:"is_existed"`

	// 是否未加入，需要用户自主确认才能加入团队
	IsUnjoin bool `json:"is_unjoin"`
}

type UserEntityAvatar struct {
	// 72*72像素头像链接
	Avatar72 string `json:"avatar_72"`

	// 240*240像素头像链接
	Avatar240 string `json:"avatar_240"`

	// 640*640像素头像链接
	Avatar640 string `json:"avatar_640"`

	// 原始头像链接
	AvatarOrigin string `json:"avatar_origin"`
}

type UserEntityOrder struct {
	// 排序信息对应的部门ID， ID值与查询参数中的department_id_type 对应
	// 不同 ID 的说明参见：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/field-overview
	DepartmentID string `json:"department_id"`

	// 用户在其直属部门内的排序，数值越大，排序越靠前
	UserOrder int `json:"user_order"`

	// 用户所属的多个部门间的排序，数值越大，排序越靠前
	DepartmentOrder int `json:"department_order"`
}

type UserEntityCustomAttr struct {
	// 自定义字段类型
	// 可选值：
	// 	TEXT：文本类型
	//  HREF：网页
	//  ENUMERATION：枚举
	//  PICTURE_ENUM：图片
	//  GENERIC_USER：用户
	Type string `json:"type"`

	// 自定义字段 ID
	ID string `json:"id"`

	// 自定义字段取值
	Value struct {
		// 字段类型为TEXT时该参数定义字段值，必填；字段类型为HREF时该参数定义网页标题，必填
		Text string `json:"text"`
		// 字段类型为 HREF 时，该参数定义默认 URL
		URL string `json:"url"`
		// 字段类型为 HREF 时，该参数定义PC端 URL
		PCURL string `json:"pc_url"`
		// 字段类型为 ENUMERATION 或 PICTURE_ENUM 时，该参数定义选项值
		OptionID string `json:"option_id"`
		// 选项值
		OptionValue string `json:"option_value"`
		// 名称
		Name string `json:"name"`
		// 图片链接
		PictureURL string `json:"picture_url"`
		// 字段类型为 GENERIC_USER 时，该参数定义引用人员
		GenericUser struct {
			// 引用人员 ID
			ID string `json:"id"`
			// 用户类型 1：用户
			Type int `json:"type"`
		} `json:"generic_user"`
	} `json:"value"`
}
