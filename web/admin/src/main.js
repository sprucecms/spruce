var Vue = require('vue');
var VueRouter = require('vue-router');

import auth from './auth/index.js';
import App from './components/App.vue';
import Home from './components/Home.vue';
import Login from './components/Login.vue';
Vue.use(VueRouter);


export var router = new VueRouter({
  routes: [
    { path: '/', component: Home },
    { path: '/login', component: Login }
  ]
});

const app = new Vue({
	router: router,
	data() {
		return {
			user: auth.user
		};
	}
}).$mount('#app');
