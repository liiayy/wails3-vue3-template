# MyApp2 — Wails 3 大型工程化最佳实践

> 基于 **Wails v3 (Alpha)** + **Go 1.25** + **Vue 3** + **TypeScript** + **Vite** + **TDesign** + **Tailwind CSS 4** 的跨平台桌面应用模版工程。

---

## 一、技术栈一览

| 层面 | 技术                          | 用途 |
|---|-----------------------------|---|
| 后端运行时 | Go 1.25+                    | 核心业务逻辑与系统调用 |
| 桌面框架 | Wails v3 (alpha.94)         | 将 Go 后端与 Webview 前端合二为一 |
| 前端框架 | Vue 3 (Composition API)     | 界面构建 |
| 国际化语言 | Vue I18n 11                 | 全局多语言方案 |
| 构建工具 | Vite 8                      | 前端热重载与打包 |
| UI 组件库 | TDesign Vue Next            | 企业级 UI 组件 |
| CSS 引擎 | Tailwind CSS 4              | 原子化样式 |
| 状态管理 | Pinia 3                     | 响应式状态与持久化 |
| 路由 | Vue Router 5 (Hash 模式)      | 多窗口路由支持 |
| 数据库 | SQLite (via GORM)           | 本地数据持久化 |
| 日志引擎 | Zap + Lumberjack            | 结构化日志与滚动切割 |
| 配置中心 | Viper                       | YAML 配置文件读写 |
| 任务自动化 | Taskfile v3                 | 跨平台构建任务编排 |
| 包管理 | pnpm (前端) / Go Modules (后端) | 依赖管理 |

---

## 二、项目目录结构

```text
myapp2/
├── main.go                          # 程序入口（依赖注入 & 框架启动）
├── go.mod / go.sum                  # Go 依赖声明
├── Taskfile.yml                     # 构建任务 (含 build:prod 流水线)
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
    │   ├── main.ts                  # 前端入口（Pinia/Router/TDesign/i18n 初始化）
    │   ├── App.vue                  # 根组件（含 t-config-provider 国际化容器）
    │   ├── api/
    │   │   └── user.ts              # API 防腐层（二次封装 Wails 绑定）
    │   ├── assets/
    │   │   └── main.css             # Tailwind 入口
    │   ├── composables/             # 通用 Vue 组合式 API 工具箱
    │   │   ├── useWailsEvent.ts     # Wails 事件安全订阅
    │   │   ├── useWindowControl.ts  # 无状态窗口控制 API
    │   │   ├── useAsyncAction.ts    # 异步 Loading/Error 封装
    │   │   └── useDebounce.ts       # 防抖响应式代理
    │   ├── layouts/
    │   │   └── DefaultLayout.vue    # 主布局（侧边栏 + 拖拽栏 + 内容区）
    │   ├── locales/
    │   │   ├── index.ts             # vue-i18n 实例
    │   │   ├── zh-CN.ts             # 中文语言包
    │   │   └── en-US.ts             # 英文语言包
    │   ├── router/
    │   │   └── index.ts             # 路由配置（含独立窗口路由）
    │   ├── stores/
    │   │   └── settings.ts          # Pinia 设置 Store（自动同步至 SQLite）
    │   └── views/
    │       ├── HomeView.vue         # 首页视图
    │       ├── UserManageView.vue   # 用户管理（完整 CRUD 范例）
    │       ├── SettingsView.vue     # 设置页面布局容器
    │       ├── settings/            # 【设置子页面】
    │       │   ├── PersonalizationView.vue  # 个性化
    │       │   └── NotificationsView.vue    # 通知设置
    │       ├── DemoView.vue         # 演示页面布局容器
    │       ├── demo/                # 【功能演示子页面】
    │       │   ├── FileDemoView.vue     # 原生文件对话框演示
    │       │   ├── ClipboardDemoView.vue # 原生剪贴板操作演示
    │       │   └── DragDropDemoView.vue  # 原生拖拽演示
    │       └── AboutView.vue        # 关于页面
    ├── vite.config.ts               # Vite 配置（含生产包 DevTools 隔离）
    └── package.json                 # 前端依赖
```

---

## 三、后端架构设计

### 3.1 分层架构（三层 + 控制器）

请求流向：**前端 JS → Binding → Service → Repository → SQLite**

```text
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

**核心优势**：更换底层数据库时，只需新增一个 Repository 实现，上层代码 **0 修改**。

### 3.2 依赖注入流水线 (`main.go`)

```go
// main.go 启动顺序：
config.InitConfig()           // 0. 读取 YAML 配置
logger.InitLogger()           // 1. 初始化日志
database.InitDB()             // 2. 连接 SQLite
  → repository.New...()       // 3. 创建数据仓储
    → service.New...()        // 4. 创建业务服务
      → binding.New...()      // 5. 创建前端桥接
application.New(bindings...)  // 6. 组装 Wails 应用
manager.NewWindowManager()    // 7. 创建窗口 & 托盘
wailsApp.Run()                // 8. 阻塞运行
```

---

## 四、核心特性沉淀

### 1. 生产级构建流水线 (Taskfile)
- 提供 `task build:prod` 和 `task package:prod` 指令。
- **自动安全增强**：在 `vite.config.ts` 判断生产模式，直接剥离 Vue DevTools，防止内部状态泄露。
- **编译时注入**：使用 `ldflags` 注入版本号和 `IsDev=false`，用户无法手动篡改系统环境标记。
- **二进制优化**：启用 Go `-trimpath` 删除绝对路径，启用 `-s -w` 压缩符号表（减小约30%体积），启用 `-H windowsgui` 隐藏控制台黑窗。

### 2. 多维度主题与国际化闭合
- **跟随系统 (Follow System)**：支持 `light` | `dark` | `auto` 三种模式。切换至 `auto` 时，应用会通过媒体查询 `prefers-color-scheme` 实时同步操作系统的深浅色设置。
- **双语国际化 (i18n)**：结合 `vue-i18n` 和 TDesign `<t-config-provider>` 实现了全局组件和私有业务文案的无缝中英切换。
- **状态持久**：使用 `settings` 响应式 store，当变更语言或主题时，会经由 Wails Call 持久化入 GORM (SQLite)。

### 3. Vue Composables 业务工具箱
专门消灭常见的前后端桥接样板代码：
- `useAsyncAction(fn)`：无需手写 try/catch，自动维护 Loading 状态、捕获 Error String、并派发强类型 Data 结果。
- `useWailsEvent(name)`：避免内存泄漏！在 `onMounted` 注册 Wails EventBus 监听，在 `onUnmounted` 自动执行 Off 取消订阅。
- `useWindowControl()`：将 Wails API（最小化、全屏、拉伸等）无状态映射为简单的 Hooks 函数供模板调用。
- `useDebounce(ref)`：防抖处理用户的即时输入（常用于搜素联想）。

### 4. 完整前后端 CRUD 范式实现
- 在 `UserManageView.vue` 落地了包含：*条件检索防抖*、*服务端分页*、*TDesign 交互表格*、*新增/编辑聚合弹窗*、*二次确认删除拦截*的整套生命周期。
- 后端完成了 GORM Offset/Limit 配合 Like 的分页聚合封装，返回标准结构 `UserListResult{Items, Total}`。

### 5. SQLite + GORM 本地持久化
- 数据库落盘 `AppData/MyApp2/data/app_data.db`。采用 `AutoMigrate` 保障迭代时不卡阻表结构变化。内置记录了包含用户管理及设置同步在内的强约束关系。

### 6. 结构化日志引擎 (Zap + Lumberjack)
- 日志文件落盘 `AppData/MyApp2/logs/app.log`。单文件 50MB 自动切割，最多备份 10 个，保留 30 天，gzip 压缩。生产仅写入特定等级，兼顾 I/O 性能。

### 7. 全局 Panic 兜底恢复
- `main.go` 引入 `defer recover`，发生致命报错时不造成静默崩溃。而是自动唤起 Wails Native Dialog 指引用户。

### 8. YAML 离线配置中心 (Viper)
- 将运维型、环境主导型参数抽离成独立的 `config.yaml`。启动检查有无此项，若无则自动注入默认结构（极大增强桌面白盒分发体验）。

### 9. 优雅停机 (Graceful Shutdown)
- 关联 Wails Shutdown 生命周期。执行包括释放本地文件锁、停止 Logger 写盘或断开 SQLite 连接等安全断线工作。

### 10. 多窗口管理器与系统托盘
- 在 Go 侧维持多独立形态的结构，包含：主控制窗、独立设置窗等。包含针对 OS 的托盘注册和右键触发菜单联调。

### 11. Pinia 状态持久化同步
- **状态单源**：所有的前端状态变更是发起点，Pinia 保存实时缓存。
- 当涉及到外观、主题、栏目折叠情况时，触发 `SettingBinding.Save()` 真正写入硬盘。下一次重启提取并还原现场。

### 12. 路由多开、状态保持与动态菜单
- **状态保持 (KeepAlive)**：主视图及设置子视图均启用了 `<keep-alive>`。切换路由时（如从用户管理切走再切回），页面状态（搜索词、滚动条、甚至未提交的表单项）将被完整保留。
- **动态菜单配置**：侧边栏菜单不再硬编码，而是通过扫描 `router/index.ts` 中的 `meta` 配置动态生成：
  - `showInMenu`: 控制是否在菜单中显示。
  - `menuSection`: 控制显示在侧边栏的「顶部主菜单」还是「底部功能菜单」。
  - `icon`: 关联图标映射表。
- **路由多开防腐**：基于 Hash `/` 规避多窗口由于 History 引引发的 HTTP Fallback 故障。并严格包装了 `api/` 目录将直接 Wails JS Binding 调用转化为带容错语义的前端 Async 函数。

### 13. 窗口状态同步控制
- **状态感知图标**：窗口控制栏的最大化按钮具备状态感知能力。当窗口已最大化时，图标自动切换为“还原（双层方框）”，反之显示“最大化（单层方框）”。
- **Wails 事件联动**：通过监听原生 `maximize` / `unmaximize` 事件，确保 UI 状态在用户通过系统边框或标题栏双击操作后也能实时对齐。

### 14. 原生能力集成演示 (Demo)
- **零胶水文件对话框**：直接调起操作系统的原生多选文件、单选文件、文件夹选择以及保存文件对话框。
- **剪贴板读写**：演示通过 `Clipboard` 接口实现与系统剪贴板的同步，支持文本写入、读取与强制清空。
- **低功耗原生拖拽**：在 Web 元素上应用 `data-file-drop-target="true"` 属性，即可让该区域具备接收操作系统文件投放的能力。通过监听 `common:WindowFilesDropped` 事件，前端可秒级获取文件物理路径，且不消耗额外的 JS 线程资源。
- **系统级原生通知**：集成 Wails 3 `notifications` 服务。支持基础桌面弹窗、副标题展示，以及超互动的“操作按钮”与“文本回复”功能。用户在通知中心的点击与回复数据，经由 Go 后端捕获后实时透传至前端 EventBus。

---

## 五、安全隔离策略

| 数据类型 | 存储位置 | 用户可改 | 示例 |
|---|---|---|---|
| 环境标识 / 版本号 | 🔒 编译期 ldflags | ❌ | `IsDev`, `Version` |
| 运维部署参数 | 📄 config.yaml | ✅ | 窗口大小、日志级别 |
| 用户偏好 | 🗂️ SQLite (Pinia) | ✅ | 主题、语言、布局 |
| 业务数据 | 🗂️ SQLite (GORM) | ✅ | 用户表 |

---

## 六、开发命令参考

```bash
# ============ 开发 ============ 
# 完整开发模式（前后端热重载）
wails3 dev
# 或
task dev

# ============ 构建 ============ 
# 一键生产构建 (注入版本、关闭控制台黑窗、极致体积压缩、屏蔽 DevTools)
wails3 task build:prod
# 指定版本构建
wails3 task build:prod APP_VERSION=2.1.0

# 打包为 Windows NSIS 安装程序
wails3 task package:prod

# ============ 辅助 ============ 
# 重新生成前端 TS 绑定 (强制清理旧版)
wails3 generate bindings --ts -clean=true

# 前端 lint & 格式化代码
cd frontend && pnpm lint && pnpm format
```

---

## 七、扩展指南

### 新增一个后端服务
1. 在 `internal/domain/` 定义领域模型和 Repo 接口。
2. 在 `internal/repository/` 撰写 GORM Repo 实现。
3. 在 `internal/service/` 编写纯业务逻辑。
4. 在 `internal/binding/` 封装可前端调用的结构。
5. 在 `main.go` 注册。
6. 行使 `wails3 generate bindings --ts -clean=true`。
7. 在 `frontend/src/api/` 提供前端函数封装。

### 新增一个前端页面
1. 在 `frontend/src/views/` 创立 `*.vue`。如果是设置子页，置于 `views/settings/`。
2. 添加进入 `frontend/src/router/index.ts`（置于 `DefaultLayout` 之内）。
3. **配置菜单元数据**：在路由的 `meta` 字段中设置：
   - `showInMenu: true`: 允许在侧边栏显示。
   - `menuSection: 'top' | 'bottom'`: 指定显示位置。
   - `title`: 指定 i18n key（如 `menu.new_page`）。
   - `icon`: 指定图标标识符（需在 `DefaultLayout.vue` 的 `menuIconMap` 中存在）。
4. 调整 `locales/` 确保多语言文案对应即可展现。

---

*本文档基于最新项目结构自动产生于当前节点，并与所有实施功能保持强对应锚点关系。*
