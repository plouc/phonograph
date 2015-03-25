var React = require('react');
var Link  = require('react-router').Link;

var SimilarArtists = React.createClass({
    render() {
        var artistNodes = this.props.artists.map(artist => {
            return (
                <Link key={artist.id} to="artist" params={{ artist_id: artist.id }} className="similar-artist">
                    {artist.name}
                </Link>
            );
        });

        return (
            <div className="similar-artists">
                <h4 className="similar-artists__title">Similar artists</h4>
                {artistNodes}
            </div>
        );
    }
});

module.exports = SimilarArtists;