package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"talkon_api/user_web/forms"
	"talkon_api/user_web/global"
	"talkon_api/user_web/global/response"
	"talkon_api/user_web/proto"
	"talkon_api/user_web/utils"
	"time"
)

// GetUserList 分页获取用户列表
func GetUserList(ctx *gin.Context) {
	zap.L().Info("[GetUserList] 被调用")
	pageNo := ctx.DefaultQuery("pageNo", "1")
	pageSize := ctx.DefaultQuery("pageSize", "10")
	pNo, err := strconv.Atoi(pageNo)
	pS, err := strconv.Atoi(pageSize)
	if err != nil {
		utils.GrpcErrToHttpErr(ctx, err)
		return
	}
	resp, err := global.NewUserSrvClient().GetUserList(context.Background(), &proto.PageInfoReq{
		PageNo:   uint32(pNo),
		PageSize: uint32(pS),
	})
	if err != nil {
		utils.GrpcErrToHttpErr(ctx, err)
		return
	}

	result := make([]response.User, 0)
	for _, item := range resp.Data {
		user := response.User{
			Id:         item.Id,
			Email:      item.Email,
			Mobile:     item.Mobile,
			Password:   item.Password,
			NickName:   item.NickName,
			Birthday:   response.JsonTimeDateOnly(time.Unix(int64(item.Birthday), 0)),
			Gender:     item.Gender,
			IdCard:     item.IdCard,
			UserType:   item.UserType,
			CreateTime: response.JsonTimeDateTime(time.Unix(int64(item.CreateTime), 0)),
			UpdateTime: response.JsonTimeDateTime(time.Unix(int64(item.UpdateTime), 0)),
			CreateUser: item.CreateUser,
			UpdateUser: item.UpdateUser,
			IsDeleted:  item.IsDeleted,
		}
		result = append(result, user)
	}
	ctx.JSON(http.StatusOK, result)
}

func EmailLogin(ctx *gin.Context) {
	zap.L().Info("[EmailLogin] 被调用")
	var form = forms.EmailLogin{}
	err := ctx.ShouldBind(&form)
	if err != nil {
		validCustomErr := []utils.CustomErrors{
			{FieldName: "Email", Tag: "required", CustomMsg: "请输入邮箱"},
			{FieldName: "Email", Tag: "email", CustomMsg: "邮箱格式错误"},
			{FieldName: "Password", Tag: "required", CustomMsg: "请输入密码"},
		}
		utils.ValidationCustomErrors(ctx, err, validCustomErr)
		return
	}

	userInfo, err := global.NewUserSrvClient().GetUserByEmail(context.Background(), &proto.EmailReq{Email: form.Email})
	if err != nil {
		utils.GrpcErrToHttpErr(ctx, err)
		return
	}
	check, err := global.NewUserSrvClient().CheckPassword(context.Background(), &proto.CheckPasswordReq{Password: form.Password, EncodedPwdSep: userInfo.Password})
	if err != nil {
		utils.GrpcErrToHttpErr(ctx, err)
		return
	}
	if check.Ok {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "登录成功",
		})
		return
	}
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"msg": "用户名或密码错误",
	})

}
