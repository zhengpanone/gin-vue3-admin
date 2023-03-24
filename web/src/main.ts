import { createApp } from 'vue'
import App from './App.vue'
import router from './routers'
import { createPinia } from 'pinia'
import elementPlus from '@/plugins/element-plus'

import './styles/index.scss'

createApp(App)
  .use(router)
  .use(createPinia())// 创建根存储库
  .use(elementPlus)
  .mount('#app')
