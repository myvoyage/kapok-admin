package common

import (
	"context"

	"kapok-admin/apps/internal/svc"
	"kapok-admin/apps/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SnsInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSnsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) SnsInfoLogic {
	return SnsInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SnsInfoLogic) SnsInfo(req types.SnsReq) (resp *types.SnsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
