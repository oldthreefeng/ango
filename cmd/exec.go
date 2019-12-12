/*
 * Copyright (c) 2019. The ango Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package cmd

import (
	"fmt"
	"github.com/oldthreefeng/ango/play"
	"github.com/unknwon/com"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	DingDingUrl string = os.Getenv("DingDingUrl")
	AllMo string = os.Getenv("DingDingMobiles")
)

const (
	WeiMo  = "17719540702"
	CardMo = "13764645987"
	Adcom  = "15862168925"
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

	var t play.Text
	//t.Title = fmt.Sprintf("%s-%s", args, Tag)
	t.Text = fmt.Sprintf("%s:%s %s成功, 请测试确认\n%s", args, Tag, Type, Comments)
	if Type == "rollback" {
		t.AtMobiles = AllMo
	} else {
		switch args {
		case "api", "yj-mall", "yj-h5", "plmall", "yj-admall":
			t.AtMobiles = WeiMo
		case "adcom", "www-ypl":
			t.AtMobiles = Adcom
		case "card":
			t.AtMobiles = CardMo
		default:
			t.AtMobiles = AllMo
		}
	}

	err = t.Dingding(DingDingUrl)
	if err != nil {
		return err
	}
	return WriteToLog(Type)
}

func WriteToLog(Type string) error {
	args := strings.Split(Config, ".")[0]
	date := time.Now().Format("2006-01-02 15:04:05")
	data := fmt.Sprintf("[INFO] %s %s-%s %s成功\n", date, args, Tag, Type)
	return com.WriteFile("fabu.log", []byte(data))
}
