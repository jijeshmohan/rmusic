'use strict';


// Declare app level module which depends on filters, and services
angular.module('raspiMusic', [
  'ngRoute',
  'raspiMusic.filters',
  'raspiMusic.services',
  'raspiMusic.directives',
  'raspiMusic.controllers'
]).
config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/songs', {templateUrl: 'partials/songslist.html', controller: 'MainCtrl'});
  $routeProvider.otherwise({redirectTo: '/songs'});
}]);
