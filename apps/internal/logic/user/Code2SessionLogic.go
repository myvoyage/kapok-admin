package user

import (
	"context"

	"kapok-admin/apps/internal/svc"
	"kapok-admin/apps/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type Code2SessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCode2SessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) Code2SessionLogic {
	return Code2SessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Code2SessionLogic) Code2Session() (resp *types.LoginAppUser, err error) {
	// todo: add your logic here and delete this line

	return
}
