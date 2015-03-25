var React        = require('react');
var Link         = require('react-router').Link;
var ArtistSkills = require('./ArtistSkills.jsx');
var ArtistStyles = require('./ArtistStyles.jsx');

var ArtisRow = React.createClass({
    render() {
        return (
            <div className="artists__list__item">
                <Link className="artists__list__item__name" to="artist" params={{ artist_id: this.props.artist.id }}>{this.props.artist.name}</Link>
                <span>
                    <ArtistSkills skills={this.props.artist.skills}/>
                    <ArtistStyles styles={this.props.artist.styles}/>
                </span>
            </div>
        );
    }
});

module.exports = ArtisRow;