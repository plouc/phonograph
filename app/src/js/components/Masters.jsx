var React          = require('react');
var Reflux         = require('reflux');
var Link           = require('react-router').Link;
var MastersActions = require('./../actions/MastersActions');
var MastersStore   = require('./../stores/MastersStore');

var Masters = React.createClass({
    mixins: [
        Reflux.ListenerMixin
    ],

    getInitialState() {
        return {
            masters: []
        };
    },

    componentWillMount() {
        this.listenTo(MastersStore, this._onStoreUpdate);

        MastersActions.list();
    },

    _onStoreUpdate(masters) {
        this.setState({ masters: masters });
    },

    render() {
        var masterNodes;
        if (this.state.masters.length > 0) {
            masterNodes = this.state.masters.map(master => {
                return <Link className="master" to="master" params={{ master_id: master.id }} key={master.id}>{master.name}</Link>
            });
        } else {
            masterNodes = <p>No item found</p>
        }

        return (
            <div>
                <h2 className="page-title">Masters</h2>
                {masterNodes}
            </div>
        );
    }
});

module.exports = Masters;