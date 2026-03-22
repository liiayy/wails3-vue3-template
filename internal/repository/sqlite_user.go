package repository

import (
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"myapp2/internal/domain"
)

// SqliteUserRepository 依靠 GORM 实现领域层的 User 仓储接口
type SqliteUserRepository struct {
	db *gorm.DB
}

// NewSqliteUserRepository 注入已经连接好的 gorm 引擎
func NewSqliteUserRepository(db *gorm.DB) *SqliteUserRepository {
	return &SqliteUserRepository{
		db: db,
	}
}

// FindByID 根据主键查询 User 对象
func (r *SqliteUserRepository) FindByID(id int) (*domain.User, error) {
	var user domain.User
	result := r.db.First(&user, id) // GORM 根据主键默认查找
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			zap.S().Warnf("数据库查询 User 失败：未找到 ID=%d 的用户", id)
			return nil, errors.New("user not found")
		}
		zap.S().Errorf("查询数据库遇见系统错: %v", result.Error)
		return nil, result.Error
	}
	return &user, nil
}

// Save 新增或更新数据
func (r *SqliteUserRepository) Save(user *domain.User) error {
	// 如果 ID 主键为 0 则创建，非 0 则执行更新（Gorm 的 Save 带有 Upsert 语义）
	result := r.db.Save(user)
	if result.Error != nil {
		zap.S().Errorf("保存数据写入 SQLite 失败: %v", result.Error)
		return result.Error
	}
	return nil
}
