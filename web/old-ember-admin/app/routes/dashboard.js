import Ember from 'ember';

export default Ember.Route.extend({
  beforeModel: function() {
    if(!this.session.get("user")) {
      return Ember.RSVP.reject({ status: 'Unauthorized', message: 'Please login to access this page' });
    }
  },
});
