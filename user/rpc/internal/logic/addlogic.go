package logic

import (
	"context"

	"go-zero-k8s-test/user/rpc/internal/svc"
	"go-zero-k8s-test/user/rpc/userpb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *userpb.AddReq) (*userpb.AddResp, error) {
	// todo: add your logic here and delete this line

	return &userpb.AddResp{}, nil
}
