package logic

import (
	"context"

	"go-zero-k8s-test/user/api/internal/svc"
	"go-zero-k8s-test/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetinfoLogic {
	return GetinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetinfoLogic) Getinfo(req types.GetInfoRequest) (resp *types.GetInfoResponse, err error) {
	if req.Uid != 1 {
		return nil,err
	}

	return &types.GetInfoResponse{
		Uid: 1,
		Username: "belm",
		Age: 18,
	}, nil
}
