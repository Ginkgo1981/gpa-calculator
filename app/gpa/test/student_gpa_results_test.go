package main_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestRegistrationOptions(t *testing.T) {
	type args struct {
		path           pathFunc
		body           requestBody
		httpMethod     string
		httpStatusCode int
	}
	path := fmt.Sprintf("%s/students/gpa", prefix)
	tests := []struct {
		name  string
		args  args
		setUp setUpFunc
		want  wantFunc
	}{
		{
			name: "return 404 if path not found",
			args: args{
				path:           func() string { return fmt.Sprintf("%s/students/abc", prefix) },
				httpMethod:     http.MethodGet,
				httpStatusCode: http.StatusNotFound,
			},
			want: func(resp *httpexpect.Response, t *testing.T) {
				resp.Status(http.StatusNotFound)
			},
		},
		{
			name: "return 200 and empty array when no student grades",
			args: args{
				path:           func() string { return path },
				httpMethod:     http.MethodGet,
				httpStatusCode: http.StatusOK,
			},
			want: func(resp *httpexpect.Response, t *testing.T) {
				resp.Status(http.StatusOK)
				resp.JSON().Object().ContainsKey("results").Value("results").Array().Length().Equal(0)
			},
		},
		{
			name: "return 200 and student gpa results",
			args: args{
				path:           func() string { return path },
				httpMethod:     http.MethodGet,
				httpStatusCode: http.StatusOK,
			},
			want: func(resp *httpexpect.Response, t *testing.T) {
				resp.Status(http.StatusOK)
				resp.JSON().Object().ContainsKey("results").Value("results").Array().Length().Equal(2)
			},
			setUp: func() func() {
				generateGradingScales()

				return func() {
					cleaner.Clean("grading_scales", "student_grades")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setUp != nil {
				t.Cleanup(tt.setUp())
			}
			e := httpexpect.New(t, baseURL)
			resp := e.Request(tt.args.httpMethod, tt.args.path()).
				WithHeader("Content-Type", "application/json").
				WithJSON(tt.args.body).
				Expect().
				Status(tt.args.httpStatusCode)
			tt.want(resp, t)
		})
	}
}
