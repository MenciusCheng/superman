package sql

import (
	"log"
	"sync/atomic"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Client继承了*gorm.DB的所有方法, 详细的使用方法请参考:
// http://gorm.io/docs/connecting_to_the_database.html
type Client struct {
	*gorm.DB
}

type Group struct {
	name string

	master  *Client
	replica []*Client
	next    uint64
	total   uint64
}

func openDB(name, address string, isMaster int, statLevel, format, logLevel string) (*Client, error) {
	db, err := gorm.Open(mysql.Open(address), &gorm.Config{})

	db = db.Debug()

	return &Client{DB: db}, err
}

// NewGroup初始化一个Group， 一个Group包含一个master实例和零个或多个slave实例
func NewGroup(d SQLGroupConfig) (*Group, error) {
	log.Printf("init sql group name [%s], master [%s], slave [%v]\n", d.Name, d.Master, d.Slaves)
	g := Group{name: d.Name}
	var err error
	g.master, err = openDB(d.Name, d.Master, 1, d.StatLevel, d.LogFormat, d.LogLevel)
	if err != nil {
		return nil, err
	}
	g.replica = make([]*Client, 0, len(d.Slaves))
	g.total = 0
	for _, slave := range d.Slaves {
		c, err := openDB(d.Name, slave, 0, d.StatLevel, d.LogFormat, d.LogLevel)
		if err != nil {
			return nil, err
		}
		g.replica = append(g.replica, c)
		g.total++

	}
	return &g, nil
}

// Master返回master实例
func (g *Group) Master() *Client {
	return g.master
}

// Slave返回一个slave实例，使用轮转算法
func (g *Group) Slave() *Client {
	if g.total == 0 {
		return g.master
	}
	next := atomic.AddUint64(&g.next, 1)
	return g.replica[next%g.total]
}

// Instance函数如果isMaster是true， 返回master实例，否则返回slave实例
func (g *Group) Instance(isMaster bool) *Client {
	if isMaster {
		return g.Master()
	}
	return g.Slave()
}
