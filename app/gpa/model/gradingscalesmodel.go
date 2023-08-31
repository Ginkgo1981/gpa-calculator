package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GradingScalesModel = (*customGradingScalesModel)(nil)

type (
	// GradingScalesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGradingScalesModel.
	GradingScalesModel interface {
		gradingScalesModel
	}

	customGradingScalesModel struct {
		*defaultGradingScalesModel
	}
)

// NewGradingScalesModel returns a model for the database table.
func NewGradingScalesModel(conn sqlx.SqlConn, c cache.CacheConf) GradingScalesModel {
	return &customGradingScalesModel{
		defaultGradingScalesModel: newGradingScalesModel(conn, c),
	}
}
