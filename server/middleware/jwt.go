package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/gin-api-learn/global"
	"github.com/zhengpanone/gin-api-learn/model/common/response"
	"github.com/zhengpanone/gin-api-learn/model/dao"
	"github.com/zhengpanone/gin-api-learn/model/request"
	"github.com/zhengpanone/gin-api-learn/utils"
)

// JWTAuthMiddleware JWT中间件
func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 获取参数中的token
		token, err := utils.GetToken(ctx)
		global.GVA_LOG.Sugar().Info("token:%s", token)
		if err != nil {
			response.ErrorWithMsg(ctx, err.Error())
			// 中断请求
			ctx.Abort()
			return
		}
		// 验证Token
		j := utils.NewJWT()
		userClaim, err := j.ParseToken(token)
		if err != nil {
			response.ErrorWithToken(ctx, "Token error:"+err.Error())
			ctx.Abort()
			return
		}
		// 设置到上下文
		setContextData(ctx, userClaim, token)
		// 继续请求后续流程
		ctx.Next()

	}
}

// 设置数据到上下文
func setContextData(ctx *gin.Context, customClaim *request.CustomClaims, token string) {
	userDao := &dao.UserDao{
		UserID: customClaim.UserID,
	}
	user, err := userDao.FindUser()
	if err != nil {
		response.ErrorWithMsg(ctx, "用户不存在")
		ctx.Abort()
		return
	}
	user.Token = token
	ctx.Set("userClaim", customClaim)
	ctx.Set("user", user)
}
