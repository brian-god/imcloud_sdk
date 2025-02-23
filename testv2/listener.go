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

type OnConnListener struct{}

func (c *OnConnListener) OnUserTokenInvalid(errMsg string) {
	//TODO implement me
	panic("implement me")
}

func (c *OnConnListener) OnConnecting() {
	// fmt.Println("OnConnecting")
}

func (c *OnConnListener) OnConnectSuccess() {
	// fmt.Println("OnConnectSuccess")
}

func (c *OnConnListener) OnConnectFailed(errCode int32, errMsg string) {
	// fmt.Println("OnConnectFailed")
}

func (c *OnConnListener) OnKickedOffline() {
	// fmt.Println("OnKickedOffline")
}

func (c *OnConnListener) OnUserTokenExpired() {
	// fmt.Println("OnUserTokenExpired")
}
