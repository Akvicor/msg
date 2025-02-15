package repository

import (
	"context"
	"gorm.io/gorm"
	"msg/cmd/app/server/common/db"
	"msg/cmd/app/server/common/preloader"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/common/sign"
	"msg/cmd/app/server/model"
	"reflect"
)

type base[T model.Baser] struct{}

// db 获取数据库
func (b *base[T]) db(c context.Context) *gorm.DB {
	d, ok := c.Value(sign.DB).(*gorm.DB)
	if !ok {
		return db.Get()
	}
	return d
}

// dbt 获取数据库并指定Table
func (b *base[T]) dbt(c context.Context) *gorm.DB {
	var a T
	d, ok := c.Value(sign.DB).(*gorm.DB)
	if !ok {
		return db.Get().Table(a.TableName())
	}
	return d.Table(a.TableName())
}

// dbm 获取数据库并指定Model
func (b *base[T]) dbm(c context.Context) *gorm.DB {
	d, ok := c.Value(sign.DB).(*gorm.DB)
	if !ok {
		return db.Get().Model(new(T))
	}
	return d.Model(new(T))
}

// paging 分页, page为nil时不分页
func (b *base[T]) paging(page *resp.PageModel, tx *gorm.DB) (results []T, err error) {
	// 判断是否分页
	if page != nil {
		// 如果未指定右界限, 则获取最新ID, 确定右界限
		if page.End == 0 {
			err = db.Get().Model(new(T)).Select("MAX(id)").Scan(&page.End).Error
			if err != nil {
				page.End = 0
			}
		}
		// 限定数据区间, 统计数据量
		tx.Where("id > ?", page.Start).Where("id <= ?", page.End).Count(&page.Total)
		if page.Limit == -1 {
			// 仅统计数量
			return nil, tx.Error
		} else if page.Limit == 0 {
			// 获取全部数据
		} else {
			// 分页获取
			tx = tx.Offset((page.Index - 1) * page.Limit).Limit(page.Limit)
		}
		// 倒序
		if page.Desc {
			tx = tx.Order("id DESC")
		} else {
			tx = tx.Order("id ASC")
		}
	}

	// 查询数据
	results = make([]T, 0)
	resultPart := make([]T, 0)
	result := tx.FindInBatches(&resultPart, 500, func(tx *gorm.DB, batch int) error {
		results = append(results, resultPart...)
		return nil
	})
	if result.Error != nil {
		return make([]T, 0), result.Error
	}
	if result.RowsAffected == 0 {
		return make([]T, 0), gorm.ErrRecordNotFound
	}

	return results, nil
}

// 活跃记录
func (b *base[T]) alive(c context.Context, alive bool) *gorm.DB {
	tx := b.dbm(c)
	if alive {
		var t T
		tx = t.Alive(tx)
	}
	return tx
}

// preload活跃记录
func (b *base[T]) preload(c context.Context, alive bool, preloader preloader.Preloader) *gorm.DB {
	tx := b.dbm(c)
	if alive {
		var t T
		tx = t.Alive(tx)
	}
	if preloader != nil && !reflect.ValueOf(preloader).IsNil() {
		tx = preloader.Preload(tx)
	}
	return tx
}

func (b *base[T]) WrapResultErr(result *gorm.DB) error {
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
