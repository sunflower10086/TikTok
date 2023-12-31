package token

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// GetUserIDAndUsernameFromCtx 从ctx中获取userID和username
func GetUserIDAndUsernameFromCtx(ctx *gin.Context) (int64, string) {
	userID, _ := ctx.Get("userID")
	username, _ := ctx.Get("username")
	fmt.Println(userID, username)
	return userID.(int64), username.(string)
}
