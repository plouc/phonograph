var React          = require('react');
var Reflux         = require('reflux');
var Router         = require('react-router');
var MastersActions = require('./../actions/MastersActions');
var MastersStore   = require('./../stores/MastersStore');

var Master = React.createClass({
    mixins: [
        Router.State,
        Reflux.ListenerMixin
    ],

    getInitialState() {
        return {
            masters: null
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
        var trackNodes = null;

        return (
            <div>
                <h2 className="page-title">{this.state.master ? this.state.master.name : ''}</h2>
                {trackNodes}
            </div>
        );
    }
});

module.exports = Master;