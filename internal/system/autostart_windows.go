//go:build windows

package system

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

const runKey = `Software\Microsoft\Windows\CurrentVersion\Run`

func setAutostart(appName string, enabled bool) error {
	k, err := registry.OpenKey(registry.CURRENT_USER, runKey, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return fmt.Errorf("无法打开注册表: %v", err)
	}
	defer k.Close()

	if enabled {
		exePath, err := os.Executable()
		if err != nil {
			return fmt.Errorf("无法获取程序路径: %v", err)
		}
		absPath, _ := filepath.Abs(exePath)
		err = k.SetStringValue(appName, fmt.Sprintf(`"%s"`, absPath))
		if err != nil {
			return fmt.Errorf("写入注册表失败: %v", err)
		}
	} else {
		_ = k.DeleteValue(appName)
	}
	return nil
}

func isAutostartEnabled(appName string) (bool, error) {
	k, err := registry.OpenKey(registry.CURRENT_USER, runKey, registry.QUERY_VALUE)
	if err != nil {
		return false, nil
	}
	defer k.Close()

	_, _, err = k.GetStringValue(appName)
	if err != nil {
		return false, nil
	}
	return true, nil
}
