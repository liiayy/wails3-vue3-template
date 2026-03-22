# Wails v3 大型工程化最佳实践架构指南

随着业务的增长，桌面应用的需求就不再仅仅是对一段代码进行简单的界面封装，而是涉及到**本地数据持久化（数据库）**、**复杂状态管理（State）**、**本地系统调用交互（OS/Filesystem）**、**日志追踪（Logging）**、**多窗口独立管理**等一系列需求。此时，我们的 Wails 3 架构必须遵循“领域驱动”和“高内聚低耦合”的原则。

下面我将以 Go + Vue 3 的体系为您详细剖析大型工程的最佳实践架构。

---

## 一、 整体目录组织结构（The Mega-Structure）

对于包含大量业务逻辑的桌面应用，我们不再把代码杂糅在根目录，而是采用标准企业级包结构组织：

```text
my_large_app/
├── cmd/
│   └── app/
│       └── main.go              # 【入口】非常轻量，仅用于组装、启动应用、依赖注入
├── internal/                    # 【私有核心】业务代码主体，外部无法 import
│   ├── app/                     # Wails 生命周期管理 (Startup, Shutdown) 和主程序状态
│   ├── config/                  # 读取和管理本地配置文件（YAML/JSON/TOML）
│   ├── database/                # SQLite/BoltDB 等本地数据库连接引擎与 ORM 初始化
│   ├── domain/                  # 【领域模型】核心业务实体模型 (Domain Models) 和接口定义
│   ├── event/                   # 事件总线定义，所有的 app.Event.Emit() 常量及封装
│   ├── repository/              # 数据持久层，负责 CRUD 和数据库的具体交互
│   ├── service/                 # 【业务服务】领域纯业务逻辑，负责将各模块粘合
│   ├── binding/                 # 【Wails 桥接层】专门暴露给前端的接口，负责接收前端参数并调用 service
│   └── logger/                  # 结构化日志引擎（如 Zap/Zerolog），输出到本地滚动文件
├── pkg/                         # 【公共包】可以被其他任何项目复用的纯工具函数 (Utils)
├── docs/                        # 架构设计文档、API 文档
└── frontend/                    # 【前端工程】
    ├── bindings/                # 自动生成的代码 (Auto Generated - 不可编辑)
    └── src/
        ├── api/                 # 对 bindings 进行二次封装，提供拦截、防抖、Mock 数据
        ├── assets/              # 静态资源 (Icon, CSS, Fonts)
        ├── components/          # 哑组件 / 通用 UI 组件 (TDesign/Tailwind 封装)
        ├── composables/         # Vue 3 组合式 API (如 useWailsEvent, useWindowHook) 
        ├── config/              # 前端环境配置
        ├── i18n/                # 国际化多语言支持
        ├── layouts/             # 多层次布局结构 (LoginLayout, MainLayout...)
        ├── router/              # Vue Router 集中管理与路由守卫
        ├── stores/              # Pinia 状态管理树，按模块拆分 (user, setting, app)
        ├── types/               # 全局 TypeScript 接口定义
        ├── utils/               # 前端纯函数工具库
        └── views/               # 页面级视图组件 / 智能组件
```

---

## 二、 后端 (Go) 架构拆解与设计模式

### 1. 隔离 Wails 层与业务层 (MVC / 三层架构)
将前端可以直接调用的方法（`binding` 目录）与真正的业务逻辑（`service` 目录）完全隔离开来。

*   **`domain`**: 定义业务对象。例えば `type User struct { ID int; Name string }`。
*   **`repository`**: 定义底层接口，比如 `SaveUser(user User) error`，屏蔽了 SQLite/文件系统 的底层实现细节。
*   **`service`**: 纯粹的业务处理。不引入 Wails SDK 依赖包，也不依赖具体的表现层。
*   **`binding` (控制器)**: 这是我们在 `main.go` 中真正注册给 Wails Application 的服务。它只负责接收参数、参数校验，并调用 `service` 处理，最后返回结果。这样即使有朝一日你需要把核心业务做成 Web 服务或者命令行，你的 `service` 和 `domain` 都不需要改代码。

### 2. 生命周期与依赖注入控制
大型应用需要在启动阶段做很多事。在 `internal/app/app.go` 里，构建统一跨层管理的结构体：

```go
package app

import (
    "context"
    "github.com/wailsapp/wails/v3/pkg/application"
)

type Application struct {
    wailsApp *application.App
    ctx      context.Context
}

// Startup 在 Wails 的 DomReady 之前/之后触发 (具体看注册点)
// 这里用于：检查环境配置、连接 SQLite 本地数据库、启动后台定时任务等
func (a *Application) Startup(ctx context.Context) {
    a.ctx = ctx
    // 初始化 logger 等...
}

// Shutdown 优雅停机
func (a *Application) Shutdown(ctx context.Context) {
    // 在此处安全关闭数据库连接池、释放文件锁、通知后台协程退出等
}
```

### 3. 规范化事件派发 (Event Bus)
对于大型项目，不要在代码各处随意硬编码字符串（如 `"time"` 或 `"user-login"`）。
在 `internal/event/keys.go` 中统一定义事件名，甚至可以将 Emit 函数直接封装在 event 包内：
```go
const (
    EventUserLoginSuccess = "event:user:login_success"
    EventSystemFileLoaded = "event:system:file_loaded"
)
```
这可有效避免拼写错误并为未来代码重构提供巨大便利。

---

## 三、 前端 (Vue 3) 大型应用实战指南

### 1. 二次封装生成的 Wails Bindings（API 层隔离）
在大型开发团队中，往往后端还没写好，前端需要先行开发。如果组件直接拉取 `bindings` 里的对象，将很难接入 Mock。建议将 API 单独抽出便于拦截与异常统一处理：
```typescript
// src/api/user.ts
import { UserBinding } from '../../bindings/myapp2/internal/binding'
import { MessagePlugin } from 'tdesign-vue-next'

export async function getUserProfile(id: number) {
  try {
    // 调用生成的 Wails API
    const res = await UserBinding.GetProfile(id);
    return res;
  } catch (error) {
    // 统一捕获处理系统错误，对接全局交互组件
    MessagePlugin.error(`获取资料失败: ${error}`);
    throw error;
  }
}
```

### 2. 拥抱 Vue 3 Composables (组合式 API) 管理事件流
Wails 具有庞大的本地通信特性（如本地文件处理进度条事件推送等）。直接写死极其容易导致事件内存泄漏。编写可复用的 `hook` 可以利用 Vue 的生命周期自动帮你扫尾：

```typescript
// src/composables/useWailsEvent.ts
import { onMounted, onUnmounted } from 'vue'
import { Events } from '@wailsio/runtime'

/**
 * 自动管理 Wails 后端事件生命周期的 hook
 */
export function useWailsEvent(eventName: string, callback: (...args: any[]) => void) {
  onMounted(() => {
    Events.On(eventName, callback)
  })
  
  onUnmounted(() => {
    // 注意：具体需调用 Wails 的 Events.Off 或者注销当前特定回调，避免卸载全部
    Events.Off(eventName) 
  })
}
```
然后在组件中引入极为优雅：`useWailsEvent('event:system:file_loaded', handler)`

### 3. 多窗口独立路由支持架构
在大型应用中（例如微信客户端的主面板和独立的独立图片查看器），需要多窗口协作。前端 `router` 建议配置不同布局的独立入口路径：
*   `/` -> MainWindowLayout (带有侧边栏和常规拖拽条)
*   `/viewer` -> ImageViewerLayout (只有关闭按钮，没有侧边栏)
*   `/settings` -> SettingsModalLayout (纯白底，固定大小和无边框布局)

后端 Go 代码中，通过创建不同的 Wails Window Options 指向不同的起始 URL：
`app.Window.NewWithOptions(..., URL: "/viewer")`。

---

## 四、 其他不可忽视的工程化刚需特征

1. **结构化本地日志持久化引擎**
   大型程序**极易**因为用户本地差异化的 OS 环境（权限、缺包、杀毒软件）导致不可预知的崩溃。**必须**引入像 `zap`，并集成 `lumberjack` 进行日志切割（Rolling log，防止几十MB卡死），把所有的 `fmt.Println` 替换为 `logger.Info`。日志通常保存在 Windows 的 `AppData` 隐藏目录下。
2. **多环境管理逻辑 (跨环境编译隔离)**
   开发阶段和生产包的环境通常不同（如本地 DevTools 日志开关状态等）。在 Go 里面使用 `//go:build dev` 和 `//go:build prod` 的 build tags 来执行条件编译；前端配合 Vite 的 `.env.development` 分开环境变量及功能。
3. **Panic 恢复与自动崩溃上报**
   桌面应用在 Go 产生的核心 Panic 会直接彻底闪退。应该利用 `recover()` 拦截 Panic，并使用一个带有弹窗提示的 Dialog (通过 Wails 包内部的原生弹窗控件)，提醒用户发生异常并将崩溃 Callstack 附带写入本地日志，最后安全结束程序。
