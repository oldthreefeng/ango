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

var (
	rollbackCmd = &cobra.Command{
		Use:     "rollback  the project",
		Short:   "rollback the project",
		Example: "api, mall, admall",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("rollback :" + strings.Join(args, " "))
		},
	}
)
