// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	gradingScalesFieldNames          = builder.RawFieldNames(&GradingScales{})
	gradingScalesRows                = strings.Join(gradingScalesFieldNames, ",")
	gradingScalesRowsExpectAutoSet   = strings.Join(stringx.Remove(gradingScalesFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	gradingScalesRowsWithPlaceHolder = strings.Join(stringx.Remove(gradingScalesFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheGradingScalesIdPrefix = "cache:gradingScales:id:"
)

type (
	gradingScalesModel interface {
		Insert(ctx context.Context, data *GradingScales) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*GradingScales, error)
		Update(ctx context.Context, data *GradingScales) error
		Delete(ctx context.Context, id int64) error
	}

	defaultGradingScalesModel struct {
		sqlc.CachedConn
		table string
	}

	GradingScales struct {
		Id            int64     `db:"id"`
		Grade         string    `db:"grade"` // grade
		MinPercentage float64   `db:"min_percentage"`
		GpaValue      float64   `db:"gpa_value"`
		CreateTime    time.Time `db:"create_time"` // created time
		UpdateTime    time.Time `db:"update_time"` // updated time
	}
)

func newGradingScalesModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultGradingScalesModel {
	return &defaultGradingScalesModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`grading_scales`",
	}
}

func (m *defaultGradingScalesModel) Delete(ctx context.Context, id int64) error {
	gradingScalesIdKey := fmt.Sprintf("%s%v", cacheGradingScalesIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, gradingScalesIdKey)
	return err
}

func (m *defaultGradingScalesModel) FindOne(ctx context.Context, id int64) (*GradingScales, error) {
	gradingScalesIdKey := fmt.Sprintf("%s%v", cacheGradingScalesIdPrefix, id)
	var resp GradingScales
	err := m.QueryRowCtx(ctx, &resp, gradingScalesIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", gradingScalesRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultGradingScalesModel) Insert(ctx context.Context, data *GradingScales) (sql.Result, error) {
	gradingScalesIdKey := fmt.Sprintf("%s%v", cacheGradingScalesIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, gradingScalesRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Grade, data.MinPercentage, data.GpaValue)
	}, gradingScalesIdKey)
	return ret, err
}

func (m *defaultGradingScalesModel) Update(ctx context.Context, data *GradingScales) error {
	gradingScalesIdKey := fmt.Sprintf("%s%v", cacheGradingScalesIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, gradingScalesRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Grade, data.MinPercentage, data.GpaValue, data.Id)
	}, gradingScalesIdKey)
	return err
}

func (m *defaultGradingScalesModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheGradingScalesIdPrefix, primary)
}

func (m *defaultGradingScalesModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", gradingScalesRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultGradingScalesModel) tableName() string {
	return m.table
}
