import { createRouter, createWebHashHistory } from 'vue-router'
import DefaultLayout from '../layouts/DefaultLayout.vue'
import HomeView from '../views/HomeView.vue'
import UserManageView from '../views/UserManageView.vue'
import SettingsView from '../views/SettingsView.vue'
import AboutView from '../views/AboutView.vue'

const router = createRouter({
  // 使用 hash 模式，因为独立窗口通过 /#/settings 导航
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    // 主布局（含侧边栏）
    {
      path: '/',
      component: DefaultLayout,
      children: [
        {
          path: '',
          name: 'home',
          component: HomeView,
          meta: { title: '首页', icon: 'home' },
        },
        {
          path: 'users',
          name: 'users',
          component: UserManageView,
          meta: { title: '用户管理', icon: 'user' },
        },
        {
          path: 'settings',
          name: 'settings',
          component: SettingsView,
          meta: { title: '系统设置', icon: 'setting' },
        },
        {
          path: 'about',
          name: 'about',
          component: AboutView,
          meta: { title: '关于', icon: 'info-circle' },
        },
      ],
    },
    // 独立窗口路由（不含侧边栏，供 WindowManager 使用）
    {
      path: '/standalone/settings',
      name: 'standalone-settings',
      component: SettingsView,
    },
    {
      path: '/standalone/about',
      name: 'standalone-about',
      component: AboutView,
    },
  ],
})

export default router
