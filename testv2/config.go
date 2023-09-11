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

import "open_im_sdk/sdk_struct"

const (
	//APIADDR = "http://43.154.157.177:10002"
	//WSADDR  = "ws://43.154.157.177:10001"
	//UserID  = "kernaltestuid2"

	//APIADDR = "http://localhost:9099"
	//WSADDR  = "ws://localhost:10001"
	//预生产
	APIADDR = "http://8.137.13.1:9099"
	WSADDR  = "ws://8.137.13.1:10001"
	//UserID       = "2688118337"
	//UserID       = "7204255074"
	//UserID = "50122626445611008"
	UserID       = "55122332112392192"
	friendUserID = "3281432310"
	// APIADDR = "http://192.168.44.128:10002"
	// WSADDR  = "ws://192.168.44.128:10001"
	// UserID  = "100"

	//APIADDR = "http://59.36.173.89:10002"
	//WSADDR  = "ws://59.36.173.89:10001"
	//UserID  = "kernaltestuid9"
	token = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJpbUNsb3VkIiwic3ViIjoie1widXNlcklkXCI6XCI1NTEyMjM1NjM3NjQ0MDgzMlwiLFwicGxhdGZvcm1cIjpcIkFuZHJvaWRcIixcInJvbGVcIjpcIlVTRVJcIn0iLCJleHAiOjE2OTQyNzI4MTgsIm5iZiI6MTY5MzkxMjgxOCwiaWF0IjoxNjkzOTEyODE4fQ.NP76Xq4dnEfxZl_DY0frWupF0r063lupcBJV7Q6RY_yIhvVjzP-ZaasMmZc1lGmOYIpwG2ry4bX7fr8rrtyjYA"
)

func getConf(APIADDR, WSADDR string) sdk_struct.IMConfig {
	var cf sdk_struct.IMConfig
	cf.ApiAddr = APIADDR
	cf.WsAddr = WSADDR
	cf.DataDir = "../../"
	cf.LogLevel = 6
	cf.IsExternalExtensions = true
	cf.PlatformID = 2
	cf.LogFilePath = ""
	cf.IsLogStandardOutput = true
	return cf
}
