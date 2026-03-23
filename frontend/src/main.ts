import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import TDesign from 'tdesign-vue-next';
import 'tdesign-vue-next/es/style/index.css';

import App from './App.vue'
import router from './router'
import { useSettingsStore } from './stores/settings'
import i18n from './locales'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(i18n)
app.use(TDesign);

// 初始化加载持久化设置
const settings = useSettingsStore()
settings.init()

app.mount('#app')
