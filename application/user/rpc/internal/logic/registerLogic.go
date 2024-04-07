package logic

import (
	"beyond/application/user/rpc/internal/code"
	"beyond/application/user/rpc/internal/model"
	"context"
	"fmt"
	"time"

	"beyond/application/user/rpc/internal/svc"
	"beyond/application/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// 当注册名字为空的时候，返回业务自定义错误码
	if len(in.Username) == 0 {
		fmt.Println("asd")
		return nil, code.RegisterNameEmpty
	}

	ret, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username:   in.Username,
		Mobile:     in.Mobile,
		Avatar:     in.Avatar,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	if err != nil {
		logx.Errorf("Register req: %v error: %v", in, err)
		return nil, err
	}
	userId, err := ret.LastInsertId()
	if err != nil {
		logx.Errorf("LastInsertId error: %v", err)
		return nil, err
	}

	return &pb.RegisterResponse{UserId: userId}, nil

}
