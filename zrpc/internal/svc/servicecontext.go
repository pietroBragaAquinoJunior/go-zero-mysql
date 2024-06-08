package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-mysql-aula/zrpc/internal/config"
	"go-zero-mysql-aula/zrpc/internal/models"
)

type ServiceContext struct {
	Config          config.Config
	BlocoNotasModel models.BlocoNotasModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSourceName)
	return &ServiceContext{
		Config:          c,
		BlocoNotasModel: models.NewBlocoNotasModel(conn, c.Cache),
	}
}
