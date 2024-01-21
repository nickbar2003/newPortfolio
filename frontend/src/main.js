import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { createRouter, createWebHistory } from 'vue-router'

import Home from './components/Home.vue'
import Portfolio from './components/Portfolio.vue'
import Timeline from './components/Timeline.vue'


const routes = [
    { path: '/', component: Home },
    { path: '/Portfolio', component: Portfolio },
    { path: '/Timeline', component: Timeline },
]

const router = createRouter({
    history: createWebHistory(),
    routes,

})


const app = createApp(App)
app.use(router)

app.mount('#app')

