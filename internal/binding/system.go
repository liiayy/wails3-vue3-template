package binding

import "context"

import (
	"myapp2/internal/system"
)

type SystemBinding struct {
	appName string
}

func NewSystemBinding(appName string) *SystemBinding {
	return &SystemBinding{
		appName: appName,
	}
}

// SetAutostart 设置开机自启动
func (b *SystemBinding) SetAutostart(ctx context.Context, enabled bool) *Result {
	err := system.SetAutostart(b.appName, enabled)
	return Failure(err)
}

// IsAutostartEnabled 检查当前是否已开启自启动
func (b *SystemBinding) IsAutostartEnabled(ctx context.Context) *Result {
	enabled, err := system.IsAutostartEnabled(b.appName)
	if err != nil {
		return Failure(err)
	}
	return Success(enabled)
}
