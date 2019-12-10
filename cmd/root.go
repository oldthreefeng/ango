package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)



var (
	project string
	tag string
	author string
	rootCmd = &cobra.Command{
		Use:   						"ango ",
		Short: 						"ango 是一个用于部署项目至生产环境的部署工具",
		Long:						`基于golang开发的一个用于部署项目至生产环境的部署工具
目前仅使用playbook部署相关业务, 文档查看: https://github.com/oldthreefeng/ango`,
		Run: 						func(cmd *cobra.Command, args []string) {
										fmt.Println("ango -h 获取帮助")
									},
	}

	projCmd = &cobra.Command{
		Use:                        "deploy [ some project to deploy ]",
		Short:                      "to deploy  project name",
		Long:                       "to deploy  project name",
		Example:					"api, mall, adamll",
		Args:						cobra.MinimumNArgs(1),
		Run:                        func(cmd *cobra.Command, args []string) {
			fmt.Println("deploy: " + strings.Join(args, " "))
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&author,"author", "a", "louis.hong", "author name for copyright attribution")
	projCmd.PersistentFlags().StringVarP(&project,"project", "p", "api", "project name for deploying")
	projCmd.PersistentFlags().StringVarP(&tag,"tag", "t", "", "tags for the project version")
	rootCmd.AddCommand(projCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}