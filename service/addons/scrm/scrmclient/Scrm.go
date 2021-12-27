// Code generated by goctl. DO NOT EDIT!
// Source: scrm.proto

package scrmclient

import (
	"context"

	"kapok-admin/service/addons/scrm/scrm"

	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ScrmConfigReq  = scrm.ScrmConfigReq
	ScrmConfigResp = scrm.ScrmConfigResp

	Scrm interface {
		GetScrmConfig(ctx context.Context, in *ScrmConfigReq, opts ...grpc.CallOption) (*ScrmConfigResp, error)
	}

	defaultScrm struct {
		cli zrpc.Client
	}
)

func NewScrm(cli zrpc.Client) Scrm {
	return &defaultScrm{
		cli: cli,
	}
}

func (m *defaultScrm) GetScrmConfig(ctx context.Context, in *ScrmConfigReq, opts ...grpc.CallOption) (*ScrmConfigResp, error) {
	client := scrm.NewScrmClient(m.cli.Conn())
	return client.GetScrmConfig(ctx, in, opts...)
}
