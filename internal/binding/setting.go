package binding

import (
	"context"
	"myapp2/internal/service"
)

// SettingBinding 将设置中心暴露给前端 JS，支持异步写入
type SettingBinding struct {
	svc *service.SettingService
}

func NewSettingBinding(svc *service.SettingService) *SettingBinding {
	return &SettingBinding{svc: svc}
}

// GetAll 获取当前 SQLite 中存储的所有设置对
func (b *SettingBinding) GetAll(ctx context.Context) *Result {
	maps, err := b.svc.GetSettings(ctx)
	if err != nil {
		return Failure(err)
	}
	return Success(maps)
}

// Save 将单个状态推送到 Go 端持久化
func (b *SettingBinding) Save(ctx context.Context, key, value string) *Result {
	err := b.svc.SaveSetting(ctx, key, value)
	return Failure(err)
}
