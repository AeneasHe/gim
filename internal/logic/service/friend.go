package service

import (
	"context"
	"gim/internal/logic/dao"
	"gim/internal/logic/model"
	"gim/pkg/gerrors"
	"gim/pkg/pb"
	"gim/pkg/rpc"
)

// 好友相关的服务

type friendService struct{}

var FriendService = new(friendService)

// List 获取好友列表
func (s *friendService) List(ctx context.Context, userId int64) ([]*pb.Friend, error) {
	// 查找我的好友，状态是已经同意
	friends, err := dao.FriendDao.List(userId, model.FriendStatusAgree)
	if err != nil {
		return nil, err
	}

	// 取出好友们的id
	userIds := make(map[int64]int32, len(friends))
	for i := range friends {
		userIds[friends[i].FriendId] = 0
	}
	// 批量获取好友们的基本信息
	resp, err := rpc.BusinessIntClient.GetUsers(ctx, &pb.GetUsersReq{UserIds: userIds})
	if err != nil {
		return nil, err
	}

	// 拼接返回数据
	var infos = make([]*pb.Friend, len(friends))
	for i := range friends {
		// 取出好友关系信息
		friend := pb.Friend{
			UserId:  friends[i].FriendId,
			Remarks: friends[i].Remarks,
			Extra:   friends[i].Extra,
		}

		// 取出好友的基础信息
		user, ok := resp.Users[friends[i].FriendId]
		if ok {
			friend.Nickname = user.Nickname
			friend.Sex = user.Sex
			friend.AvatarUrl = user.AvatarUrl
			friend.UserExtra = user.Extra
		}
		// 将结果放入info
		infos[i] = &friend
	}

	return infos, nil
}

// AddFriend 添加好友
func (*friendService) AddFriend(ctx context.Context, userId, friendId int64, remarks, description string) error {
	// 查找好友表中的申请记录
	friend, err := dao.FriendDao.Get(userId, friendId)
	if err != nil {
		return err
	}
	if friend != nil {
		// 已经申请
		if friend.Status == model.FriendStatusApply {
			return nil
		}
		// 已经同意了服务
		if friend.Status == model.FriendStatusAgree {
			return gerrors.ErrAlreadyIsFriend
		}
	}

	// 添加新申请
	err = dao.FriendDao.Add(model.Friend{
		UserId:   userId,
		FriendId: friendId,
		Remarks:  remarks,
		Status:   model.FriendStatusApply,
	})
	if err != nil {
		return err
	}

	// 查询发起申请的用户信息
	resp, err := rpc.BusinessIntClient.GetUser(ctx, &pb.GetUserReq{UserId: userId})
	if err != nil {
		return err
	}

	// 推送服务，将发起申请用户的信息推送给被申请者
	err = PushService.PushToUser(ctx, friendId, pb.PushCode_PC_ADD_FRIEND, &pb.AddFriendPush{
		FriendId:    userId,
		Nickname:    resp.User.Nickname,
		AvatarUrl:   resp.User.AvatarUrl,
		Description: description,
	}, true)
	if err != nil {
		return err
	}
	return nil
}

// AgreeAddFriend 同意添加好友
func (*friendService) AgreeAddFriend(ctx context.Context, userId, friendId int64, remarks string) error {
	// 一段好友关系在数据库中有两条记录：
	// 一条是我发起申请的记录， userId是我，friendId是好友
	// 一条是朋友发起申请的记录，userId是好友，freindId是我
	// 两条记录的状态status都是1时，我们是好友

	// 查找好友申请记录：
	// 【注意】
	// 这里userId是我自己，friendId是好友
	// 申请记录时，发起者是好友，所以这里先查找好友的申请记录
	friend, err := dao.FriendDao.Get(friendId, userId)
	if err != nil {
		return err
	}
	if friend == nil {
		return gerrors.ErrBadRequest
	}

	if friend.Status == model.FriendStatusAgree {
		return nil
	}
	// 更新好友关系状态为同意
	err = dao.FriendDao.UpdateStatus(friendId, userId, model.FriendStatusAgree)
	if err != nil {
		return err
	}

	// 添加好友关系记录：
	// 【注意】
	// 这里添加的申请记录是指我向好友申请的记录
	// 发起者是我，对象是好友，默认是同意
	err = dao.FriendDao.Add(model.Friend{
		UserId:   userId,
		FriendId: friendId,
		Remarks:  remarks,
		Status:   model.FriendStatusAgree,
	})
	if err != nil {
		return err
	}

	// 获取我的信息
	resp, err := rpc.BusinessIntClient.GetUser(ctx, &pb.GetUserReq{UserId: userId})
	if err != nil {
		return err
	}

	// 将我已经同意好友的申请，用消息推送给好友
	err = PushService.PushToUser(ctx, friendId, pb.PushCode_PC_AGREE_ADD_FRIEND, &pb.AgreeAddFriendPush{
		FriendId:  userId,
		Nickname:  resp.User.Nickname,
		AvatarUrl: resp.User.AvatarUrl,
	}, true)
	if err != nil {
		return err
	}
	return nil
}
