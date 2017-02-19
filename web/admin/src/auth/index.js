import {router} from '../main';

export default {

  user: {
    authenticated: false
  },

  login(context, creds, redirect) {
		this.user.authenticated = true;
		console.log(creds);
    router.push('/');
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
