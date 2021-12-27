package logic

import (
	"context"

	"kapok/kapok-admin/service/basic/user/internal/svc"
	"kapok/kapok-admin/service/basic/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type SnsLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSnsLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SnsLoginLogic {
	return &SnsLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SnsLoginLogic) SnsLogin(in *user.AppConfigReq) (*user.AppUserResp, error) {
	// todo: add your logic here and delete this line

	return &user.AppUserResp{}, nil
}
