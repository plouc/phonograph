var React          = require('react');
var Reflux         = require('reflux');
var Link           = require('react-router').Link;
var MastersActions = require('./../actions/MastersActions');
var MastersStore   = require('./../stores/MastersStore');
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
                return (
                    <li>
                        <Link className="master" to="master" params={{ master_id: master.id }} key={master.id}>{master.name}</Link>
                    </li>
                );
            });
        } else {
            masterNodes = <li>No item found</li>
        }

        var pagerNode = null;
        if (this.state.pager) {
            pagerNode = <Pager pager={this.state.pager} handler={this._onPageUpdate}/>
        }

        return (
            <div>
                <h2 className="page-title">Masters</h2>
                <ul>
                    {masterNodes}
                </ul>
                {pagerNode}
            </div>
        );
    }
});

module.exports = Masters;