package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"myapp2/internal/domain"
)

// UserService 处理关于用户的纯粹核心逻辑
type UserService struct {
	repo      domain.UserRepository
	logger    domain.Logger
	validator *validator.Validate
}

func NewUserService(repo domain.UserRepository, logger domain.Logger) *UserService {
	v := validator.New()
	return &UserService{repo: repo, logger: logger, validator: v}
}

// validate 内部通用校验工具
func (s *UserService) validateStruct(data interface{}) error {
	err := s.validator.Struct(data)
	if err == nil {
		return nil
	}

	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		return domain.ErrInternal("数据校验底层异常", err)
	}

	// 将多条错误转为一条易读的信息
	var errMsgs []string
	for _, ve := range validationErrors {
		// 这里可以根据 ve.Tag() 实现更多的国际化映射
		msg := fmt.Sprintf("字段 [%s] 校验不通过 (%s)", ve.Field(), ve.Tag())
		errMsgs = append(errMsgs, msg)
	}

	return domain.ErrValidation(strings.Join(errMsgs, "; "))
}

// RegisterUser 新增用户
func (s *UserService) RegisterUser(ctx context.Context, req domain.UserRegisterRequest) (*domain.User, error) {
	if err := s.validateStruct(req); err != nil {
		return nil, err
	}

	user := &domain.User{Name: req.Name, Email: req.Email}
	err := s.repo.Save(ctx, user)
	if err != nil {
		return nil, domain.ErrInternal("注册入库失败", err)
	}
	return user, nil
}

// GetUserProfile 查询单个用户
func (s *UserService) GetUserProfile(ctx context.Context, id int) (*domain.User, error) {
	return s.repo.FindByID(ctx, id)
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(ctx context.Context, req domain.UserUpdateRequest) (*domain.User, error) {
	if err := s.validateStruct(req); err != nil {
		return nil, err
	}

	user, err := s.repo.FindByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}

	if err := s.repo.Save(ctx, user); err != nil {
		return nil, domain.ErrInternal("更新用户失败", err)
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
		return domain.ErrInternal("JSON 序列化失败", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return domain.ErrInternal("写入文件失败", err)
	}
	return nil
}
