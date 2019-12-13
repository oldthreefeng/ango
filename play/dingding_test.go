/*
 * Copyright (c) 2019. The ango Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package play

import (
	"fmt"
	"os"
	"testing"
)

// TODO
func Test_dingding(t *testing.T) {
	body := fmt.Sprintf(TextTemplate,"hello ango", os.Getenv("DingDingMobiles"))
	fmt.Println(body)
	//type args struct {
	//	DingDingUrl string
	//	baseBody    string
	//}
	//tests := []struct {
	//	name    string
	//	args    args
	//	wantErr bool
	//}{
	//	{"test01",args{
	//		DingDingUrl: os.Getenv("DingDingUrl"),
	//		baseBody:    body,
	//	}, true},
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if err := dingding(tt.args.DingDingUrl, tt.args.baseBody); (err != nil) != tt.wantErr {
	//			t.Errorf("dingding() error = %v, wantErr %v", err, tt.wantErr)
	//		}
	//	})
	//}
}