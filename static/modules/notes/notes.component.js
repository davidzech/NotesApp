var app = angular.module('notes');


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

app.component('notesList', {
    templateUrl: "modules/notes/notes.template.html",
    controller: function NotesController($http) {
        var $ctrl = this
        $http.get('/notes').then(function(response) {
            $ctrl.notes = response.data; // != this.notes
        });
    }
})