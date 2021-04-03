package api

import (
	"context"
	"fmt"
	"gim/pkg/pb"
	"gim/pkg/util"
	"strconv"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// 获取客户端
func getLogicExtClient() pb.LogicExtClient {
	conn, err := grpc.Dial("127.0.0.1:50001", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return pb.NewLogicExtClient(conn)
}

// 上下文，即登录的身份，设备，token等信息
func getCtx() context.Context {
	token := "0"
	return metadata.NewOutgoingContext(context.TODO(), metadata.Pairs(
		"user_id", "2",
		"device_id", "2",
		"token", token,
		"request_id", strconv.FormatInt(time.Now().UnixNano(), 10)))
}

// 测试注册设备
func TestLogicExtServer_RegisterDevice(t *testing.T) {
	resp, err := getLogicExtClient().RegisterDevice(context.TODO(),
		&pb.RegisterDeviceReq{
			Type:          1,
			Brand:         "huawei",
			Model:         "huawei P30",
			SystemVersion: "1.0.0",
			SdkVersion:    "1.0.0",
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

// 测试发送消息
func TestLogicExtServer_SendMessage(t *testing.T) {
	buf, err := proto.Marshal(&pb.Text{
		Text: "hello alber 4",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	// 用getCtx()得到的身份2，给身份1的用户发送消息
	resp, err := getLogicExtClient().SendMessage(getCtx(),
		&pb.SendMessageReq{
			ReceiverType:   pb.ReceiverType_RT_USER,
			ReceiverId:     1, // 消息接收者，身份1
			ToUserIds:      nil,
			MessageType:    pb.MessageType_MT_TEXT,
			MessageContent: buf,
			IsPersist:      true,
			SendTime:       util.UnixMilliTime(time.Now()),
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func TestLogicExtServer_SendImageMessage(t *testing.T) {
	buf, err := proto.Marshal(&pb.Image{
		Id:           "",
		Width:        0,
		Height:       0,
		Url:          "https://img.iplaysoft.com/wp-content/uploads/2019/free-images/free_stock_photo.jpg",
		ThumbnailUrl: "",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := getLogicExtClient().SendMessage(getCtx(),
		&pb.SendMessageReq{
			ReceiverType:   pb.ReceiverType_RT_USER,
			ReceiverId:     1,
			ToUserIds:      nil,
			MessageType:    pb.MessageType_MT_IMAGE,
			MessageContent: buf,
			IsPersist:      true,
			SendTime:       util.UnixMilliTime(time.Now()),
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func TestLogicExtServer_CreateGroup(t *testing.T) {
	resp, err := getLogicExtClient().CreateGroup(getCtx(),
		&pb.CreateGroupReq{
			Name:         "10",
			Introduction: "10",
			Type:         1,
			Extra:        "10",
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func TestLogicExtServer_UpdateGroup(t *testing.T) {
	resp, err := getLogicExtClient().UpdateGroup(getCtx(),
		&pb.UpdateGroupReq{
			GroupId:      2,
			Name:         "11",
			Introduction: "11",
			Extra:        "11",
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func TestLogicExtServer_GetGroup(t *testing.T) {
	resp, err := getLogicExtClient().GetGroup(getCtx(),
		&pb.GetGroupReq{
			GroupId: 2,
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func TestLogicExtServer_GetUserGroups(t *testing.T) {
	resp, err := getLogicExtClient().GetGroups(getCtx(), &pb.GetGroupsReq{})
	if err != nil {
		fmt.Println(err)
		return
	}
	// todo 不能获取用户所在的超大群组
	fmt.Printf("%+v\n", resp)
}

func TestLogicExtServer_AddGroupMember(t *testing.T) {
	resp, err := getLogicExtClient().AddGroupMembers(getCtx(),
		&pb.AddGroupMembersReq{
			GroupId: 2,
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func TestLogicExtServer_UpdateGroupMember(t *testing.T) {
	resp, err := getLogicExtClient().UpdateGroupMember(getCtx(),
		&pb.UpdateGroupMemberReq{
			GroupId: 2,
			UserId:  3,
			Remarks: "2",
			Extra:   "2",
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func TestLogicExtServer_DeleteGroupMember(t *testing.T) {
	resp, err := getLogicExtClient().DeleteGroupMember(getCtx(),
		&pb.DeleteGroupMemberReq{
			GroupId: 10,
			UserId:  1,
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func TestLogicExtServer_GetGroupMembers(t *testing.T) {
	resp, err := getLogicExtClient().GetGroupMembers(getCtx(),
		&pb.GetGroupMembersReq{
			GroupId: 2,
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}
