package logic

import (
	"context"

	"github.com/45427186/DouSheng/rpc/publish/internal/svc"
	"github.com/45427186/DouSheng/rpc/publish/types/publish"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishActionLogic) PublishAction(in *publish.PublishActionReq) (*publish.PublishActionRes, error) {
	// todo: add your logic here and delete this line

	return &publish.PublishActionRes{}, nil
}
