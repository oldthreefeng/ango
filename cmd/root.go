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
	DingDingToken string = os.Getenv("DingDingToken")
	Config 		  string
	Tag           string
	Author        string
	rootCmd       = &cobra.Command{
		Use:   "ango ",
		Short: "ango 是一个用于部署项目至生产环境的部署工具",
		Long: `基于golang开发的一个用于部署项目至生产环境的部署工具
目前仅使用playbook部署相关业务, 文档查看: https://github.com/oldthreefeng/ango`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("ango 是一个用于部署项目至生产环境的部署工具\nango -h 获取帮助")
		},
	}
)

func init() {
	if DingDingToken == "" {
		DingDingToken = "01bc245b59a337090fca147c123488de188d00cc56e60c77c3c573ddfae655b9"
	}
	rootCmd.PersistentFlags().StringVarP(&Author, "author", "a", "louis.hong", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&Tag, "tag", "t", "", "tags for the project version")
	rootCmd.PersistentFlags().StringVarP(&Config, "config", "f", "", "ansible-playbook for yml config")
	rootCmd.AddCommand(projCmd)
	rootCmd.AddCommand(rollbackCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
