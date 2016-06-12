import Ember from 'ember';
import UnauthenticatedRouteMixin from 'ember-simple-auth/mixins/unauthenticated-route-mixin';

export default UnauthenticatedRouteMixin.extend({
  actions: {
    login: function() {
      var email = this.controller.get('email');
      var password = this.controller.get('password');

      this.get('session').authenticate('authenticator:oauth2', email, password).catch((reason) => {
        this.set('errorMessage', reason.error || reason);
      });
    }
  }
});
