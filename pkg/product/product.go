package product

import (
	"context"
	"database/sql"
)

type IResource interface {
	UpdateAllProduct(ctx context.Context) error
}

type module struct {
	persistent persistent
}

func New(dbSource *sql.DB, dbDest *sql.DB) IResource {
	return module{
		persistent: newPersistent(dbSource, dbDest),
	}
}

func (m module) UpdateAllProduct(ctx context.Context) error {
	return m.persistent.updateAllProduct(ctx)
}
