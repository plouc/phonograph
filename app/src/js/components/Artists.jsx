var React          = require('react');
var Reflux         = require('reflux');
var ArtistsActions = require('./../actions/ArtistsActions');
var ArtistsStore   = require('./../stores/ArtistsStore');
var ArtistRow      = require('./ArtistRow.jsx');
var Pager          = require('./Pager.jsx');
var Router         = require('react-router');
var Link           = Router.Link;


var Artists = React.createClass({
    mixins: [
        Reflux.ListenerMixin
    ],

    getInitialState() {
        return {
            artists: [],
            pager:   null
        }
    },

    componentWillMount() {
        this.listenTo(ArtistsStore, this._onArtistsUpdate);

        ArtistsActions.list();
    },

    _onArtistsUpdate(data) {
        this.setState({
            artists: data.results,
            pager:   data.pager
        });
    },

    _onPageUpdate(page) {
        ArtistsActions.list({
            page: page
        });
    },

    render() {
        var artistNodes = this.state.artists.map(artist => {
            return <ArtistRow artist={artist} key={artist.id} />
        });

        var pagerNode = null;
        if (this.state.pager) {
            pagerNode = <Pager pager={this.state.pager} handler={this._onPageUpdate}/>
        }

        return (
            <div>
                <h2 className="page-title">
                    Artists
                    <Link to="artist_create" className="page-title__action">
                        <i className="fa fa-plus"/>
                    </Link>
                </h2>
                {pagerNode}
                <table>
                    <tbody>
                        {artistNodes}
                    </tbody>
                </table>
            </div>
        );
    }
});

module.exports = Artists;