var React          = require('react');
var Reflux         = require('reflux');
var Router         = require('react-router');
var Link           = Router.Link;
var ReleaseNodes   = require('./MasterReleases.jsx');
var MasterTracks   = require('./MasterTracks.jsx');
var Api            = require('./../stores/Api');

var Master = React.createClass({
    mixins: [
        Router.State
    ],

    statics: {
        fetchData(params) {
            return Api.getMaster(params.master_id);
        }
    },

    render() {
        var {master} = this.props.data;

        var artistNodes = master.artists.map(artist => {
            return (
                <Link ref={artist.id}
                    to="artist"
                    params={{ artist_id: artist.id }}
                    className="master__artist"
                >
                        {artist.name}
                </Link>
            );
        });

        return (
            <div>
                <div className="breadcrumbs">
                    <Link to="masters">
                        <i className="fa fa-angle-left"/> masters
                    </Link>
                </div>
                <h2 className="page-title">{master.name}</h2>
                <div className="master__artists">by {artistNodes} - {master.year}</div>
                <MasterTracks tracks={master.tracks}/>
                <ReleaseNodes releases={master.releases}/>
            </div>
        );
    }
});

module.exports = Master;