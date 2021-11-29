package dragons

import (
	"context"
	"github.com/MenciusCheng/superman/util/golang/sql"
)

var Default = New()

func Init(options ...Option) {
	Default.Init(options...)
}

func Shutdown() error {
	return Default.Shutdown()
}

func Scan(v interface{}) error {
	return Default.Scan(v)
}

func SQLClient(ctx context.Context, name string) *sql.Group {
	return Default.SQLClient(name)
}
