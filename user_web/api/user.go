package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"talkon_api/user_web/global/response"
	"talkon_api/user_web/proto"
	"talkon_api/user_web/utils"
	"time"
)

func GetUserList(ctx *gin.Context) {
	zap.L().Info("[GetUserList] 被调用")
	ip := "127.0.0.1"
	port := "8088"
	dial, err := grpc.Dial(fmt.Sprintf("%s:%s", ip, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.L().Error("[GetUserList] 连接用户服务失败", zap.String("msg", err.Error()))
	}
	client := proto.NewUserClient(dial)
	resp, err := client.GetUserList(context.Background(), &proto.PageInfoReq{
		PageNo:   1,
		PageSize: 10,
	})
	if err != nil {
		utils.GrpcErrToHttpErr(err, ctx)
		zap.L().Error("[GetUserList] 调用失败", zap.String("msg", err.Error()))
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
