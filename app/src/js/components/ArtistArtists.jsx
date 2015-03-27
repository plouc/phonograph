var React = require('react');
var Link  = require('react-router').Link;

var ArtistArtists = React.createClass({
    render() {
        var artistNodes = <p>Nothing found.</p>;

        if (this.props.artists.length > 0) {
            artistNodes = this.props.artists.map(artist => {
                return <Link to="artist" params={{ artist_id: artist.id }} className="group" key={artist.id}>{artist.name}</Link>
            });
        }

        return (
            <div className="artist__groups">
                <h4 className="artist__groups__title">{this.props.title}</h4>
                {artistNodes}
            </div>
        );
    }
});

module.exports = ArtistArtists;