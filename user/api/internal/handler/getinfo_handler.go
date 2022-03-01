package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-k8s-test/user/api/internal/logic"
	"go-zero-k8s-test/user/api/internal/svc"
	"go-zero-k8s-test/user/api/internal/types"
)

func getinfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetinfoLogic(r.Context(), svcCtx)
		resp, err := l.Getinfo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
