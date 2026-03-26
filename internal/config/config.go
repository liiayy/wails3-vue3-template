package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// AppConfig 代表应用的全部 YAML 配置项（结构化映射）
// 注意：is_dev 和 version 是编译期注入的，不在此处管理
type AppConfig struct {
	App      AppInfo      `mapstructure:"app"`
	Database DatabaseConf `mapstructure:"database"`
	Log      LogConf      `mapstructure:"log"`
	Window   WindowConf   `mapstructure:"window"`
}

type AppInfo struct {
	Name string `mapstructure:"name"`
}

type DatabaseConf struct {
	Driver string `mapstructure:"driver"` // sqlite, mysql（预留）
	DBName string `mapstructure:"db_name"`
}

type LogConf struct {
	Level      string `mapstructure:"level"`        // debug, info, warn, error
	MaxSizeMB  int    `mapstructure:"max_size_mb"`  // 单文件最大体积
	MaxBackups int    `mapstructure:"max_backups"`  // 最多保留备份数
	MaxAgeDays int    `mapstructure:"max_age_days"` // 备份保留天数
	Compress   bool   `mapstructure:"compress"`     // 是否 gzip 压缩
}

type WindowConf struct {
	Width       int    `mapstructure:"width"`
	Height      int    `mapstructure:"height"`
	X           int    `mapstructure:"x"`
	Y           int    `mapstructure:"y"`
	IsMaximized bool   `mapstructure:"is_maximized"`
	Title       string `mapstructure:"title"`
}

// defaultYAML 是首次启动时自动写出的默认配置内容
// 【重要】is_dev 和 version 不在此处，它们通过编译参数 -ldflags 注入，用户无法篡改
const defaultYAML = `# ============================
# MyApp2 应用配置文件
# 修改后重启应用即可生效
# ============================

app:
  name: MyApp2

database:
  driver: sqlite
  db_name: app_data.db

log:
  level: debug
  max_size_mb: 50
  max_backups: 10
  max_age_days: 30
  compress: true

window:
  width: 1280
  height: 800
  x: -1
  is_maximized: false
  title: "Wails 3 Mega-Structure Dashboard"
`

// Cfg 全局配置单例（初始化后可在任何地方通过 config.Cfg 访问）
var Cfg AppConfig

// InitConfig 初始化配置中心
// 1. 定位配置文件存放目录（AppData/appName/）
// 2. 若不存在则自动生成带注释的默认 YAML
// 3. 反序列化进 Cfg 结构体
func InitConfig(appName string) error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = "."
	}

	appConfigDir := filepath.Join(configDir, appName)
	if err := os.MkdirAll(appConfigDir, 0755); err != nil {
		return fmt.Errorf("创建配置目录失败: %w", err)
	}

	configFilePath := filepath.Join(appConfigDir, "config.yaml")

	// 首次启动：自动生成默认配置文件
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		zap.S().Infof("首次启动，生成默认配置文件: %s", configFilePath)
		if err := os.WriteFile(configFilePath, []byte(defaultYAML), 0644); err != nil {
			return fmt.Errorf("写入默认配置失败: %w", err)
		}
	}

	// 配置 Viper
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType("yaml")

	// 设置环境变量前缀，支持 MYAPP2_APP_NAME 等覆盖方式
	viper.SetEnvPrefix("MYAPP2")
	viper.AutomaticEnv()

	// 设置默认值（防止升级用户配置文件缺失字段导致 0 坐标到左上角）
	viper.SetDefault("window.x", -1)
	viper.SetDefault("window.y", -1)

	// 读取
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 反序列化
	if err := viper.Unmarshal(&Cfg); err != nil {
		return fmt.Errorf("解析配置结构失败: %w", err)
	}

	zap.S().Infof("配置中心加载成功: %s", viper.ConfigFileUsed())
	zap.S().Debugf("当前配置: %+v", Cfg)

	return nil
}

// SaveConfig 将当前内存中的 Cfg 状态持久化回磁盘文件 (YAML)
func SaveConfig() error {
	// 将结构体同步回 Viper 内存
	viper.Set("window.width", Cfg.Window.Width)
	viper.Set("window.height", Cfg.Window.Height)
	viper.Set("window.x", Cfg.Window.X)
	viper.Set("window.y", Cfg.Window.Y)
	viper.Set("window.is_maximized", Cfg.Window.IsMaximized)

	if err := viper.WriteConfig(); err != nil {
		zap.S().Errorf("写入配置文件失败: %v", err)
		return err
	}
	zap.S().Debugf("配置文件保存成功")
	return nil
}

// UpdateWindowSize 快捷更新窗口尺寸并保存
func UpdateWindowSize(width, height int) {
	if Cfg.Window.Width == width && Cfg.Window.Height == height {
		return
	}
	Cfg.Window.Width = width
	Cfg.Window.Height = height
	_ = SaveConfig()
}

// UpdateWindowPosition 快捷更新窗口位置并保存
func UpdateWindowPosition(x, y int) {
	if Cfg.Window.X == x && Cfg.Window.Y == y {
		return
	}
	Cfg.Window.X = x
	Cfg.Window.Y = y
	_ = SaveConfig()
}

// UpdateWindowMaximizedState 更新并保存窗口最大化状态
func UpdateWindowMaximizedState(isMaximized bool) {
	if Cfg.Window.IsMaximized == isMaximized {
		return
	}
	Cfg.Window.IsMaximized = isMaximized
	_ = SaveConfig()
}

// GetConfigDir 返回配置文件所在目录（供其他模块定位同级文件）
func GetConfigDir(appName string) string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "."
	}
	return filepath.Join(configDir, appName)
}
