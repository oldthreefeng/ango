/*
 * Copyright (c) 2020. The ango Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */
package cmd

import (
	"fmt"
	"strings"
	"testing"
)

func TestWriteToLog(t *testing.T) {
	type args struct {
		Type string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test01",args{"deploy"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteToLog(tt.args.Type)
		})
	}

	var config = "test/test.yml"
	project := strings.Split(config,"/")
	fmt.Println(project[len(project)-1])
	fmt.Println(project[len(project)-2])
}