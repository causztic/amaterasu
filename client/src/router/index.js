import Vue from 'vue';
import Router from 'vue-router';
import Main from '@/components/Main';
import axios from 'axios';
import VueAxios from 'vue-axios';
import VueAuth from '@websanova/vue-auth';

import Bearer from '@websanova/vue-auth/drivers/auth/bearer';
import HttpDriver from '@websanova/vue-auth/drivers/http/axios.1.x';
import RouterDriver from '@websanova/vue-auth/drivers/router/vue-router.2.x';

Vue.use(Router);
Vue.use(VueAxios, axios);

const router = new Router({
  routes: [
    {
      path: '/',
      name: 'Main',
      component: Main,
    },
  ],
});

Vue.router = router;

Vue.axios.defaults.baseURL = 'localhost:9000/api/v1';
Vue.use(VueAuth, {
  auth: Bearer,
  http: HttpDriver,
  router: RouterDriver,
});

export default router;
