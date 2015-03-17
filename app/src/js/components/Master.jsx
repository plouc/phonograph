var React          = require('react');
var Reflux         = require('reflux');
var Router         = require('react-router');
var Link           = Router.Link;
var MastersActions = require('./../actions/MastersActions');
var MastersStore   = require('./../stores/MastersStore');
var ReleaseNodes   = require('./MasterReleases.jsx');

var Master = React.createClass({
    mixins: [
        Router.State,
        Reflux.ListenerMixin
    ],

    getInitialState() {
        return {
            master: null
        };
    },

    componentWillMount() {
        this.listenTo(MastersStore, this._onStoreUpdate);

        MastersActions.get(this.getParams().master_id);
    },

    _onStoreUpdate(master) {
        this.setState({
            master: master
        });
    },

    render() {
        var trackNodes  = null;
        var releaseNode = null;
        var artistNode  = null;

        if (this.state.master !== null) {
            releaseNode = <ReleaseNodes releases={this.state.master.releases}/>

            var artistNodes = this.state.master.artists.map(artist => {
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

            artistNode = <div className="master__artists">by {artistNodes}</div>;
        }



        return (
            <div>
                <div className="breadcrumbs">
                    <Link to="masters">
                        <i className="fa fa-angle-left"/> masters
                    </Link>
                </div>
                <h2 className="page-title">{this.state.master ? this.state.master.name : ''}</h2>
                {artistNode}
                {trackNodes}
                {releaseNode}
            </div>
        );
    }
});

module.exports = Master;