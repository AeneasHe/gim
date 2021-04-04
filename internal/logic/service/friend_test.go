package service

import (
	"context"
	"testing"
)

func TestAddFriend(t *testing.T) {
	userId := 1
	friendId := 2
	err := FriendService.AddFriend(context.TODO(), int64(userId), int64(friendId), "hello", "new friend")
	if err != nil {
		t.Error("添加好友失败")
	}
}

func TestAgreeAddFriend(t *testing.T) {
	userId := 1
	friendId := 2
	err := FriendService.AgreeAddFriend(context.TODO(), int64(userId), int64(friendId), "hello")
	if err != nil {
		t.Error("添加好友失败")
	}
}
