var React  = require('react');
var Reflux = require('reflux');

var Styles = React.createClass({
    mixins: [
        Reflux.ListenerMixin
    ],

    render() {
        return (
            <div>
                <h2 className="page-title">Styles</h2>
            </div>
        );
    }
});

module.exports = Styles;