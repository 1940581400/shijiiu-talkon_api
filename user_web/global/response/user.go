package response

type User struct {
	// id
	Id int32 `json:"id"`
	// 用户名
	Email string `json:"email"`
	// 手机号
	Mobile string `json:"mobile"`
	// 密码
	Password string `json:"password"`
	// 昵称
	NickName string `json:"nickname"`
	// 出生日期
	Birthday JsonTimeDateOnly `json:"birthday"`
	// 性别
	Gender int32 `json:"gender"`
	// 身份证号
	IdCard string `json:"idCard"`
	// 用户类型 0普通用户，1管理员用户
	UserType int32 `json:"userType"`
	// 创建时间
	CreateTime JsonTimeDateTime `json:"createTime"`
	// 更新时间
	UpdateTime JsonTimeDateTime `json:"updateTime"`
	// 创建人
	CreateUser int32 `json:"createUser"`
	// 更新人
	UpdateUser int32 `json:"updateUser"`
	// 是否删除，0否，1是
	IsDeleted int32 `json:"isDeleted"`
}
