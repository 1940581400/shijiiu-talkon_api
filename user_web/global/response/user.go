package response

type User struct {
	// id
	Id int32 `json:"Id"`
	// 用户名
	Email string `json:"Email"`
	// 手机号
	Mobile string `json:"Mobile"`
	// 密码
	Password string `json:"Password"`
	// 昵称
	NickName string `json:"NickName"`
	// 出生日期
	Birthday JsonTimeDateOnly `json:"Birthday"`
	// 性别
	Gender int32 `json:"Gender"`
	// 身份证号
	IdCard string `json:"IdCard"`
	// 用户类型 0普通用户，1管理员用户
	UserType int32 `json:"UserType"`
	// 创建时间
	CreateTime JsonTimeDateTime `json:"CreateTime"`
	// 更新时间
	UpdateTime JsonTimeDateTime `json:"UpdateTime"`
	// 创建人
	CreateUser int32 `json:"CreateUser"`
	// 更新人
	UpdateUser int32 `json:"UpdateUser"`
	// 是否删除，0否，1是
	IsDeleted int32 `json:"IsDeleted"`
}
