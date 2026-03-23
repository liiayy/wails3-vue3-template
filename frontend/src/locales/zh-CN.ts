export default {
  // ========== 通用 ==========
  common: {
    confirm: '确认',
    cancel: '取消',
    save: '保存',
    create: '创建',
    edit: '编辑',
    delete: '删除',
    search: '搜索',
    reset: '重置',
    loading: '加载中...',
    success: '操作成功',
    failed: '操作失败',
    required: '此项为必填',
  },

  // ========== 侧边栏菜单 ==========
  menu: {
    home: '首页',
    users: '用户管理',
    settings: '系统设置',
    about: '关于',
  },

  // ========== 首页 ==========
  home: {
    title: 'Wails 3 Mega-Structure Demo',
    subtitle: '全链路架构验证: App → Binding → Service → Repository',
    ready: '准备与 Go 后端交互 🚀',
    registerSuccess: '注册成功',
    found: '查到',
    registerSection: '1. 新增 (Write)',
    lookupSection: '2. 查询 (Read)',
    namePlaceholder: '用户名',
    emailPlaceholder: '邮箱',
    registerBtn: '注册用户',
    fetchBtn: '查询用户',
    eventBusTitle: 'Wails 事件总线 (后台 Goroutine):',
    eventBusWaiting: '等待 time 事件...',
    persistTheme: '持久化主题:',
    themeLight: '明亮',
    themeDark: '黑暗',
  },

  // ========== 用户管理 ==========
  users: {
    title: '用户管理',
    searchPlaceholder: '搜索姓名或邮箱...',
    addUser: '新增用户',
    editUser: '编辑用户',
    colId: 'ID',
    colName: '用户名',
    colEmail: '邮箱',
    colAction: '操作',
    deleteConfirmTitle: '确认删除',
    deleteConfirmBody: '确定要删除用户「{name}」吗？此操作不可撤销。',
    createSuccess: '新增成功',
    updateSuccess: '更新成功',
    deleteSuccess: '删除成功',
    loadFailed: '加载用户列表失败',
    deleteFailed: '删除失败',
    nameRequired: '用户名不能为空',
    emailRequired: '邮箱不能为空',
    namePlaceholder: '请输入用户名',
    emailPlaceholder: '请输入邮箱地址',
  },

  // ========== 系统设置 ==========
  settings: {
    title: '系统设置',
    themeLabel: '主题模式',
    themeLight: '明亮',
    themeDark: '黑暗',
    languageLabel: '语言',
    sidebarLabel: '侧边栏默认折叠',
  },

  // ========== 关于 ==========
  about: {
    title: '关于 MyApp2',
    version: '版本',
    description: '基于 Wails 3 + Vue 3 + Go 的大型工程化桌面应用模板',
    copyright: '© 2026 MyApp2 Studio. All rights reserved.',
  },

  // ========== 标题栏 ==========
  titlebar: {
    toggleDark: '切换到黑暗模式',
    toggleLight: '切换到明亮模式',
  },
}
