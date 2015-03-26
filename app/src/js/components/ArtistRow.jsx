var React        = require('react');
var Link         = require('react-router').Link;
var ArtistSkills = require('./ArtistSkills.jsx');
var ArtistStyles = require('./ArtistStyles.jsx');

var ArtisRow = React.createClass({
    render() {
        var imgNode = null;
        if (this.props.artist.picture !== '') {
            imgNode = (
                <div className="artists__list__item__picture">
                    <img src={`/images/${ this.props.artist.picture }`}/>
                </div>
            );
        }

        return (
            <div className="artists__list__item">
                {imgNode}
                <Link className="artists__list__item__name" to="artist" params={{ artist_id: this.props.artist.id }}>{this.props.artist.name}</Link>
                <span>
                    <ArtistSkills skills={this.props.artist.skills} mode="list"/>
                    <ArtistStyles styles={this.props.artist.styles} mode="list"/>
                </span>
            </div>
        );
    }
});

module.exports = ArtisRow;