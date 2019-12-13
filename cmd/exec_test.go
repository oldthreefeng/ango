/*
 * Copyright (c) 2019. The ango Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package cmd

import "testing"

// test the writeToLog is ok~
func TestWriteToLog(t *testing.T) {
	type args struct {
		Type string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test01", args{"rollback"}, true},
		{"test02", args{"deploy"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteToLog(tt.args.Type); (err != nil) == tt.wantErr {
				t.Errorf("WriteToLog() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
