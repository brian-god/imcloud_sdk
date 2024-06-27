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

package sdk_params_callback

import (
	"github.com/brian-god/imcloud_sdk/pkg/server_api_params"
	"github.com/brian-god/imcloud_sdk/sdk_struct"
)

type FindMessageListParams []*ConversationArgs
type ConversationArgs struct {
	ConversationID  string   `json:"conversationID"`
	ClientMsgIDList []string `json:"clientMsgIDList"`
}
type FindMessageListCallback struct {
	TotalCount      int                           `json:"totalCount"`
	FindResultItems []*SearchByConversationResult `json:"findResultItems"`
}
type GetHistoryMessageListParams struct {
	UserID           string `json:"userID"`
	GroupID          string `json:"groupID"`
	ConversationID   string `json:"conversationID"`
	StartClientMsgID string `json:"startClientMsgID"`
	Count            int    `json:"count"`
}
type GetHistoryMessageListCallback []*sdk_struct.MsgStruct
type GetAdvancedHistoryMessageListParams struct {
	LastMinSeq       int64  `json:"lastMinSeq"`
	ConversationID   string `json:"conversationID"`
	StartClientMsgID string `json:"startClientMsgID"`
	Count            int    `json:"count"`
}
type GetAdvancedHistoryMessageListCallback struct {
	MessageList []*sdk_struct.MsgStruct `json:"messageList"`
	LastMinSeq  int64                   `json:"lastMinSeq"`
	IsEnd       bool                    `json:"isEnd"`
	ErrCode     int32                   `json:"errCode"`
	ErrMsg      string                  `json:"errMsg"`
}

type SetConversationStatusParams struct {
	UserId string `json:"userID" validate:"required"`
	Status int    `json:"status" validate:"required"`
}
type SearchLocalMessagesParams struct {
	ConversationID       string   `json:"conversationID"`
	KeywordList          []string `json:"keywordList"`
	KeywordListMatchType int      `json:"keywordListMatchType"`
	SenderUserIDList     []string `json:"senderUserIDList"`
	MessageTypeList      []int    `json:"messageTypeList"`
	SearchTimePosition   int64    `json:"searchTimePosition"`
	SearchTimePeriod     int64    `json:"searchTimePeriod"`
	PageIndex            int      `json:"pageIndex"`
	Count                int      `json:"count"`
}
type SearchLocalMessagesCallback struct {
	TotalCount        int                           `json:"totalCount"`
	SearchResultItems []*SearchByConversationResult `json:"searchResultItems"`
}
type SearchByConversationResult struct {
	ConversationID    string                  `json:"conversationID"`
	ConversationType  int32                   `json:"conversationType"`
	ShowName          string                  `json:"showName"`
	FaceURL           string                  `json:"faceURL"`
	LatestMsgSendTime int64                   `json:"latestMsgSendTime,omitempty"`
	MessageCount      int                     `json:"messageCount"`
	MessageList       []*sdk_struct.MsgStruct `json:"messageList"`
}
type SetMessageReactionExtensionsParams []*server_api_params.KeyValue

type SetMessageReactionExtensionsCallback struct {
	Key     string `json:"key" validate:"required"`
	Value   string `json:"value" validate:"required"`
	ErrCode int32  `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}

type AddMessageReactionExtensionsParams []*server_api_params.KeyValue

type AddMessageReactionExtensionsCallback struct {
	Key     string `json:"key" validate:"required"`
	Value   string `json:"value" validate:"required"`
	ErrCode int32  `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}
type DeleteMessageReactionExtensionsParams []string

type GetTypekeyListResp struct {
	TypeKeyInfoList []*SingleTypeKeyInfoSum `json:"TypeKeyListInfo"`
}
type SingleTypeKeyInfoSum struct {
	TypeKey       string  `json:"typeKey"`
	Counter       int64   `json:"counter"`
	InfoList      []*Info `json:"infoList"`
	IsContainSelf bool    `json:"isContainSelf"`
}

type SingleTypeKeyInfo struct {
	TypeKey     string           `json:"typeKey"`
	Counter     int64            `json:"counter"`
	IsCanRepeat bool             `json:"isCanRepeat"`
	Index       int              `json:"index"`
	InfoList    map[string]*Info `json:"infoList"`
}
type Info struct {
	UserID string `json:"userID"`
	Ex     string `json:"ex"`
}
