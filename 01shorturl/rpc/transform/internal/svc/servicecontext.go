package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"shorturl/rpc/transform/internal/config"
	"shorturl/rpc/transform/model"
)

type ServiceContext struct {
	Config config.Config
	//初始化模型，同时把缓存也初始化
	Model model.ShorturlModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//初始化模型，同时把缓存也初始化
		Model: model.NewShorturlModel(sqlx.NewMysql(c.DataSource), c.CacheRedis123),
	}
}
