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

package open_im_sdk_callback

import (
	"context"

	"github.com/brian-god/imcloud_sdk/pkg/db/model_struct"
	"github.com/brian-god/imcloud_sdk/pkg/utils"

	"github.com/OpenIMSDK/tools/log"
)

type OnFriendshipListenerSdk interface {
	OnFriendApplicationAdded(friendApplication model_struct.LocalFriendRequest)
	OnFriendApplicationDeleted(friendApplication model_struct.LocalFriendRequest)
	OnFriendApplicationAccepted(friendApplication model_struct.LocalFriendRequest)
	OnFriendApplicationRejected(friendApplication model_struct.LocalFriendRequest)
	OnFriendAdded(friendInfo model_struct.LocalFriend)
	OnFriendDeleted(friendInfo model_struct.LocalFriend)
	OnFriendInfoChanged(friendInfo model_struct.LocalFriend)
	OnBlackAdded(blackInfo model_struct.LocalBlack)
	OnBlackDeleted(blackInfo model_struct.LocalBlack)
}

type onFriendshipListener struct {
	onFriendshipListener func() OnFriendshipListener
}

func NewOnFriendshipListenerSdk(listener func() OnFriendshipListener) OnFriendshipListenerSdk {
	return &onFriendshipListener{listener}
}

func (o *onFriendshipListener) OnFriendApplicationAdded(friendApplication model_struct.LocalFriendRequest) {
	log.ZDebug(context.Background(), "OnFriendApplicationAdded", "friendApplication", friendApplication)
	o.onFriendshipListener().OnFriendApplicationAdded(utils.StructToJsonString(friendApplication))
}

func (o *onFriendshipListener) OnFriendApplicationDeleted(friendApplication model_struct.LocalFriendRequest) {
	log.ZDebug(context.Background(), "OnFriendApplicationDeleted", "friendApplication", friendApplication)
	o.onFriendshipListener().OnFriendApplicationDeleted(utils.StructToJsonString(friendApplication))
}

func (o *onFriendshipListener) OnFriendApplicationAccepted(friendApplication model_struct.LocalFriendRequest) {
	log.ZDebug(context.Background(), "OnFriendApplicationAccepted", "friendApplication", friendApplication)
	o.onFriendshipListener().OnFriendApplicationAccepted(utils.StructToJsonString(friendApplication))
}

func (o *onFriendshipListener) OnFriendApplicationRejected(friendApplication model_struct.LocalFriendRequest) {
	log.ZDebug(context.Background(), "OnFriendApplicationRejected", "friendApplication", friendApplication)
	o.onFriendshipListener().OnFriendApplicationRejected(utils.StructToJsonString(friendApplication))
}

func (o *onFriendshipListener) OnFriendAdded(friendInfo model_struct.LocalFriend) {
	log.ZDebug(context.Background(), "OnFriendAdded", "friendInfo", friendInfo)
	o.onFriendshipListener().OnFriendAdded(utils.StructToJsonString(friendInfo))
}

func (o *onFriendshipListener) OnFriendDeleted(friendInfo model_struct.LocalFriend) {
	log.ZDebug(context.Background(), "OnFriendDeleted", "friendInfo", friendInfo)
	o.onFriendshipListener().OnFriendDeleted(utils.StructToJsonString(friendInfo))
}

func (o *onFriendshipListener) OnFriendInfoChanged(friendInfo model_struct.LocalFriend) {
	log.ZDebug(context.Background(), "OnFriendInfoChanged", "friendInfo", friendInfo)
	o.onFriendshipListener().OnFriendInfoChanged(utils.StructToJsonString(friendInfo))
}

func (o *onFriendshipListener) OnBlackAdded(blackInfo model_struct.LocalBlack) {
	log.ZDebug(context.Background(), "OnBlackAdded", "blackInfo", blackInfo)
	o.onFriendshipListener().OnBlackAdded(utils.StructToJsonString(blackInfo))
}

func (o *onFriendshipListener) OnBlackDeleted(blackInfo model_struct.LocalBlack) {
	log.ZDebug(context.Background(), "OnBlackDeleted", "blackInfo", blackInfo)
	o.onFriendshipListener().OnBlackDeleted(utils.StructToJsonString(blackInfo))
}
