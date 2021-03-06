package grpclib

import (
	"context"
	"fmt"
	"gim/pkg/gerrors"
	"gim/pkg/logger"
	"strconv"

	"google.golang.org/grpc/metadata"
)

/********

本文件主要处理上下文，
从metadata中解析出以下参数：
	user_id,device_id,token,request_id

*********/

const (
	CtxUserId    = "user_id"
	CtxDeviceId  = "device_id"
	CtxToken     = "token"
	CtxRequestId = "request_id"
)

func ContextWithRequstId(ctx context.Context, requestId int64) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.Pairs(CtxRequestId, strconv.FormatInt(requestId, 10)))
}

// GetCtxAppId 获取ctx的app_id
func GetCtxRequstId(ctx context.Context) int64 {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0
	}

	requstIds, ok := md[CtxRequestId]
	if !ok && len(requstIds) == 0 {
		return 0
	}
	requstId, err := strconv.ParseInt(requstIds[0], 10, 64)
	if err != nil {
		return 0
	}
	return requstId
}

// GetCtxData 获取ctx的用户数据，依次返回user_id,device_id
func GetCtxData(ctx context.Context) (int64, int64, error) {
	// 从上下文获取用户传过来的metadata, 主要检查用户是否登录
	// 注意：metadata的数据都是字符串，使用时需要转换

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("========>3.1 auth fail: 没有metadata")
		return 0, 0, gerrors.ErrUnauthorized
	}

	var (
		userId   int64
		deviceId int64
		err      error
	)

	userIdStrs, ok := md[CtxUserId]
	if !ok && len(userIdStrs) == 0 {
		fmt.Println("========>3.2 auth fail:没有user_id")
		return 0, 0, gerrors.ErrUnauthorized
	}
	userId, err = strconv.ParseInt(userIdStrs[0], 10, 64)
	if err != nil {
		logger.Sugar.Error(err)
		fmt.Println("========>3.2 auth fail:解析user_id失败")

		return 0, 0, gerrors.ErrUnauthorized
	}

	deviceIdStrs, ok := md[CtxDeviceId]
	if !ok && len(deviceIdStrs) == 0 {
		fmt.Println("========>3.3 auth fail:没有device_id")

		return 0, 0, gerrors.ErrUnauthorized
	}
	deviceId, err = strconv.ParseInt(deviceIdStrs[0], 10, 64)
	if err != nil {
		logger.Sugar.Error(err)
		fmt.Println("========>3.3 auth fail:解析device_id失败")

		return 0, 0, gerrors.ErrUnauthorized
	}
	return userId, deviceId, nil
}

// GetCtxDeviceId 获取ctx的设备id
func GetCtxDeviceId(ctx context.Context) (int64, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("========> auth3.6")
		return 0, gerrors.ErrUnauthorized
	}

	deviceIdStrs, ok := md[CtxDeviceId]
	if !ok && len(deviceIdStrs) == 0 {
		fmt.Println("========> auth3.7")
		return 0, gerrors.ErrUnauthorized
	}
	deviceId, err := strconv.ParseInt(deviceIdStrs[0], 10, 64)
	if err != nil {
		logger.Sugar.Error(err)
		fmt.Println("========> auth3.8")
		return 0, gerrors.ErrUnauthorized
	}
	return deviceId, nil
}

// GetCtxToken 获取ctx的token
func GetCtxToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("========> auth3.9")
		return "", gerrors.ErrUnauthorized
	}

	tokens, ok := md[CtxToken]
	if !ok && len(tokens) == 0 {
		fmt.Println("========> auth3.10")
		return "", gerrors.ErrUnauthorized
	}

	return tokens[0], nil
}

// NewAndCopyRequestId 创建一个context,并且复制RequestId
func NewAndCopyRequestId(ctx context.Context) context.Context {
	newCtx := context.TODO()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return newCtx
	}

	requstIds, ok := md[CtxRequestId]
	if !ok && len(requstIds) == 0 {
		return newCtx
	}
	return metadata.NewOutgoingContext(newCtx, metadata.Pairs(CtxRequestId, requstIds[0]))
}
