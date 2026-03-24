package service

import (
	"embed"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed locales/*.json
var localeFS embed.FS

type I18nService struct {
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
	lang      string
	mu        sync.RWMutex
}

var (
	instance *I18nService
	once     sync.Once
)

// GetI18n 获取国际化服务单例
func GetI18n() *I18nService {
	once.Do(func() {
		bundle := i18n.NewBundle(language.Chinese)
		bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

		s := &I18nService{
			bundle: bundle,
			lang:   "zh-CN",
		}
		s.loadResources()
		s.updateLocalizer()
		instance = s
	})
	return instance
}

func (s *I18nService) loadResources() {
	// 加载嵌入的 JSON 文件
	files := []string{"zh-CN.json", "en-US.json"}
	for _, f := range files {
		data, err := localeFS.ReadFile("locales/" + f)
		if err != nil {
			fmt.Printf("Error reading locale file %s: %v\n", f, err)
			continue
		}
		s.bundle.MustParseMessageFileBytes(data, f)
	}
}

func (s *I18nService) updateLocalizer() {
	s.localizer = i18n.NewLocalizer(s.bundle, s.lang)
}

// SetLanguage 切换语言 (由前端调用同步)
func (s *I18nService) SetLanguage(lang string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lang = lang
	s.updateLocalizer()
}

// T 核心翻译方法 (常用写法)
func (s *I18nService) T(messageID string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	translation, err := s.localizer.Localize(&i18n.LocalizeConfig{
		MessageID: messageID,
	})
	if err != nil {
		return messageID // 翻译失败则返回原始 ID
	}
	return translation
}

// TP 支持占位符的翻译方法
func (s *I18nService) TP(messageID string, templateData map[string]interface{}) string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	translation, err := s.localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: templateData,
	})
	if err != nil {
		return messageID
	}
	return translation
}
