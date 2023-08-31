package logic

import (
	"context"
	"gpa/app/gpa/internal/svc"
	"gpa/app/gpa/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GpaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGpaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GpaLogic {
	return &GpaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GpaLogic) Gpa(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
