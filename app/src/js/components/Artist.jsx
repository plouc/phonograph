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
                    {groupsNode}
                    <ArtistMasters masters={masters}/>
                    <SimilarArtists artists={similars}/>
                </div>
            </div>
        )
    }
});

module.exports = Artist;