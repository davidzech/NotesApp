var webtutorialApp = angular.module('webTutorial', []);


// TODO Build out the NotesService
// Make NotesController an individual component to include in the home page

// webtutorialApp.factory('NotesService', function($http) {
//     var notes = {};
//     // Or put this in a method on the returned object
//     // if you want to invoke at some point other than load
//     $http.get('/notes').then(function(response) {
//         notes = response;
//     });
//     return {
//         getNotes: function() {
//             return notes;
//         }
//     };
// })

webtutorialApp.controller('NotesController', function NotesController($scope, $http) {
    $scope.world = "World!"

    $http.get('/notes').then(function(response) {
        $scope.notes = response.data;
    })
    
})

webtutorialApp.controller("OtherController", function OtherController($scope) {
    $scope.world = "No"
})