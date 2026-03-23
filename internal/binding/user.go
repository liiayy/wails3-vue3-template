package binding

import (
	"myapp2/internal/domain"
	"myapp2/internal/service"
)

// UserBinding 专门暴露给前端的 JS/TS 调用入口
type UserBinding struct {
	svc *service.UserService
}

func NewUserBinding(svc *service.UserService) *UserBinding {
	return &UserBinding{svc: svc}
}

// GetProfile 查询单个用户
func (b *UserBinding) GetProfile(id int) (*domain.User, error) {
	if id <= 0 {
		return nil, nil
	}
	return b.svc.GetUserProfile(id)
}

// Register 新增用户
func (b *UserBinding) Register(name, email string) (*domain.User, error) {
	return b.svc.RegisterUser(name, email)
}

// Update 更新用户
func (b *UserBinding) Update(id int, name, email string) (*domain.User, error) {
	return b.svc.UpdateUser(id, name, email)
}

// Delete 删除用户
func (b *UserBinding) Delete(id int) error {
	return b.svc.DeleteUser(id)
}

// List 分页查询用户列表
func (b *UserBinding) List(keyword string, page, pageSize int) (*domain.UserListResult, error) {
	return b.svc.ListUsers(keyword, page, pageSize)
}
