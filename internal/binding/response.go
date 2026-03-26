package binding

import (
	"myapp2/internal/domain"
)

// Result 是 Wails 暴露给前端的统一 JSON 响应格式。
// 无论成功还是失败，都返回 200 (Wails 层级)，业务逻辑通过 success 字段判断。
type Result struct {
	Success bool             `json:"success"`         // 业务执行是否成功
	Data    interface{}      `json:"data,omitempty"`  // 成功时携带的数据
	Error   *domain.AppError `json:"error,omitempty"` // 失败时携带的业务错误对象
}

// Success 封装成功响应
func Success(data interface{}) *Result {
	return &Result{
		Success: true,
		Data:    data,
	}
}

// Failure 封装失败响应
func Failure(err error) *Result {
	if err == nil {
		return Success(nil)
	}

	// 转换领域错误为 Result 结构
	if ae, ok := err.(*domain.AppError); ok {
		return &Result{
			Success: false,
			Error:   ae,
		}
	}

	// 降级处理：将普通错误转为内部错误
	return &Result{
		Success: false,
		Error:   domain.ErrInternal(err.Error(), err),
	}
}
