var React  = require('react');
var Router = require('react-router');
var Link   = Router.Link;
var Pager  = require('./Pager.jsx');
var Api    = require('./../stores/Api');

var Skills = React.createClass({
    mixins: [
        Router.State,
    ],

    contextTypes: {
        router: React.PropTypes.func
    },

    statics: {
        fetchData(params, query) {
            return Api.getSkills({
                page: query.p || 1
            });
        }
    },

    _onPageUpdate(page) {
        this.context.router.transitionTo('skills', {}, {
            p: page
        });
    },

    render() {
        var {results, pager} = this.props.data.skills;

        var skillNodes;
        if (results.length > 0) {
            skillNodes = results.map(skill => {
                return (
                    <div className="list__item">
                        <Link to="skill" params={{ skill_id: skill.id }} key={skill.id}>{skill.name}</Link>
                    </div>
                );
            });
        } else {
            skillNodes = <li>No item found</li>
        }

        return (
            <div className="container">
                <h2 className="page-title">Skills</h2>
                <Pager pager={pager} handler={this._onPageUpdate}/>
                <div className="list">
                    {skillNodes}
                </div>
            </div>
        );
    }
});

module.exports = Skills;