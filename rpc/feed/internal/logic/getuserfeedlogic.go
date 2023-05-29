package logic

import (
	"context"

	"github.com/45427186/DouSheng/rpc/feed/internal/svc"
	"github.com/45427186/DouSheng/rpc/feed/types/feed"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFeedLogic {
	return &GetUserFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFeedLogic) GetUserFeed(in *feed.FeedReq) (*feed.FeedRes, error) {
	// todo: add your logic here and delete this line

	return &feed.FeedRes{}, nil
}
