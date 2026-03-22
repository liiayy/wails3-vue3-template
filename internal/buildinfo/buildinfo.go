package buildinfo

// 以下变量在编译期通过 -ldflags 注入，用户无法修改。
// 编译命令示例：
//
//	go build -ldflags "-X myapp2/internal/buildinfo.Version=1.0.0 -X myapp2/internal/buildinfo.IsDev=false"
var (
	// Version 应用程序版本号
	Version = "dev"

	// IsDev 是否为开发模式（仅在编译时通过 ldflags 设定）
	// 默认值 "true" 表示未经正式构建流程直接 go run 时自动为开发模式
	IsDev = "true"
)

// IsDevMode 返回当前是否处于开发模式
func IsDevMode() bool {
	return IsDev == "true"
}
