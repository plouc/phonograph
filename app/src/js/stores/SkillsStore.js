var Reflux        = require('reflux');
var SkillsActions = require('./../actions/SkillsActions');
var request       = require('superagent');

var _skill;
var _skills = [];

var SkillsStore = Reflux.createStore({
    listenables: SkillsActions,

    list(params) {
        params = params || {};

        var req = request.get('http://localhost:2000/skills');
        if (params.page) {
            req.query({ page: params.page });
        }

        req.end((err, res) => {
            _skills = res.body;

            this.trigger(_skills);
        });
    },

    get(id) {
        request.get('http://localhost:2000/skills/' + id).end((err, res) => {
            _skill = res.body;

            this.trigger(_skill);
        });
    }
});

module.exports = SkillsStore;