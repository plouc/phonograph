var Reflux         = require('reflux');
var ArtistsActions = require('./../actions/ArtistsActions');
var request        = require('superagent');

var _artists = [];

var SimilarArtistsStore = Reflux.createStore({
    init: function () {
        this.listenTo(ArtistsActions.similars, this.getSimilarArtists);
    },

    getSimilarArtists: function (id) {
        request.get('http://localhost:2000/artists/' + id + '/similars').end(function (err, res) {
            _artists = res.body;

            this.trigger(_artists);
        }.bind(this));
    }
});

module.exports = SimilarArtistsStore;