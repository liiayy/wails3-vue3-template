package domain

// User 代表系统用户的核心业务实体
type User struct {
	ID    int    `json:"id" gorm:"primaryKey;autoIncrement"` // 数据库主键
	Name  string `json:"name" gorm:"size:255;not null"`      // 字符长度或不为空
	Email string `json:"email" gorm:"unique;not null"`       // 唯一索引约束
}

// UserRepository 定义了用户的数据存取抽象接口
// 不论底层使用的是 SQLite、MySQL、Redis 或是远程 API，Service 层仅依赖此接口
type UserRepository interface {
	FindByID(id int) (*User, error)
	Save(user *User) error
}
