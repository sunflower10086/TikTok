// Code generated by goctl. DO NOT EDIT.
// Source: relation.proto

package server

import (
	"context"

	"github.com/sunflower10086/TikTok/relation/internal/logic"
	"github.com/sunflower10086/TikTok/relation/internal/svc"
	___relation "github.com/sunflower10086/TikTok/relation/pb"
)

type RelationServer struct {
	svcCtx *svc.ServiceContext
	___relation.UnimplementedRelationServer
}

func NewRelationServer(svcCtx *svc.ServiceContext) *RelationServer {
	return &RelationServer{
		svcCtx: svcCtx,
	}
}

func (s *RelationServer) Action(ctx context.Context, in *___relation.ActionRequest) (*___relation.Empty, error) {
	l := logic.NewActionLogic(ctx, s.svcCtx)
	return l.Action(in)
}

func (s *RelationServer) FollowList(ctx context.Context, in *___relation.FollowListRequest) (*___relation.FollowListResponse, error) {
	l := logic.NewFollowListLogic(ctx, s.svcCtx)
	return l.FollowList(in)
}

func (s *RelationServer) FollowerList(ctx context.Context, in *___relation.FollowerListRequest) (*___relation.FollowerListResponse, error) {
	l := logic.NewFollowerListLogic(ctx, s.svcCtx)
	return l.FollowerList(in)
}

func (s *RelationServer) FriendList(ctx context.Context, in *___relation.FriendListRequest) (*___relation.FriendListResponse, error) {
	l := logic.NewFriendListLogic(ctx, s.svcCtx)
	return l.FriendList(in)
}

func (s *RelationServer) MessageChat(ctx context.Context, in *___relation.MessageChatRequest) (*___relation.MessageChatResponse, error) {
	l := logic.NewMessageChatLogic(ctx, s.svcCtx)
	return l.MessageChat(in)
}

func (s *RelationServer) MessageAction(ctx context.Context, in *___relation.MessageActionRequest) (*___relation.MessageActionResponse, error) {
	l := logic.NewMessageActionLogic(ctx, s.svcCtx)
	return l.MessageAction(in)
}
