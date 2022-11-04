import {createApp} from 'vue'
import App from './App.vue'
import naive from 'naive-ui'
const vueApp = createApp(App)
vueApp.use(naive)
vueApp.mount('#app')
