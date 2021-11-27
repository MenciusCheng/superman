package manager

import (
	"context"
	"github.com/MenciusCheng/superman/conf"
)

// Manager represents middleware component
// such as, kafka, http client or rpc client, etc.
type Manager struct {
	c *conf.Config
}

func New(conf *conf.Config) *Manager {
	return &Manager{
		c: conf,
	}
}

func (m *Manager) Ping(ctx context.Context) error {
	return nil
}

func (m *Manager) Close() error {
	return nil
}
