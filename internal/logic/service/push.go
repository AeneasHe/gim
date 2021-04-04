package service

import (
	"context"
	"gim/internal/logic/model"
	"gim/pkg/gerrors"
	"gim/pkg/grpclib"
	"gim/pkg/logger"
	"gim/pkg/pb"
	"gim/pkg/util"
	"time"

	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
)

// 推送服务
type pushService struct{}

var PushService = new(pushService)

// 消息推送给用户
func (s *pushService) PushToUser(ctx context.Context, userId int64, code pb.PushCode, message proto.Message, isPersist bool) error {
	// 记录日志
	logger.Logger.Debug("push",
		zap.Int64("request_id", grpclib.GetCtxRequstId(ctx)),
		zap.Int64("user_id", userId),
		zap.Int32("code", int32(code)),
		zap.Any("message", message))

	// 消息序列化
	messageBuf, err := proto.Marshal(message)
	if err != nil {
		return gerrors.WrapError(err)
	}

	// 命令序列化
	commandBuf, err := proto.Marshal(&pb.Command{Code: int32(code), Data: messageBuf})
	if err != nil {
		return gerrors.WrapError(err)
	}

	// 发送给用户
	_, err = MessageService.SendToUser(ctx,
		model.Sender{
			SenderType: pb.SenderType_ST_SYSTEM,
			SenderId:   0,
			DeviceId:   0,
		},
		userId,
		&pb.SendMessageReq{
			ReceiverType:   pb.ReceiverType_RT_USER,
			ReceiverId:     userId,
			ToUserIds:      nil,
			MessageType:    pb.MessageType_MT_COMMAND,
			MessageContent: commandBuf,
			SendTime:       util.UnixMilliTime(time.Now()),
			IsPersist:      isPersist,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// 消息推送给群
func (s *pushService) PushToGroup(ctx context.Context, groupId int64, code pb.PushCode, message proto.Message, isPersist bool) error {
	logger.Logger.Debug("push_to_group",
		zap.Int64("request_id", grpclib.GetCtxRequstId(ctx)),
		zap.Int64("group_id", groupId),
		zap.Int32("code", int32(code)),
		zap.Any("message", message))

	messageBuf, err := proto.Marshal(message)
	if err != nil {
		return gerrors.WrapError(err)
	}

	commandBuf, err := proto.Marshal(&pb.Command{Code: int32(code), Data: messageBuf})
	if err != nil {
		return gerrors.WrapError(err)
	}

	_, err = MessageService.SendToGroup(ctx,
		model.Sender{
			SenderType: pb.SenderType_ST_SYSTEM,
			SenderId:   0,
			DeviceId:   0,
		},
		&pb.SendMessageReq{
			ReceiverType:   pb.ReceiverType_RT_SMALL_GROUP,
			ReceiverId:     groupId,
			ToUserIds:      nil,
			MessageType:    pb.MessageType_MT_COMMAND,
			MessageContent: commandBuf,
			SendTime:       util.UnixMilliTime(time.Now()),
			IsPersist:      isPersist,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
