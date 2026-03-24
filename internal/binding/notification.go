package binding

import (
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/services/notifications"
	"go.uber.org/zap"
)

type NotificationBinding struct {
	notifier *notifications.NotificationService
}

func NewNotificationBinding(notifier *notifications.NotificationService) *NotificationBinding {
	// 初始化通知类别（针对交互式通知）
	categoryID := "demo-reply"
	category := notifications.NotificationCategory{
		ID: categoryID,
		Actions: []notifications.NotificationAction{
			{
				ID:    "APPROVE",
				Title: "同意",
			},
			{
				ID:          "REJECT",
				Title:       "显绝",
				Destructive: true,
			},
		},
		HasReplyField:    true,
		ReplyPlaceholder: "输入回复内容...",
		ReplyButtonTitle: "发送回复",
	}

	notifier.RegisterNotificationCategory(category)

	// 监听通知响应
	notifier.OnNotificationResponse(func(result notifications.NotificationResult) {
		if result.Error != nil {
			zap.S().Errorf("[Notification] 响应错误: %v", result.Error)
			return
		}

		resp := result.Response
		zap.S().Infof("[Notification] 收到交互响应 ID=%s, Action=%s, Text=%s",
			resp.ID, resp.ActionIdentifier, resp.UserText)

		// 转发给前端
		application.Get().Event.Emit("notification-clicked", resp)
	})

	return &NotificationBinding{
		notifier: notifier,
	}
}

// SendBasic 发送最基础的通知
func (b *NotificationBinding) SendBasic(title, body string) error {
	zap.S().Infof("[Notification] 发送基础通知: %s - %s", title, body)
	err := b.notifier.SendNotification(notifications.NotificationOptions{
		ID:    "basic-demo",
		Title: title,
		Body:  body,
	})
	if err != nil {
		zap.S().Errorf("[Notification] SendBasic Error: %v", err)
	}
	return err
}

// SendWithSubtitle 发送带副标题的通知 (macOS/Linux)
func (b *NotificationBinding) SendWithSubtitle(title, subtitle, body string) error {
	zap.S().Infof("[Notification] 发送带副标题通知: %s - %s - %s", title, subtitle, body)
	err := b.notifier.SendNotification(notifications.NotificationOptions{
		ID:       "subtitle-demo",
		Title:    title,
		Subtitle: subtitle,
		Body:     body,
		Data: map[string]interface{}{
			"sender": "Antigravity AI",
		},
	})
	if err != nil {
		zap.S().Errorf("[Notification] SendWithSubtitle Error: %v", err)
	}
	return err
}

// SendInteractive 发送带有操作按钮和回复框的通知
func (b *NotificationBinding) SendInteractive(title, body string) error {
	zap.S().Infof("[Notification] 发送交互式通知: %s - %s", title, body)
	err := b.notifier.SendNotificationWithActions(notifications.NotificationOptions{
		ID:         "interactive-demo",
		Title:      title,
		Body:       body,
		CategoryID: "demo-reply", // 对应上面注册的类别
		Data: map[string]interface{}{
			"sender": "Antigravity AI",
		},
	})
	if err != nil {
		zap.S().Errorf("[Notification] SendInteractive Error: %v", err)
	}
	return err
}
