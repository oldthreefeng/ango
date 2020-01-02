/*
 * Copyright (c) 2019. The ango Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	Buildstamp = ""
	Githash    = ""
	Goversion  = ""
	Config     string
	Tag        string
	Author     string
	Comments   string
	Detail	   bool
	Real       bool
	rootCmd    = &cobra.Command{
		Use:   "ango ",
		Short: "ango 是一个用于部署项目至生产环境的部署工具",
		Long: `基于golang开发的一个用于部署项目至生产环境的部署工具
目前仅使用playbook部署/回滚相关业务并使用钉钉的"webhook"通知, 文档查看: https://github.com/oldthreefeng/ango`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf(`ango is cli tools to running Ansible playbooks from Golang.
run "ango -h" get more help, more see https://github.com/oldthreefeng/ango
`)
			fmt.Printf("ango version :       %s\n", Version)
			fmt.Printf("Git Commit Hash:     %s\n", Githash)
			fmt.Printf("UTC Build Time :     %s\n", Buildstamp)
			fmt.Printf("Go Version:          %s\n", Goversion)
			fmt.Printf("Author :             %s\n", Author)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&Author, "author", "", "louis.hong", "author name for copyright attribution")
	rootCmd.PersistentFlags().BoolVarP(&Detail, "verbose", "v", false, "verbose mode to see more detail infomation")
	rootCmd.PersistentFlags().StringVarP(&Tag, "tags", "t", "", "tags for the project version")
	rootCmd.PersistentFlags().StringVarP(&Config, "filename", "f", "", "ansible-playbook for yml config(requried)")
	projCmd.PersistentFlags().StringVarP(&Comments, "comments", "m", "", "add comments when send message to dingding")
	rollbackCmd.PersistentFlags().BoolVarP(&Real, "real", "r", false, "really to rollback this version")
	rootCmd.AddCommand(projCmd)
	rootCmd.AddCommand(rollbackCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
