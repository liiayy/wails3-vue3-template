package main

import (
	"embed"
	_ "embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
	"go.uber.org/zap"

	"myapp2/internal/app"
	"myapp2/internal/binding"
	"myapp2/internal/database"
	"myapp2/internal/logger"
	"myapp2/internal/repository"
	"myapp2/internal/service"
)

//go:embed all:frontend/dist
var assets embed.FS

// main 程序的唯一入口点，它的职责极度专注：仅仅负责对象的实例化、依赖组装和框架启动。
func main() {
	// 【0. 初始化全局日志系统】
	// isDev 设为 true 时可以提供控制台彩色显示（可从环境变量动态获取）
	if err := logger.InitLogger(true, "MyApp2"); err != nil {
		log.Fatalf("无法初始化日志系统: %v", err)
	}
	defer zap.L().Sync() // 确保程序退出前刷新磁盘IO

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

	// -- 1.2 初始化业务逻辑服务层 (注入 Repo) --
	userSvc := service.NewUserService(userRepo)

	// -- 1.3 初始化 Wails 控制器 (暴露给前端 JS 的接口层，注入业务服务) --
	userBinding := binding.NewUserBinding(userSvc)

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
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// 【3. 关联 Wails 全局生命周期事件】
	// Wails v3 alpha版本中生命周期事件挂载 API 有变动，
	// 实际项目中可在此处挂载 coreApp.Startup 等HOOK。

	// 【4. 创建主进程界面窗口】
	wailsApp.Window.NewWithOptions(application.WebviewWindowOptions{
		Title: "Wails 3 Mega-Structure Dashboard",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
	})

	// 【5. 阻塞式运行启动】
	zap.S().Info("Wails主进程启动中...")
	err = wailsApp.Run()
	if err != nil {
		zap.S().Fatal("运行中崩溃退出: ", err)
	}
}
