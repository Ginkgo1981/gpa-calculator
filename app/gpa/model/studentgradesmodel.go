package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StudentGradesModel = (*customStudentGradesModel)(nil)

type StudentGrading struct {
	StudentId int64
	Gpa       float64
}
type (
	// StudentGradesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStudentGradesModel.
	StudentGradesModel interface {
		studentGradesModel
		GetAllStudentGpa(ctx context.Context) ([]*StudentGrading, error)
	}

	customStudentGradesModel struct {
		*defaultStudentGradesModel
	}
)

// NewStudentGradesModel returns a model for the database table.
func NewStudentGradesModel(conn sqlx.SqlConn, c cache.CacheConf) StudentGradesModel {
	return &customStudentGradesModel{
		defaultStudentGradesModel: newStudentGradesModel(conn, c),
	}
}

func (c *customStudentGradesModel) GetAllStudentGpa(ctx context.Context) ([]*StudentGrading, error) {
	var data []*StudentGrading
	sql := "SELECT sg.student_id, SUM(gs.gpa_value) / COUNT(sg.course_id) AS gpa FROM student_grades sg JOIN grading_scales gs ON sg.grade_received = gs.grade GROUP BY sg.student_id ORDER BY sg.student_id ASC;"

	if err := c.QueryRowsNoCacheCtx(ctx, &data, sql); err != nil {
		return nil, err
	}

	return data, nil
}
