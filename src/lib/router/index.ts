import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '@/pages/HomePage.vue';

const routes: any = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('@/pages/LoginPage.vue'),
  },
  {
    path: '/upload_file',
    name: 'upload_file',
    component: () => import('@/pages/UpLoadPage.vue'),
  },
  {
    path: '/about',
    name: 'about',
    component: () => import('@/pages/About/IndexPage.vue'),
    children: [
      {
        path: '',
        description: 'list',
        component: () => import('@/pages/About/ListPage.vue'),
      },
      {
        path: 'pwa',
        name: 'pwa',
        description: 'PWA应用安装指南',
        component: () => import('@/pages/About/PWA.vue'),
      },
      {
        path: 'duty',
        name: 'duty',
        description: '用户协议',
        component: () => import('@/pages/About/DutyPage.vue'),
      },
      {
        path: 'release_notes',
        name: 'release_notes',
        description: '版本说明',
        component: () => import('@/pages/About/ReleaseNotes.vue'),
      },
    ],
  },
  {
    path: '/daily_cover',
    name: 'daily_cover',
    component: () => import('@/pages/DailyCover.vue'),
  },
  {
    path: '/personal',
    name: 'personal',
    component: () => import('@/pages/PersonalPage.vue'),
  },
  {
    path: '/:pathMatch(.*)',
    name: 'NotFound',
    component: () => import('@/pages/NotFound.vue'),
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export { router, routes };
