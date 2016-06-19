import Ember from 'ember';

export default Ember.Route.extend({
  sessoion: Ember.inject.service('session'),

  actions: {
    login: function() {
      var email = this.controller.get('email');
      var password = this.controller.get('password');

      this.get('session').authenticate('authenticator:oauth2', email, password).catch(function(reason) {
        alert(reason.error || error);
        this.set('errorMessage', reason.error || reason);
      });
    }
  }
});
