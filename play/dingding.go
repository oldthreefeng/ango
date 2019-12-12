/*
 * Copyright (c) 2019. The ango Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */
package play

import (
	"bytes"
	"fmt"
	"net/http"
)


type Alarm interface {
	Dingding(Dingdingurl string) error
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

func (m MarkDowning) Dingding(DingDingUrl string) error {
	baseBody:= fmt.Sprintf(`{"msgtype": "markdown", 
    "markdown": {
        "title": "%s",
        "text": "%s"
     },
	"at":{
       "atMobiles": %s
     },
     "isAtAll": false
	}`,m.Markdown.Title, m.Markdown.Text, m.At.AtMobiles)
	req := bytes.NewBuffer([]byte(baseBody))
	_, err := http.DefaultClient.Post(DingDingUrl, "application/json", req)
	return err
}

func (m Linking) Dingding( DingDingUrl string) error {
	baseBody := fmt.Sprintf(`{"msgtype": "link", 
    "link": {
        "title": "%s",
        "text": "%s"
        "messageUrl": "%s"
        "picUrl": "http://icons.iconarchive.com/icons/paomedia/small-n-flat/1024/sign-check-icon.png"
     }`,m.Link.Title, m.Link.Text, m.Link.MessageUrl)
	req := bytes.NewBuffer([]byte(baseBody))
	_, err := http.DefaultClient.Post(DingDingUrl, "application/json", req)
	return err
}
