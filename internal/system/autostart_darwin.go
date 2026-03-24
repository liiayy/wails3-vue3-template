//go:build darwin

package system

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

const plistTemplateData = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>{{.AppName}}</string>
    <key>ProgramArguments</key>
    <array>
        <string>{{.ExePath}}</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
</dict>
</plist>`

func setAutostart(appName string, enabled bool) error {
	home, _ := os.UserHomeDir()
	plistPath := filepath.Join(home, "Library/LaunchAgents", appName+".plist")

	if enabled {
		exePath, _ := os.Executable()
		absPath, _ := filepath.Abs(exePath)

		f, err := os.Create(plistPath)
		if err != nil {
			return fmt.Errorf("无法创建 plist 文件: %v", err)
		}
		defer f.Close()

		tmpl, _ := template.New("plist").Parse(plistTemplateData)
		err = tmpl.Execute(f, map[string]string{
			"AppName": appName,
			"ExePath": absPath,
		})
		if err != nil {
			return fmt.Errorf("写入 plist 失败: %v", err)
		}
	} else {
		_ = os.Remove(plistPath)
	}
	return nil
}

func isAutostartEnabled(appName string) (bool, error) {
	home, _ := os.UserHomeDir()
	plistPath := filepath.Join(home, "Library/LaunchAgents", appName+".plist")
	_, err := os.Stat(plistPath)
	return err == nil, nil
}
