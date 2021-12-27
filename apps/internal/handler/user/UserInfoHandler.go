package user

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"kapok/kapok-admin/apps/internal/logic/user"
	"kapok/kapok-admin/apps/internal/svc"
	"kapok/kapok-admin/apps/internal/types"
)

func UserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserInfoLogic(r.Context(), ctx)
		resp, err := l.UserInfo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
