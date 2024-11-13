import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from './routers'


const app = createApp(App)
app.use(ElementPlus)
app.use(router)
app.mount('#app')