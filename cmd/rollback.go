/*
 * Copyright (c) 2019. The ango Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	rollbackType = "回滚成功"
	rollbackCmd = &cobra.Command{
		Use:     "rollback [flags]",
		Short:   "rollback the project",
		Long:    "rollback 回退版本, 需要指定回退版本的yml文件及要回退的version",
		Example: "  ango rollback -f roll_api.yml -t v1.2",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 && Config != "" && Tag != "" {
				err := RollBack()
				if err != nil {
					fmt.Println(err)
					return
				}
			} else {
				fmt.Println(`Use "ango rollback -h" to get help `)
			}
		},
	}
)

func RollBack() error {
	cmdStr := fmt.Sprintf("%s %s -e version=%s", AnsibleBin, Config, Tag)
	return Exec(cmdStr, rollbackType)
}