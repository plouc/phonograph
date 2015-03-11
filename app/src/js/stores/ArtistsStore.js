var Reflux         = require('reflux');
var ArtistsActions = require('./../actions/ArtistsActions');
var request        = require('superagent');

var _artist;
var _artists = [];

var ArtistsStore = Reflux.createStore({
    listenables: ArtistsActions,

    list: function () {
        request.get('http://localhost:2000/artists').end(function (err, res) {
            _artists = res.body;

            this.trigger(_artists);
        }.bind(this));
    },

    get: function (id) {
        request.get('http://localhost:2000/artists/' + id).end(function (err, res) {
            _artist = res.body;

            console.log(_artist);

            this.trigger(_artist);
        }.bind(this));
    }
});

module.exports = ArtistsStore;