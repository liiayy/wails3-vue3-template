package repository

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"myapp2/internal/domain"
)

// SqliteUserRepository 依靠 GORM 实现领域层的 User 仓储接口
type SqliteUserRepository struct {
	db     *gorm.DB
	logger domain.Logger
}

// NewSqliteUserRepository 注入已经连接好的 gorm 引擎与 Logger
func NewSqliteUserRepository(db *gorm.DB, logger domain.Logger) *SqliteUserRepository {
	return &SqliteUserRepository{db: db, logger: logger}
}

// FindByID 根据主键查询 User 对象
func (r *SqliteUserRepository) FindByID(ctx context.Context, id int) (*domain.User, error) {
	var user domain.User
	result := r.db.WithContext(ctx).First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			r.logger.Warnf("数据库查询 User 失败：未找到 ID=%d 的用户", id)
			return nil, domain.ErrNotFound(fmt.Sprintf("未找到 ID=%d 的用户", id), result.Error)
		}
		r.logger.Errorf("查询数据库遇见系统错: %v", result.Error)
		return nil, domain.ErrInternal("查询数据库遇见系统错", result.Error)
	}
	return &user, nil
}

// Save 新增或更新数据
func (r *SqliteUserRepository) Save(ctx context.Context, user *domain.User) error {
	result := r.db.WithContext(ctx).Save(user)
	if result.Error != nil {
		r.logger.Errorf("保存数据写入 SQLite 失败: %v", result.Error)
		return domain.ErrInternal("保存数据写入 SQLite 失败", result.Error)
	}
	return nil
}

// Delete 根据主键删除用户
func (r *SqliteUserRepository) Delete(ctx context.Context, id int) error {
	result := r.db.WithContext(ctx).Delete(&domain.User{}, id)
	if result.Error != nil {
		r.logger.Errorf("删除用户失败 ID=%d: %v", id, result.Error)
		return domain.ErrInternal("删除用户失败", result.Error)
	}
	if result.RowsAffected == 0 {
		return domain.ErrNotFound(fmt.Sprintf("未找到 ID=%d 的用户", id), nil)
	}
	return nil
}

// List 分页 + 关键词模糊搜索
func (r *SqliteUserRepository) List(ctx context.Context, keyword string, page, pageSize int) (*domain.UserListResult, error) {
	var users []*domain.User
	var total int64

	query := r.db.WithContext(ctx).Model(&domain.User{})

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

// GetAll 获取所有用户数据 (用于导出等全量场景)
func (r *SqliteUserRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User
	if err := r.db.WithContext(ctx).Order("id ASC").Find(&users).Error; err != nil {
		r.logger.Errorf("查询全量 User 数据错误: %v", err)
		return nil, err
	}
	return users, nil
}
