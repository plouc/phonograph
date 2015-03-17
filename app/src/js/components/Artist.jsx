var React          = require('react');
var Reflux         = require('reflux');
var Router         = require('react-router');
var Link           = Router.Link;
var ArtistSkills   = require('./ArtistSkills.jsx');
var ArtistGroups   = require('./ArtistGroups.jsx');
var ArtistMasters  = require('./ArtistMasters.jsx');
var SimilarArtists = require('./SimilarArtists.jsx');
var ArtistsActions = require('./../actions/ArtistsActions');
var ArtistsStore   = require('./../stores/ArtistsStore');

var Artist = React.createClass({
    mixins: [
        Router.State,
        Reflux.ListenerMixin
    ],

    getInitialState() {
        return {
            artist: null
        };
    },

    componentWillReceiveProps() {
        ArtistsActions.get(this.getParams().artist_id);
    },

    componentWillMount() {
        this.listenTo(ArtistsStore, this._onArtistUpdate);
        ArtistsActions.get(this.getParams().artist_id);
    },

    _onArtistUpdate(artist) {
        this.setState({
            artist: artist
        });
    },

    render() {
        if (this.state.artist === null) {
            return <p>loading</p>
        }

        var groupsNode = null;
        if (this.state.artist.groups.length > 0) {
            groupsNode = <ArtistGroups groups={this.state.artist.groups}/>
        }

        return (
            <div>
                <div className="breadcrumbs">
                    <Link to="index">
                        <i className="fa fa-angle-left"/> artists
                    </Link>
                </div>
                <h2 className="page-title">{this.state.artist.name}</h2>
                <div>
                    <ArtistSkills skills={this.state.artist.skills}/>
                </div>
                {groupsNode}
                <ArtistMasters artist={this.state.artist}/>
                <SimilarArtists artist={this.state.artist}/>
            </div>
        )
    }
});

module.exports = Artist;