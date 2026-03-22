# MyApp2 — Wails 3 大型工程化最佳实践

> 基于 **Wails v3 (Alpha)** + **Go 1.25** + **Vue 3** + **TypeScript** + **Vite** + **TDesign** + **Tailwind CSS 4** 的跨平台桌面应用模板工程。

---

## 一、技术栈一览

| 层面 | 技术 | 用途 |
|---|---|---|
| 后端运行时 | Go 1.25+ | 核心业务逻辑与系统调用 |
| 桌面框架 | Wails v3 (alpha.74) | 将 Go 后端与 Webview 前端合二为一 |
| 前端框架 | Vue 3 (Composition API) | 界面构建 |
| 构建工具 | Vite 8 | 前端热重载与打包 |
| UI 组件库 | TDesign Vue Next | 企业级 UI 组件 |
| CSS 引擎 | Tailwind CSS 4 | 原子化样式 |
| 状态管理 | Pinia 3 | 响应式状态与持久化 |
| 路由 | Vue Router 5 (Hash 模式) | 多窗口路由支持 |
| 数据库 | SQLite (via GORM) | 本地数据持久化 |
| 日志引擎 | Zap + Lumberjack | 结构化日志与滚动切割 |
| 配置中心 | Viper | YAML 配置文件读写 |
| 任务自动化 | Taskfile v3 | 跨平台构建任务编排 |
| 包管理 | pnpm (前端) / Go Modules (后端) | 依赖管理 |

---

## 二、项目目录结构

```text
myapp2/
├── main.go                          # 程序入口（依赖注入 & 框架启动）
├── go.mod / go.sum                  # Go 依赖声明
├── Taskfile.yml                     # 构建任务
├── docs/
│   └── architecture.md              # 架构设计文档
│
├── internal/                        # 【Go 私有核心包】
│   ├── app/
│   │   └── app.go                   # Wails 生命周期管理 (Startup/Shutdown)
│   ├── binding/
│   │   ├── user.go                  # 用户服务桥接层（暴露给前端 JS）
│   │   └── setting.go               # 设置服务桥接层
│   ├── buildinfo/
│   │   └── buildinfo.go             # 编译期注入版本号与环境标识（ldflags）
│   ├── config/
│   │   └── config.go                # Viper YAML 配置中心
│   ├── database/
│   │   └── database.go              # GORM + SQLite 连接与自动迁移
│   ├── domain/
│   │   ├── user.go                  # User 领域模型 + Repository 接口
│   │   └── setting.go               # Setting 领域模型 + Repository 接口
│   ├── logger/
│   │   └── logger.go                # Zap + Lumberjack 日志引擎
│   ├── manager/
│   │   ├── window_manager.go        # 多窗口管理器 + 系统托盘
│   │   └── tray_icon.ico            # 嵌入式托盘图标
│   ├── repository/
│   │   ├── user.go                  # InMemory 实现（已弃用，保留参考）
│   │   ├── sqlite_user.go           # SQLite User 仓储实现
│   │   └── sqlite_setting.go        # SQLite Setting 仓储实现
│   └── service/
│       ├── user.go                  # 用户业务逻辑
│       └── setting.go               # 设置业务逻辑
│
├── build/                           # Wails 构建配置 & 平台资源
│
└── frontend/                        # 【Vue 3 前端工程】
    ├── bindings/                    # Wails 自动生成的 TS 绑定（勿手动编辑）
    ├── src/
    │   ├── main.ts                  # 前端入口（Pinia/Router/TDesign 初始化）
    │   ├── App.vue                  # 根组件（仅 <router-view />）
    │   ├── api/
    │   │   └── user.ts              # API 防腐层（二次封装 Wails 绑定）
    │   ├── assets/
    │   │   └── main.css             # Tailwind 入口
    │   ├── composables/             # （预留）Vue 组合式 API 工具库
    │   ├── layouts/
    │   │   └── DefaultLayout.vue    # 主布局（侧边栏 + 标题栏 + 内容区）
    │   ├── router/
    │   │   └── index.ts             # 路由配置（含独立窗口路由）
    │   ├── stores/
    │   │   └── settings.ts          # Pinia 设置 Store（自动同步至 SQLite）
    │   └── views/
    │       ├── HomeView.vue         # 首页（用户注册/查询演示）
    │       ├── SettingsView.vue     # 设置页面
    │       └── AboutView.vue        # 关于页面
    ├── vite.config.ts               # Vite 配置
    └── package.json                 # 前端依赖
```

---

## 三、后端架构设计

### 3.1 分层架构（三层 + 控制器）

请求流向：**前端 JS → Binding → Service → Repository → SQLite**

```
┌──────────────┐
│   Frontend   │  Vue 3 / TypeScript
│  (Webview)   │
└──────┬───────┘
       │  Wails JS SDK 自动调用
┌──────▼───────┐
│   Binding    │  internal/binding/    接收参数、参数校验、防呆拦截
└──────┬───────┘                       （= Controller 控制器层）
       │
┌──────▼───────┐
│   Service    │  internal/service/    纯业务逻辑，不依赖 Wails SDK
└──────┬───────┘                       可单独进行单元测试
       │
┌──────▼───────┐
│  Repository  │  internal/repository/ GORM 数据库读写操作
└──────┬───────┘
       │
┌──────▼───────┐
│    Domain    │  internal/domain/     实体模型 + 抽象接口
└──────────────┘                       零依赖，全项目的核心
```

**核心优势**：更换底层数据库（如从 SQLite 迁移到 MySQL）时，只需新增一个 Repository 实现，上层 Service/Binding 代码 **0 修改**。

### 3.2 依赖注入流水线 (`main.go`)

```go
// main.go 启动顺序：
config.InitConfig()           // 0. 读取 YAML 配置
logger.InitLogger()           // 1. 初始化日志
database.InitDB()             // 2. 连接 SQLite
  → repository.New...()       // 3. 创建数据仓储
    → service.New...()        // 4. 创建业务服务
      → binding.New...()      // 5. 创建前端桥接
application.New(bindins...)   // 6. 组装 Wails 应用
manager.NewWindowManager()    // 7. 创建窗口 & 托盘
wailsApp.Run()                // 8. 阻塞运行
```

---

## 四、已实现的核心功能

### 4.1 SQLite + GORM 本地持久化

- **位置**：`internal/database/database.go`
- 数据库文件存放在 OS 安全目录 (`AppData/MyApp2/data/app_data.db`)
- 启动时自动执行 `AutoMigrate`，结构体新增字段时自动同步表结构
- 已注册的领域模型：`User`、`Setting`

### 4.2 结构化日志引擎 (Zap + Lumberjack)

- **位置**：`internal/logger/logger.go`
- 日志文件：`AppData/MyApp2/logs/app.log`
- **开发模式**：控制台彩色输出 + 文件 JSON 双写
- **生产模式**：仅写 JSON 文件，级别 >= Info
- 单文件 50MB 自动切割，最多 10 个备份，保留 30 天，gzip 压缩

### 4.3 全局 Panic 兜底恢复

- **位置**：`main.go` 顶部 `defer func()`
- 捕获所有未处理的 panic，写入日志
- 弹出 Wails 原生 Error Dialog 通知用户，而非静默闪退
- 引导用户将日志目录发送给开发团队

### 4.4 编译期环境隔离 (buildinfo)

- **位置**：`internal/buildinfo/buildinfo.go`
- `IsDev` 和 `Version` 通过编译参数 `-ldflags` 注入，用户**无法**修改
- 开发时默认 `IsDev=true`，正式打包时注入 `false`
- 构建命令示例：
  ```bash
  go build -ldflags "-X myapp2/internal/buildinfo.Version=1.0.0 -X myapp2/internal/buildinfo.IsDev=false"
  ```

### 4.5 YAML 配置中心 (Viper)

- **位置**：`internal/config/config.go`
- 首次启动自动生成带注释的 `AppData/MyApp2/config.yaml`
- 支持环境变量前缀 `MYAPP2_` 覆盖任意配置项
- 配置项包括：应用名称、数据库驱动、日志策略、窗口尺寸等
- **安全敏感项（is_dev/version）不在此文件中**

### 4.6 优雅停机 (Graceful Shutdown)

- **位置**：`internal/app/app.go` + `main.go` 中的 `defer`
- 程序退出前自动调用 `coreApp.Shutdown()`
- 可在此处安全关闭数据库连接池、释放文件锁、终止后台协程

### 4.7 多窗口管理器 (WindowManager)

- **位置**：`internal/manager/window_manager.go`
- `CreateMainWindow()` — 主窗口（尺寸从配置文件读取）
- `CreateSettingsWindow()` — 独立设置窗口（防重复创建）
- `CreateAboutWindow()` — 独立关于窗口
- 窗口参数（宽高、标题）从 YAML 配置中心注入

### 4.8 系统托盘 (System Tray)

- **位置**：`internal/manager/window_manager.go` 下半部分
- 图标通过 `//go:embed` 嵌入编译产物
- 右键菜单：显示主窗口 / 打开设置 / 关于 / 退出
- 单击托盘图标：唤出主窗口

### 4.9 Pinia 状态持久化同步

- **前端**：`frontend/src/stores/settings.ts`
- **后端**：`binding/setting.go` → `service/setting.go` → `repository/sqlite_setting.go`
- 应用启动时从 SQLite 加载所有设置项到 Pinia Store
- 用户修改设置（如主题切换）时，自动异步同步至 Go 后端 SQLite
- 已实现的持久化设置项：
  - `theme` (light / dark)
  - `language` (zh-CN / en-US)
  - `isSidebarCollapsed` (true / false)

### 4.10 侧边栏导航布局

- **位置**：`frontend/src/layouts/DefaultLayout.vue`
- TDesign `<t-menu>` 驱动的可折叠侧边栏
- 导航页面：首页 / 系统设置 / 关于
- 折叠状态持久化（关联 Pinia → SQLite）
- 顶部标题栏动态显示当前页面名称
- 页面切换带 Fade 渐变动画
- 窗口拖拽区域与可交互区域精确分离

### 4.11 深色模式 / 明亮模式切换

- 通过 Pinia `settings.theme` 控制
- 修改 `<html>` 的 `theme-mode` 属性，TDesign 自动切换全局配色
- 同时添加 Tailwind `dark` 类名
- 切换后自动持久化到 SQLite，重启程序自动恢复

### 4.12 前端 API 防腐层

- **位置**：`frontend/src/api/user.ts`
- 二次封装 Wails 自动生成的 Binding 接口
- 统一错误捕获与日志记录
- 便于接入 Mock 数据进行独立前端测试

### 4.13 多窗口路由架构

- **位置**：`frontend/src/router/index.ts`
- 使用 Hash 模式 (`createWebHashHistory`)，兼容 Wails Webview
- **主窗口路由**（含侧边栏）：`/`、`/settings`、`/about`
- **独立窗口路由**（无侧边栏）：`/standalone/settings`、`/standalone/about`
- Go 端 WindowManager 通过 URL 参数选择路由模式

---

## 五、安全隔离策略

| 数据类型 | 存储位置 | 用户可改 | 示例 |
|---|---|---|---|
| 环境标识 / 版本号 | 🔒 编译期 ldflags | ❌ | `is_dev`, `version` |
| 运维部署参数 | 📄 config.yaml | ✅ | 窗口大小、日志级别 |
| 用户偏好 | 🗂️ SQLite (Pinia) | ✅ | 主题、语言、布局 |
| 业务数据 | 🗂️ SQLite (GORM) | ✅ | 用户表 |

---

## 六、开发命令参考

```bash
# 完整开发模式（前后端热重载）
wails3 dev
# 或
task dev

# 仅前端开发
cd frontend && pnpm dev

# 生产构建
wails3 build

# 生产构建 + 环境注入
go build -ldflags "-X myapp2/internal/buildinfo.Version=1.0.0 -X myapp2/internal/buildinfo.IsDev=false"

# 重新生成前端绑定
wails3 generate bindings --ts

# 前端 lint & 格式化
cd frontend && pnpm lint && pnpm format
```

---

## 七、扩展指南

### 新增一个后端服务

1. 在 `internal/domain/` 定义领域模型和 Repository 接口
2. 在 `internal/repository/` 实现 GORM 版 Repository
3. 在 `internal/service/` 编写业务逻辑（注入 Repository）
4. 在 `internal/binding/` 创建前端桥接接口（注入 Service）
5. 在 `main.go` 中组装依赖链并注册到 `application.NewService()`
6. 运行 `wails3 generate bindings --ts` 自动生成前端 API
7. 在 `frontend/src/api/` 中二次封装

### 新增一个前端页面

1. 在 `frontend/src/views/` 创建 `XxxView.vue`
2. 在 `frontend/src/router/index.ts` 添加路由（放在 DefaultLayout children 下）
3. 如需独立窗口模式，额外在 `/standalone/xxx` 下注册同一组件
4. 在 Go 端 `WindowManager` 中添加 `CreateXxxWindow()` 方法

---

*本文档自动生成于 2026-03-23，与项目代码同步维护。*
