package repository

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"msg/cmd/app/server/common/db"
	"msg/cmd/app/server/common/sign"
)

var Common = new(common)

type common struct {
}

func (*common) WithTransaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return db.Get().Transaction(fc, opts...)
}

func (*common) Context(db *gorm.DB) context.Context {
	return context.WithValue(context.Background(), sign.DB, db)
}

func (*common) InTransaction(ctx context.Context) bool {
	_, ok := ctx.Value(sign.DB).(*gorm.DB)
	return ok
}

func (com *common) Transaction(ctx context.Context, f func(ctx context.Context) error) error {
	if !com.InTransaction(ctx) {
		return com.WithTransaction(func(tx *gorm.DB) error {
			c := com.Context(tx)
			return f(c)
		})
	}
	return f(ctx)
}
