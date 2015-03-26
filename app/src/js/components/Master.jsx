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

        var stylesNode = null;
        if (master.styles.length > 0) {
            var styleNodes = [];
            master.styles.map((style, i) => {
                styleNodes.push(<span>{style.name}</span>);
                if (i <  master.styles.length - 1) {
                    styleNodes.push(<span>,</span>);
                    styleNodes.push(<span>&nbsp;</span>);
                }
            });

            stylesNode = <div>{styleNodes}</div>;
        }

        var img = null;
        if (master.picture !== '') {
            img = (
                <div className="master__header__picture">
                    <img src={`/images/${ master.picture }`}/>
                </div>
            );
        }


        return (
            <div>
                <div className="breadcrumbs">
                    <div className="container">
                        <Link to="masters">
                            <i className="fa fa-angle-left"/> masters
                        </Link>
                    </div>
                </div>
                <div className="master__header">
                    <div className="master__header__container">
                        {img}
                        <h2 className="page-title">{master.name}</h2>
                        <div className="master__artists">by {artistNodes} - {master.year}</div>
                        {stylesNode}
                    </div>
                </div>
                <div className="container">
                    <MasterTracks tracks={master.tracks}/>
                    <ReleaseNodes releases={master.releases}/>
                </div>
            </div>
        );
    }
});

module.exports = Master;