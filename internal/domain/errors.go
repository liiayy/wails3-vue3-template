package domain

import (
	"errors"
	"fmt"
)

// ErrorCode 用于前端判断错误的具体类型，实现精细化提示
type ErrorCode string

const (
	CodeInternal       ErrorCode = "INTERNAL_ERROR"    // 内部未预期错误
	CodeNotFound       ErrorCode = "NOT_FOUND"         // 资源未找到
	CodeValidation     ErrorCode = "VALIDATION_FAILED" // 输入参数校验失败
	CodeConflict       ErrorCode = "CONFLICT"          // 资源冲突（如 Email 已存在）
	CodeAuthFailed     ErrorCode = "AUTH_FAILED"       // 认证失败
	CodePermissionDeny ErrorCode = "PERMISSION_DENIED" // 权限不足
)

// AppError 是领域定义的统一错误结构，其 JSON 字段将直接暴露给前端
type AppError struct {
	Code    ErrorCode `json:"code"`    // 机器可读的错误码
	Message string    `json:"message"` // 用户可读的详细提示
	Err     error     `json:"-"`       // 原始错误对象（内部日志记录使用，不传给前端）
}

// Error 实现 Go 内置的 error 接口
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// Unwrap 返回原始错误（用于 errors.Is/As）
func (e *AppError) Unwrap() error {
	return e.Err
}

// NewError 创建一个新的业务异常
func NewError(code ErrorCode, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// ---------------- 常用便捷方法 ----------------

func ErrNotFound(msg string, err error) *AppError {
	return NewError(CodeNotFound, msg, err)
}

func ErrInternal(msg string, err error) *AppError {
	return NewError(CodeInternal, msg, err)
}

func ErrValidation(msg string) *AppError {
	return NewError(CodeValidation, msg, nil)
}

func ErrConflict(msg string) *AppError {
	return NewError(CodeConflict, msg, nil)
}

// IsAppError 检查一个错误是否为业务定义的异常
func IsAppError(err error) bool {
	var e *AppError
	return errors.As(err, &e)
}
