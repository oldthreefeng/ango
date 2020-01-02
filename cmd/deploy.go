/*
 * Copyright (c) 2019. The ango Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

const (
	AnsibleBin = "/usr/bin/ansible-playbook "
	Version    = "1.0.0"
)

var (
	DeployType = "deploy"
	projCmd    = &cobra.Command{
		Use:     "deploy [flags]",
		Short:   "to deploy project",
		Long:    "use ango to deploy project with webhook to dingding",
		Example: "  ango deploy -f api.yml -t v1.2.0",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if Config != ""  {
				err := Deploy()
				if err != nil {
					fmt.Println(err)
					return
				}
			} else {
				fmt.Println(`Use "ango deploy -h" to get help `)
			}
		},
	}
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
	cmdStr := fmt.Sprintf("%s %s.yml -e version=%s -f 1 ", AnsibleBin, args, Tag)
	return Exec(cmdStr, "deploy 部署")
}

func Deploy() error {
	var cmdStr string
	if Tag == "" {
		cmdStr = fmt.Sprintf("%s %s  -f 1", AnsibleBin, Config)
	} else  {
		cmdStr = fmt.Sprintf("%s %s -e version=%s -f 1", AnsibleBin, Config, Tag)
	}

	if Detail {
		cmdStr += " -vv"
	}
	return Exec(cmdStr, DeployType)
}
