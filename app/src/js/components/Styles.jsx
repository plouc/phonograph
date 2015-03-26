var React  = require('react');
var Router = require('react-router');
var Pager  = require('./Pager.jsx');
var Row    = require('./StyleRow.jsx');
var Api    = require('./../stores/Api');

var Styles = React.createClass({
    mixins: [
        Router.State
    ],

    contextTypes: {
        router: React.PropTypes.func
    },

    statics: {
        fetchData(params, query) {
            return Api.getStyles({
                page: query.p || 1
            });
        }
    },

    _onPageUpdate(page) {
        this.context.router.transitionTo('styles', {}, {
            p: page
        });
    },

    render() {
        var {results, pager} = this.props.data.styles;

        var rows = results.map(style => {
            return <Row style={style}/>;
        });

        return (
            <div className="container">
                <h2 className="page-title">Styles</h2>
                <Pager pager={pager} handler={this._onPageUpdate}/>
                {rows}
            </div>
        );
    }
});

module.exports = Styles;