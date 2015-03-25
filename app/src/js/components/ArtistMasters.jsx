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
        var contentNode;
        if (this.state.masters.length > 0) {
            var masterNodes = this.state.masters.map(master => {
                return (
                    <div className="list__item list__item--sm">
                        <Link key={master.id} to="master" params={{ master_id: master.id }} className="artist__master">
                            {master.name}
                        </Link>
                    </div>
                );
            });
            contentNode = (
                <div className="list">
                    {masterNodes}
                </div>
            );
        } else {
            contentNode = <p>No item found.</p>
        }

        return (
            <div className="artist__masters">
                <h4 className="artist__masters__title">Played in</h4>
                {contentNode}
            </div>
        );
    }
});

module.exports = ArtistMasters;