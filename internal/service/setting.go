package service

import (
	"context"
	"go.uber.org/zap"
	"myapp2/internal/domain"
)

// SettingService 执行关于应用设置的核心逻辑处理 (比如校验、转换、批量保存)
type SettingService struct {
	repo domain.SettingRepository
}

func NewSettingService(repo domain.SettingRepository) *SettingService {
	return &SettingService{repo: repo}
}

// GetSettings 返回所有的应用设置供前端初始化
func (s *SettingService) GetSettings(ctx context.Context) (map[string]string, error) {
	return s.repo.GetAll(ctx)
}

// SaveSetting 异步或即时保存单个设置项
func (s *SettingService) SaveSetting(ctx context.Context, key, value string) error {
	zap.S().Debugf("正在异步持久化设置项: %s = %s", key, value)
	return s.repo.Set(ctx, key, value)
}
