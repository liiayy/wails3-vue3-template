package service

import (
	"errors"
	"fmt"

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
func (s *UserService) RegisterUser(name, email string) (*domain.User, error) {
	if name == "" {
		return nil, errors.New("用户名不能为空")
	}
	if email == "" {
		return nil, errors.New("邮箱不能为空")
	}
	user := &domain.User{Name: name, Email: email}
	err := s.repo.Save(user)
	if err != nil {
		return nil, fmt.Errorf("注册入库失败: %w", err)
	}
	return user, nil
}

// GetUserProfile 查询单个用户
func (s *UserService) GetUserProfile(id int) (*domain.User, error) {
	return s.repo.FindByID(id)
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(id int, name, email string) (*domain.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}
	if err := s.repo.Save(user); err != nil {
		return nil, fmt.Errorf("更新用户失败: %w", err)
	}
	return user, nil
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}

// ListUsers 分页列表
func (s *UserService) ListUsers(keyword string, page, pageSize int) (*domain.UserListResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	return s.repo.List(keyword, page, pageSize)
}

// GetAllUsers 返回所有用户（全量导出场景）
func (s *UserService) GetAllUsers() ([]*domain.User, error) {
	return s.repo.GetAll()
}
