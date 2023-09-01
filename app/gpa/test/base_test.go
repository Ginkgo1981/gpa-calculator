package main_test

import (
	"context"
	"gpa/app/gpa/internal/config"
	"gpa/app/gpa/internal/handler"
	"gpa/app/gpa/internal/svc"
	"gpa/app/gpa/model"
	"path/filepath"
	"testing"
	"time"

	"github.com/gavv/httpexpect"
	"github.com/khaiql/dbcleaner"
	"github.com/khaiql/dbcleaner/engine"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

type setUpFunc func() func()
type pathFunc func() string
type wantFunc func(*httpexpect.Response, *testing.T)
type requestBody map[string]any

var testServer *rest.Server
var testSvcCtx *svc.ServiceContext
var baseURL = "http://localhost:8888"
var prefix = "/api/v1"
var cleaner = dbcleaner.New()

func pkgSetUp() func() {
	path := filepath.Join("./", "gpa-api.test.yaml")

	var c config.Config
	conf.MustLoad(path, &c, conf.UseEnv())
	testSvcCtx = svc.NewServiceContext(c)
	testServer = rest.MustNewServer(c.RestConf)
	handler.RegisterHandlers(testServer, testSvcCtx)
	mysql := engine.NewMySQLEngine(c.GpaDB.DataSource)
	cleaner.SetEngine(mysql)

	go testServer.Start()

	return func() {
		testServer.Stop()
		cleaner.Clean("student_grades", "grading_scales")
	}
}

func generateGradingScales() {
	testSvcCtx.GradingScalesModel.Insert(context.TODO(), &model.GradingScales{Grade: "A", MinPercentage: 80.0, GpaValue: 4.0})
	testSvcCtx.GradingScalesModel.Insert(context.TODO(), &model.GradingScales{Grade: "B", MinPercentage: 80.0, GpaValue: 4.0})
	testSvcCtx.StudentGradesModel.Insert(context.TODO(), &model.StudentGrades{StudentId: 1, CourseId: 1, GradeReceived: "A"})
	testSvcCtx.StudentGradesModel.Insert(context.TODO(), &model.StudentGrades{StudentId: 2, CourseId: 2, GradeReceived: "B"})
}

func TestMain(m *testing.M) {
	defer pkgSetUp()()

	time.Sleep(1 * time.Second)
	m.Run()
}
