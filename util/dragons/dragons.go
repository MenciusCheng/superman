package dragons

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/MenciusCheng/superman/util/golang/sql"
	"io/ioutil"
	"log"
	"sync"
)

type Dragons struct {
	Name         string
	ConfigPath   string
	initOnce     sync.Once
	config       dragonsConfig
	configFile   []byte
	mysqlClients sync.Map
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

		if len(d.configFile) > 0 {
			_ = d.Scan(&d.config)

			// init middleware client
			if err := d.initMiddleware(); err != nil {
				log.Fatalf("init middleware fatal error:%v", err)
			}
		}

	})

}

func (d *Dragons) Scan(v interface{}) error {
	_, err := toml.Decode(string(d.configFile), v)
	return err
}

func (d *Dragons) SQLClient(name string) *sql.Group {
	if client, ok := d.mysqlClients.Load(name); ok {
		if v, ok1 := client.(*sql.Group); ok1 {
			return v
		}
	}
	return nil
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
