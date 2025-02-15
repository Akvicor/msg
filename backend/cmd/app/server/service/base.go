package service

import (
	"context"
	"gorm.io/gorm"
	"msg/cmd/app/server/repository"
)

type base struct{}

// context 创建携带DB的Context
func (b *base) context(db *gorm.DB) context.Context {
	return repository.Common.Context(db)
}

// Transaction 发起事务
func (b *base) transaction(ctx context.Context, f func(ctx context.Context) error) error {
	return repository.Common.Transaction(ctx, f)
}

// 判断是否处在事务中
func (b *base) inTransaction(ctx context.Context) bool {
	return repository.Common.InTransaction(ctx)
}
