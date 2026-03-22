package binding

import (
	"myapp2/internal/domain"
	"myapp2/internal/service"
)

// UserBinding 专门暴漏给前端的 JS/TS 调用的入口
// 它充当 Controller，主要接收解析参数及进行防错拦截
type UserBinding struct {
	svc *service.UserService
}

func NewUserBinding(svc *service.UserService) *UserBinding {
	return &UserBinding{
		svc: svc,
	}
}

// GetProfile 供前端通过 Wails JS SDK 调用以获取用户信息
func (b *UserBinding) GetProfile(id int) (*domain.User, error) {
	// 如果前端传入负数等边界条件，可在此 Controller 层先拦截，不漏给业务纯层
	if id <= 0 {
		return nil, nil // 或 return errors.New("参数不合法")
	}

	return b.svc.GetUserProfile(id)
}

// Register 处理前端传来的注册请求
func (b *UserBinding) Register(name, email string) (*domain.User, error) {
	return b.svc.RegisterUser(name, email)
}
