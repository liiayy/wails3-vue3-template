# CLAUDE.md

此文件为 Claude Code (claude.ai/code) 在此仓库中工作时提供指导。

## 项目概述

这是一个 **Wails v3** 应用程序 - 一个跨平台桌面应用框架，结合了：
- **后端**: Go 1.25+ 使用 Wails v3（目前是 alpha 版本）
- **前端**: Vue 3 + TypeScript + Vite，使用 Pinia 进行状态管理，Vue Router 进行路由

## 开发命令

### 主要开发命令
```bash
# 完整开发模式（同时运行前端开发服务器和 Go 后端，支持热重载）
wails3 dev

# 或通过 Task（与 wails3 dev 相同）
task dev
```

### 构建
```bash
# 生产构建
wails3 build

# 或通过 Task（特定平台）
task build

# 开发构建（更快，不压缩）
wails3 build DEV=true

# 打包分发
task package
```

### 仅前端开发
```bash
cd frontend

# 安装依赖（使用 pnpm）
pnpm install

# 前端开发服务器（独立运行）
pnpm dev

# 构建前端（生产模式）
pnpm build

# 构建前端（开发模式）
pnpm build:dev

# 运行单元测试
pnpm test:unit

# 运行 linter（oxlint + eslint）
pnpm lint

# 格式化代码
pnpm format
```

### 服务器模式（无 GUI）
```bash
# 构建服务器模式（仅 HTTP 服务器，无 GUI）
task build:server

# 运行服务器模式
task run:server

# 构建 Docker 镜像
task build:docker

# 在 Docker 中运行
task run:docker
```

### 代码生成
```bash
# 从 Go 服务生成 TypeScript/JavaScript 绑定
wails3 generate bindings

# 从源图片生成应用图标
wails3 generate icons
```

## 架构

### 后端 (Go)
- **入口**: `main.go` - 创建应用、窗口并注册服务
- **服务**: 暴露给前端的 Go 结构体（如 `greetservice.go` 中的 `GreetService`）
- **事件**: 可以在 `init()` 中注册自定义事件，并通过 `app.Event.Emit()` 发出
- **资源**: `frontend/dist/` 中的前端文件使用 `//go:embed` 嵌入到二进制文件中

### 前端 (Vue 3 + TypeScript)
- **入口**: `frontend/src/main.ts` - 创建 Vue 应用，配置 Pinia 和路由
- **路由**: `frontend/src/router/index.ts` - Vue Router 配置
- **状态**: `frontend/src/stores/` - Pinia stores 用于响应式状态管理
- **组件**: `frontend/src/components/` - 可复用的 Vue 组件
- **视图**: `frontend/src/views/` - 路由的页面级组件

### 绑定系统
- **位置**: `frontend/bindings/`
- **重要**: 这些文件是从 Go 服务**自动生成**的 - 请勿编辑
- **模块名**: 基于 Go 模块名（当前为 `myapp2`）
- **用法**: 从 `@wailsio/runtime` 和生成的绑定导入以调用 Go 方法
- **示例**: Go 服务方法 `GreetService.Greet(name string)` 在 TypeScript 中变为 `GreetService.Greet(name)`

### 开发模式配置
- **配置文件**: `build/config.yml`
- **监听文件**: Go 文件、JS/TS 文件（排除 `frontend/` 目录）
- **前端端口**: 默认 9245（可通过 `WAILS_VITE_PORT` 环境变量配置）
- **开发模式命令**: 在 `build/config.yml` 的 `dev_mode.executes` 中定义

## 项目结构

```
├── main.go                 # 应用入口点，创建应用和窗口
├── greetservice.go         # 示例服务（将方法暴露给前端）
├── go.mod / go.sum         # Go 依赖
├── Taskfile.yml            # 根任务文件（包含平台特定任务）
├── build/
│   ├── Taskfile.yml        # 通用构建任务
│   ├── config.yml          # 开发模式和构建配置
│   ├── windows/            # Windows 特定任务
│   ├── darwin/             # macOS 特定任务
│   ├── linux/              # Linux 特定任务
│   └── ...                 # 其他平台特定构建
├── frontend/
│   ├── bindings/           # 自动生成的 Go→TypeScript 绑定（请勿编辑）
│   ├── src/
│   │   ├── main.ts         # 前端入口点
│   │   ├── App.vue         # 根 Vue 组件
│   │   ├── router/         # Vue Router 配置
│   │   ├── stores/         # Pinia 状态存储
│   │   ├── views/          # 页面组件
│   │   └── components/     # 可复用组件
│   ├── index.html          # HTML 模板
│   ├── vite.config.ts      # Vite 配置（含 Wails 插件）
│   └── package.json        # 前端依赖
└── bin/                    # 构建输出目录
```

## 添加新的后端服务

1. 创建新的 Go 结构体并添加导出方法：
```go
// myservice.go
package main

type MyService struct{}

func (m *MyService) DoSomething(input string) string {
    return "处理: " + input
}
```

2. 在 `main.go` 中注册服务：
```go
app := application.New(application.Options{
    Services: []application.Service{
        application.NewService(&GreetService{}),
        application.NewService(&MyService{}),  // 在这里添加
    },
    // ... 其他选项
})
```

3. 运行 `wails3 dev` - 绑定会自动生成并在 `frontend/bindings/` 中可用

4. 在 TypeScript 中使用：
```typescript
import { MyService } from '@/bindings/changeme';

const result = await MyService.DoSomething("hello");
```

## 添加事件

1. 在 `init()` 中注册事件类型：
```go
application.RegisterEvent[string]("myEvent")
```

2. 发出事件：
```go
app.Event.Emit("myEvent", "data")
```

3. 在前端监听：
```typescript
import { EventsOn } from "@wailsio/runtime";

EventsOn("myEvent", (data: string) => {
    console.log("收到:", data);
});
```

## 重要说明

- **Go 模块名**: 目前在 `go.mod` 中设置为 `changeme`。请更新为你的实际模块路径。
- **绑定是只读的**: 永远不要手动编辑 `frontend/bindings/` 中的文件 - 它们在每次构建时都会重新生成。
- **Node 版本**: 前端需要 Node.js ^20.19.0 或 >=22.12.0
- **包管理器**: 前端使用 `pnpm` 进行依赖管理
- **Wails v3**: 目前处于 alpha 阶段 - 在稳定版本发布前 API 可能会发生变化
