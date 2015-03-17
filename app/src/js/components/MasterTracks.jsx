var React = require('react');

function formatDuration(duration) {
    var seconds = duration % 60;
    if (seconds < 10) {
        seconds = `0${ seconds}`;
    }
    return `${ Math.floor(duration / 60) }:${ seconds }`;
}

var MasterTracks = React.createClass({
    render() {
        var trackRows = this.props.tracks.map(track => {
            return (
                <tr>
                    <td>
                        {track.name}
                    </td>
                    <td>
                        {formatDuration(track.duration)}
                    </td>
                </tr>
            );
        });

        return (
            <div className="master__tracks">
                <h2 className="master__tracks__title">Tracks</h2>
                <table className="master__tracks__table">
                    <tbody>
                        {trackRows}
                    </tbody>
                </table>
            </div>
        );
    }
});

module.exports = MasterTracks;