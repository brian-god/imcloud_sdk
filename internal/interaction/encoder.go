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

package interaction

import (
	"bytes"
	"encoding/gob"

	"github.com/openimsdk/openim-sdk-core/pkg/utils"
)

type Encoder interface {
	Encode(data interface{}) ([]byte, error)
	Decode(encodeData []byte, decodeData interface{}) error
}

type GobEncoder struct {
}

func NewGobEncoder() *GobEncoder {
	return &GobEncoder{}
}
func (g *GobEncoder) Encode(data interface{}) ([]byte, error) {
	buff := bytes.Buffer{}
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
func (g *GobEncoder) Decode(encodeData []byte, decodeData interface{}) error {
	buff := bytes.NewBuffer(encodeData)
	dec := gob.NewDecoder(buff)
	err := dec.Decode(decodeData)
	if err != nil {
		return utils.Wrap(err, "")
	}
	return nil
}
