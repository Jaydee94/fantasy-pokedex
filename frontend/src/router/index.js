// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router'
import Create from '../views/Create.vue'
import Home from '../views/Home.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/create', component: Create },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
