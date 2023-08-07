// Copyright © 2023 OpenIM SDK. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testv2

import (
	"open_im_sdk/open_im_sdk"
	"open_im_sdk/pkg/constant"
	"open_im_sdk/pkg/sdk_params_callback"
	"testing"
	"time"

	groupv1 "github.com/imCloud/api/group/v1"
)

func Test_CreateGroupV2(t *testing.T) {
	req := &groupv1.CrateGroupReq{
		MemberList:   []string{"1463426311015", "1463426512456", "1463426515762", "1463426574029"},
		GroupName:    "圣光裁决",
		GroupType:    2,
		Notification: "公告：这是一个荣誉",
		Introduction: "洗脑群",
		FaceURL:      "https://dfsjk/djfhsd/5d1f5562d/154452.jpg",
	}
	info, err := open_im_sdk.UserForSDK.Group().CreateGroup(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("group info: %s", info.String())
}

func Test_JoinGroup(t *testing.T) {
	err := open_im_sdk.UserForSDK.Group().JoinGroup(ctx,
		"50069143293952",
		"进群收钱呀",
		1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("JoinGroup success")
}

func Test_QuitGroup(t *testing.T) {
	err := open_im_sdk.UserForSDK.Group().QuitGroup(ctx, "106366035300352")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("QuitGroup success")
}

func Test_DismissGroup(t *testing.T) {
	err := open_im_sdk.UserForSDK.Group().DismissGroup(ctx, "50069143293952")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("DismissGroup success")
}

func Test_ChangeGroupMute(t *testing.T) {
	err := open_im_sdk.UserForSDK.Group().ChangeGroupMute(ctx, "50069143293952", true)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ChangeGroupMute success", ctx.Value("operationID"))
}

func Test_CancelMuteGroup(t *testing.T) {
	err := open_im_sdk.UserForSDK.Group().ChangeGroupMute(ctx, "50069143293952", false)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ChangeGroupMute success", ctx.Value("operationID"))
}

func Test_ChangeGroupMemberMute(t *testing.T) {
	err := open_im_sdk.UserForSDK.Group().ChangeGroupMemberMute(
		ctx,
		"50069143293952",
		"1463426574029",
		10000)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ChangeGroupMute success", ctx.Value("operationID"))
}

func Test_CancelChangeGroupMemberMute(t *testing.T) {
	err := open_im_sdk.UserForSDK.Group().ChangeGroupMemberMute(ctx,
		"50069143293952",
		"1463426574029",
		0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("CancelChangeGroupMemberMute success", ctx.Value("operationID"))
}

func Test_SetGroupMemberRoleLevel(t *testing.T) {
	// 1:普通成员 2:群主 3:管理员
	err := open_im_sdk.UserForSDK.Group().SetGroupMemberRoleLevel(
		ctx,
		"50069143293952",
		"1463426574029",
		3)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("SetGroupMemberRoleLevel success", ctx.Value("operationID"))
}

func Test_SetGroupMemberNickname(t *testing.T) {
	err := open_im_sdk.UserForSDK.Group().SetGroupMemberNickname(
		ctx,
		"50069143293952",
		"1463426574029",
		"不灭之祸")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("SetGroupMemberNickname success", ctx.Value("operationID"))
}

func Test_SetGroupMemberInfo(t *testing.T) {
	// 1:普通成员 2:群主 3:管理员
	err := open_im_sdk.UserForSDK.Group().SetGroupMemberInfo(ctx, &groupv1.SetGroupMemberInfo{
		GroupID:  "50069143293952",
		UserID:   "1463426574029",
		FaceURL:  "https://doc.rentsoft.cn/images/logo.png",
		Nickname: "熔火之心",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("SetGroupMemberNickname success", ctx.Value("operationID"))
}

func Test_GetJoinedGroupList(t *testing.T) {
	info, err := open_im_sdk.UserForSDK.Group().GetJoinedGroupList(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("GetJoinedGroupList: %d\n", len(info))
	for _, localGroup := range info {
		t.Logf("%#v", localGroup)
	}
}

func Test_GetSpecifiedGroupsInfo(t *testing.T) {
	info, err := open_im_sdk.UserForSDK.Group().GetSpecifiedGroupsInfo(
		ctx, []string{"50069143293952"})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("GetGroupsInfo: %d\n", len(info))
	for _, localGroup := range info {
		t.Logf("%#v", localGroup)
	}
}

func Test_GetGroupApplicationListAsRecipient(t *testing.T) {
	info, err := open_im_sdk.UserForSDK.Group().GetGroupApplicationListAsRecipient(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("GetRecvGroupApplicationList: %d\n", len(info))
	for _, localGroup := range info {
		t.Logf("%#v", localGroup)
	}
}

func Test_GetGroupApplicationListAsApplicant(t *testing.T) {
	info, err := open_im_sdk.UserForSDK.Group().GetGroupApplicationListAsApplicant(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("GetSendGroupApplicationList: %d\n", len(info))
	for _, localGroup := range info {
		t.Logf("%#v", localGroup)
	}
}

func Test_AcceptGroupApplication(t *testing.T) {
	err := open_im_sdk.UserForSDK.Group().AcceptGroupApplication(ctx,
		"50069143293952",
		"1463426515762",
		"test accept")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("AcceptGroupApplication success", ctx.Value("operationID"))
}

func Test_RefuseGroupApplication(t *testing.T) {
	t.Log("operationID:", ctx.Value("operationID"))
	err := open_im_sdk.UserForSDK.Group().RefuseGroupApplication(
		ctx,
		"171491979169792",
		"1463426311015",
		"test refuse")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("AcceptGroupApplication success")
}

func Test_HandlerGroupApplication(t *testing.T) {
	err := open_im_sdk.UserForSDK.Group().HandlerGroupApplication(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("AcceptGroupApplication success", ctx.Value("operationID"))
}

func Test_SearchGroupMembers(t *testing.T) {
	info, err := open_im_sdk.UserForSDK.Group().SearchGroupMembers(ctx, &sdk_params_callback.SearchGroupMembersParam{
		GroupID:                "171491979169792",
		KeywordList:            []string{"之"},
		IsSearchUserID:         false,
		IsSearchMemberNickname: true,
		Offset:                 0,
		Count:                  10,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("SearchGroupMembers: %d\n", len(info))
	for _, localGroup := range info {
		t.Logf("%#v", localGroup)
	}
}

func Test_KickGroupMember(t *testing.T) {
	err := open_im_sdk.UserForSDK.Group().KickGroupMember(
		ctx,
		"171491979169792",
		"我要踢人",
		[]string{"1463426512456"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("InviteUserToGroup success", ctx.Value("operationID"))
}

func Test_TransferGroupOwner(t *testing.T) {
	err := open_im_sdk.UserForSDK.Group().TransferGroupOwner(
		ctx,
		"171491979169792",
		"1463426574029")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TransferGroupOwner success", ctx.Value("operationID"))
}

func Test_InviteUserToGroup(t *testing.T) {
	err := open_im_sdk.UserForSDK.Group().InviteUserToGroup(ctx, "3459296007", "test", []string{"45644221123"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("InviteUserToGroup success", ctx.Value("operationID"))
}

func Test_SyncGroup(t *testing.T) {
	err := open_im_sdk.UserForSDK.Group().SyncGroupMember(ctx, "3179997540")
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 100000)
}

func Test_GetGroup(t *testing.T) {
	t.Log("--------------------------")
	infos, err := open_im_sdk.UserForSDK.Group().GetSpecifiedGroupsInfo(ctx, []string{"3179997540"})
	if err != nil {
		t.Fatal(err)
	}
	for i, info := range infos {
		t.Logf("%d: %#v", i, info)
	}
	//time.Sleep(time.Second * 100000)
}

func Test_IsJoinGroup(t *testing.T) {
	t.Log("--------------------------")
	join, err := open_im_sdk.UserForSDK.Group().IsJoinGroup(ctx, "1875806101")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("join:", join)
}

func Test_GetGroupMemberList(t *testing.T) {
	t.Log("--------------------------")
	m := map[int32]string{
		constant.GroupOwner:         "群主",
		constant.GroupAdmin:         "管理员",
		constant.GroupOrdinaryUsers: "成员",
	}

	members, err := open_im_sdk.UserForSDK.Group().GetGroupMemberList(ctx, "2246086342", 0, 0, 9999999)
	if err != nil {
		panic(err)
	}
	for i, member := range members {
		name := m[member.RoleLevel]
		t.Log(i, member.UserID, member.Nickname, name)
	}

	t.Log("--------------------------")
}
