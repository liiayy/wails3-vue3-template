//go:build linux

package system

import (
	"fmt"
	"os"
	"path/filepath"
)

func setAutostart(appName string, enabled bool) error {
	home, _ := os.UserHomeDir()
	desktopDir := filepath.Join(home, ".config/autostart")
	desktopPath := filepath.Join(desktopDir, appName+".desktop")

	if enabled {
		os.MkdirAll(desktopDir, 0755)
		exePath, _ := os.Executable()
		absPath, _ := filepath.Abs(exePath)

		content := fmt.Sprintf(`[Desktop Entry]
Type=Application
Version=1.0
Name=%s
Comment=Start MyApp2 on login
Exec="%s"
StartupNotify=false
Terminal=false
`, appName, absPath)

		err := os.WriteFile(desktopPath, []byte(content), 0644)
		if err != nil {
			return fmt.Errorf("无法创建 desktop 文件: %v", err)
		}
	} else {
		_ = os.Remove(desktopPath)
	}
	return nil
}

func isAutostartEnabled(appName string) (bool, error) {
	home, _ := os.UserHomeDir()
	desktopPath := filepath.Join(home, ".config/autostart", appName+".desktop")
	_, err := os.Stat(desktopPath)
	return err == nil, nil
}
