var Reflux        = require('reflux');
var LabelsActions = require('./../actions/LabelsActions');
var request       = require('superagent');

var _label;
var _labels = [];

var LabelsStore = Reflux.createStore({
    listenables: LabelsActions,

    list(params) {
        params = params || {};

        var req = request.get('http://localhost:2000/labels');
        if (params.page) {
            req.query({ page: params.page });
        }

        req.end((err, res) => {
            _labels = res.body;

            this.trigger(_labels);
        });
    },

    get(id) {
        request.get('http://localhost:2000/labels/' + id).end((err, res) => {
            _label = res.body;

            this.trigger(_label);
        });
    }
});

module.exports = LabelsStore;