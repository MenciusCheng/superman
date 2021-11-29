package dragons

import (
	"context"
	"fmt"
	"time"
)

func (d *Dragons) initMiddleware() error {
	mysqlClientInit := func() error {
		return d.InitSqlClient(d.config.Database)
	}

	middlewares := map[string]func() error{
		"mysqlClientInit": mysqlClientInit,
	}
	for name, fn := range middlewares {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		fnDone := make(chan error)
		go func() {
			fnDone <- fn()
		}()
	INNER:
		for {
			select {
			case <-ctx.Done():
				cancel()
				return fmt.Errorf("doing %s timeout, please check your config", name)
			case err := <-fnDone:
				if err != nil {
					cancel()
					return err
				}
				break INNER
			}
		}
		cancel()
	}

	return nil
}
