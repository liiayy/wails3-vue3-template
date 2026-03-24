package binding

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
func (b *SystemBinding) SetAutostart(enabled bool) error {
	return system.SetAutostart(b.appName, enabled)
}

// IsAutostartEnabled 检查当前是否已开启自启动
func (b *SystemBinding) IsAutostartEnabled() (bool, error) {
	return system.IsAutostartEnabled(b.appName)
}
