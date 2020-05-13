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
	PathName = os.Getenv("AngoBaseDir")
	listCmd = &cobra.Command{
		Use:     "list [flags]",
		Short:   "to list project i can deploy with ango",
		Long:    "use ango to list current path yml file, i can deploy with ango",
		Example: "ango list",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			List()
		},
	}
)

func List() error {
	if PathName == "" {
		PathName = "/opt/playbook/prod"
	}
	f, err := WalkDir(PathName, ".yml")
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
		if fi.IsDir() {                                                                  // 忽略目录
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}

		return nil
	})

	return files, err
}

func GetProjectName(config string) (yml, baseYml, baseProject string) {
	if PathName == "" {
		PathName = "/opt/playbook/prod"
	}
	files , _ := WalkDir(PathName,".yml")
	for _, v := range files {
		if  strings.Contains(v, config) {
			yml = v
			base := strings.Split(yml, "/")
			baseYml = base[len(base)-1]
			baseProject = base[len(base)-2]
			break
		}
	}
	return
}
