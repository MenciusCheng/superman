package proxy

import (
	"context"
	"github.com/MenciusCheng/superman/util/dragons"
	"github.com/MenciusCheng/superman/util/golang/sql"
)

type SQL struct {
	name []string
}

func InitSQL(name ...string) *SQL {
	if len(name) == 0 {
		return nil
	}
	return &SQL{name}
}

func (s *SQL) Master(ctx context.Context, name ...string) *sql.Client {
	var gName string
	if len(name) == 0 {
		gName = s.name[0]
	} else {
		gName = name[0]
	}
	return dragons.SQLClient(ctx, gName).Master()
}

func (s *SQL) Slave(ctx context.Context, name ...string) *sql.Client {
	var gName string
	if len(name) == 0 {
		gName = s.name[0]
	} else {
		gName = name[0]
	}
	return dragons.SQLClient(ctx, gName).Slave()
}
