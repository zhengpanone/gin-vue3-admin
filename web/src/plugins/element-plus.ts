import { App } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

export default {
  install (app:App) {
    app.use(ElementPlus, { size: 'min', zIndex: 2000 })
  }
}
