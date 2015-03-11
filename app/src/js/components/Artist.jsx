var React  = require('react');
var Reflux = require('reflux');
var Router = require('react-router');

var ArtistSkills   = require('./ArtistSkills.jsx');
var ArtistsActions = require('./../actions/ArtistsActions');
var ArtistsStore   = require('./../stores/ArtistsStore');

var Artist = React.createClass({
    mixins: [
        Router.State,
        Reflux.ListenerMixin
    ],

    getInitialState: function () {
        return {
            artist: null
        };
    },

    componentWillMount: function () {
        this.listenTo(ArtistsStore, this._onArtistUpdate);
        ArtistsActions.get(this.getParams().artist_id);
    },

    _onArtistUpdate: function (artist) {
        this.setState({ artist: artist });
    },

    render: function () {
        if (this.state.artist === null) {
            return <p>loading</p>
        }

        return (
            <div>
                <h2>{this.state.artist.name}</h2>
                <div>
                    <ArtistSkills skills={this.state.artist.skills}/>
                </div>
            </div>
        )
    }
});

module.exports = Artist;