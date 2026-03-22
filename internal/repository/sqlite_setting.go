package repository

import (
	"errors"
	"gorm.io/gorm"
	"myapp2/internal/domain"
)

// SqliteSettingRepository 使用 GORM 存储设置数据
type SqliteSettingRepository struct {
	db *gorm.DB
}

func NewSqliteSettingRepository(db *gorm.DB) *SqliteSettingRepository {
	return &SqliteSettingRepository{db: db}
}

func (r *SqliteSettingRepository) Get(key string) (string, error) {
	var s domain.Setting
	if err := r.db.First(&s, "key = ?", key).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil // 未设置项不算由于驱动错，返回空
		}
		return "", err
	}
	return s.Value, nil
}

func (r *SqliteSettingRepository) Set(key, value string) error {
	s := domain.Setting{Key: key, Value: value}
	// Upsert: 若存在则更新，若不存在则新增
	return r.db.Save(&s).Error
}

func (r *SqliteSettingRepository) GetAll() (map[string]string, error) {
	var settings []domain.Setting
	if err := r.db.Find(&settings).Error; err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, s := range settings {
		result[s.Key] = s.Value
	}
	return result, nil
}
