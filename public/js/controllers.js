'use strict';

/* Controllers */

angular.module('raspiMusic.controllers', []).
  controller('MainCtrl', ['$scope', '$http', function($scope,$http) {
  	
  	$http.get('/playlist').success(function (data) {
  		$scope.playlist=data;
  	});

  	$http.get('/songs').success(function (data) {
  		$scope.songs=data;
  	});

  	$scope.stop=function(){
  		$http({method: 'POST', url: '/stop'});
  	};

  	$scope.inPlayList=function(song){
  		for(var key in $scope.playlist) {
        	if(song.file === $scope.playlist[key].file){
        		return "selected";
        		break;
        	}
    	}
    	return "";
  	};

  	$scope.play=function(song){
  		var data={};
  		data.path = song.file;
  		$http({method: 'POST', url: '/songs/play', params: data });
  	};

  }]);