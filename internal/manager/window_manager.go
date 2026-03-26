package manager

import (
	_ "embed"
	"myapp2/internal/config"
	"myapp2/internal/service"

	"go.uber.org/zap"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

//go:embed tray_icon.ico
var trayIcon []byte

// WindowManager 统一管控所有 Wails 窗口的创建、显示、隐藏与销毁。同时托管系统托盘的生命周期。
type WindowManager struct {
	app        *application.App
	mainWindow *application.WebviewWindow
	tray       *application.SystemTray
	// 从配置中心注入
	winWidth  int
	winHeight int
	winTitle  string
}

// NewWindowManager 在 Wails App 创建之后调用
func NewWindowManager(app *application.App, width, height int, title string) *WindowManager {
	return &WindowManager{
		app:       app,
		winWidth:  width,
		winHeight: height,
		winTitle:  title,
	}
}

// ---------- 窗口管理 ----------

// CreateMainWindow 创建主窗口
func (wm *WindowManager) CreateMainWindow() *application.WebviewWindow {
	zap.S().Info("[WindowManager] 创建主窗口...")

	// 确定初始位置控制逻辑
	initialPos := application.WindowCentered
	winX, winY := 0, 0
	restoreX, restoreY := config.Cfg.Window.X, config.Cfg.Window.Y

	if restoreX != -1 && restoreY != -1 {
		// 【增强】坐标有效性校验：检查保存的坐标是否在任何当前活跃的屏幕范围内
		screens := wm.app.Screen.GetAll()
		foundValidScreen := false

		// 如果在启动早期还没拿到屏幕信息（数组为空），我们暂时信任坐标
		if len(screens) == 0 {
			zap.S().Debug("[WindowManager] 启动早期暂未获取到屏幕列表，信任保存坐标")
			foundValidScreen = true
		} else {
			for _, s := range screens {
				// 如果保存的 X,Y 在任意屏幕矩形区域内，则视为合法
				if restoreX >= s.X && restoreX < (s.X+s.Bounds.Width) &&
					restoreY >= s.Y && restoreY < (s.Y+s.Bounds.Height) {
					foundValidScreen = true
					break
				}
			}
		}

		if foundValidScreen && restoreX > -10000 && restoreY > -10000 {
			initialPos = application.WindowXY
			winX = restoreX
			winY = restoreY
			zap.S().Infof("[WindowManager] 恢复窗口位置: (%d, %d)", winX, winY)
		} else {
			if !foundValidScreen {
				zap.S().Warnf("[WindowManager] 检测到保存的坐标 (%d, %d) 已超出当前显示器范围，将重置居中", restoreX, restoreY)
			} else {
				zap.S().Warnf("[WindowManager] 检测到保存的坐标为 Windows 最小化状态 (-32000)，将重置居中以防止窗口丢失")
			}
			// initialPos 保持 WindowCentered 即可
		}
	}

	wm.mainWindow = wm.app.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:             "main",
		Title:            wm.winTitle,
		Width:            wm.winWidth,
		Height:           wm.winHeight,
		MinWidth:         1024,
		MinHeight:        750,
		InitialPosition:  initialPos,
		X:                winX,
		Y:                winY,
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
		EnableFileDrop:   true,
		Frameless:        true,
	})

	// 监听原生文件拖放事件
	wm.mainWindow.OnWindowEvent(events.Common.WindowFilesDropped, func(ev *application.WindowEvent) {
		files := ev.Context().DroppedFiles()
		zap.S().Infof("[WindowManager] 原生文件拖入: %v", files)
		// 发送给前端通用事件总线
		wm.app.Event.Emit("files-dropped", files)
	})

	// 【新增】监听窗口从最大化/最小化还原事件
	wm.mainWindow.OnWindowEvent(events.Common.WindowRestore, func(ev *application.WindowEvent) {
		zap.S().Debug("[WindowManager] 窗口已还原，重新应用尺寸约束并更新状态")
		wm.mainWindow.SetMinSize(1024, 750)
		config.UpdateWindowMaximizedState(false)
	})

	// 【新增】监听窗口最大化事件
	wm.mainWindow.OnWindowEvent(events.Common.WindowMaximise, func(ev *application.WindowEvent) {
		zap.S().Debug("[WindowManager] 窗口已最大化，更新状态")
		config.UpdateWindowMaximizedState(true)
	})

	// 【新增】监听窗口位移事件，保存坐标
	wm.mainWindow.OnWindowEvent(events.Common.WindowDidMove, func(ev *application.WindowEvent) {
		x, y := wm.mainWindow.Position()
		// 关键过滤：Windows 最小化时坐标会变为 (-32000, -32000)，绝对不能保存这个状态
		if x < -10000 || y < -10000 {
			zap.S().Debugf("[WindowManager] 窗口移动事件: 忽略最小化状态坐标 (%d, %d)", x, y)
			return
		}
		zap.S().Debugf("[WindowManager] 窗口移动，新坐标: (%d, %d)", x, y)
		config.UpdateWindowPosition(x, y)
	})

	// 【修改】监听窗口缩放结束事件，保存尺寸到配置文件
	wm.mainWindow.OnWindowEvent(events.Common.WindowDidResize, func(ev *application.WindowEvent) {
		// 重要：如果窗口当前是最大化状态，不保存尺寸（防止把全屏尺寸存为默认尺寸）
		if wm.mainWindow.IsMaximised() {
			return
		}
		w, h := wm.mainWindow.Size()
		zap.S().Infof("[WindowManager] 窗口缩放结束，保存新尺寸: %dx%d", w, h)
		config.UpdateWindowSize(w, h)
	})

	// 【新增】根据配置决定是否启动即最大化
	if config.Cfg.Window.IsMaximized {
		zap.S().Info("[WindowManager] 根据配置，启动即最大化主窗口")
		wm.mainWindow.Maximise()
	}

	return wm.mainWindow
}

// CreateAboutWindow 创建一个"关于"窗口
func (wm *WindowManager) CreateAboutWindow() *application.WebviewWindow {
	zap.S().Info("[WindowManager] 创建关于窗口...")

	if w, ok := wm.app.Window.GetByName("about"); ok {
		w.Show()
		return w.(*application.WebviewWindow)
	}

	aboutWin := wm.app.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:             "about",
		Title:            service.GetI18n().T("about.title"),
		Width:            400,
		Height:           400,
		URL:              "/#/standalone/about",
		BackgroundColour: application.NewRGB(27, 38, 54),
		Frameless:        true,
		DisableResize:    true,
		HideOnEscape:     true,
	})

	return aboutWin
}

// ShowMainWindow 把主窗口调到前景
func (wm *WindowManager) ShowMainWindow() {
	if wm.mainWindow != nil {
		wm.mainWindow.Show()
		wm.mainWindow.Focus()
	}
}

// ---------- 系统托盘管理 ----------

// SetupSystemTray 配置系统托盘图标和右键菜单
func (wm *WindowManager) SetupSystemTray() {
	zap.S().Info("[WindowManager] 初始化系统托盘...")

	wm.tray = wm.app.SystemTray.New()
	wm.tray.SetIcon(trayIcon)

	// 监听前端发出的语言变更
	wm.app.Event.On("app:settings-changed", func(ev *application.CustomEvent) {
		zap.S().Infof("[Tray] 捕获到设置变更事件: %v", ev.Data)
		data, ok := ev.Data.(map[string]interface{})
		if !ok {
			zap.S().Warn("[Tray] 事件数据类型转换失败")
			return
		}

		if data["key"] == "language" {
			lang := data["value"].(string)
			zap.S().Infof("[Tray] 正在将后端语言同步至: %s 并刷新菜单", lang)
			// 核心修复：在这里也显式同步一次，确保单例状态最新
			service.GetI18n().SetLanguage(lang)
			wm.RefreshTrayMenu()
		}
	})

	wm.RefreshTrayMenu()

	// 单击托盘图标
	wm.tray.OnClick(func() {
		wm.ShowMainWindow()
	})
}

// RefreshTrayMenu 构建并重置托盘右键菜单 (用于 I18n 反馈)
func (wm *WindowManager) RefreshTrayMenu() {
	i18n := service.GetI18n()
	trayMenu := application.NewMenu()

	// 动态显示主窗口
	trayMenu.Add(i18n.T("tray_show")).OnClick(func(ctx *application.Context) {
		wm.ShowMainWindow()
	})

	// 动态关于
	trayMenu.Add(i18n.T("tray_about")).OnClick(func(ctx *application.Context) {
		wm.CreateAboutWindow()
	})

	trayMenu.AddSeparator()

	// 彻底退出
	trayMenu.Add(i18n.T("tray_exit")).OnClick(func(ctx *application.Context) {
		zap.S().Info("[Tray] 用户点击了退出菜单...")
		wm.app.Quit()
	})

	wm.tray.SetMenu(trayMenu)
}
