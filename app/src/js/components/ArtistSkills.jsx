var React  = require('react');
var Router = require('react-router');
var Link   = Router.Link;

var ArtistSkills = React.createClass({
    render() {
        if (this.props.skills.length === 0) {
            return null;
        }

        var skillNodes = [];
        this.props.skills.forEach((skill, i) => {
            skillNodes.push(<Link to="skill" params={{ skill_id: skill.id }} className="artists__list__skill" key={skill.id}>{skill.name}</Link>);
            if (i < this.props.skills.length - 1) {
                skillNodes.push(<span>,</span>);
                skillNodes.push(<span>&nbsp;</span>);
            }
        });

        var classes = 'artist__skills';
        if (this.props.mode === 'list') {
            classes += ' artist__skills--list';
        }

        return (
            <div className={classes}>
                <span className="artist__skills__title">skills</span>
                {skillNodes}
            </div>
        );
    }
});

module.exports = ArtistSkills;