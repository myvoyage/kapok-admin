package logic

import (
	"context"

	"kapok-admin/service/addons/cms/cms"
	"kapok-admin/service/addons/cms/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetCmsConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCmsConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCmsConfigLogic {
	return &GetCmsConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCmsConfigLogic) GetCmsConfig(in *cms.CmsConfigReq) (*cms.CmsConfigResp, error) {
	// todo: add your logic here and delete this line

	return &cms.CmsConfigResp{}, nil
}
