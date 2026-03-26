package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"myapp2/internal/domain"
)

// UserService 处理关于用户的纯粹核心逻辑
type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// RegisterUser 新增用户
func (s *UserService) RegisterUser(ctx context.Context, name, email string) (*domain.User, error) {
	if name == "" {
		return nil, errors.New("用户名不能为空")
	}
	if email == "" {
		return nil, errors.New("邮箱不能为空")
	}
	user := &domain.User{Name: name, Email: email}
	err := s.repo.Save(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("注册入库失败: %w", err)
	}
	return user, nil
}

// GetUserProfile 查询单个用户
func (s *UserService) GetUserProfile(ctx context.Context, id int) (*domain.User, error) {
	return s.repo.FindByID(ctx, id)
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(ctx context.Context, id int, name, email string) (*domain.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}
	if err := s.repo.Save(ctx, user); err != nil {
		return nil, fmt.Errorf("更新用户失败: %w", err)
	}
	return user, nil
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

// ListUsers 分页列表
func (s *UserService) ListUsers(ctx context.Context, keyword string, page, pageSize int) (*domain.UserListResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	return s.repo.List(ctx, keyword, page, pageSize)
}

// GetAllUsers 返回所有用户（全量导出场景）
func (s *UserService) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	return s.repo.GetAll(ctx)
}

// ExportToJSON 将用户数据导出至指定路径
func (s *UserService) ExportToJSON(ctx context.Context, path string) error {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
