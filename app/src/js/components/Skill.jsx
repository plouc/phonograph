var React     = require('react');
var Reflux    = require('reflux');
var Router    = require('react-router');
var Link      = Router.Link;
var Api       = require('./../stores/Api');
var Pager     = require('./Pager.jsx');
var ArtistRow = require('./ArtistRow.jsx');

var Artist = React.createClass({
    mixins: [
        Reflux.ListenerMixin,
        Router.State
    ],

    statics: {
        fetchData(params, query) {
            return Api.getSkillFull(params.skill_id, {
                page: query.p || 1
            });
        }
    },

    _onPageUpdate(page) {
        var {router} = this.context;
        console.log('_onPageUpdate', router.getCurrentParams().skill_id, page);
        router.transitionTo('skill', {
            skill_id: router.getCurrentParams().skill_id
        }, {
            p: page
        });
    },

    render() {
        var {skill, artists} = this.props.data.skill;

        var artistNodes = artists.results.map(artist => {
            return <ArtistRow artist={artist} key={artist.id} />
        });

        return (
            <div>
                <div className="breadcrumbs">
                    <Link to="skills">
                        <i className="fa fa-angle-left"/> skills
                    </Link>
                </div>
                <h2 className="page-title">{skill.name}</h2>
                <Pager pager={artists.pager} handler={this._onPageUpdate}/>
                <div className="artists__list">
                    {artistNodes}
                </div>
            </div>
        )
    }
});

module.exports = Artist;