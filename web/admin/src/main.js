var Vue = require('vue');
var VueRouter = require('vue-router');

import auth from './auth/index.js';
import App from './components/App.vue';
import Home from './components/Home.vue';
import Login from './components/Login.vue';
import SpruceMenu from './components/SpruceMenu.vue';

Vue.use(VueRouter);


export var router = new VueRouter({
  routes: [
    { path: '/', component: Home },
    { path: '/login', component: Login, meta: { allowAnon: true } }
  ]
});

router.beforeEach(function(to, from, next) {
	console.log("Route change to ", to);
	if(!to.meta.allowAnon && !auth.user.authenticated) {
		// if route requires auth and user isn't authenticated
		next('/login');
	} else {
		next();
	}
});


const app = new Vue({
	router: router,
  components: {
    'spruce-menu': SpruceMenu
  },
	data() {
		return {
			user: auth.user
		};
	}
}).$mount('#app');
