var React         = require('react');
var Reflux        = require('reflux');
var Router        = require('react-router');
var Link          = Router.Link;
var LabelsActions = require('./../actions/LabelsActions');
var LabelsStore   = require('./../stores/LabelsStore');

var Label = React.createClass({
    mixins: [
        Router.State,
        Reflux.ListenerMixin
    ],

    getInitialState() {
        return {
            label: null
        };
    },

    componentWillReceiveProps() {
        LabelsActions.get(this.getParams().label_id);
    },

    componentWillMount() {
        this.listenTo(LabelsStore, this._onLabelUpdate);
        LabelsActions.get(this.getParams().label_id);
    },

    _onLabelUpdate(label) {
        this.setState({
            label: label
        });
    },

    render() {
        if (this.state.label === null) {
            return <p>loading</p>
        }

        return (
            <div>
                <div className="breadcrumbs">
                    <Link to="index">
                        <i className="fa fa-angle-left"/> labels
                    </Link>
                </div>
                <h2 className="page-title">{this.state.label.name}</h2>
            </div>
        )
    }
});

module.exports = Label;