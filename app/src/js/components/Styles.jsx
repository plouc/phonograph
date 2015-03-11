var React  = require('react');
var Reflux = require('reflux');

var Styles = React.createClass({
    mixins: [
        Reflux.ListenerMixin
    ],

    componentWillMount() {

    },

    render() {
        return (
            <div>
                <h2 className="page-title">Styles</h2>
                <table>
                    <tbody>
                    </tbody>
                </table>
            </div>
        );
    }
});

module.exports = Styles;