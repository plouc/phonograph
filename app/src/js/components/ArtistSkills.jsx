var React = require('react');

var ArtistSkills = React.createClass({
    render() {
        var skillNodes = this.props.skills.map(skill => {
            return <span className="skill" key={skill.id}>{skill.name}</span>
        });

        return (
            <span>
                {skillNodes}
            </span>
        );
    }
});

module.exports = ArtistSkills;