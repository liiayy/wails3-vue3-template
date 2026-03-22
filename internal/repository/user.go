package repository

import (
	"errors"
	"myapp2/internal/domain"
	"sync"
)

// InMemoryUserRepository 实现了 domain.UserRepository 接口 (使用内存模拟数据库)
// 在生产环境中，此处可以替换为 SqliteRepository 注入 GORM
type InMemoryUserRepository struct {
	mu    sync.RWMutex
	store map[int]*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		store: map[int]*domain.User{
			1: {ID: 1, Name: "Admin", Email: "admin@myapp.com"},
		},
	}
}

func (r *InMemoryUserRepository) FindByID(id int) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if user, exists := r.store[id]; exists {
		// 返回对象的拷贝避免并发修改问题
		dbUser := *user
		return &dbUser, nil
	}
	return nil, errors.New("user not found")
}

func (r *InMemoryUserRepository) Save(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 若没有分配 ID，则自动分配
	if user.ID == 0 {
		user.ID = len(r.store) + 1
	}

	// 存入 Map
	u := *user
	r.store[user.ID] = &u
	return nil
}
