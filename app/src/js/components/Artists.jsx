var React          = require('react');
var Reflux         = require('reflux');
var ArtistsActions = require('./../actions/ArtistsActions');
var ArtistsStore   = require('./../stores/ArtistsStore');
var ArtistRow      = require('./ArtistRow.jsx');

var Artists = React.createClass({
    mixins: [
        Reflux.ListenerMixin
    ],

    getInitialState() {
        return {
            artists: []
        }
    },

    componentWillMount() {
        this.listenTo(ArtistsStore, this._onArtistsUpdate);

        ArtistsActions.list();
    },

    _onArtistsUpdate(artists) {
        this.setState({ artists: artists });
    },

    render() {
        var artistNodes = this.state.artists.map(artist => {
            return <ArtistRow artist={artist} key={artist.id} />
        });

        return (
            <div>
                <h2 className="page-title">Artists</h2>
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