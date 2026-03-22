package logger

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// InitLogger 初始化全局的结构化日志组件 (Zap + Lumberjack)
// isDev 决定是往控制台输出带有颜色的可读日志，还是纯粹的 JSON 格式
// appName 用于在本地构建存放日志的目录，比如 "%AppData%/appName/logs"
func InitLogger(isDev bool, appName string) error {
	// 1. 获取本地日志保存根路径 (Windows: AppData, Mac: Library/Application Support)
	// 大型应用绝对不可以把日志和可执行文件放在一起，通常会由于读写权限不足崩溃
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = "." // Fallback 到当前目录
	}

	// 如：C:\Users\Admin\AppData\Roaming\MyApp2\logs\app.log
	logFileDir := filepath.Join(configDir, appName, "logs")
	if err := os.MkdirAll(logFileDir, 0755); err != nil {
		return err
	}
	logFilePath := filepath.Join(logFileDir, "app.log")

	// 2. 配置 Lumberjack 轮转日志
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    50,   // 每个日志文件最大 50 MB
		MaxBackups: 10,   // 最多保留 10 个备份
		MaxAge:     30,   // 文件最多保留 30 天
		Compress:   true, // 是否压缩旧日志(gzip)
	}

	// 3. 配置输出介质和格式
	fileWriter := zapcore.AddSync(lumberjackLogger)
	consoleWriter := zapcore.AddSync(os.Stdout)

	var encoderConfig zapcore.EncoderConfig
	var core zapcore.Core

	if isDev {
		// 开发模式: 控制台彩色输出 + 文件记录
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

		core = zapcore.NewTee(
			zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), consoleWriter, zap.DebugLevel),
			zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), fileWriter, zap.DebugLevel),
		)
	} else {
		// 生产模式: 仅文件记录 JSON，限制日志级别为 Info 及以上以提高性能
		encoderConfig = zap.NewProductionEncoderConfig()
		core = zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), fileWriter, zap.InfoLevel)
	}

	// 4. 生成 Logger 实例并替换全局
	// 添加 caller (触发日志的文件和行号) 与 stacktrace (若为 Error 强行打出堆栈)
	log := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	// 替换由于调用 zap.L() 和 zap.S() 产生的全局单例实例
	zap.ReplaceGlobals(log)

	zap.S().Infof("日志系统初始化完毕, 存储路径: %s", logFilePath)
	return nil
}
