package sdk_user_simulator

import (
	"context"
	"fmt"

	"github.com/OpenIMSDK/tools/log"
	"github.com/brian-god/imcloud_sdk/internal/login"
	"github.com/brian-god/imcloud_sdk/pkg/ccontext"
	"github.com/brian-god/imcloud_sdk/pkg/constant"
	"github.com/brian-god/imcloud_sdk/pkg/utils"
	"github.com/brian-god/imcloud_sdk/sdk_struct"
)

var (
	UserMessageMap = make(map[string]*MsgListenerCallBak)
	timeOffset     int64
)

var (
	TESTIP     = "125.124.195.201"
	APIADDR    = fmt.Sprintf("http://%v:10002", TESTIP)
	WSADDR     = fmt.Sprintf("ws://%v:10001", TESTIP)
	SECRET     = "openIM123"
	PLATFORMID = constant.WindowsPlatformID
	LogLevel   = uint32(5)
)

func SetServerTimeOffset(offset int64) {
	timeOffset = offset
}
func GetRelativeServerTime() int64 {
	return utils.GetCurrentTimestampByMill() + timeOffset
}

func InitSDKAndLogin(userID, token string) error {
	userForSDK := login.NewLoginMgr()
	var cf sdk_struct.IMConfig
	cf.ApiAddr = APIADDR
	cf.PlatformID = int32(PLATFORMID)
	cf.WsAddr = WSADDR
	cf.DataDir = "./"
	cf.LogLevel = LogLevel
	cf.IsExternalExtensions = true
	cf.IsLogStandardOutput = true
	cf.LogFilePath = ""
	var testConnListener testConnListener
	userForSDK.InitSDK(cf, &testConnListener)
	if err := log.InitFromConfig(userID+"_open-im-sdk-core", "", int(LogLevel), true, false, cf.DataDir, 0, 24); err != nil {
		return err
	}
	ctx := ccontext.WithOperationID(userForSDK.BaseCtx(), utils.OperationIDGenerator())
	SetListener(userForSDK, userID)
	err := userForSDK.Login(ctx, userID, token)
	if err != nil {
		return err
	}

	return nil
}

func SetListener(userForSDK *login.LoginMgr, userID string) {
	var testConversation conversationCallBack
	userForSDK.SetConversationListener(&testConversation)

	var testUser userCallback
	userForSDK.SetUserListener(testUser)

	msgCallBack := NewMsgListenerCallBak(userID)
	UserMessageMap[userID] = msgCallBack
	userForSDK.SetAdvancedMsgListener(msgCallBack)

	var friendListener testFriendListener
	userForSDK.SetFriendListener(friendListener)

	var groupListener testGroupListener
	userForSDK.SetGroupListener(groupListener)
}
func CheckMessageDelay(singleMessageCount map[string]int, groupMessageCount map[string]int) {
	ctx := context.Background()
	log.ZDebug(ctx, "chat checking....")
	var sAllDelay, gAllDelay int64
	var sAllPercent, gAllPercent float64
	for userID, bak := range UserMessageMap {
		delay, percent := calculate(singleMessageCount, bak.SingleDelay)
		sAllDelay += delay
		sAllPercent += percent
		log.ZDebug(ctx, fmt.Sprintf("single chat %v delay %v ms,success rate %v/100", userID, delay,
			percent))
		gDelay, gPercent := calculate(groupMessageCount, bak.GroupDelay)
		gAllDelay += gDelay
		gAllPercent += gPercent
		log.ZDebug(ctx, fmt.Sprintf("group chat %v delay %v ms,success rate %v/100", userID, gDelay, gPercent))
	}
	log.ZDebug(ctx, fmt.Sprintf("single chat all delay %v ms,success rate %v/100", sAllDelay/int64(len(UserMessageMap)), sAllPercent/float64(len(UserMessageMap))))
	log.ZDebug(ctx, fmt.Sprintf("group chat all delay %v ms,success rate %v/100", gAllDelay/int64(len(UserMessageMap)), gAllPercent/float64(len(UserMessageMap))))
	log.ZDebug(ctx, fmt.Sprintf("all chat all delay %v ms,success rate %v/100", (sAllDelay+gAllDelay)/2, (sAllPercent+gAllPercent)/float64(2)))

	log.ZDebug(ctx, "chat checking end....")

}

func calculate(singleMessageCount map[string]int, data map[string][]*SingleMessage) (delay int64, percent float64) {
	var allDelay int64
	var SuccessRate float64
	var successCount int
	for sendIDOrGroupID, messages := range data {
		if count, ok := singleMessageCount[sendIDOrGroupID]; ok {
			SuccessRate += float64(len(messages)) / float64(count)
			successCount++
		}
		var singleUserOrGroupDelay int64
		for _, message := range messages {
			singleUserOrGroupDelay += message.Delay
		}
		allDelay += singleUserOrGroupDelay / int64(len(messages))

	}
	return allDelay / int64(len(data)), SuccessRate / float64(successCount) * 100

}
