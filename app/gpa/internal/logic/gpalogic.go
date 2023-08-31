package logic

import (
	"context"
	"fmt"

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

func (l *GpaLogic) Gpa() (resp *types.StudentGpaResponse, err error) {
	studentGpa, err := l.svcCtx.StudentGradesModel.GetAllStudentGpa(l.ctx)
	if err != nil {
		return nil, err
	}

	items := make([]types.StudentGpaItem, len(studentGpa))

	for i, student := range studentGpa {
		items[i] = types.StudentGpaItem{
			StudentId: fmt.Sprintf("%d", student.StudentId),
			Gpa:       fmt.Sprintf("%.0f", student.Gpa),
		}
	}

	resp = &types.StudentGpaResponse{
		Results: items,
	}

	return
}
