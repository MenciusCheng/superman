package dragons

import "github.com/MenciusCheng/superman/util/golang/sql"

type dragonsConfig struct {
	Server struct {
		Port int `toml:"port"`
	} `toml:"server"`

	Log struct {
		Level string `toml:"level"`
	} `toml:"log"`

	Database []sql.SQLGroupConfig `toml:"database"`
}
