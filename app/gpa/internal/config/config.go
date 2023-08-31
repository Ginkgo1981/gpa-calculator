package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	GpaDB struct {
		DataSource string
	}

	CacheRedis cache.CacheConf
}
