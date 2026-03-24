package binding

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

type SystemBinding struct {
	appName string
}

func NewSystemBinding(appName string) *SystemBinding {
	return &SystemBinding{
		appName: appName,
	}
}

const runKey = `Software\Microsoft\Windows\CurrentVersion\Run`

// SetAutostart 设置开机自启动
func (b *SystemBinding) SetAutostart(enabled bool) error {
	k, err := registry.OpenKey(registry.CURRENT_USER, runKey, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return fmt.Errorf("无法打开注册表: %v", err)
	}
	defer k.Close()

	if enabled {
		// 获取当前程序可执行文件路径
		exePath, err := os.Executable()
		if err != nil {
			return fmt.Errorf("无法获取程序路径: %v", err)
		}
		// 转换成绝对路径
		absPath, _ := filepath.Abs(exePath)
		// 写入注册表 (格式: "C:\path\to\app.exe")
		err = k.SetStringValue(b.appName, fmt.Sprintf(`"%s"`, absPath))
		if err != nil {
			return fmt.Errorf("写入注册表失败: %v", err)
		}
	} else {
		// 删除注册表项
		_ = k.DeleteValue(b.appName)
	}

	return nil
}

// IsAutostartEnabled 检查当前是否已开启自启动
func (b *SystemBinding) IsAutostartEnabled() (bool, error) {
	k, err := registry.OpenKey(registry.CURRENT_USER, runKey, registry.QUERY_VALUE)
	if err != nil {
		return false, nil
	}
	defer k.Close()

	_, _, err = k.GetStringValue(b.appName)
	if err != nil {
		return false, nil
	}
	return true, nil
}
