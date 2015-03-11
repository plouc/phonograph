var Reflux = require('reflux');

var ArtistsActions = Reflux.createActions([
    'list',
    'get',
    'similars'
]);

module.exports = ArtistsActions;