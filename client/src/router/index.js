import Vue from 'vue';
import Router from 'vue-router';
import Login from '@/pages/Login';
import Main from '@/pages/Main';
import axios from 'axios';
import VueAxios from 'vue-axios';

Vue.use(Router);
Vue.use(VueAxios, axios);
Vue.axios.defaults.baseURL = 'http://localhost:9000/api/v1';

const router = new Router({
  routes: [
    { path: '/', name: 'main', component: Main, meta: { requiresAuth: true } },
    {
      path: '/login',
      name: 'Login',
      component: Login,
      meta: { redirectIfAuth: true },
    },
  ],
});

function checkIfAuth() {
  const hasToken = (localStorage.getItem('amaterasu_token') !== null);
  const expiryDate = localStorage.getItem('amaterasu_token_expire');
  if (hasToken && expiryDate !== null) {
    return new Date(expiryDate) > new Date();
  }
  localStorage.removeItem('amaterasu_token_expire');
  return false;
}

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (checkIfAuth()) {
      next();
    } else {
      next({
        path: '/login',
        query: {
          redirect: to.fullPath,
        },
      });
    }
  } else if (to.matched.some(record => record.meta.redirectIfAuth)) {
    // redirect to main page if already authenticated.
    if (checkIfAuth()) {
      next({
        path: '/',
      });
    } else {
      next();
    }
  } else {
    next();
  }
});

export default router;
