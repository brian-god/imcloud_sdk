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
	"context"
	"fmt"
	"github.com/imCloud/im/pkg/proto/sdkws"
	"open_im_sdk/open_im_sdk"
	"open_im_sdk/pkg/log"
	"open_im_sdk/pkg/sdk_params_callback"
	"open_im_sdk/pkg/utils"
	"open_im_sdk/sdk_struct"
	"strings"
	"sync"
	"testing"
	"time"
)

type SendCallback struct {
	clientMsgID string
}

func (b *SendCallback) OnError(errCode int32, errMsg string) {
	log.Info("", "!!!!!!!OnError ")
}

func (b *SendCallback) OnSuccess(data string) {
	log.Info("", "!!!!!!!OnSuccess ")
}

func (s *SendCallback) OnProgress(progress int) {
	log.Info("", "上传进度", progress)
}
func Test_GetAllConversationList(t *testing.T) {
	conversations, err := open_im_sdk.UserForSDK.Conversation().GetAllConversationList(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for _, conversation := range conversations {
		t.Log(conversation)
	}
}

func Test_GetConversationListSplit(t *testing.T) {
	conversations, err := open_im_sdk.UserForSDK.Conversation().GetConversationListSplit(ctx, 0, 20)
	if err != nil {
		t.Fatal(err)
	}
	for _, conversation := range conversations {
		t.Log(conversation)
	}
}

//func Test_SetConversationRecvMessageOpt(t *testing.T) {
//	err := open_im_sdk.UserForSDK.Conversation().SetConversationRecvMessageOpt(ctx, []string{"asdasd"}, 1)
//	if err != nil {
//		t.Fatal(err)
//	}
//}

func Test_SetSetGlobalRecvMessageOpt(t *testing.T) {
	err := open_im_sdk.UserForSDK.Conversation().SetGlobalRecvMessageOpt(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}
}

// 隐藏会话
func Test_HideConversation(t *testing.T) {
	err := open_im_sdk.UserForSDK.Conversation().HideConversation(ctx, "asdasd")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_GetConversationRecvMessageOpt(t *testing.T) {
	opts, err := open_im_sdk.UserForSDK.Conversation().GetConversationRecvMessageOpt(ctx, []string{"asdasd"})
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range opts {
		t.Log(v.ConversationID, *v.Result)
	}
}

func Test_GetGlobalRecvMessageOpt(t *testing.T) {
	opt, err := open_im_sdk.UserForSDK.Conversation().
		GetOneConversation(ctx, 3, "27951173210112")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(*opt)
}

func Test_GetGetMultipleConversation(t *testing.T) {
	conversations, err := open_im_sdk.UserForSDK.Conversation().GetMultipleConversation(ctx, []string{"asdasd"})
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range conversations {
		t.Log(v)
	}
}

func Test_DeleteConversation(t *testing.T) {
	err := open_im_sdk.UserForSDK.Conversation().DeleteConversationAndDeleteAllMsg(ctx, "sg_486854059102208")
	if err != nil {
		if !strings.Contains(err.Error(), "no update") {
			t.Fatal(err)
		}
	}
}

func Test_DeleteAllConversationFromLocal(t *testing.T) {
	err := open_im_sdk.UserForSDK.Conversation().DeleteAllConversationFromLocal(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_SetConversationDraft(t *testing.T) {
	err := open_im_sdk.UserForSDK.Conversation().SetConversationDraft(ctx, "group_17729585012", "draft")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_ResetConversationGroupAtType(t *testing.T) {
	err := open_im_sdk.UserForSDK.Conversation().ResetConversationGroupAtType(ctx,
		"sg_27951173210112")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_PinConversation(t *testing.T) {
	err := open_im_sdk.UserForSDK.Conversation().PinConversation(ctx, "group_17729585012", true)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_SetOneConversationPrivateChat(t *testing.T) {
	err := open_im_sdk.UserForSDK.Conversation().SetOneConversationPrivateChat(ctx, "single_3411008330", true)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_SetOneConversationBurnDuration(t *testing.T) {
	err := open_im_sdk.UserForSDK.Conversation().SetOneConversationBurnDuration(ctx, "single_3411008330", 10)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_SetOneConversationRecvMessageOpt(t *testing.T) {
	err := open_im_sdk.UserForSDK.Conversation().SetOneConversationRecvMessageOpt(ctx, "single_3411008330", 1)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_GetTotalUnreadMsgCount(t *testing.T) {
	count, err := open_im_sdk.UserForSDK.Conversation().GetTotalUnreadMsgCount(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(count)
}

func Test_SendMessage(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	msg, _ := open_im_sdk.UserForSDK.Conversation().CreateTextMessage(ctx, "textMsg")
	_, err := open_im_sdk.UserForSDK.Conversation().SendMessage(ctx, msg, "1463426512456", "", nil)
	if err != nil {
		t.Fatal(err)
	}
}
func Test_SendMessage1(t *testing.T) {
	ids := []string{"55122332112392192", "55122332229832704", "55122332330496000", "55122332447936512", "55122332565377024", "55122332682817536", "55122332783480832", "55122332884144128", "55122332984807424", "55122333085470720", "55122333186134016", "55122333286797312"}
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func(userId string) {
			defer wg.Done()
			for i := 0; i < 2000; i++ {
				ctx = context.WithValue(ctx, "callback", TestSendMsg{})
				msg, _ := open_im_sdk.UserForSDK.Conversation().CreateTextMessage(ctx, fmt.Sprintf("textMsg_%d", i))
				open_im_sdk.UserForSDK.Conversation().SendMessage(ctx, msg, userId, "", nil)
				//if err != nil {
				//	t.Fatal(err)
				//}
			}
		}(id)
	}
	wg.Wait()
}
func Test_SendMessageNotOss(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	msg, _ := open_im_sdk.UserForSDK.Conversation().CreateTextMessage(ctx, "textMsg")
	_, err := open_im_sdk.UserForSDK.Conversation().SendMessageNotOss(ctx, msg, "1463426512456", "", nil)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_SendMessageByBuffer(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	msg, _ := open_im_sdk.UserForSDK.Conversation().CreateTextMessage(ctx, "textMsg")
	_, err := open_im_sdk.UserForSDK.Conversation().SendMessageByBuffer(ctx, msg, "3411008330", "", &sdkws.OfflinePushInfo{}, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_FindMessageList(t *testing.T) {
	msgs, err := open_im_sdk.UserForSDK.Conversation().FindMessageList(ctx, []*sdk_params_callback.ConversationArgs{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(msgs.TotalCount)
	for _, v := range msgs.FindResultItems {
		t.Log(v)
	}
}

func Test_GetHistoryMessageList(t *testing.T) {
	msgs, err := open_im_sdk.UserForSDK.Conversation().GetHistoryMessageList(ctx, sdk_params_callback.GetHistoryMessageListParams{})
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range msgs {
		t.Log(v)
	}
}

func Test_GetAdvancedHistoryMessageList(t *testing.T) {
	msgs, err := open_im_sdk.UserForSDK.Conversation().GetAdvancedHistoryMessageList(ctx, sdk_params_callback.GetAdvancedHistoryMessageListParams{
		LastMinSeq:     0,
		UserID:         `55224175421296640`,
		Count:          20,
		ConversationID: `si_55223677259616256_55224175421296640`,
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range msgs.MessageList {
		t.Log(v)
	}
}

func Test_GetAdvancedHistoryMessageListReverse(t *testing.T) {
	msgs, err := open_im_sdk.UserForSDK.Conversation().GetAdvancedHistoryMessageListReverse(ctx, sdk_params_callback.GetAdvancedHistoryMessageListParams{})
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range msgs.MessageList {
		t.Log(v)
	}
}

func Test_InsertSingleMessageToLocalStorage(t *testing.T) {
	_, err := open_im_sdk.UserForSDK.Conversation().InsertSingleMessageToLocalStorage(ctx, &sdk_struct.MsgStruct{}, "3411008330", "")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_InsertGroupMessageToLocalStorage(t *testing.T) {
	_, err := open_im_sdk.UserForSDK.Conversation().InsertGroupMessageToLocalStorage(ctx, &sdk_struct.MsgStruct{}, "group_17729585012", "")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_SearchLocalMessages(t *testing.T) {
	msgs, err := open_im_sdk.UserForSDK.Conversation().SearchLocalMessages(ctx, &sdk_params_callback.SearchLocalMessagesParams{})
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range msgs.SearchResultItems {
		t.Log(v)
	}
}

// // delete
// funcation Test_DeleteMessageFromLocalStorage(t *testing.T) {
//	err := open_im_sdk.UserForSDK.Conversation().DeleteMessageFromLocalStorage(ctx, &sdk_struct.MsgStruct{SessionType: 1, ContentType: 1203,
//		ClientMsgID: "ef02943b05b02d02f92b0e92516099a3", Seq: 31, SendID: "kernaltestuid8", RecvID: "kernaltestuid9"})
//	if err != nil {
//		t.Fatal(err)
//	}
// }
//
// funcation Test_DeleteMessage(t *testing.T) {
//	err := open_im_sdk.UserForSDK.Conversation().DeleteMessage(ctx, &sdk_struct.MsgStruct{SessionType: 1, ContentType: 1203,
//		ClientMsgID: "ef02943b05b02d02f92b0e92516099a3", Seq: 31, SendID: "kernaltestuid8", RecvID: "kernaltestuid9"})
//	if err != nil {
//		t.Fatal(err)
//	}
// }

func Test_DeleteAllMessage(t *testing.T) {
	err := open_im_sdk.UserForSDK.Conversation().DeleteAllMessage(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_DeleteAllMessageFromLocalStorage(t *testing.T) {
	err := open_im_sdk.UserForSDK.Conversation().DeleteAllMessageFromLocalStorage(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_ClearConversationAndDeleteAllMsg(t *testing.T) {
	err := open_im_sdk.UserForSDK.Conversation().ClearConversationAndDeleteAllMsg(ctx, "si_3271407977_7152307910")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_RevokeMessage(t *testing.T) {
	conId := open_im_sdk.GetConversationIDBySessionType("ssss", "1463426512456", 1)
	err := open_im_sdk.UserForSDK.Conversation().RevokeMessage(ctx, conId, "bf70f6d012eb3254c03595cc2c2e0dc2")
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 10)
}

func Test_MarkConversationMessageAsRead(t *testing.T) {
	err := open_im_sdk.UserForSDK.Conversation().MarkConversationMessageAsRead(ctx, "si_2688118337_7249315132")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_MarkMsgsAsRead(t *testing.T) {
	conId := open_im_sdk.GetConversationIDBySessionType("ssss", "50122626445611008", 1)
	err := open_im_sdk.UserForSDK.Conversation().MarkMessagesAsReadByMsgID(ctx, conId, []string{"e664dbd03600e798ea1c2351d6989b10"})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_MarkMessagesAsReadByMsgID(t *testing.T) {
	conId := "si_50122626445611008_50891326056566784" //open_im_sdk.GetConversationIDBySessionType("ssss", "50891326056566784", 1)
	open_im_sdk.MarkMessagesAsReadByMsgID(&SendCallback{}, utils.OperationIDGenerator(), conId, "[\"873e51d049f3aa28fc02bd0f6e0a47bc\",\"d81411a19c13f6783145db74e8e7f229\",\"9c6a75550045cab4cc7128b4de370222\",\"3f6bf3105ef5d8292ab952da54b36505\",\"ef6f904dc5c7fb9a996ee12d9b80c572\",\"db8e2564a911e25a89decced4f7da564\",\"e720dc443fcbed03566467932fbd6496\",\"9fec40158070b27dbf1734a977536571\"]")
	time.Sleep(time.Second * 15)
}

func Test_SendImgMsg(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	msg, err := open_im_sdk.UserForSDK.Conversation().CreateImageMessageFromFullPath(ctx, "/Users/tang/workspace/icon.png")
	if err != nil {
		t.Fatal(err)
	}
	res, err := open_im_sdk.UserForSDK.Conversation().SendMessage(ctx, msg, "49395156675203072", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("send smg => %+v\n", res)
}

func Test_SendFileMsg(t *testing.T) {
	ctx = context.WithValue(ctx, "callback", TestSendMsg{})
	msg, err := open_im_sdk.UserForSDK.Conversation().CreateFileMessageFromFullPath(ctx, "/Users/tang/workspace/weisancloud.zip", "weisancloud.zip")
	if err != nil {
		t.Fatal(err)
	}
	res, err := open_im_sdk.UserForSDK.Conversation().SendMessage(ctx, msg, "49395156675203072", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("send smg => %+v\n", res)
}

func Test_GetConversationIDBySessionType(t *testing.T) {
	conId := open_im_sdk.GetConversationIDBySessionType("ssss", "1463426512456", 1) //open_im_sdk.UserForSDK.Conversation().GetConversationIDBySessionType(context.Background(), "12312", 1)
	t.Logf("send conId => %s\n", conId)
}
