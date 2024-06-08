package logic

import (
	"context"

	"go-zero-mysql-aula/common/pb"
	"go-zero-mysql-aula/zrpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ZrpcMethodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewZrpcMethodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ZrpcMethodLogic {
	return &ZrpcMethodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ZrpcMethodLogic) ZrpcMethod(in *__.ZrpcRequest) (*__.ZrpcResponse, error) {
	// todo: add your logic here and delete this line

	bloco, err := l.svcCtx.BlocoNotasModel.FindOne(l.ctx, "98dab354-a9b0-4b2b-9b3b-2a7c311882ab")
	if err != nil {
		return nil, err
	}

	return &__.ZrpcResponse{Msg: bloco.Titulo + bloco.Descricao}, nil
}
