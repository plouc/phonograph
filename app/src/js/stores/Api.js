var request = require('superagent-bluebird-promise');
var Promise = require('bluebird');

var API_URL = 'http://localhost:2000';

var Api = {
    getArtists(params) {
        params = params || {};
        var req = request.get(`${ API_URL }/artists`);

        if (params.page) {
            req.query({ page: params.page });
        }

        return req.then(res => res.body);
    },

    getArtist(id) {
        var req = request.get(`${ API_URL }/artists/${ id }`);

        return req.then(res => res.body);
    },

    getSimilarArtists(id) {
        var req = request.get(`${ API_URL }/artists/${ id }/similars`);

        return req.then(res => res.body);
    },

    getArtistMasters(id) {
        var req = request.get(`${ API_URL }/artists/${ id }/masters`);

        return req.then(res => res.body);
    },

    getArtistFull(id) {
        return Promise.props({
            artist:   Api.getArtist(id),
            similars: Api.getSimilarArtists(id),
            masters:  Api.getArtistMasters(id)
        });
    },

    getSkill(id) {
        var req = request.get(`${ API_URL }/skills/${ id }`);

        return req.then(res => res.body);
    },

    getSkillArtists(id, params) {
        params = params || {};
        var req = request.get(`${ API_URL }/skills/${ id }/artists`);

        if (params.page) {
            req.query({ page: params.page });
        }

        return req.then(res => res.body);
    },

    getSkillFull(id, params) {
        return Promise.props({
            skill:   Api.getSkill(id),
            artists: Api.getSkillArtists(id, params)
        });
    },

    getSkills(params) {
        params = params || {};
        var req = request.get(`${ API_URL }/skills`);

        if (params.page) {
            req.query({ page: params.page });
        }

        return req.then(res => res.body);
    },

    getMaster(id) {
        var req = request.get(`${ API_URL }/masters/${ id }`);

        return req.then(res => res.body);
    }
};

module.exports = Api;