package binding

import (
	"context"

	"github.com/wailsapp/wails/v3/pkg/application"

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
func (b *UserBinding) GetProfile(ctx context.Context, id int) *Result {
	if id <= 0 {
		return Success(nil)
	}
	user, err := b.svc.GetUserProfile(ctx, id)
	if err != nil {
		return Failure(err)
	}
	return Success(user)
}

// Register 新增用户
func (b *UserBinding) Register(ctx context.Context, name, email string) *Result {
	user, err := b.svc.RegisterUser(ctx, name, email)
	if err != nil {
		return Failure(err)
	}
	return Success(user)
}

// Update 更新用户
func (b *UserBinding) Update(ctx context.Context, id int, name, email string) *Result {
	user, err := b.svc.UpdateUser(ctx, id, name, email)
	if err != nil {
		return Failure(err)
	}
	return Success(user)
}

// Delete 删除用户
func (b *UserBinding) Delete(ctx context.Context, id int) *Result {
	err := b.svc.DeleteUser(ctx, id)
	return Failure(err)
}

// List 分页查询用户列表
func (b *UserBinding) List(ctx context.Context, keyword string, page, pageSize int) *Result {
	res, err := b.svc.ListUsers(ctx, keyword, page, pageSize)
	if err != nil {
		return Failure(err)
	}
	return Success(res)
}

// GetModels 仅用于强制 Wails 在生成 Bindings 时包含这些模型定义
func (b *UserBinding) GetModels() (*domain.User, *domain.UserListResult) {
	return nil, nil
}

// Export 导出所有用户数据到本地文件 (原生对话框演示)
func (b *UserBinding) Export(ctx context.Context) *Result {
	// 1. 调用 Wails 原生保存文件对话框 (V3 Alpha API)
	dialog := application.Get().Dialog.SaveFile()
	dialog.SetOptions(&application.SaveFileDialogOptions{
		Title:    "导出用户数据",
		Filename: "users_export.json",
		Filters: []application.FileFilter{
			{
				DisplayName: "JSON Files (*.json)",
				Pattern:     "*.json",
			},
		},
	})

	path, err := dialog.PromptForSingleSelection()
	if err != nil {
		return Failure(err)
	}

	// 2. 检查用户是否取消了操作
	if path == "" {
		return Success(nil)
	}

	// 3. 将数据序列化并写入文件 (业务逻辑已下沉至 Service)
	return Failure(b.svc.ExportToJSON(ctx, path))
}
