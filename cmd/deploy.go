/*
 * Copyright (c) 2019. The ango Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */
package cmd

import (
	"fmt"
	"github.com/oldthreefeng/ango/play"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

var (
	projCmd = &cobra.Command{
		Use:     "deploy [flags]",
		Short:   "to deploy project",
		Long:    "use ango to deploy project with webhook to dingding",
		Example: "ango deploy -f api.yml -t v1.2.0",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				err := PlaybookForFile()
				if err != nil {
					return
				}
			}
			for _, arg := range  args {
				err := PlayBook(arg)
				if err != nil {
					return
				}
			}
		},
	}
)

const (
	AnsibleBin  = "/usr/bin/ansible-playbook "
	MallApiUrl  = "https://mall.youpenglai.com/apis/version"
	AdMallUrl   = "https://admall.youpenglai.com"
	AdComApiUrl = "https://ad.youpenglai.com/Public/version"
	CardApiUrl  = "https://card.youpenglai.com/card/nologin/version"
	WWWUrl      = "https://www.youpenglai.com/"
	PlMall      = "https://plmall.youpenglai.com/"
)

func PlayBook(args string) error {
	//运行完了才打印. 不方便查看
	//cmd := AnsibleBin + args + ".yml" + " -e version=" + Tag
	//output, err := exec.Command("sh", "-c", cmd).Output()
	//if err != nil {
	//	return err
	//}
	//fmt.Printf("%s", output)

	if Config != "" {
		args = strings.Split(Config, ".")[0]
	}
	cmdStr := AnsibleBin + args + ".yml" + " -e version=" + Tag
	return Exec(cmdStr)
}

func PlaybookForFile() error {
	cmdStr := AnsibleBin + Config + " -e version=" + Tag
	return Exec(cmdStr)
}

func Exec(cmdStr string) error {
	fmt.Println(cmdStr)
	// yj-admall.yml ==> yj-admall
	if Config == "" {
		Config = strings.Split(cmdStr, " ")[1]
	}
	args := strings.Split(Config, ".")[0]
	fmt.Printf("%s,%s", args, Config)
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
	link.Link.Title = args
	link.Link.Text = args + ":" + Tag + "部署成功"
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
	return nil
}
