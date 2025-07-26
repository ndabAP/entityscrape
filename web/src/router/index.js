import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/isob',
      name: 'home',
     component: () => import('../views/Isob.vue'),
    },
  ],
})

export default router
