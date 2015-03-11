var React  = require('react');
var Reflux = require('reflux');
var Link   = require('react-router').Link;

var MastersActions = require('./../actions/MastersActions');
var MastersStore   = require('./../stores/MastersStore');

var ArtistMasters = React.createClass({
    mixins: [
        Reflux.ListenerMixin
    ],

    getInitialState() {
        return {
            masters: []
        };
    },

    componentWillMount() {
        MastersActions.playedBy(this.props.artist.id);
        this.listenTo(MastersStore, this._onStoreUpdate);
    },

    _onStoreUpdate(masters) {
        this.setState({
            masters: masters
        });
    },

    render() {
        var masterNodes
        if (this.state.masters.length > 0) {
            masterNodes = this.state.masters.map(master => {
                return (
                    <Link key={master.id} to="artist" params={{ artist_id: master.id }} className="artist__master">
                        {master.name}
                    </Link>
                );
            });
        } else {
            masterNodes = <p>No item found.</p>
        }

        return (
            <div className="artist__masters">
                <h4 className="artist__masters__title">Played in</h4>
                {masterNodes}
            </div>
        );
    }
});

module.exports = ArtistMasters;