package svc

import (
	"github.com/tal-tech/go-zero/rest"
	"kapok/kapok-admin/apps/internal/config"
	"kapok/kapok-admin/apps/internal/middleware"
)

type ServiceContext struct {
	Config    config.Config
	Usercheck rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Usercheck: middleware.NewUsercheckMiddleware().Handle,
	}
}
