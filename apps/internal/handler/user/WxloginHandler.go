package user

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"kapok-admin/apps/internal/logic/user"
	"kapok-admin/apps/internal/svc"
	"kapok-admin/apps/internal/types"
)

func WxloginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WxLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewWxloginLogic(r.Context(), ctx)
		resp, err := l.Wxlogin(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
