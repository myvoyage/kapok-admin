package logic

import (
	"context"

	"kapok/kapok-admin/service/basic/user/internal/svc"
	"kapok/kapok-admin/service/basic/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type SetMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetMobileLogic {
	return &SetMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetMobileLogic) SetMobile(in *user.SetMobileReq) (*user.UserReply, error) {
	// todo: add your logic here and delete this line

	return &user.UserReply{}, nil
}
