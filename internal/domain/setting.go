package domain

// Setting 领域模型：用于存放应用级的键值对设置信息
type Setting struct {
	Key   string `json:"key" gorm:"primaryKey"` // 设置项的名称
	Value string `json:"value"`                 // 设置项的值（通常为 JSON 字符串或简单字串）
}

// SettingRepository 定义了设置信息的持久化抽象
type SettingRepository interface {
	Get(key string) (string, error)
	Set(key, value string) error
	GetAll() (map[string]string, error)
}
