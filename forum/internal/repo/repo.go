package repo

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	conf "github.com/Skijetler/alphinium/forum/internal/config"
	"github.com/Skijetler/alphinium/pkg/ent"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewRepo, NewForumRepo)

// Repo .
type Repo struct {
	db *ent.Client
}

// NewRepo .
func NewRepo(conf *conf.Config, logger *logrus.Logger) (*Repo, func(), error) {
	drv, err := sql.Open(
		conf.Storage.Database.Driver,
		conf.Storage.Database.Source,
	)
	if err != nil {
		logger.Errorf("failed opening connection to database: %v", err)
		return nil, nil, err
	}

	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		logger.WithContext(ctx).Info(i...)
	})

	client := ent.NewClient(ent.Driver(sqlDrv))
	if err := client.Schema.Create(context.Background()); err != nil {
		logger.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	r := &Repo{
		db: client,
	}

	cleanup := func() {
		logger.Info("closing the data resources")
		if err := r.db.Close(); err != nil {
			logger.Errorf("failed closing the data resources: %v", err)
		}
	}
	return r, cleanup, nil
}
