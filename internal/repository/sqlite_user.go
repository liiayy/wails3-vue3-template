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
	return &SqliteUserRepository{db: db}
}

// FindByID 根据主键查询 User 对象
func (r *SqliteUserRepository) FindByID(id int) (*domain.User, error) {
	var user domain.User
	result := r.db.First(&user, id)
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
	result := r.db.Save(user)
	if result.Error != nil {
		zap.S().Errorf("保存数据写入 SQLite 失败: %v", result.Error)
		return result.Error
	}
	return nil
}

// Delete 根据主键删除用户
func (r *SqliteUserRepository) Delete(id int) error {
	result := r.db.Delete(&domain.User{}, id)
	if result.Error != nil {
		zap.S().Errorf("删除用户失败 ID=%d: %v", id, result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

// List 分页 + 关键词模糊搜索
func (r *SqliteUserRepository) List(keyword string, page, pageSize int) (*domain.UserListResult, error) {
	var users []*domain.User
	var total int64

	query := r.db.Model(&domain.User{})

	// 模糊搜索（名称或邮箱）
	if keyword != "" {
		like := "%" + keyword + "%"
		query = query.Where("name LIKE ? OR email LIKE ?", like, like)
	}

	// 先查总数
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Order("id DESC").Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, err
	}

	return &domain.UserListResult{
		Items: users,
		Total: total,
	}, nil
}
