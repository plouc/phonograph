var React  = require('react');
var Reflux = require('reflux');
var Link   = require('react-router').Link;

var MenuStore = require('./../stores/MenuStore');

var Menu = React.createClass({
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

    render() {
        var classes = 'menu';
        if (this.state.active === true) {
            classes += ' _is-open';
        }

        return (
            <div className={classes}>
                <Link key="index"         className="menu__item" to="index">Artists</Link>
                <Link key="artist_create" className="menu__item" to="artist_create">New artist</Link>
                <Link key="masters"       className="menu__item" to="masters">Masters</Link>
                <Link key="styles"        className="menu__item" to="styles">Styles</Link>
                <Link key="skills"        className="menu__item" to="skills">Skills</Link>
                <Link key="labels"        className="menu__item" to="skills">Labels</Link>
            </div>
        );
    }
});

module.exports = Menu;