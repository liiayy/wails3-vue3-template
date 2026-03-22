package database

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"myapp2/internal/domain"
)

// InitDB 初始化本地 SQLite 数据库连接并自动迁移结构体
// 返回 GORM DB 对象供各 Repository 层直接使用
func InitDB(appName string) (*gorm.DB, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = "."
	}

	// 保证数据库文件夹存在 C:\Users\Admin\AppData\Roaming\MyApp2\data
	dataDir := filepath.Join(configDir, appName, "data")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		zap.S().Fatalf("创建数据库目录失败: %v", err)
		return nil, err
	}

	dbPath := filepath.Join(dataDir, "app_data.db")
	zap.S().Infof("正在连接本地 SQLite 数据库: %s", dbPath)

	// 配置 GORM：将 Gorm 自身的日志级别设为 Warn，避免污染你的业务日志
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		zap.S().Fatalf("连接数据库引擎失败: %v", err)
		return nil, err
	}

	// 执行自动迁移（Auto Migrate）
	// Gorm 会自动对比现在的 domain.User 结构体，如果没有表就建表，缺字段就加字段
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		zap.S().Fatalf("数据库结构迁移失败: %v", err)
		return nil, err
	}

	zap.S().Info("数据库表结构迁移与加载成功！")
	return db, nil
}
