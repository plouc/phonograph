var React  = require('react');
var Reflux = require('reflux');
var Link   = require('react-router').Link;

var ArtistsActions      = require('./../actions/ArtistsActions');
var SimilarArtistsStore = require('./../stores/SimilarArtistsStore');

var SimilarArtists = React.createClass({
    mixins: [
        Reflux.ListenerMixin
    ],

    getInitialState() {
        return {
            artists: []
        };
    },

    componentWillMount() {
        ArtistsActions.similars(this.props.artist.id);
        this.listenTo(SimilarArtistsStore, this._onStoreUpdate);
    },

    _onStoreUpdate(artists) {
        this.setState({
            artists: artists
        });
    },

    render() {
        var artistNodes = this.state.artists.map(artist => {
            return (
                <Link key={artist.id} to="artist" params={{ artist_id: artist.id }} className="similar-artist">
                    {artist.name}
                </Link>
            );
        });

        return (
            <div className="similar-artists">
                <h4 className="similar-artists__title">Similar artists</h4>
                {artistNodes}
            </div>
        );
    }
});

module.exports = SimilarArtists;