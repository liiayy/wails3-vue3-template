export default {
  // ========== Common ==========
  common: {
    confirm: 'Confirm',
    cancel: 'Cancel',
    save: 'Save',
    create: 'Create',
    edit: 'Edit',
    delete: 'Delete',
    search: 'Search',
    reset: 'Reset',
    loading: 'Loading...',
    success: 'Success',
    failed: 'Failed',
    required: 'This field is required',
  },

  // ========== Sidebar Menu ==========
  menu: {
    home: 'Home',
    users: 'User Manage',
    settings: 'Settings',
    about: 'About',
  },

  // ========== Home ==========
  home: {
    title: 'Wails 3 Mega-Structure Demo',
    subtitle: 'Full-stack Architecture: App → Binding → Service → Repository',
    ready: 'Ready to interact with Go Backend 🚀',
    registerSuccess: 'Registered',
    found: 'Found',
    registerSection: '1. Register (Write)',
    lookupSection: '2. Lookup (Read)',
    namePlaceholder: 'Name',
    emailPlaceholder: 'Email',
    registerBtn: 'Register User',
    fetchBtn: 'Fetch User',
    eventBusTitle: 'Wails Event Bus (Background Goroutine):',
    eventBusWaiting: 'Waiting for time event...',
    persistTheme: 'Persist Theme:',
    themeLight: 'Light',
    themeDark: 'Dark',
  },

  // ========== User Management ==========
  users: {
    title: 'User Management',
    searchPlaceholder: 'Search by name or email...',
    addUser: 'Add User',
    editUser: 'Edit User',
    colId: 'ID',
    colName: 'Name',
    colEmail: 'Email',
    colAction: 'Actions',
    deleteConfirmTitle: 'Confirm Delete',
    deleteConfirmBody:
      'Are you sure you want to delete user "{name}"? This action cannot be undone.',
    createSuccess: 'Created successfully',
    updateSuccess: 'Updated successfully',
    deleteSuccess: 'Deleted successfully',
    loadFailed: 'Failed to load user list',
    deleteFailed: 'Delete failed',
    nameRequired: 'Name is required',
    emailRequired: 'Email is required',
    namePlaceholder: 'Enter name',
    emailPlaceholder: 'Enter email address',
  },

  // ========== Settings ==========
  settings: {
    title: 'Settings',
    themeLabel: 'Theme Mode',
    themeLight: 'Light',
    themeDark: 'Dark',
    languageLabel: 'Language',
    sidebarLabel: 'Collapse Sidebar by Default',
  },

  // ========== About ==========
  about: {
    title: 'About MyApp2',
    version: 'Version',
    description: 'A large-scale desktop application template built with Wails 3 + Vue 3 + Go',
    copyright: '© 2026 MyApp2 Studio. All rights reserved.',
  },

  // ========== Titlebar ==========
  titlebar: {
    toggleDark: 'Switch to dark mode',
    toggleLight: 'Switch to light mode',
  },
}
