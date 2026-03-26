package domain

import "context"

// User 代表系统用户的核心业务实体
type User struct {
	ID    int    `json:"id" gorm:"primaryKey;autoIncrement"` // 数据库主键
	Name  string `json:"name" gorm:"size:255;not null"`      // 字符长度或不为空
	Email string `json:"email" gorm:"unique;not null"`       // 唯一索引约束
}

// UserListResult 分页查询结果
type UserListResult struct {
	Items []*User `json:"items"`
	Total int64   `json:"total"`
}

// UserRegisterRequest 注册接口入参
type UserRegisterRequest struct {
	Name  string `json:"name" validate:"required,min=2,max=20"` // 必填，2-20字符
	Email string `json:"email" validate:"required,email"`       // 必填，合法邮箱
}

// UserUpdateRequest 更新接口入参
type UserUpdateRequest struct {
	ID    int    `json:"id" validate:"required"`                 // 必填 ID
	Name  string `json:"name" validate:"omitempty,min=2,max=20"` // 可选，若有则 2-20 字符
	Email string `json:"email" validate:"omitempty,email"`       // 可选，若有则合法邮箱
}

// UserRepository 定义了用户的数据存取抽象接口
// 不论底层使用的是 SQLite、MySQL、Redis 或是远程 API，Service 层仅依赖此接口
type UserRepository interface {
	FindByID(ctx context.Context, id int) (*User, error)
	Save(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, keyword string, page, pageSize int) (*UserListResult, error)
	GetAll(ctx context.Context) ([]*User, error)
}
