package user

import (
	"context"

	"kapok-admin/apps/internal/svc"
	"kapok-admin/apps/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type WxloginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxloginLogic(ctx context.Context, svcCtx *svc.ServiceContext) WxloginLogic {
	return WxloginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxloginLogic) Wxlogin(req types.WxLoginReq) (resp *types.LoginAppUser, err error) {
	// todo: add your logic here and delete this line

	return
}
