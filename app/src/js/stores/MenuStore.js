var Reflux      = require('reflux');
var MenuActions = require('./../actions/MenuActions');

var _active = false;

var MenuStore = Reflux.createStore({
    listenables: MenuActions,

    toggle() {
        _active = !_active;

        this.trigger(_active);
    },

    close() {
        _active = false;

        this.trigger(_active);
    }
});

module.exports = MenuStore;