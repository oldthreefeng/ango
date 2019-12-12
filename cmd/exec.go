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
	link := play.Linking{}
	link.Msgtype = "link"
	link.Link.Title = fmt.Sprintf("%s-%s", args, Tag)
	link.Link.Text = fmt.Sprintf("%s:%s %s成功", args, Tag, Type)
	link.Link.PicUrl = "http://icons.iconarchive.com/icons/paomedia/small-n-flat/1024/sign-check-icon.png"
	switch args {
	case "api", "yj-mall", "yj-h5":
		link.Link.MessageUrl = MallApiUrl
	case "card":
		link.Link.MessageUrl = CardApiUrl
	case "adcom":
		link.Link.MessageUrl = AdComApiUrl
	case "www-ypl":
		link.Link.MessageUrl = WWWUrl
	case "yj-admall":
		link.Link.MessageUrl = AdMallUrl
	case "plmall":
		link.Link.MessageUrl = PlMall
	default:
		link.Link.MessageUrl = MallApiUrl
	}
	fmt.Println(link)
	err = link.Dingding(DingDingToken)
	if err != nil {
		return err
	}
	err = WriteToLog(Type)
	if err != nil {
		return err
	}
	return nil
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
