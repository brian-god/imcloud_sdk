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

package main

import (
	"errors"
	"flag"

	"github.com/OpenIMSDK/tools/log"
	"github.com/brian-god/imcloud_sdk/test"
)

func main() {
	var senderNum *int          //Number of users sending messages
	var singleSenderMsgNum *int //Number of single user send messages
	var intervalTime *int       //Sending time interval, in millisecond

	senderNum = flag.Int("sn", 200, "sender num")
	singleSenderMsgNum = flag.Int("mn", 100, "single sender msg num")
	intervalTime = flag.Int("t", 10, "interval time mill second")
	flag.Parse()
	test.InitMgr(*senderNum)
	log.ZInfo(ctx, "logName", test.LogName, "logLevel", uint32(test.LogLevel))
	log.ZWarn(ctx, "reliability test start ", errors.New(""), "sender num", *senderNum, " single sender msg num", *singleSenderMsgNum, " send msg total num ", *senderNum**singleSenderMsgNum)

	test.ReliabilityTest(*singleSenderMsgNum, *intervalTime, 10, *senderNum)
}
