import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory('/entityscrape/'),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('../views/HomeView.vue')
    },
    {
      path: '/',
      name: 'about',
      component: () => import('../views/AboutView.vue')
    },
    {
      path: '/isopf',
      name: 'isopf',
      component: () => import('../views/IsopfView.vue')
    },
    {
      path: '/nsops',
      name: 'nsops',
      component: () => import('../views/NsopsView.vue')
    },
    {
      path: '/rvomg',
      name: 'rvomg',
      component: () => import('../views/RvomgView.vue')
    }
  ]
})

export default router
