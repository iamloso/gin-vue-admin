import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [{
  path: '/',
  redirect: '/login'
},
{
  path: '/init',
  name: 'Init',
  component: () => import('@/view/init/index.vue')
},
{
  path: '/login',
  name: 'Login',
  component: () => import('@/view/login/index.vue')
},
{
  path: '/home',
  name: 'Home',
  component: () => import('@/view/home/home.vue')
},
{
  path: '/print',
  name: 'Print',
  component: () => import('@/view/home/printReport.vue')
},
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
