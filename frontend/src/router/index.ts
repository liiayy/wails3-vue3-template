import { createRouter, createWebHashHistory } from 'vue-router'
import DefaultLayout from '@/layouts/DefaultLayout.vue'
import HomeView from '@/views/HomeView.vue'
import UserManageView from '@/views/UserManageView.vue'
import SettingsView from '@/views/SettingsView.vue'
import AboutView from '@/views/AboutView.vue'

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
          meta: { title: 'menu.home', icon: 'home', showInMenu: true, menuSection: 'top' },
        },
        {
          path: 'users',
          name: 'users',
          component: UserManageView,
          meta: { title: 'menu.users', icon: 'user', showInMenu: true, menuSection: 'top' },
        },
        {
          path: 'demo',
          component: () => import('@/views/demo/DemoView.vue'),
          meta: { title: 'menu.demo', icon: 'play-circle', showInMenu: true, menuSection: 'top' },
          children: [
            {
              path: '',
              redirect: '/demo/file',
            },
            {
              path: 'file',
              name: 'demo-file',
              component: () => import('@/views/demo/FileDemoView.vue'),
              meta: { title: 'demo.fileOps', showInMenu: true },
            },
            {
              path: 'clipboard',
              name: 'demo-clipboard',
              component: () => import('@/views/demo/ClipboardDemoView.vue'),
              meta: { title: 'demo.clipboardOps', showInMenu: true },
            },
          ],
        },
        {
          path: 'settings',
          component: SettingsView,
          meta: {
            title: 'menu.settings',
            icon: 'setting',
            showInMenu: true,
            menuSection: 'bottom',
          },
          children: [
            {
              path: '',
              redirect: '/settings/personalization',
            },
            {
              path: 'personalization',
              name: 'settings',
              component: () => import('@/views/settings/PersonalizationView.vue'),
              meta: { title: 'settings.personalization', showInMenu: true },
            },
            {
              path: 'notifications',
              name: 'settings-notifications',
              component: () => import('@/views/settings/NotificationsView.vue'),
              meta: { title: 'settings.notifications', showInMenu: true },
            },
          ],
        },
        {
          path: 'about',
          name: 'about',
          component: AboutView,
          meta: {
            title: 'menu.about',
            icon: 'info-circle',
            showInMenu: false,
            menuSection: 'bottom',
          },
        },
      ],
    },
    // 独立窗口路由（不含侧边栏，供 WindowManager 使用）
    {
      path: '/standalone/settings',
      component: SettingsView,
      children: [
        {
          path: '',
          redirect: '/standalone/settings/personalization',
        },
        {
          path: 'personalization',
          name: 'standalone-settings',
          component: () => import('@/views/settings/PersonalizationView.vue'),
        },
        {
          path: 'notifications',
          name: 'standalone-settings-notifications',
          component: () => import('@/views/settings/NotificationsView.vue'),
        },
      ],
    },
    {
      path: '/standalone/about',
      name: 'standalone-about',
      component: AboutView,
    },
  ],
})

export default router
