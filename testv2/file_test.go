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
	"path/filepath"
	"testing"

	"github.com/brian-god/imcloud_sdk/internal/file"
	"github.com/brian-god/imcloud_sdk/open_im_sdk"
)

func TestUploadFile(t *testing.T) {

	fp := `C:\Users\openIM\Desktop\dist.zip`

	resp, err := open_im_sdk.UserForSDK.File().UploadFile(ctx, &file.UploadFileReq{
		Filepath: fp,
		Name:     filepath.Base(fp),
		Cause:    "test",
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)

}
