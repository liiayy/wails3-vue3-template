package manager

import (
	_ "embed"

	"go.uber.org/zap"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed tray_icon.ico
var trayIcon []byte

// WindowManager 统一管控所有 Wails 窗口的创建、显示、隐藏与销毁。
// 同时托管系统托盘的生命周期。
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

// CreateMainWindow 创建主窗口，只应调用一次
func (wm *WindowManager) CreateMainWindow() *application.WebviewWindow {
	zap.S().Info("[WindowManager] 创建主窗口...")

	wm.mainWindow = wm.app.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:   "main",
		Title:  wm.winTitle,
		Width:  wm.winWidth,
		Height: wm.winHeight,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
	})

	return wm.mainWindow
}

// CreateSettingsWindow 创建一个独立的设置窗口
func (wm *WindowManager) CreateSettingsWindow() *application.WebviewWindow {
	zap.S().Info("[WindowManager] 创建设置窗口...")

	// 检查是否已经有名为 settings 的窗口
	if w, ok := wm.app.Window.GetByName("settings"); ok {
		w.Show()
		return w.(*application.WebviewWindow)
	}

	settingsWin := wm.app.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:             "settings",
		Title:            "系统设置",
		Width:            720,
		Height:           520,
		URL:              "/#/standalone/settings",
		BackgroundColour: application.NewRGB(27, 38, 54),
		// 按 Escape 键自动隐藏此窗口
		HideOnEscape: true,
	})

	return settingsWin
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
		Title:            "关于 MyApp2",
		Width:            400,
		Height:           300,
		URL:              "/#/standalone/about",
		BackgroundColour: application.NewRGB(27, 38, 54),
		HideOnEscape:     true,
	})

	return aboutWin
}

// ShowMainWindow 把主窗口调到前景
func (wm *WindowManager) ShowMainWindow() {
	if wm.mainWindow != nil {
		wm.mainWindow.Show()
	}
}

// ---------- 系统托盘管理 ----------

// SetupSystemTray 配置系统托盘图标和右键菜单
func (wm *WindowManager) SetupSystemTray() {
	zap.S().Info("[WindowManager] 初始化系统托盘...")

	wm.tray = wm.app.SystemTray.New()
	wm.tray.SetIcon(trayIcon)

	// 构建托盘右键菜单
	trayMenu := application.NewMenu()

	trayMenu.Add("显示主窗口").OnClick(func(ctx *application.Context) {
		wm.ShowMainWindow()
	})

	trayMenu.Add("打开设置").OnClick(func(ctx *application.Context) {
		wm.CreateSettingsWindow()
	})

	trayMenu.Add("关于").OnClick(func(ctx *application.Context) {
		wm.CreateAboutWindow()
	})

	trayMenu.AddSeparator()

	trayMenu.Add("退出程序").OnClick(func(ctx *application.Context) {
		zap.S().Info("[Tray] 用户点击了退出菜单...")
		wm.app.Quit()
	})

	wm.tray.SetMenu(trayMenu)

	// 单击托盘图标：显示/隐藏主窗口
	wm.tray.OnClick(func() {
		wm.ShowMainWindow()
	})
}
