/*
 * Copyright (c) 2019. The ango Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */
package play

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Alarm interface {
	Dingding(DingDingToken string) error
}

type MarkDowning struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
	At struct {
		AtMobiles string `json:"atMobiles"` //应该是[]string,图方便,改成这个
		IsAtAll   bool   `json:"isAtAll"`
	} `json:"at"`
}

type Linking struct {
	Msgtype string `json:"msgtype"`
	Link    struct {
		Text       string `json:"text"`
		Title      string `json:"title"`
		PicUrl     string `json:"picUrl"`
		MessageUrl string `json:"messageUrl"`
	}
}

func (m MarkDowning) Dingding(DingDingToken string) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req := bytes.NewBuffer(data)
	DingDingUrl := "https://oapi.dingtalk.com/robot/send?access_token=" + DingDingToken
	_, err = http.DefaultClient.Post(DingDingUrl, "application/json", req)
	if err != nil {
		return err
	}
	return nil
}

func (m Linking) Dingding(DingDingToken string) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(data)
	DingDingUrl := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", DingDingToken)
	req, err := http.NewRequest("POST", DingDingUrl, reader)
	if err != nil {
		return err
	}
	client := &http.Client{}
	req.Header.Set("content-Typr", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	return nil
}