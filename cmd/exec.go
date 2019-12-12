/*
 * Copyright (c) 2019. The ango Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package cmd

import (
	"fmt"
	"github.com/oldthreefeng/ango/play"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	DingDingUrl string = os.Getenv("DingDingUrl")
)


func Exec(cmdStr, Type string) error {
	fmt.Println(cmdStr)
	// yj-admall.yml ==> yj-admall
	args := strings.Split(Config, ".")[0]
	//fmt.Printf("%s,%s", args, Config)
	cmd := exec.Command("sh", "-c", cmdStr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}

	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}

	if err = cmd.Wait(); err != nil {
		return err
	}
	var l play.Linking
	l.Title = fmt.Sprintf("%s-%s", args, Tag)
	l.Text = fmt.Sprintf("%s:%s %s成功", args, Tag, Type)
	switch args {
	case "api", "yj-mall", "yj-h5":
		l.MessageUrl = MallApiUrl
	case "card":
		l.MessageUrl = CardApiUrl
	case "adcom":
		l.MessageUrl = AdComApiUrl
	case "www-ypl":
		l.MessageUrl = WWWUrl
	case "yj-admall":
		l.MessageUrl = AdMallUrl
	case "plmall":
		l.MessageUrl = PlMall
	default:
		l.MessageUrl = MallApiUrl
	}
	if DingDingUrl == "" {
		DingDingUrl = `https://oapi.dingtalk.com/robot/send?access_token=01bc245b59a337090fca147c123488de188d00cc56e60c77c3c573ddfae655b9`
	}
	err = l.Dingding(DingDingUrl)
	if err != nil {
		return err
	}
	err = WriteToLog(Type)
	return err
}

func WriteToLog(Type string) error {
	f, err := os.OpenFile("fabu.log", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	args := strings.Split(Config, ".")[0]
	date := time.Now().Format("2006-01-02 15:04:05")
	_, err = fmt.Fprintf(f, "[INFO] %s %s-%s %s成功\n", date, args, Tag, Type)
	defer f.Close()
	if err != nil {
		return err
	}
	return nil
}
