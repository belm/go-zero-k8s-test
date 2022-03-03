package logic

import (
	"context"
	"go-zero-k8s-test/user/rpc/user"

	"go-zero-k8s-test/user/api/internal/svc"
	"go-zero-k8s-test/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddLogic {
	return AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req types.AddRequest) (resp *types.AddResponse, err error) {
	rsp, err := l.svcCtx.UserRpc.Add(l.ctx, &user.AddReq{
		A: req.A,
		B: req.B,
	})

	if err != nil {
		return nil,err
	}

	return &types.AddResponse{
		Sum: rsp.Sum,
	},nil
}
