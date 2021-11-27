package dragons

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"sync"
)

type Dragons struct {
	Name       string
	ConfigPath string
	initOnce   sync.Once
	config     dragonsConfig
	configFile []byte
}

func New() *Dragons {
	return &Dragons{}
}

func (d *Dragons) Init(options ...Option) {
	d.initOnce.Do(func() {
		for _, opt := range options {
			opt(d)
		}

		// 读取本地配置
		d.configFile = d.loadLocalConfig()

	})

}

func (d *Dragons) Scan(v interface{}) error {
	_, err := toml.Decode(string(d.configFile), v)
	return err
}

func (d *Dragons) loadLocalConfig() []byte {
	if len(d.ConfigPath) == 0 {
		panic(fmt.Sprintf("ConfigPath is empty"))
	}

	file, err := ioutil.ReadFile(d.ConfigPath)
	if err != nil {
		panic(fmt.Sprintf("LoadFile fail, err: %+v", err))
	}

	return file
}
