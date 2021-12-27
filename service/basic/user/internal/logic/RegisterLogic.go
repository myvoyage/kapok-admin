package logic

import (
	"context"

	"kapok-admin/service/basic/user/internal/svc"
	"kapok-admin/service/basic/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.UserReply, error) {
	// todo: add your logic here and delete this line

	return &user.UserReply{}, nil
}
