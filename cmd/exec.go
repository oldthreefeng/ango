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
	AllMo       string = os.Getenv("DingDingMobiles")
)

const (
	WeiMo  = "17719540702"
	CardMo = "13764645987"
	Adcom  = "15862168925"
	Hudong = "13258335315"
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
		case "weimall/api", "weimall/yj-mall", "weimall/yj-h5", "weimall/plmall", "weimall/yj-admall":
			t.AtMobiles = WeiMo
		case "penglai/adcom", "penglai/www-ypl":
			t.AtMobiles = Adcom
		case "card":
			t.AtMobiles = CardMo
		case "hudong/fotoup-server", "hudong/order-server", "hudong/sign-router", "hudong/service-apis",
		"hudong/sign-server", "hudong/hudong-sign-client", "hudong/service-msite",
		"hudong/hudong-sign-manager", "hudong/service-user":
			t.AtMobiles = Hudong
		default:
			t.AtMobiles = AllMo
		}
	}
	WriteToLog(Type)
	return  t.Dingding(DingDingUrl)
}

func WriteToLog(Type string)  {
	filename := "fabu.log"
	args := strings.Split(Config, ".")[0]
	date := time.Now().Format("2006-01-02 15:04:05")
	data := fmt.Sprintf("[INFO] %s %s-%s %s成功", date, args, Tag, Type)
	file, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	file.WriteString(data)
	defer file.Close()
}

func IsFile(filePath string) bool {
	f, e := os.Stat(filePath)
	if e != nil {
		return false
	}
	return !f.IsDir()
}
