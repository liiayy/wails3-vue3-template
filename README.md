# wails3-vue3-template - Wails 3 Enterprise Desktop Application Template

<div align="center">

**[English](README.md) | [简体中文](README_zh.md)**

**A production-grade, enterprise-ready desktop application template built with Wails v3**

[![Wails Version](https://img.shields.io/badge/Wails-v3.0.0--alpha.94-blue)](https://wails.io)
[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go)](https://go.dev)
[![Vue Version](https://img.shields.io/badge/Vue-3.5+-4FC08D?logo=vue.js)](https://vuejs.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

[Features](#features) • [Screenshots](#screenshots) • [Quick Start](#quick-start) • [Architecture](#architecture) • [Documentation](#documentation)

</div>

---

## Introduction

wails3-vue3-template is a comprehensive desktop application template that showcases **Wails v3 best practices** for building cross-platform desktop applications. It combines Go's powerful backend capabilities with Vue 3's modern frontend ecosystem, providing a solid foundation for enterprise-grade applications.

### Why This Project?

- 🚀 **Production Ready**: Includes logging, configuration, database migration, error handling, and graceful shutdown
- 🏗️ **Clean Architecture**: Layered design (Binding → Service → Repository → Domain) for easy maintenance
- 🎨 **Modern UI**: Built with Vue 3, TDesign component library, and Tailwind CSS 4
- 🌍 **Internationalization**: Full Chinese and English multilingual support
- 🌓 **Theme System**: Light/Dark/Auto themes synced with system preferences
- 💾 **Persistence**: SQLite + GORM with auto migration
- 📦 **Standalone Distribution**: Single binary with all resources embedded

---

## Screenshots

<table>
  <tr>
    <td align="center"><b>Home Page</b></td>
    <td align="center"><b>Theme Mode</b></td>
  </tr>
  <tr>
    <td><img src="docs/images/首页.png" alt="Home Page" width="450"/></td>
    <td><img src="docs/images/主题模式.png" alt="Theme Mode" width="450"/></td>
  </tr>
  <tr>
    <td align="center"><b>Configuration Center</b></td>
    <td align="center"><b>File Operations</b></td>
  </tr>
  <tr>
    <td><img src="docs/images/配置中心.png" alt="Configuration Center" width="450"/></td>
    <td><img src="docs/images/文件操作.png" alt="File Operations" width="450"/></td>
  </tr>
  <tr>
    <td align="center"><b>Clipboard Operations</b></td>
    <td align="center"><b>Native Drag & Drop</b></td>
  </tr>
  <tr>
    <td><img src="docs/images/剪切板操作.png" alt="Clipboard Operations" width="450"/></td>
    <td><img src="docs/images/原生拖拽.png" alt="Native Drag & Drop" width="450"/></td>
  </tr>
  <tr>
    <td align="center"><b>Native Notifications</b></td>
    <td align="center"><b>Multi-window State Sharing</b></td>
  </tr>
  <tr>
    <td><img src="docs/images/原生通知.png" alt="Native Notifications" width="450"/></td>
    <td><img src="docs/images/多窗口状态共享.gif" alt="Multi-window State Sharing" width="450"/></td>
  </tr>
</table>

---

## Features

### Core Capabilities

- **Multi-window Management**: Create and manage multiple independent windows with system tray integration
- **User Management**: Complete CRUD example with pagination, search, and form validation
- **Settings Persistence**: Real-time settings sync with automatic persistence to database
- **Native Integration**:
  - File dialogs (open/save/multi-select)
  - Clipboard operations
  - Native file drag & drop
  - System notifications with actions and replies
- **State Management**: Pinia state stores with route state preservation
- **Dynamic Routing**: Auto-generated sidebar menu from route metadata
- **Developer Tools**: Hot reload, Vue DevTools, TypeScript strict mode

### Developer Experience

- **Type Safety**: Full TypeScript coverage with auto-generated bindings
- **Code Quality**: ESLint, Oxlint, Prettier, and auto-formatting
- **Build Pipeline**: Automated production build optimization with Taskfile
- **Logging**: Structured logging with rotation support (Lumberjack)
- **Configuration**: Viper-based YAML configuration
- **Error Handling**: Unified error responses and panic recovery

---

## Tech Stack

| Layer | Technology | Purpose |
|------|-----------|---------|
| **Backend Runtime** | Go 1.25+ | Core business logic & system calls |
| **Desktop Framework** | Wails v3 (alpha.94) | Bridge Go backend with WebView frontend |
| **Frontend Framework** | Vue 3 (Composition API) | Reactive UI components |
| **Internationalization** | Vue I18n 11 | Multi-language support |
| **Build Tool** | Vite 8 | Fast HMR and optimized builds |
| **UI Component Library** | TDesign Vue Next | Enterprise component library |
| **CSS Engine** | Tailwind CSS 4 | Utility-first styling |
| **State Management** | Pinia 3 | Reactive state & persistence |
| **Router** | Vue Router 5 (Hash Mode) | Multi-window routing support |
| **Database** | SQLite (via GORM) | Local data persistence |
| **Logging** | Zap + Lumberjack | Structured logging & rotation |
| **Configuration** | Viper | YAML configuration management |
| **Task Automation** | Taskfile v3 | Cross-platform build tasks |

---

## Quick Start

### Prerequisites

- **Go**: 1.25 or higher
- **Node.js**: ^20.19.0 or >=22.12.0
- **pnpm**: Latest (recommended) or npm/yarn

### Installation

```bash
# Clone the repository
git clone https://github.com/liiayy/wails3-vue3-template.git
cd wails3-vue3-template

# Install frontend dependencies
cd frontend
pnpm install
cd ..

# Run development mode
wails3 dev
# Or
task dev
```

That's it! The application will open automatically, and both frontend and backend code changes will hot-reload.

### Production Build

```bash
# Standard production build
wails3 build

# Optimized build (inject version number)
task build:prod

# Build with specific version
task build:prod APP_VERSION=2.1.0
```

Output files are in the `bin/` directory.

---

## Project Structure

```
myapp2/
├── main.go                          # Application entry point (DI & startup)
├── go.mod / go.sum                  # Go dependencies
├── Taskfile.yml                     # Build task automation
│
├── internal/                        # Go private core packages
│   ├── app/                         # Wails lifecycle management
│   ├── binding/                     # Frontend bridge layer (controllers)
│   ├── buildinfo/                   # Compile-time version injection
│   ├── config/                      # Viper YAML configuration
│   ├── database/                    # GORM + SQLite connection
│   ├── domain/                      # Domain models + Repository interfaces
│   ├── logger/                      # Zap + Lumberjack logging engine
│   ├── manager/                     # Multi-window manager + system tray
│   ├── repository/                  # SQLite repository implementation
│   └── service/                     # Business logic layer
│
├── build/                           # Wails build configuration & assets
│
└── frontend/                        # Vue 3 frontend
    ├── bindings/                    # Auto-generated TS bindings (do not edit)
    ├── src/
    │   ├── main.ts                  # Frontend entry point
    │   ├── App.vue                  # Root component
    │   ├── api/                     # API abstraction layer
    │   ├── assets/                  # Static assets
    │   ├── composables/             # Vue composables toolkit
    │   ├── layouts/                 # Layout components
    │   ├── locales/                 # i18n language files
    │   ├── router/                  # Vue Router configuration
    │   ├── stores/                  # Pinia state stores
    │   └── views/                   # Page components
    ├── vite.config.ts               # Vite configuration
    └── package.json                 # Frontend dependencies
```

---

## Architecture

### Layered Architecture

```
┌──────────────┐
│   Frontend   │  Vue 3 / TypeScript
│   (WebView)  │
└──────┬───────┘
       │  Wails JS SDK
┌──────▼───────┐
│   Binding    │  Parameter validation, error handling
└──────┬───────┘  (Controller layer)
       │
┌──────▼───────┐
│   Service    │  Pure business logic, testable
└──────┬───────┘
       │
┌──────▼───────┐
│  Repository  │  GORM database operations
└──────┬───────┘
       │
┌──────▼───────┐
│    Domain    │  Entity models + interfaces
└──────────────┘  Zero dependencies, application core
```

**Core Advantage**: Simply add a new Repository implementation to switch databases — upper layers **require no changes**.

### Dependency Injection Flow

```go
// main.go startup order:
config.InitConfig()           // 0. Load YAML configuration
logger.InitLogger()           // 1. Initialize logger
database.InitDB()             // 2. Connect to SQLite
  → repository.New...()       // 3. Create repositories
    → service.New...()        // 4. Create services
      → binding.New...()      // 5. Create bindings
application.New(bindings...)  // 6. Assemble Wails app
manager.NewWindowManager()    // 7. Create windows & tray
wailsApp.Run()                // 8. Run event loop
```

---

## Documentation

- **[Architecture Guide](ARCHITECTURE.md)** - Detailed architecture documentation
- **[Development Guide](CLAUDE.md)** - Contributor development notes
- **[Wails Documentation](https://wails.io/docs/next/introduction)** - Official Wails v3 documentation

---

## Core Features Deep Dive

### 1. Production Build Pipeline

The `task build:prod` command provides:
- **Security**: Automatically removes Vue DevTools in production
- **Compile-time Injection**: Injects version number and environment via ldflags
- **Binary Optimization**: `-trimpath`, `-s -w` for smaller footprint
- **Hidden Console**: Uses `-H windowsgui` on Windows

### 2. Theme & Internationalization

- **System-aware**: Light/Dark/Auto modes synced with OS settings
- **Bilingual Support**: Chinese/English switching via Vue I18n + TDesign
- **Persistent State**: Settings auto-synced to SQLite (via Pinia)

### 3. Vue Composables Toolkit

Eliminate boilerplate with reusable composables:
- `useAsyncAction(fn)`: Auto loading/error state wrapper
- `useWailsEvent(name)`: Memory-leak-free event subscription
- `useWindowControl()`: Stateless window control API
- `useDebounce(ref)`: Reactive debounce

### 4. Complete CRUD Example

`UserManageView.vue` demonstrates:
- Server-side pagination (Offset/Limit)
- Debounced search (LIKE query)
- TDesign data table with inline actions
- Create/Edit modal dialogs
- Delete confirmation

### 5. SQLite + GORM

Database location: `AppData/MyApp2/data/app_data.db`
- AutoMigrate
- Foreign key constraints
- Transaction support
- Connection pooling

### 6. Structured Logging

Log file: `AppData/MyApp2/logs/app.log`
- 50MB per-file auto rotation
- Max 10 backup files retained
- 30-day retention with gzip compression
- Level-based filtering for performance

### 7. Graceful Shutdown

Cleanup on application exit:
- Release file locks
- Flush log buffers
- Close database connections
- Unregister system tray

### 8. Multi-window & Tray

- Independent window management
- System tray with context menu
- Window state persistence
- Cross-platform tray icons

---

## Development Tasks

### Adding a Backend Service

1. Define domain models in `internal/domain/`
2. Create repository in `internal/repository/`
3. Implement business logic in `internal/service/`
4. Expose API in `internal/binding/`
5. Register the service in `main.go`
6. Run `wails3 generate bindings --ts -clean=true`
7. Add frontend wrapper in `frontend/src/api/`

### Adding a Frontend Page

1. Create `*.vue` in `frontend/src/views/`
2. Add route in `frontend/src/router/index.ts`
3. Configure menu metadata:
   ```typescript
   meta: {
     showInMenu: true,
     menuSection: 'top', // or 'bottom'
     title: 'menu.new_page',
     icon: 'desktop'
   }
   ```
4. Add i18n key-value pairs in `locales/`

### Common Commands

```bash
# Development
wails3 dev              # Full dev mode (hot reload)
task dev                # Same as above

# Build
wails3 build           # Standard build
task build:prod        # Optimized build
task package:prod      # Create installer package

# Frontend only
cd frontend
pnpm dev               # Frontend dev server
pnpm build             # Production build
pnpm lint              # Code linting
pnpm format            # Code formatting

# Code generation
wails3 generate bindings --ts -clean=true  # Regenerate bindings
wails3 generate icons                     # Generate app icons
```

---

## Security Policy

| Data Type | Storage Location | User Modifiable | Example |
|-----------|-----------------|-----------------|---------|
| Environment/Version | 🔒 Compile-time ldflags | ❌ | `IsDev`, `Version` |
| Deployment Config | 📄 config.yaml | ✅ | Window size, log level |
| User Preferences | 🗂️ SQLite (Pinia) | ✅ | Theme, language, layout |
| Business Data | 🗂️ SQLite (GORM) | ✅ | User records |

---

## Contributing

Contributions are welcome! Please follow these guidelines:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Code Standards

- **Go**: Follow standard Go conventions, run `gofmt`
- **TypeScript/Vue**: Use ESLint and Prettier (`pnpm lint && pnpm format`)
- **Commit Messages**: Use conventional commit format

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

- [Wails](https://wails.io) - Excellent desktop application framework
- [Vue.js](https://vuejs.org) - Progressive JavaScript framework
- [TDesign](https://tdesign.tencent.com/) - Enterprise Vue component library
- [Tailwind CSS](https://tailwindcss.com) - Utility-first CSS framework

---

<div align="center">

**Built with Wails v3 ❤️**

[Report Issue](https://gitee.com/liiayy/wails3-vue3-template/issues) · [Feature Request](https://gitee.com/liiayy/wails3-vue3-template/issues)

</div>