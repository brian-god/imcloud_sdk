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

package common

import (
	"encoding/json"
	"runtime"

	"github.com/brian-god/imcloud_sdk/open_im_sdk_callback"
	"github.com/brian-god/imcloud_sdk/pkg/sdkerrs"
	"github.com/brian-god/imcloud_sdk/pkg/utils"
)

//var validate *validator.Validate

//funcation init() {
//	validate = validator.New()
//}

func CheckAnyErrCallback(callback open_im_sdk_callback.Base, errCode int32, err error, operationID string) {
	if err != nil {
		errInfo := "operationID[" + operationID + "], " + "info[" + err.Error() + "]"
		//log.NewError(operationID, "checkErr ", errInfo)
		callback.OnError(errCode, errInfo)
		runtime.Goexit()
	}
}
func CheckConfigErrCallback(callback open_im_sdk_callback.Base, err error, operationID string) {
	CheckAnyErrCallback(callback, sdkerrs.ArgsError, err, operationID)
}

//func CheckTokenErrCallback(callback open_im_sdk_callback.Base, err error, operationID string) {
//	CheckAnyErrCallback(callback, sdkerrs.TokenInvalidError, err, operationID)
//}

func CheckDBErrCallback(callback open_im_sdk_callback.Base, err error, operationID string) {
	CheckAnyErrCallback(callback, sdkerrs.SdkInternalError, err, operationID)
}

func CheckDataErrCallback(callback open_im_sdk_callback.Base, err error, operationID string) {
	CheckAnyErrCallback(callback, sdkerrs.SdkInternalError, err, operationID)
}

func CheckArgsErrCallback(callback open_im_sdk_callback.Base, err error, operationID string) {
	CheckAnyErrCallback(callback, sdkerrs.ArgsError, err, operationID)
}

//
//funcation CheckErrAndResp2(err error, resp []byte, output interface{}) error {
//	if err != nil {
//		return utils.Wrap(err, "api resp failed")
//	}
//	var c server_api_params.CommDataResp
//	err = json.Unmarshal(resp, &c)
//	if err == nil {
//		if c.ErrCode != 0 {
//			return utils.Wrap(errors.New(c.ErrMsg), "")
//		}
//		if output != nil {
//			err = mapstructure.Decode(c.Data, output)
//			if err != nil {
//				goto one
//			}
//			return nil
//		}
//		return nil
//	}
//
//	unMarshaler := jsonpb.Unmarshaler{}
//	unMarshaler.Unmarshal()
//	s, _ := marshaler.MarshalToString(pb)
//	out := make(map[string]interface{})
//	json.Unmarshal([]byte(s), &out)
//	if idFix {
//		if _, ok := out["id"]; ok {
//			out["_id"] = out["id"]
//			delete(out, "id")
//		}
//	}
//	return out
//
//one:
//	var c2 server_api_params.CommDataRespOne
//
//	err = json.Unmarshal(resp, &c2)
//	if err != nil {
//		return utils.Wrap(err, "")
//	}
//	if c2.ErrCode != 0 {
//		return utils.Wrap(errors.New(c2.ErrMsg), "")
//	}
//	if output != nil {
//		err = mapstructure.Decode(c2.Data, output)
//		if err != nil {
//			return utils.Wrap(err, "")
//		}
//		return nil
//	}
//	return nil
//}

func JsonUnmarshalAndArgsValidate(s string, args interface{}, callback open_im_sdk_callback.Base, operationID string) error {
	err := json.Unmarshal([]byte(s), args)
	if err != nil {
		if callback != nil {
			//log.NewError(operationID, "Unmarshal failed ", err.Error(), s)
			callback.OnError(sdkerrs.ArgsError, err.Error())
			runtime.Goexit()
		} else {
			return utils.Wrap(err, "json Unmarshal failed")
		}
	}
	//err = validate.Struct(args)
	//if err != nil {
	//	if callback != nil {
	//		log.NewError(operationID, "validate failed ", err.Error(), s)
	//		callback.OnError(constant.ErrArgs.ErrCode, constant.ErrArgs.ErrMsg)
	//		runtime.Goexit()
	//	}
	//}
	//return utils.Wrap(err, "args check failed")
	return nil
}

func JsonUnmarshalCallback(s string, args interface{}, callback open_im_sdk_callback.Base, operationID string) error {
	err := json.Unmarshal([]byte(s), args)
	if err != nil {
		if callback != nil {
			//log.NewError(operationID, "Unmarshal failed ", err.Error(), s)
			callback.OnError(sdkerrs.ArgsError, err.Error())
			runtime.Goexit()
		} else {
			return utils.Wrap(err, "json Unmarshal failed")
		}
	}
	return nil
}
