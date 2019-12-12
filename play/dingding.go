/*
 * Copyright (c) 2019. The ango Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */
package play

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	TextTemplate = `{
    "msgtype": "text", 
    "text": {
        "content": "%s"
    }, 
    "at": {
        "atMobiles": [
            "%s"
        ], 
        "isAtAll": false
    }
}`
	LinkTemplate = `{
    "msgtype": "link", 
    "link": {
        "text": "%s", 
        "title": "%s", 
        "picUrl": "http://icons.iconarchive.com/icons/paomedia/small-n-flat/1024/sign-check-icon.png", //jenkins 发布的对勾
        "messageUrl": "%s"
    }
}`
	MarkTemplate = `{
     "msgtype": "markdown",
     "markdown": {
         "title":"%s",
         "text": "%s"
     },
    "at": {
        "atMobiles": [
            "%s"
        ], 
        "isAtAll": false
    }
 }`
)

type Alarm interface {
	Dingding(Dingdingurl string) error
}

type MarkDowning struct {
	Title     string `json:"title"`
	Text      string `json:"text"`
	AtMobiles string `json:"atMobiles"` //应该是[]string,图方便,改成这个

}

type Linking struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

type Text struct {
	MarkDowning
}

func (m Text) Dingding(DingDingUrl string) error {
	baseBody := fmt.Sprintf(TextTemplate, m.Title, m.Text, m.AtMobiles)
	req, err := http.NewRequest("POST", DingDingUrl, strings.NewReader(baseBody))
	if err != nil {
		return err
	}
	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-agent", "firefox")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return err
}

func (m MarkDowning) Dingding(DingDingUrl string) error {
	baseBody := fmt.Sprintf(MarkTemplate, m.Title, m.Text, m.AtMobiles)
	req, err := http.NewRequest("POST", DingDingUrl, strings.NewReader(baseBody))
	if err != nil {
		return err
	}
	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-agent", "firefox")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return err
}

func (m Linking) Dingding(DingDingUrl string) error {
	baseBody := fmt.Sprintf(LinkTemplate, m.Title, m.Text, m.MessageUrl)
	req, err := http.NewRequest("POST", DingDingUrl, strings.NewReader(baseBody))
	if err != nil {
		return err
	}
	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-agent", "firefox")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return err
}
