var Vue = require('vue');
import {router} from '../main';

export default {

  user: {
    authenticated: false
  },

  login(context, creds, redirect) {
		console.log(creds);
    var formData = new FormData();
    formData.append('username', creds.username);
    formData.append('password', creds.password);
		Vue.http.post(
			'/api/v1/oauth2/token', formData
		).then(response => {
        console.log(response.body);
        console.log(response.data);
				context.error = false;
				localStorage.setItem('id_token', response.data.access_token);
				Vue.http.headers.common.Authorization = 'Bearer ' + localStorage.getItem('id_token');
				this.user.authenticated = true;
				router.push('/');
		}, response => {
				context.error = true;
				router.push('/login');
		});
  },

  logout() {
  },

  checkAuth() {
    if(!user.authenticated) {
      router.push('/login');
    }
  },

  getAuthHeader() {
  }
};
