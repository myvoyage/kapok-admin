package user

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"kapok/kapok-admin/apps/internal/logic/user"
	"kapok/kapok-admin/apps/internal/svc"
	"kapok/kapok-admin/apps/internal/types"
)

func LoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), ctx)
		resp, err := l.Login(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
