package service

import (
	"context"
	"fmt"
	"gim/internal/business/cache"
	"gim/internal/business/dao"
	"gim/internal/business/model"
	"gim/pkg/gerrors"
	"gim/pkg/pb"
	"gim/pkg/rpc"
	"time"
)

type authService struct{}

var AuthService = new(authService)

// SignIn 长连接登录
// 登录需要三个参数
// phoneNumber: 手机号码
// code: 手机验证码,开发测试阶段可以任意验证码
// deviceId: 设备id
func (*authService) SignIn(ctx context.Context, phoneNumber, code string, deviceId int64) (bool, int64, string, error) {

	// 检查手机验证码是否正确
	if !Verify(phoneNumber, code) {
		return false, 0, "", gerrors.ErrBadCode
	}

	// 根据手机号，从数据库中查询用户
	fmt.Println("登录手机号:", phoneNumber)
	user, err := dao.UserDao.GetByPhoneNumber(phoneNumber)
	if err != nil {
		return false, 0, "", err
	}

	// 是否新用户
	var isNew = false
	if user == nil {
		// 如果没有找到用户，则认为是新用户
		user = &model.User{PhoneNumber: phoneNumber}
		// 新用户添加到数据库
		id, err := dao.UserDao.Add(*user)
		if err != nil {
			return false, 0, "", err
		}
		user.Id = id
		isNew = true
	}

	// 查询用户的设备
	resp, err := rpc.LogicIntClient.GetDevice(ctx, &pb.GetDeviceReq{DeviceId: deviceId})
	if err != nil {
		return false, 0, "", err
	}

	// 登录token
	//token := util.RandString(40)
	// 开发测试阶段，为了方便token设定为"0"
	token := "0"

	// 缓存中添加用户的认证信息
	err = cache.AuthCache.Set(user.Id, resp.Device.DeviceId, model.Device{
		Type:   resp.Device.Type,
		Token:  token,
		Expire: time.Now().AddDate(0, 3, 0).Unix(),
	})

	if err != nil {
		return false, 0, "", err
	}

	// 返回登录结果
	return isNew, user.Id, token, nil
}

// 检查手机验证码是否正确
func Verify(phoneNumber, code string) bool {
	// 假装他成功了
	return true
}

// Auth 验证用户是否登录
func (*authService) Auth(ctx context.Context, userId, deviceId int64, token string) error {
	fmt.Println("========> auth1")

	device, err := cache.AuthCache.Get(userId, deviceId)
	if err != nil {
		fmt.Println("========> auth1.0")
		return err
	}

	if device == nil {
		fmt.Println("========> auth1.1")
		return gerrors.ErrUnauthorized
	}

	if device.Expire < time.Now().Unix() {
		fmt.Println("========> auth1.2")

		return gerrors.ErrUnauthorized
	}

	if device.Token != token {
		fmt.Println("========> auth1.3")

		return gerrors.ErrUnauthorized
	}
	fmt.Println("========> auth1.4 登录成功")
	return nil
}
