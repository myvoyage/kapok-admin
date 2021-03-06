// Code generated by goctl. DO NOT EDIT!
// Source: cms.proto

package server

import (
	"context"

	"kapok-admin/service/addons/cms/cms"
	"kapok-admin/service/addons/cms/internal/logic"
	"kapok-admin/service/addons/cms/internal/svc"
)

type CmsServer struct {
	svcCtx *svc.ServiceContext
}

func NewCmsServer(svcCtx *svc.ServiceContext) *CmsServer {
	return &CmsServer{
		svcCtx: svcCtx,
	}
}

func (s *CmsServer) GetCmsConfig(ctx context.Context, in *cms.CmsConfigReq) (*cms.CmsConfigResp, error) {
	l := logic.NewGetCmsConfigLogic(ctx, s.svcCtx)
	return l.GetCmsConfig(in)
}
