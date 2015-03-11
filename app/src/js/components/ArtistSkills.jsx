var React = require('react');

var ArtistSkills = React.createClass({
    render: function () {
        var skillNodes = this.props.skills.map(function (skill) {
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