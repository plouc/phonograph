var Reflux         = require('reflux');
var ArtistsActions = require('./../actions/ArtistsActions');
var request        = require('superagent');

var _artist;
var _artists = [];

var ArtistsStore = Reflux.createStore({
    listenables: ArtistsActions,

    list(params) {
        params = params || {};

        var req = request.get('http://localhost:2000/artists');
        if (params.page) {
            req.query({ page: params.page });
        }

        req.end((err, res) => {
            _artists = res.body;

            this.trigger(_artists);
        });
    },

    get(id) {
        request.get('http://localhost:2000/artists/' + id).end((err, res) => {
            _artist = res.body;

            this.trigger(_artist);
        });
    }
});

module.exports = ArtistsStore;