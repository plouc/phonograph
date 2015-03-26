var React     = require('react');
var Router    = require('react-router');
var Link      = Router.Link;
var Api       = require('./../stores/Api');
var Pager     = require('./Pager.jsx');
var ArtistRow = require('./ArtistRow.jsx');

var Style = React.createClass({
    mixins: [
        Router.State
    ],

    statics: {
        fetchData(params, query) {
            return Api.getStyleFull(params.style_id, {
                page: query.p || 1
            });
        }
    },

    _onPageUpdate(page) {
        var {router} = this.context;
        router.transitionTo('style', {
            style_id: router.getCurrentParams().style_id
        }, {
            p: page
        });
    },

    render() {
        var {style, artists} = this.props.data.style;

        var artistNodes = artists.results.map(artist => {
            return <ArtistRow artist={artist} key={artist.id} />
        });

        return (
            <div>
                <div className="breadcrumbs">
                    <div className="container">
                        <Link to="styles">
                            <i className="fa fa-angle-left"/> styles
                        </Link>
                    </div>
                </div>
                <div className="container">
                    <h2 className="page-title">{style.name}</h2>
                    <Pager pager={artists.pager} handler={this._onPageUpdate}/>
                    <div className="artists__list">
                        {artistNodes}
                    </div>
                </div>
            </div>
        )
    }
});

module.exports = Style;