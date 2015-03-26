var React          = require('react');
var Reflux         = require('reflux');
var Link           = require('react-router').Link;
var MastersActions = require('./../actions/MastersActions');
var MastersStore   = require('./../stores/MastersStore');
var MasterRow      = require('./MasterRow.jsx');
var Pager          = require('./Pager.jsx');

var Masters = React.createClass({
    mixins: [
        Reflux.ListenerMixin
    ],

    getInitialState() {
        return {
            masters: [],
            pager:   null
        };
    },

    componentWillMount() {
        this.listenTo(MastersStore, this._onStoreUpdate);

        MastersActions.list();
    },

    _onStoreUpdate(data) {
        this.setState({
            masters: data.results,
            pager:   data.pager
        });
    },

    _onPageUpdate(page) {
        MastersActions.list({
            page: page
        });
    },

    render() {
        var masterNodes;
        if (this.state.masters.length > 0) {
            masterNodes = this.state.masters.map(master => {
                return <MasterRow key={master.id} master={master}/>;
            })
        } else {
            masterNodes = <li>No item found</li>
        }

        var pagerNode = null;
        if (this.state.pager) {
            pagerNode = <Pager pager={this.state.pager} handler={this._onPageUpdate}/>
        }

        return (
            <div className="container">
                <h2 className="page-title">Masters</h2>
                <div className="list">
                    {masterNodes}
                </div>
                {pagerNode}
            </div>
        );
    }
});

module.exports = Masters;