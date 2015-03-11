var React        = require('react');
var Link         = require('react-router').Link;
var ArtistSkills = require('./ArtistSkills.jsx');

var ArtisRow = React.createClass({
    render: function () {
        return (
            <tr>
                <td>
                    <Link to="artist" params={{ artist_id: this.props.artist.id }}>{this.props.artist.name}</Link>
                </td>
                <td>
                    <ArtistSkills skills={this.props.artist.skills} />
                </td>
            </tr>
        );
    }
});

module.exports = ArtisRow;