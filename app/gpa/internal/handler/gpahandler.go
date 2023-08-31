package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gpa/app/gpa/internal/logic"
	"gpa/app/gpa/internal/svc"
)

func GpaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGpaLogic(r.Context(), svcCtx)
		resp, err := l.Gpa()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
