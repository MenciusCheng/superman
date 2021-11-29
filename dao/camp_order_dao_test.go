package dao

import (
	"context"
	"github.com/MenciusCheng/superman/conf"
	"github.com/MenciusCheng/superman/util/dragons"
	"testing"
)

var cod *CampOrderDao

func TestMain(m *testing.M) {
	dragons.Init(
		dragons.ConfigPath("/Users/chengmengwei/goProject/superman-config/config.toml"),
	)

	cfg, err := conf.Init()
	if err != nil {
		return
	}
	cod = NewCampOrderDao(cfg)

	m.Run()
}

func TestCampOrderDao_GetOrder(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{ctx: context.Background()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cod.GetOrder(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("GetOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
