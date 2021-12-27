package logic

import (
	"context"

	"kapok-admin/service/addons/scrm/internal/svc"
	"kapok-admin/service/addons/scrm/scrm"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetScrmConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetScrmConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScrmConfigLogic {
	return &GetScrmConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetScrmConfigLogic) GetScrmConfig(in *scrm.ScrmConfigReq) (*scrm.ScrmConfigResp, error) {
	// todo: add your logic here and delete this line

	return &scrm.ScrmConfigResp{}, nil
}
