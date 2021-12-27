package common

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"kapok/kapok-admin/apps/internal/logic/common"
	"kapok/kapok-admin/apps/internal/svc"
	"kapok/kapok-admin/apps/internal/types"
)

func AppInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Beid
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := common.NewAppInfoLogic(r.Context(), ctx)
		resp, err := l.AppInfo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
