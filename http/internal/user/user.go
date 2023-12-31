package user

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
)

type Service interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	GetInfo(*gin.Context, *GetInfoRequest) (*GetInfoResponse, error)
}

type LoginRequest struct {
	Username string `json:"username" binding:"required,email" form:"username"`
	Password string `json:"password" binding:"required,min=6,alphanum" form:"password"`
}

type LoginResponse struct {
	result.Response
	Token  *string `json:"token"`   // 用户鉴权token
	UserID *int64  `json:"user_id"` // 用户id
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,email" form:"username"`          // 登录用户名
	Password string `json:"password" binding:"required,min=6,alphanum" form:"password"` // 登录密码
}

type RegisterResponse struct {
	result.Response
	Token  *string `json:"token"`   // 用户鉴权token
	UserID *int64  `json:"user_id"` // 用户id
}

type GetInfoRequest struct {
	UserID uint64  `json:"user_id" query:"user_id" binding:"required" form:"user_id"` // 用户id
	Token  *string `json:"token"`
}

type GetInfoResponse struct {
	result.Response
	User *User `json:"user"` // 用户信息
}

// User
type User struct {
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	ID              int64  `json:"id"`               // 用户id
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Name            string `json:"name"`             // 用户名称
	Signature       string `json:"signature"`        // 个人简介
	TotalFavorited  int64  `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数
}
