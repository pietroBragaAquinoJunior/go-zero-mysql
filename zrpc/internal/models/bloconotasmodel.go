package models

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BlocoNotasModel = (*customBlocoNotasModel)(nil)

type (
	// BlocoNotasModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBlocoNotasModel.
	BlocoNotasModel interface {
		blocoNotasModel
	}

	customBlocoNotasModel struct {
		*defaultBlocoNotasModel
	}
)

// NewBlocoNotasModel returns a model for the database table.
func NewBlocoNotasModel(conn sqlx.SqlConn, c cache.CacheConf) BlocoNotasModel {
	return &customBlocoNotasModel{
		defaultBlocoNotasModel: newBlocoNotasModel(conn, c),
	}
}
