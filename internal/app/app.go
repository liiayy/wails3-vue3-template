package app

import (
	"context"

	"github.com/wailsapp/wails/v3/pkg/application"
	"myapp2/internal/domain"
)

// App 结构体负责整合 Wails 整个应用的生命周期（Startup, Shutdown...）
// 业务状态管理，如托盘控制、崩溃拦截等也可以放在此类
type App struct {
	wailsApp *application.App
	ctx      context.Context
	logger   domain.Logger
}

// NewApp 构造注入
func NewApp(logger domain.Logger) *App {
	return &App{logger: logger}
}

// Startup 用于捕获 Wails 生命周期事件。此处可做如: 初始化 DB、配置加载、定时任务。
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.logger.Info("[Wails App Startup] 执行生命周期 hook，初始化基础设施...")
}

// Shutdown 拦截退出信号，如果连有数据库，需要优雅断开池
func (a *App) Shutdown(ctx context.Context) {
	a.logger.Info("[Wails App Shutdown] 执行清理和优雅退出流程...")
}

// OnWindowResized 如果需要跟踪主窗口重塑尺寸的自定义事件：
func (a *App) OnWindowResized() {
	// ...保存窗口尺寸落盘记录...
}
