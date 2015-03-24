var React          = require('react');
var Reflux         = require('reflux');
var SkillsActions  = require('./../actions/SkillsActions');
var SkillsStore    = require('./../stores/SkillsStore');
var ArtistsActions = require('./../actions/ArtistsActions');

var ArtistForm = React.createClass({
    mixins: [
        Reflux.ListenerMixin
    ],

    getInitialState() {
        return {
            data: {
                name:   '',
                skills: []
            },
            skills: []
        };
    },

    componentWillMount() {
        this.listenTo(SkillsStore, this._onSkillsUpdate);
        SkillsActions.list();
    },

    _onSkillsUpdate(skills) {
        this.setState({
            skills: skills
        });
    },

    _handleChange(event) {
        this.setState({
            data: {
                name: event.target.value
            }
        });
    },

    _onSubmit() {
        console.log(this.state);
    },

    render() {
        var data = this.state.data;

        var skillsSelector = null;
        if (this.state.skills.length > 0) {
            var optionNodes = this.state.skills.map(skill => {
                return (
                    <option value={skill.id}>{skill.name}</option>
                );
            });

            skillsSelector = (
                <select>
                {optionNodes}
                </select>
            );
        }

        return (
            <div>
                <p>{name}</p>
                <input className="form__control" type="text" value={data.name} onChange={this._handleChange} />
                {skillsSelector}
                <button onClick={this._onSubmit}>
                    submit
                </button>
            </div>
        );
    }
});

module.exports = ArtistForm;