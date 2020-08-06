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
	//PathName = getDefaultPathname("AngoBaseDir", "/opt/playbook/prod")
	PathName string
	listExample = `
	# first to base where your yml to store .
	# "export AngoBaseDir='/usr/local/'" 
	# use "ango list" to recursion list  all *.yml 
`
	//PathName = os.Getenv("AngoBaseDir")
	listCmd = &cobra.Command{
		Use:     "list [flags]",
		Short:   "to list project i can deploy with ango",
		Long:    "use ango to list current path yml file, i can deploy with ango",
		Example: listExample,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			List()
		},
	}
)


func List() error {
	f := list()
	if f == nil {
		return nil
	}
	for _, v := range f {
		fmt.Println(v)
	}
	return nil
}

func list() (f []string) {
	f, _ = WalkDir(PathName, ".yml")
	f1,_ := WalkDir(PathName, ".yaml")
	f = append(f,f1...)
	return f
}

func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 100)
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
	files := list()
	for _, v := range files {
		if  strings.Contains(v, config) {
			yml = v
			// yml = /opt/playbook/prod/hudong/talk-server.yml
			base := strings.Split(yml, "/")
			// base = [opt,playbook,prod,hudong,talk-server.yml]
			baseYml = base[len(base)-1]
			// baseYml = talk-server.yml
			baseProject = base[len(base)-2]
			// baseProject = hudong
			break
		}
	}
	return
}
