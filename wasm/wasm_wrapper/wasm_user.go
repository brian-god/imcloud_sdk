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

//go:build js && wasm
// +build js,wasm

package wasm_wrapper

import (
	"syscall/js"

	"github.com/brian-god/imcloud_sdk/open_im_sdk"
	"github.com/brian-god/imcloud_sdk/pkg/utils"
	"github.com/brian-god/imcloud_sdk/wasm/event_listener"
)

// ------------------------------------group---------------------------
type WrapperUser struct {
	*WrapperCommon
}

func NewWrapperUser(wrapperCommon *WrapperCommon) *WrapperUser {
	return &WrapperUser{WrapperCommon: wrapperCommon}
}

func (w *WrapperUser) GetSelfUserInfo(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.GetSelfUserInfo, callback, &args).AsyncCallWithCallback()
}

func (w *WrapperUser) SetSelfInfo(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.SetSelfInfo, callback, &args).AsyncCallWithCallback()
}
func (w *WrapperUser) SetSelfInfoEx(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.SetSelfInfoEx, callback, &args).AsyncCallWithCallback()
}
func (w *WrapperUser) GetUsersInfo(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.GetUsersInfo, callback, &args).AsyncCallWithCallback()
}

func (w *WrapperUser) GetUsersInfoWithCache(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.GetUsersInfoWithCache, callback, &args).AsyncCallWithCallback()
}

func (w *WrapperUser) SubscribeUsersStatus(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.SubscribeUsersStatus, callback, &args).AsyncCallWithCallback()
}

func (w *WrapperUser) UnsubscribeUsersStatus(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.UnsubscribeUsersStatus, callback, &args).AsyncCallWithCallback()
}

func (w *WrapperUser) GetSubscribeUsersStatus(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.GetSubscribeUsersStatus, callback, &args).AsyncCallWithCallback()
}

func (w *WrapperUser) GetUserStatus(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.GetUserStatus, callback, &args).AsyncCallWithCallback()
}
func (w *WrapperUser) AddUserCommand(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.AddUserCommand, callback, &args).AsyncCallWithCallback()
}
func (w *WrapperUser) DeleteUserCommand(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.DeleteUserCommand, callback, &args).AsyncCallWithCallback()
}
func (w *WrapperUser) GetAllUserCommands(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.GetAllUserCommands, callback, &args).AsyncCallWithCallback()
}
