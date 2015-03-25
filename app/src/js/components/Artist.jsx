var React          = require('react');
var Router         = require('react-router');
var Link           = Router.Link;
var ArtistSkills   = require('./ArtistSkills.jsx');
var ArtistStyles   = require('./ArtistStyles.jsx');
var ArtistGroups   = require('./ArtistGroups.jsx');
var ArtistMasters  = require('./ArtistMasters.jsx');
var SimilarArtists = require('./SimilarArtists.jsx');
var Api            = require('./../stores/Api');

var Artist = React.createClass({
    mixins: [
        Router.State
    ],

    statics: {
        fetchData(params) {
            return Api.getArtistFull(params.artist_id);
        }
    },

    render() {
        var {artist, similars, masters} = this.props.data.artist;

        var groupsNode = null;
        if (artist.groups.length > 0) {
            groupsNode = <ArtistGroups groups={artist.groups}/>
        }

        return (
            <div>
                <div className="breadcrumbs">
                    <Link to="index">
                        <i className="fa fa-angle-left"/> artists
                    </Link>
                </div>
                <h2 className="page-title">{artist.name}</h2>
                <div>
                    <ArtistSkills skills={artist.skills}/>
                    <ArtistStyles styles={artist.styles}/>
                </div>
                {groupsNode}
                <ArtistMasters masters={masters}/>
                <SimilarArtists artists={similars}/>
            </div>
        )
    }
});

module.exports = Artist;