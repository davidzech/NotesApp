angular.module('notes').factory("Note", ['$resource', 
function Note($resource) {
    return $resource('notes/:noteId', {
        noteId: '@id'
    }, {
        query: {
            method: 'GET',
            isArray: true
        }
    })
}])