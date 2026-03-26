package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"myapp2/internal/domain"
)

// SqliteSettingRepository 使用 GORM 存储设置数据
type SqliteSettingRepository struct {
	db     *gorm.DB
	logger domain.Logger
}

func NewSqliteSettingRepository(db *gorm.DB, logger domain.Logger) *SqliteSettingRepository {
	return &SqliteSettingRepository{db: db, logger: logger}
}

func (r *SqliteSettingRepository) Get(ctx context.Context, key string) (string, error) {
	var s domain.Setting
	if err := r.db.WithContext(ctx).First(&s, "key = ?", key).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil // 未设置项不算由于驱动错，返回空
		}
		return "", domain.ErrInternal("获取设置项失败", err)
	}
	return s.Value, nil
}

func (r *SqliteSettingRepository) Set(ctx context.Context, key, value string) error {
	s := domain.Setting{Key: key, Value: value}
	// Upsert: 若存在则更新，若不存在则新增
	if err := r.db.WithContext(ctx).Save(&s).Error; err != nil {
		return domain.ErrInternal("保存设置项失败", err)
	}
	return nil
}

func (r *SqliteSettingRepository) GetAll(ctx context.Context) (map[string]string, error) {
	var settings []domain.Setting
	if err := r.db.WithContext(ctx).Find(&settings).Error; err != nil {
		return nil, domain.ErrInternal("获取全量设置项失败", err)
	}

	result := make(map[string]string)
	for _, s := range settings {
		result[s.Key] = s.Value
	}
	return result, nil
}
