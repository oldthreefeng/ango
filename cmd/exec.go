/*
 * Copyright (c) 2019. The ango Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package cmd

import (
	"fmt"
	"github.com/oldthreefeng/ango/play"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var (
	log = logrus.New()
	DingDingUrl string = os.Getenv("DingDingUrl")
	AllMo       string = os.Getenv("DingDingMobiles")
)

const (
	WeiMa  = "17719540702"
	CardMo = "13764645987"
	Adcom  = "15862168925"
	Hudong = "13258335315"
	XiaoKe = "16621186818"
)

func init()  {
	log.SetFormatter(&logrus.JSONFormatter{})
}

func Exec(cmdStr, Type, project string) error {
	// yj-admall.yml ==> yj-admall
	args := strings.Split(FileConfig, ".")[0]
	// find the project name
	//fmt.Printf("%s,%s", args, Config)
	var filename string
	switch runtime.GOOS {
	case "windows":
		filename = "fabu.log"
	case "linux":
		filename = PathName + "/fabu.log"
	default:
		filename = PathName + "/fabu.log"
	}
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		log.Out = file
	} else {
		log.WithFields(logrus.Fields{
			"name": Config,
			"tag": Tag,
			"deployType": Type,
			"comments": Comments,
		}).Info("Failed to log to file, using default stderr")
		return err
	}
	err = execCmd("/bin/sh", "-c", cmdStr)
	if err != nil {
		log.WithFields(logrus.Fields{
			"name": Config,
			"tag": Tag,
			"deployType": Type,
			"comments": Comments,
		}).Fatalf("%s exec err: %s",cmdStr, err)
	}
	var t play.Text
	//t.Title = fmt.Sprintf("%s-%s", args, Tag)
	t.Text = fmt.Sprintf("%s:%s %s 成功, 请测试确认", args, Tag, Type)
	log.WithFields(logrus.Fields{
		"name": args,
		"tag": Tag,
		"deployType": Type,
		"comments": Comments,
	}).Infof("%s exec success,please confirm", cmdStr)

	switch project {
	case "weimall":
		t.AtMobiles = WeiMa
	case "penglai":
		t.AtMobiles = Adcom
	case "card":
		t.AtMobiles = CardMo
	case "hudong":
		t.AtMobiles = Hudong
	case "xiaoke":
		t.AtMobiles = XiaoKe
	default:
		t.AtMobiles = AllMo
	}
	if DingDingUrl == "" {
		return nil
	}
	return t.Dingding(DingDingUrl)
}

func execCmd(name string, arg ...string) error {
	fmt.Println("[os]exec cmd is :", name, arg)
	cmd := exec.Command(name, arg[:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("[os]os call error.", err)
		return err
	}
	return nil
}

//func execCmd(cmdStr string) error {
//	cmd := exec.Command("sh", "-c", cmdStr)
//	stdout, err := cmd.StdoutPipe()
//	if err != nil {
//		return err
//	}
//	if err = cmd.Start(); err != nil {
//		return err
//	}
//
//	for {
//		tmp := make([]byte, 1024)
//		_, err := stdout.Read(tmp)
//		fmt.Print(string(tmp))
//		if err != nil {
//			break
//		}
//	}
//
//	if err = cmd.Wait(); err != nil {
//		return err
//	}
//	return nil
//}

func WriteToLog(Type string) {
	var filename string
	switch runtime.GOOS {
	case "windows":
		filename = "fabu.log"
	case "linux":
		filename = PathName + "/fabu.log"
	default:
		filename = PathName + "/fabu.log"
	}
	args := strings.Split(FileConfig, ".")[0]
	date := time.Now().Format("2006-01-02 15:04:05")
	data := fmt.Sprintf("[INFO] %s %s-%s %s成功\n", date, args, Tag, Type)
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

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}