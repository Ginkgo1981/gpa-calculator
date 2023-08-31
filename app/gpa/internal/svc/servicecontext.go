package svc

import (
	"gpa/app/gpa/internal/config"
	"gpa/app/gpa/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config             config.Config
	StudentGradesModel model.StudentGradesModel
	GradingScalesModel model.GradingScalesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.GpaDB.DataSource)

	return &ServiceContext{
		Config:             c,
		StudentGradesModel: model.NewStudentGradesModel(sqlConn, c.CacheRedis),
		GradingScalesModel: model.NewGradingScalesModel(sqlConn, c.CacheRedis),
	}
}
