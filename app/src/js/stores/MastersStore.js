var Reflux         = require('reflux');
var MastersActions = require('./../actions/MastersActions');
var request        = require('superagent');

var _master;
var _masters = [];

var MastersStore = Reflux.createStore({
    listenables: MastersActions,

    list() {
        request.get('http://localhost:2000/masters').end((err, res) => {
            _masters = res.body;

            this.trigger(_masters);
        });
    },

    get(id) {
        request.get('http://localhost:2000/masters/' + id).end((err, res) => {
            _master = res.body;

            this.trigger(_master);
        });
    },

    playedBy(artistId) {
        request.get('http://localhost:2000/artists/' + artistId + '/masters').end((err, res) => {
            _masters = res.body;

            this.trigger(_masters);
        });
    }
});

module.exports = MastersStore;