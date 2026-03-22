package main

import (
	"context"
	"embed"
	_ "embed"
	"fmt"
	"log"
	"runtime/debug"

	"github.com/wailsapp/wails/v3/pkg/application"
	"go.uber.org/zap"

	"myapp2/internal/app"
	"myapp2/internal/binding"
	"myapp2/internal/database"
	"myapp2/internal/logger"
	"myapp2/internal/manager"
	"myapp2/internal/repository"
	"myapp2/internal/service"
)

//go:embed all:frontend/dist
var assets embed.FS

// IsDev dynamically dictates the environment (true for dev, false for prod)
const IsDev = true

// main 程序的唯一入口点，它的职责极度专注：仅仅负责对象的实例化、依赖组装和框架启动。
func main() {
	// 【0. 全局 Panic 兜底恢复与原生错误弹窗弹出】
	defer func() {
		if r := recover(); r != nil {
			errStr := fmt.Sprintf("进程遇到致命错误崩溃: %v\n\n%s", r, string(debug.Stack()))
			// 1. 落盘记录供排查
			zap.S().Errorf("【Fatal Panic】:\n%s", errStr)
			// 2. 调用 Wails 3 原生 Dialog 弹窗通知用户，而不是静默闪退
			if wailsApp := application.Get(); wailsApp != nil {
				wailsApp.Dialog.Error().
					SetTitle("App Critical Crash").
					SetMessage("糟糕，程序崩溃了！\n请将您的日志(AppData目录下)发送给我们的支持邮箱。\n详情: " + fmt.Sprintf("%v", r)).
					Show()
			}
		}
	}()

	// 【1. 动态环境隔绝与全局日志系统】
	// isDev 决定屏蔽或开启某些工具 (结合 Wails 构建参数可自动化)
	if err := logger.InitLogger(IsDev, "MyApp2"); err != nil {
		log.Fatalf("无法初始化日志系统: %v", err)
	}
	defer zap.L().Sync() // 确保程序退出前刷新磁盘 IO

	// 【1. 依赖注入与装配期】
	// -- 1.0 初始化 SQLite 数据库及 GORM 对象 --
	db, err := database.InitDB("MyApp2")
	if err != nil {
		zap.S().Fatalf("核心数据库引擎启动失败，终止此应用: %v", err)
	}

	// -- 1.1 初始化底层数据仓储 --
	// 只需要把旧的 NewInMemoryUserRepository() 替换掉，上层的纯代码 0 修改！
	// userRepo := repository.NewInMemoryUserRepository()
	userRepo := repository.NewSqliteUserRepository(db)
	settingRepo := repository.NewSqliteSettingRepository(db)

	// -- 1.2 初始化业务逻辑服务层 (注入 Repo) --
	userSvc := service.NewUserService(userRepo)
	settingSvc := service.NewSettingService(settingRepo)

	// -- 1.3 初始化 Wails 控制器 (暴露给前端 JS 的接口层，注入业务服务) --
	userBinding := binding.NewUserBinding(userSvc)
	settingBinding := binding.NewSettingBinding(settingSvc)

	// -- 1.4 初始化主应用生命周期管家 --
	coreApp := app.NewApp()
	_ = coreApp

	// 【2. 构建 Wails 应用实例】
	wailsApp := application.New(application.Options{
		Name:        "myapp2",
		Description: "A demo application with large-scale architecture best-practices",
		// 【注册所有想要暴露给前台调用的 Bindings】
		Services: []application.Service{
			application.NewService(userBinding),
			application.NewService(settingBinding),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// 【3. 关联 Wails 全局生命周期事件】
	// 在退出主函数前，调用了我们自定义的优雅停机代码
	defer coreApp.Shutdown(context.Background())

	// 【4. 窗口管理器 & 系统托盘初始化】
	winManager := manager.NewWindowManager(wailsApp)
	winManager.CreateMainWindow() // 创建主窗口
	winManager.SetupSystemTray()  // 挂载系统托盘图标和菜单
	_ = winManager

	// 【5. 阻塞式运行启动】
	zap.S().Info("Wails主进程启动中...")
	err = wailsApp.Run()
	if err != nil {
		zap.S().Fatal("运行中崩溃退出: ", err)
	}
}
