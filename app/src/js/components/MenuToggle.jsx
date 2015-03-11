var React  = require('react');
var Reflux = require('reflux');

var MenuActions = require('./../actions/MenuActions');
var MenuStore   = require('./../stores/MenuStore');

var MenuToggle = React.createClass({
    mixins: [
        Reflux.ListenerMixin
    ],

    getInitialState() {
        return {
            active: false
        };
    },

    componentWillMount() {
        this.listenTo(MenuStore, this._onStoreUpdate);
    },

    _onStoreUpdate(active) {
        this.setState({
            active: active
        });
    },

    _onClick() {
        MenuActions.toggle();
    },

    render() {
        var classes = 'header__menu';
        if (this.state.active === true) {
            classes += ' _is-open';
        }

        return (
            <div className={classes} onClick={this._onClick}>
                <i className="fa fa-close"/>
                <i className="fa fa-bars"/>
            </div>
        );
    }
});

module.exports = MenuToggle;