package dragons

type dragonsConfig struct {
	Server struct {
		Port int `toml:"port"`
	} `toml:"server"`

	Log struct {
		Level string `toml:"level"`
	} `toml:"log"`
}
