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

package network

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"unsafe"

	"github.com/brian-god/imcloud_sdk/pkg/utils"
)

type HttpCli struct {
	httpClient  *http.Client
	httpRequest *http.Request
	Error       error
}

func newHttpClient() *http.Client {
	return &http.Client{Timeout: 30 * time.Second}
}

func PostWithTimeOut(url string, data interface{}, token string, timeout time.Duration) (content []byte, err error) {
	return Post(url).BodyWithJson(data).SetTimeOut(timeout).SetHeader("Authorization", token).ToBytes()
}

func PostWithTimeOutByte(url string, data []byte, token, language string, timeout time.Duration) (content []byte, err error) {
	setTimeOut := Post(url).BodyWithBytes(data).SetTimeOut(timeout)
	header := setTimeOut.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	if language == "" {
		language = "zh"
	}
	header = header.SetHeader("lang", language)
	toBytes, err := header.ToBytes()
	if err != nil {
		if err.Error() == http.ErrHandlerTimeout.Error() || strings.Contains(err.Error(), "network") {
			return []byte{}, errors.New("网络错误！")
		}
	}
	return toBytes, err
}

func Get(url string) *HttpCli {
	request, err := http.NewRequest("GET", url, nil)
	return &HttpCli{
		httpClient:  newHttpClient(),
		httpRequest: request,
		Error:       err,
	}
}

func Post(url string) *HttpCli {
	request, err := http.NewRequest("POST", url, nil)
	return &HttpCli{
		httpClient:  newHttpClient(),
		httpRequest: request,
		Error:       utils.Wrap(err, "newRequest failed, url"),
	}
}

func (c *HttpCli) SetTimeOut(timeout time.Duration) *HttpCli {
	c.httpClient.Timeout = timeout
	return c
}

func (c *HttpCli) SetHeader(key, value string) *HttpCli {
	c.httpRequest.Header.Set(key, value)
	return c
}

func (c *HttpCli) BodyWithJson(obj interface{}) *HttpCli {
	if c.Error != nil {
		return c
	}

	buf, err := json.Marshal(obj)
	if err != nil {
		c.Error = utils.Wrap(err, "marshal failed, url")
		return c
	}
	c.httpRequest.Body = ioutil.NopCloser(bytes.NewReader(buf))
	c.httpRequest.ContentLength = int64(len(buf))
	c.httpRequest.Header.Set("Content-Type", "application/json")
	return c
}

func (c *HttpCli) BodyWithBytes(buf []byte) *HttpCli {
	if c.Error != nil {
		return c
	}

	c.httpRequest.Body = ioutil.NopCloser(bytes.NewReader(buf))
	c.httpRequest.ContentLength = int64(len(buf))
	return c
}

func (c *HttpCli) BodyWithForm(form map[string]string) *HttpCli {
	if c.Error != nil {
		return c
	}

	var value url.Values = make(map[string][]string, len(form))
	for k, v := range form {
		value.Add(k, v)
	}
	buf := Str2bytes(value.Encode())

	c.httpRequest.Body = ioutil.NopCloser(bytes.NewReader(buf))
	c.httpRequest.ContentLength = int64(len(buf))
	c.httpRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return c
}

func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func (c *HttpCli) ToBytes() (content []byte, err error) {
	if c.Error != nil {
		return nil, c.Error
	}

	resp, err := c.httpClient.Do(c.httpRequest)
	if err != nil {
		return nil, utils.Wrap(err, "client.Do failed, url")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, utils.Wrap(errors.New(resp.Status), "status code failed ")
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, utils.Wrap(err, "ioutil.ReadAll failed, url")
	}

	return buf, nil
}

func (c *HttpCli) ToJson(obj interface{}) error {
	if c.Error != nil {
		return c.Error
	}

	resp, err := c.httpClient.Do(c.httpRequest)
	if err != nil {
		return utils.Wrap(err, "client.Do failed, url")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return utils.Wrap(errors.New(resp.Status), "status code failed ")
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return utils.Wrap(err, "ioutil.ReadAll failed, url")
	}
	err = json.Unmarshal(buf, obj)
	if err != nil {
		return utils.Wrap(err, "marshal failed, url")
	}
	return nil
}
