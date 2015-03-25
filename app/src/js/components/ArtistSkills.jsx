var React = require('react');

var ArtistSkills = React.createClass({
    render() {
        if (this.props.skills.length === 0) {
            return null;
        }

        var skillNodes = [];
        this.props.skills.forEach((skill, i) => {
            skillNodes.push(<span className="artists__list__skill" key={skill.id}>{skill.name}</span>);
            if (i < this.props.skills.length - 1) {
                skillNodes.push(<span>,</span>);
                skillNodes.push(<span>&nbsp;</span>);
            }
        });

        return (
            <div className="artists__list__skills">
                skills: {skillNodes}
            </div>
        );
    }
});

module.exports = ArtistSkills;