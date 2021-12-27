// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"kapok/kapok-admin/service/basic/user/internal/logic"
	"kapok/kapok-admin/service/basic/user/internal/svc"
	"kapok/kapok-admin/service/basic/user/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Ping(ctx context.Context, in *user.Request) (*user.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *UserServer) Register(ctx context.Context, in *user.RegisterReq) (*user.UserReply, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginReq) (*user.UserReply, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) UserInfo(ctx context.Context, in *user.UserReq) (*user.UserReply, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

func (s *UserServer) SnsLogin(ctx context.Context, in *user.AppConfigReq) (*user.AppUserResp, error) {
	l := logic.NewSnsLoginLogic(ctx, s.svcCtx)
	return l.SnsLogin(in)
}

func (s *UserServer) SetMobile(ctx context.Context, in *user.SetMobileReq) (*user.UserReply, error) {
	l := logic.NewSetMobileLogic(ctx, s.svcCtx)
	return l.SetMobile(in)
}
