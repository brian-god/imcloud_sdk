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

//go:build !js
// +build !js

package db

import (
	"context"

	"github.com/brian-god/imcloud_sdk/pkg/db/model_struct"
	"github.com/brian-god/imcloud_sdk/pkg/utils"
)

func (d *DataBase) InsertSendingMessage(ctx context.Context, message *model_struct.LocalSendingMessages) error {
	d.mRWMutex.Lock()
	defer d.mRWMutex.Unlock()
	return utils.Wrap(d.conn.WithContext(ctx).Create(message).Error, "InsertSendingMessage failed")
}

func (d *DataBase) DeleteSendingMessage(ctx context.Context, conversationID, clientMsgID string) error {
	d.groupMtx.Lock()
	defer d.groupMtx.Unlock()
	localSendingMessage := model_struct.LocalSendingMessages{ConversationID: conversationID, ClientMsgID: clientMsgID}
	return utils.Wrap(d.conn.WithContext(ctx).Delete(&localSendingMessage).Error, "DeleteSendingMessage failed")
}
func (d *DataBase) GetAllSendingMessages(ctx context.Context) (friendRequests []*model_struct.LocalSendingMessages, err error) {
	d.friendMtx.Lock()
	defer d.friendMtx.Unlock()
	err = utils.Wrap(d.conn.WithContext(ctx).Find(&friendRequests).Error, "GetAllSendingMessages failed")
	return friendRequests, utils.Wrap(err, "GetAllSendingMessages failed")
}
