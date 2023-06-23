package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

// Config 这个文件为主，这个文件写的任何key内容，都需要去etc/yaml里面去找，也就是配置的key以这个文件为主
// 比如下面的CacheRedis123，那么在配置文件里面就是CacheRedis123
type Config struct {
	zrpc.RpcServerConf
	//增加如下
	DataSource    string
	Table         string
	CacheRedis123 cache.CacheConf
}
