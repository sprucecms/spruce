'use strict';

function WeaveDesktopController($rootScope, $scope, $http, $timeout, $filter, $templateCache) {
 
  $scope.editorData = mockData.xj7kr2;

}

// This is just some scratch code of an eventual complete UI
// system (a module? directive?). It accepts a JSON key and value pair,
// looks for specific UI instructions based on the key, or falls back
// on default instructions based on the type of data.
var UIControlDefaults = [];
var UIControlPrepare = function (key, value) {

  var UIControl = {
    'type' : 'text',
    'label' : key,
    'value' : value
  };

  if (key in UIControlDefaults) {
    // Use specific field settings
  }

};

var myApp = angular.module('weavedesktop', []);

// @end