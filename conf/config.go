package conf

import (
	"github.com/MenciusCheng/superman/util/dragons"
)

// toml 转 go struct 工具 https://xuri.me/toml-to-go/

type Config struct {
}

func Init() (*Config, error) {
	// parse Config from config file
	cfg := &Config{}
	err := dragons.Scan(cfg)
	return cfg, err
}
