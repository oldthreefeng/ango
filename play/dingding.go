package play

import (
	"bytes"
	"encoding/json"
	"net/http"
)

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
	req := bytes.NewBuffer(data)
	DingDingUrl := "https://oapi.dingtalk.com/robot/send?access_token=" + DingDingToken
	_, err = http.DefaultClient.Post(DingDingUrl, "application/json", req)
	if err != nil {
		return err
	}
	return nil
}
