package common

import (
	"context"

	"kapok-admin/apps/internal/svc"
	"kapok-admin/apps/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AppInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) AppInfoLogic {
	return AppInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppInfoLogic) AppInfo(req types.Beid) (resp *types.Application, err error) {
	// todo: add your logic here and delete this line

	return
}
