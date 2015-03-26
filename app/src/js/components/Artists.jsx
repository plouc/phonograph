var React     = require('react');
var ArtistRow = require('./ArtistRow.jsx');
var Pager     = require('./Pager.jsx');
var Router    = require('react-router');
var Link      = Router.Link;
var Api       = require('./../stores/Api');

var Artists = React.createClass({
    mixins: [
        Router.State
    ],

    contextTypes: {
        router: React.PropTypes.func
    },

    statics: {
        fetchData(params, query) {
            return Api.getArtists({
                page: query.p || 1
            });
        }
    },

    _onPageUpdate(page) {
        this.context.router.transitionTo('index', {}, {
            p: page
        });
    },

    render() {
        var {results, pager} = this.props.data.index;

        var artistNodes = results.map(artist => {
            return <ArtistRow artist={artist} key={artist.id} />
        });

        return (
            <div className="container">
                <h2 className="page-title">
                    Artists
                    <Link to="artist_create" className="page-title__action">
                        <i className="fa fa-plus"/>
                    </Link>
                </h2>
                <Pager pager={pager} handler={this._onPageUpdate}/>
                <div className="artists__list">
                    {artistNodes}
                </div>
            </div>
        );
    }
});

module.exports = Artists;