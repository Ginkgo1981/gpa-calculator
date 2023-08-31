package handler

import (
	"net/http"

	"gpa/internal/logic"
	"gpa/internal/svc"
	"gpa/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GpaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGpaLogic(r.Context(), svcCtx)
		resp, err := l.Gpa(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
