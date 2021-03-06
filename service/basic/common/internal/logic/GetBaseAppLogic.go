package logic

import (
	"context"

	"kapok-admin/service/basic/common/common"
	"kapok-admin/service/basic/common/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetBaseAppLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBaseAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBaseAppLogic {
	return &GetBaseAppLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBaseAppLogic) GetBaseApp(in *common.BaseAppReq) (*common.BaseAppResp, error) {
	// todo: add your logic here and delete this line

	return &common.BaseAppResp{}, nil
}
