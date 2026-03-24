package binding

import (
	"myapp2/internal/service"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/services/notifications"
	"go.uber.org/zap"
)

type NotificationBinding struct {
	notifier *notifications.NotificationService
}

func NewNotificationBinding(notifier *notifications.NotificationService) *NotificationBinding {
	b := &NotificationBinding{
		notifier: notifier,
	}

	// 初始化翻译并注册分类
	b.RefreshCategories()

	// 监听通知响应
	notifier.OnNotificationResponse(func(result notifications.NotificationResult) {
		if result.Error != nil {
			zap.S().Errorf("[Notification] 响应错误: %v", result.Error)
			return
		}

		resp := result.Response
		zap.S().Infof("[Notification] 收到交互响应 ID=%s, Action=%s, Text=%s",
			resp.ID, resp.ActionIdentifier, resp.UserText)

		// 转发给前端处理
		application.Get().Event.Emit("notification-clicked", resp)
	})

	return b
}

// SetLanguage 供前端调用同步语言环境
func (b *NotificationBinding) SetLanguage(lang string) {
	zap.S().Infof("[I18n] 后端语言切换至: %s", lang)
	service.GetI18n().SetLanguage(lang)
	b.RefreshCategories() // 核心：重新注册分类以刷新按钮标题
}

// RefreshCategories 根据当前语言环境注册/更新通知类别
func (b *NotificationBinding) RefreshCategories() {
	i18n := service.GetI18n()
	categoryID := "demo-reply"

	category := notifications.NotificationCategory{
		ID: categoryID,
		Actions: []notifications.NotificationAction{
			{
				ID:    "APPROVE",
				Title: i18n.T("notif_approve"), // "同意审批" / "Approve"
			},
			{
				ID:          "REJECT",
				Title:       i18n.T("notif_reject"), // "驳回申请" / "Reject"
				Destructive: true,
			},
		},
		HasReplyField:    true,
		ReplyPlaceholder: i18n.T("notif_reply_placeholder"), // "在此输入您的意见..."
		ReplyButtonTitle: i18n.T("notif_reply_button"),      // "提交变更"
	}

	b.notifier.RegisterNotificationCategory(category)
}

// SendBasic 发送最基础的通知
func (b *NotificationBinding) SendBasic(title, body string) error {
	i18n := service.GetI18n()
	// 如果前端没传参数，使用后端默认的翻译 Key
	if title == "" {
		title = i18n.T("notif_basic_title")
	}
	if body == "" {
		body = i18n.T("notif_basic_body")
	}

	zap.S().Infof("[Notification] 发送基础通知: %s", title)
	err := b.notifier.SendNotification(notifications.NotificationOptions{
		ID:    "basic-demo",
		Title: title,
		Body:  body,
	})
	return err
}

// SendWithSubtitle 发送带副标题的通知
func (b *NotificationBinding) SendWithSubtitle(title, subtitle, body string) error {
	i18n := service.GetI18n()
	if subtitle == "" {
		subtitle = i18n.T("notif_subtitle_demo")
	}

	err := b.notifier.SendNotification(notifications.NotificationOptions{
		ID:       "subtitle-demo",
		Title:    title,
		Subtitle: subtitle,
		Body:     body,
	})
	return err
}

// SendInteractive 发送带有操作按钮和回复框的通知
func (b *NotificationBinding) SendInteractive(title, body string) error {
	err := b.notifier.SendNotificationWithActions(notifications.NotificationOptions{
		ID:         "interactive-demo",
		Title:      title,
		Body:       body,
		CategoryID: "demo-reply",
	})
	return err
}
