package logic

import (
	"context"

	"github.com/45427186/DouSheng/rpc/publish/internal/svc"
	"github.com/45427186/DouSheng/rpc/publish/types/publish"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishListLogic) PublishList(in *publish.PublishListReq) (*publish.PublishListRes, error) {
	// todo: add your logic here and delete this line

	return &publish.PublishListRes{}, nil
}
