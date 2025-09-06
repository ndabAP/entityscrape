import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('../views/Home.vue')
    },
    {
      path: '/',
      name: 'about',
      component: () => import('../views/About.vue')
    },
    {
      path: '/isob',
      name: 'isob',
      component: () => import('../views/Isob.vue')
    },
    {
      path: '/nsops',
      name: 'nsops',
      component: () => import('../views/Nsops.vue')
    },
    {
      path: '/rvomg',
      name: 'rvomg',
      component: () => import('../views/Rvomg.vue')
    }
  ]
})

export default router
