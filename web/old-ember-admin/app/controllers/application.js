import Ember from 'ember';

export default Ember.Controller.extend({
  session: Ember.inject.service('session'),
  firstName: "Aaron",
  lastName: "Longwell",
  actions: {
    invalidateSession() {
      this.get('session').invalidate();
    }
  },
});
