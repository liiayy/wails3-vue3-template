package service

import (
	"errors"
	"fmt"
	"myapp2/internal/domain"
)

// UserService 处理关于用户的纯粹核心逻辑
// 它隔离了具体的数据库技术，也不清楚究竟是 Wails 还是 HTTP 请求调用它
type UserService struct {
	repo domain.UserRepository
}

// 依赖注入 Repository，以后即便要换 MySQL，UserService 也不需修改半行代码
func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// RegisterUser 执行业务逻辑：比如校验邮箱格式、落库持久化
func (s *UserService) RegisterUser(name, email string) (*domain.User, error) {
	if name == "" {
		return nil, errors.New("用户名不能为空")
	}
	if email == "" {
		return nil, errors.New("邮箱不能为空")
	}

	user := &domain.User{
		Name:  name,
		Email: email,
	}

	// 调用仓储层
	err := s.repo.Save(user)
	if err != nil {
		return nil, fmt.Errorf("注册入库失败: %w", err)
	}

	return user, nil
}

// GetUserProfile 专门用于展示特定的用户信息组装
func (s *UserService) GetUserProfile(id int) (*domain.User, error) {
	return s.repo.FindByID(id)
}
