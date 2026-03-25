package binding

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

type ScreenService struct {
	app *application.App
}

func NewScreenService(app *application.App) *ScreenService {
	return &ScreenService{app: app}
}

type ScreenInfo struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	X           int     `json:"x"`       // 逻辑 X
	Y           int     `json:"y"`       // 逻辑 Y
	Width       int     `json:"width"`   // 逻辑宽
	Height      int     `json:"height"`  // 逻辑高
	PX          int     `json:"px"`      // 物理 X
	PY          int     `json:"py"`      // 物理 Y
	PWidth      int     `json:"pWidth"`  // 物理宽
	PHeight     int     `json:"pHeight"` // 物理高
	ScaleFactor float32 `json:"scaleFactor"`
	IsPrimary   bool    `json:"isPrimary"`
}

// GetAllScreens 获取所有显示器详细信息
func (s *ScreenService) GetAllScreens() []ScreenInfo {
	screens := s.app.Screen.GetAll()
	var result []ScreenInfo

	for _, screen := range screens {
		result = append(result, ScreenInfo{
			ID:          screen.ID,
			Name:        screen.Name,
			X:           screen.Bounds.X,
			Y:           screen.Bounds.Y,
			Width:       screen.Bounds.Width,
			Height:      screen.Bounds.Height,
			PX:          screen.PhysicalBounds.X,
			PY:          screen.PhysicalBounds.Y,
			PWidth:      screen.PhysicalBounds.Width,
			PHeight:     screen.PhysicalBounds.Height,
			ScaleFactor: screen.ScaleFactor,
			IsPrimary:   screen.IsPrimary,
		})
	}
	return result
}

// GetPrimaryScreen 获取主显示器
func (s *ScreenService) GetPrimaryScreen() ScreenInfo {
	screen := s.app.Screen.GetPrimary()
	return ScreenInfo{
		ID:          screen.ID,
		Name:        screen.Name,
		X:           screen.Bounds.X,
		Y:           screen.Bounds.Y,
		Width:       screen.Bounds.Width,
		Height:      screen.Bounds.Height,
		PX:          screen.PhysicalBounds.X,
		PY:          screen.PhysicalBounds.Y,
		PWidth:      screen.PhysicalBounds.Width,
		PHeight:     screen.PhysicalBounds.Height,
		ScaleFactor: screen.ScaleFactor,
		IsPrimary:   screen.IsPrimary,
	}
}
