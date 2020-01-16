/*
 * Copyright (c) 2020. The ango Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var (
	pathName = "/opt/playbook/prod"
	listjCmd    = &cobra.Command{
		Use:     "list [flags]",
		Short:   "to list project",
		Long:    "use ango to lsit project with webhook to dingding",
		Example: "  ango list",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			List(pathName)
		},
	}
)

func List(pathname string ) error {
	f,err := WalkDir(pathname, ".yml")
	if err != nil {
		return err
	}
	for _, v := range f {
		fmt.Println(v)
	}
	return nil
}

func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}

		if fi.IsDir() { // 忽略目录
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}

		return nil
	})

	return files, err
}