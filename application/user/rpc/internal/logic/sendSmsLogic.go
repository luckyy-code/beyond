package logic

import (
	"context"

	"beyond/application/user/rpc/internal/svc"
	"beyond/application/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendSmsLogic) SendSms(in *pb.SendSmsRequest) (*pb.SendSmsResponse, error) {
	// TODO send sms
	return &pb.SendSmsResponse{}, nil

}
