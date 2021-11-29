package dragons

import (
	"github.com/MenciusCheng/superman/util/golang/sql"
	"strings"
)

func (d *Dragons) Shutdown() error {
	return nil
}

func (d *Dragons) InitSqlClient(sqlList []sql.SQLGroupConfig) error {
	for _, c := range sqlList {
		if _, ok := d.mysqlClients.Load(c.Name); ok {
			continue
		}
		if len(c.LogLevel) == 0 {
			c.LogLevel = strings.ToLower(d.config.Log.Level)
		}
		g, err := sql.NewGroup(c)
		if err != nil {
			return err
		}
		_ = sql.SQLGroupManager.Add(c.Name, g)
		d.mysqlClients.LoadOrStore(c.Name, g)
	}
	return nil
}
