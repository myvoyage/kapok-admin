// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userclient

import (
	"context"

	"kapok/kapok-admin/service/basic/user/user"

	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AppConfigReq     = user.AppConfigReq
	AppConfigResp    = user.AppConfigResp
	AppUserLoginResp = user.AppUserLoginResp
	AppUserResp      = user.AppUserResp
	JwtToken         = user.JwtToken
	LoginReq         = user.LoginReq
	RegisterReq      = user.RegisterReq
	Request          = user.Request
	Response         = user.Response
	SetMobileReq     = user.SetMobileReq
	UserReply        = user.UserReply
	UserReq          = user.UserReq

	User interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*UserReply, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*UserReply, error)
		UserInfo(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserReply, error)
		SnsLogin(ctx context.Context, in *AppConfigReq, opts ...grpc.CallOption) (*AppUserResp, error)
		SetMobile(ctx context.Context, in *SetMobileReq, opts ...grpc.CallOption) (*UserReply, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*UserReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*UserReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) UserInfo(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

func (m *defaultUser) SnsLogin(ctx context.Context, in *AppConfigReq, opts ...grpc.CallOption) (*AppUserResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.SnsLogin(ctx, in, opts...)
}

func (m *defaultUser) SetMobile(ctx context.Context, in *SetMobileReq, opts ...grpc.CallOption) (*UserReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.SetMobile(ctx, in, opts...)
}