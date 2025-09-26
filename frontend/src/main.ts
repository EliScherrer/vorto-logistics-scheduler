import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import './style.css'
import App from './App.vue'

// Import components
import Home from './components/Home.vue'
import AdminLogin from './components/AdminLogin.vue'
import DriverLogin from './components/DriverLogin.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/admin/login', component: AdminLogin },
  { path: '/driver/login', component: DriverLogin },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

const app = createApp(App)
app.use(router)
app.mount('#app')