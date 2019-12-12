/*
 * Copyright (c) 2019. The ango Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */
package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	rollbackType = "rollback"
	rollbackCmd  = &cobra.Command{
		Use:     "rollback [flags]",
		Short:   "rollback the project",
		Long:    "rollback 回退版本, 需要指定回退版本的yml文件及要回退的version",
		Args:    cobra.NoArgs,
		Example: "  ango rollback -f roll_api.yml -t v1.2",
		Run: func(cmd *cobra.Command, args []string) {
			if Config != "" && Tag != "" {
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
	cmdStr := fmt.Sprintf("%s %s -e version=%s -f 1", AnsibleBin, Config, Tag)
	if Detail {
		cmdStr += " -vv"
	}
	if Real {
		return Exec(cmdStr, rollbackType)
	} else {
		return errors.New("use extra flags -r or --real to run rollback")
	}

}
