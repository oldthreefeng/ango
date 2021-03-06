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
	Version    = "latest"
	Config     string
	Tag        string
	Author     string
	Comments   string
	Verbose    bool
	Real       bool
	rootCmd    = &cobra.Command{
		Use:   "ango ",
		Short: "ango 是一个用于部署项目至生产环境的部署工具",
		Long: `基于golang开发的一个用于部署项目至生产环境的部署工具
目前仅使用playbook部署/回滚相关业务并使用钉钉的"webhook"通知, 文档查看: https://github.com/oldthreefeng/ango
run "ango list" to find playbook to deploy`,
		Run: func(cmd *cobra.Command, args []string) {
			VersionStr()
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&Author, "author", "", "louis.hong", "author name for copyright attribution")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose mode to see more detail infomation")
	rootCmd.PersistentFlags().StringVarP(&Tag, "tags", "t", "", "tags for the project version")
	rootCmd.PersistentFlags().StringVarP(&Config, "filename", "f", "", "ansible-playbook for yml config(requried)")
	rootCmd.PersistentFlags().StringVarP(&PathName, "path", "p", "", "ango basedir")
	projCmd.PersistentFlags().StringVarP(&Comments, "comments", "m", "", "add comments when send message to dingding")
	rollbackCmd.PersistentFlags().BoolVarP(&Real, "real", "", false, "really to rollback this version")
	rootCmd.AddCommand(projCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(rollbackCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func VersionStr() {
	fmt.Printf(`ango is cli tools to running Ansible playbooks from Golang.
run "ango -h" get more help, more see https://github.com/oldthreefeng/ango
`)
	fmt.Printf("ango version :       %s\n", Version)
	fmt.Printf("Git Commit Hash:     %s\n", Githash)
	fmt.Printf("Build Time :         %s\n", Buildstamp)
	fmt.Printf("Go Version:          %s\n", Goversion)
	fmt.Printf("BuildBy :            %s\n", Author)
}

func init() {
	if PathName == "" {
		PathName = getDefaultPathname("AngoBaseDir", "/opt/playbook/")
	}
	if !FileExist(PathName) {
		err := os.Mkdir(PathName, 0770)
		if err != nil {
			fmt.Println("mkdir ango Base dir err")
			os.Exit(-1)
		}
	}
}

func getDefaultPathname(key, defVal string)  string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}