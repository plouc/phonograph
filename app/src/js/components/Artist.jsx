var React          = require('react');
var Router         = require('react-router');
var Link           = Router.Link;
var ArtistSkills   = require('./ArtistSkills.jsx');
var ArtistStyles   = require('./ArtistStyles.jsx');
var ArtistArtists  = require('./ArtistArtists.jsx');
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
            groupsNode = <ArtistArtists artists={artist.groups} title="Member of groups"/>
        }

        var membersNode = null;
        if (artist.members.length > 0) {
            membersNode = <ArtistArtists artists={artist.members} title="Members"/>
        }

        var img = null;
        if (artist.picture !== '') {
            img = (
                <div className="artist__header__picture">
                    <img src={`/images/${ artist.picture }`}/>
                </div>
            );
        }

        return (
            <div>
                <div className="breadcrumbs">
                    <div className="container">
                        <Link to="index">
                            <i className="fa fa-angle-left"/> artists
                        </Link>
                    </div>
                </div>
                <div className="artist__header">
                    <div className="artist__header__container">
                        {img}
                        <h2 className="page-title">{artist.name}</h2>
                        <div className="cf">
                            <ArtistSkills skills={artist.skills}/>
                            <ArtistStyles styles={artist.styles}/>
                        </div>
                    </div>
                </div>
                <div className="container">
                    {membersNode}
                    {groupsNode}
                    <ArtistMasters masters={masters}/>
                    <SimilarArtists artists={similars}/>
                </div>
            </div>
        )
    }
});

module.exports = Artist;