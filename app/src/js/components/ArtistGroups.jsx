var React = require('react');
var Link  = require('react-router').Link;

var ArtistGroups = React.createClass({
    render() {
        var groupNodes = <p>This artist does not belong to any group.</p>

        if (this.props.groups.length > 0) {
            groupNodes = this.props.groups.map(artist => {
                return <Link to="artist" params={{ artist_id: artist.id }} className="group" key={artist.id}>{artist.name}</Link>
            });
        }

        return (
            <div className="artist__groups">
                <h4 className="artist__groups__title">Member of groups</h4>
                {groupNodes}
            </div>
        );
    }
});

module.exports = ArtistGroups;