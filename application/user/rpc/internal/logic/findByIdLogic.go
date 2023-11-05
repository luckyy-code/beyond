package logic

import (
	"context"

	"beyond/application/user/rpc/internal/svc"
	"beyond/application/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByIdLogic {
	return &FindByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByIdLogic) FindById(in *pb.FindByIdRequest) (*pb.FindByIdResponse, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		logx.Errorf("FindById userId: %d error: %v", in.UserId, err)
		return nil, err
	}

	return &pb.FindByIdResponse{
		UserId:   user.Id,
		Username: user.Username,
		Avatar:   user.Avatar,
	}, nil

	return &pb.FindByIdResponse{}, nil
}
