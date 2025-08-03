package dao

import (
	"context"
	"testing"
)

func TestAcessMysqlSelect(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantRsp string
		wantErr bool
	}{// TODO: Add test cases.
		{name: "test1",args: args{ctx: context.Background()}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := AcessMysqlSelect(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AcessMysqlSelect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRsp != tt.wantRsp {
				t.Errorf("AcessMysqlSelect() gotRsp = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}
