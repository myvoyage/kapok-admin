package user

import (
	"context"

	"kapok/kapok-admin/apps/internal/svc"
	"kapok/kapok-admin/apps/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SetMobileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) SetMobileLogic {
	return SetMobileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetMobileLogic) SetMobile(req types.SetMobileReq) (resp *types.UserReply, err error) {
	// todo: add your logic here and delete this line

	return
}
