package svc

import (
	"beyond/application/article/mq/internal/config"
	"beyond/application/article/mq/internal/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	ArticleModel model.ArticleModel
	BizRedis     *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds, err := redis.NewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})
	if err != nil {
		panic(err)
	}

	conn := sqlx.NewMysql(c.Datasource)
	return &ServiceContext{
		Config:       c,
		ArticleModel: model.NewArticleModel(conn),
		BizRedis:     rds,
	}
}
